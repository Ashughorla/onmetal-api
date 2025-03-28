/*
 * Copyright (c) 2022 by the OnMetal authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package validation

import (
	"fmt"

	"github.com/onmetal/onmetal-api/admission/plugin/machinevolumedevices/device"
	onmetalapivalidation "github.com/onmetal/onmetal-api/api/validation"
	"github.com/onmetal/onmetal-api/apis/compute"
	"github.com/onmetal/onmetal-api/apis/storage"
	storagevalidation "github.com/onmetal/onmetal-api/apis/storage/validation"
	corev1 "k8s.io/api/core/v1"
	apivalidation "k8s.io/apimachinery/pkg/api/validation"
	metav1validation "k8s.io/apimachinery/pkg/apis/meta/v1/validation"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

// ValidateMachine validates a Machine object.
func ValidateMachine(machine *compute.Machine) field.ErrorList {
	var allErrs field.ErrorList

	allErrs = append(allErrs, apivalidation.ValidateObjectMetaAccessor(machine, true, apivalidation.NameIsDNSLabel, field.NewPath("metadata"))...)
	allErrs = append(allErrs, validateMachineSpec(&machine.Spec, field.NewPath("spec"))...)

	return allErrs
}

// ValidateMachineUpdate validates a Machine object before an update.
func ValidateMachineUpdate(newMachine, oldMachine *compute.Machine) field.ErrorList {
	var allErrs field.ErrorList

	allErrs = append(allErrs, apivalidation.ValidateObjectMetaAccessorUpdate(newMachine, oldMachine, field.NewPath("metadata"))...)
	allErrs = append(allErrs, validateMachineSpecUpdate(&newMachine.Spec, &oldMachine.Spec, newMachine.DeletionTimestamp != nil, field.NewPath("spec"))...)
	allErrs = append(allErrs, ValidateMachine(newMachine)...)

	return allErrs
}

// validateMachineSpec validates the spec of a Machine object.
func validateMachineSpec(machineSpec *compute.MachineSpec, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList

	if machineSpec.MachineClassRef == (corev1.LocalObjectReference{}) {
		allErrs = append(allErrs, field.Required(fldPath.Child("machineClassRef"), "must specify a machine class ref"))
	}

	for _, msg := range apivalidation.NameIsDNSLabel(machineSpec.MachineClassRef.Name, false) {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("machineClassRef").Child("name"), machineSpec.MachineClassRef.Name, msg))
	}

	if machineSpec.MachinePoolRef != nil {
		for _, msg := range ValidateMachinePoolName(machineSpec.MachinePoolRef.Name, false) {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("machinePoolRef").Child("name"), machineSpec.MachinePoolRef.Name, msg))
		}
	}

	if machineSpec.IgnitionRef != nil && machineSpec.IgnitionRef.Name != "" {
		for _, msg := range apivalidation.NameIsDNSLabel(machineSpec.IgnitionRef.Name, false) {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("ignitionRef").Child("name"), machineSpec.IgnitionRef.Name, msg))
		}
	}

	if machineSpec.ImagePullSecretRef != nil {
		for _, msg := range apivalidation.NameIsDNSLabel(machineSpec.ImagePullSecretRef.Name, false) {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("imagePullSecretRef").Child("name"), machineSpec.ImagePullSecretRef.Name, msg))
		}
	}

	seenNames := sets.NewString()
	seenDevices := sets.NewString()
	for i, vol := range machineSpec.Volumes {
		if seenNames.Has(vol.Name) {
			allErrs = append(allErrs, field.Duplicate(fldPath.Child("volume").Index(i).Child("name"), vol.Name))
		} else {
			seenNames.Insert(vol.Name)
		}
		if seenDevices.Has(vol.Device) {
			allErrs = append(allErrs, field.Duplicate(fldPath.Child("volume").Index(i).Child("device"), vol.Device))
		} else {
			seenDevices.Insert(vol.Device)
		}
		allErrs = append(allErrs, validateVolume(&vol, fldPath.Child("volume").Index(i))...)
	}

	allErrs = append(allErrs, metav1validation.ValidateLabels(machineSpec.MachinePoolSelector, fldPath.Child("machinePoolSelector"))...)

	return allErrs
}

func isReservedDeviceName(prefix string, idx int) bool {
	return idx == 0
}

func validateVolume(volume *compute.Volume, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList

	for _, msg := range apivalidation.NameIsDNSLabel(volume.Name, false) {
		allErrs = append(allErrs, field.Invalid(fldPath.Child("name"), volume.Name, msg))
	}

	if volume.Device == "" {
		allErrs = append(allErrs, field.Required(fldPath.Child("device"), "must specify device"))
	} else {
		// TODO: Improve validation on prefix.
		prefix, idx, err := device.ParseName(volume.Device)
		if err != nil {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("device"), volume.Device, fmt.Sprintf("invalid device name: %v", err)))
		} else {
			if isReservedDeviceName(prefix, idx) {
				allErrs = append(allErrs, field.Forbidden(fldPath.Child("device"), fmt.Sprintf("device name %s is reserved", volume.Device)))
			}
		}
	}

	allErrs = append(allErrs, validateVolumeSource(&volume.VolumeSource, fldPath)...)

	return allErrs
}

func validateVolumeSource(source *compute.VolumeSource, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList

	var numDefs int
	if source.VolumeRef != nil {
		numDefs++
		for _, msg := range apivalidation.NameIsDNSLabel(source.VolumeRef.Name, false) {
			allErrs = append(allErrs, field.Invalid(fldPath.Child("volumeRef").Child("name"), source.VolumeRef.Name, msg))
		}
	}
	if source.EmptyDisk != nil {
		if numDefs > 0 {
			allErrs = append(allErrs, field.Forbidden(fldPath.Child("emptyDisk"), "must only specify one volume source"))
		} else {
			numDefs++
			allErrs = append(allErrs, validateEmptyDiskVolumeSource(source.EmptyDisk, fldPath.Child("emptyDisk"))...)
		}
	}
	if source.Ephemeral != nil {
		if numDefs > 0 {
			allErrs = append(allErrs, field.Forbidden(fldPath.Child("ephemeral"), "must only specify one volume source"))
		} else {
			numDefs++
			allErrs = append(allErrs, validateEphemeralVolumeSource(source.Ephemeral, fldPath.Child("ephemeral"))...)
		}
	}
	if numDefs == 0 {
		allErrs = append(allErrs, field.Invalid(fldPath, source, "must specify at least one volume source"))
	}

	return allErrs
}

func validateEmptyDiskVolumeSource(source *compute.EmptyDiskVolumeSource, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList

	if sizeLimit := source.SizeLimit; sizeLimit != nil {
		allErrs = append(allErrs, onmetalapivalidation.ValidateNonNegativeQuantity(*sizeLimit, fldPath.Child("sizeLimit"))...)
	}

	return allErrs
}

func validateEphemeralVolumeSource(source *compute.EphemeralVolumeSource, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList

	if source.VolumeTemplate == nil {
		allErrs = append(allErrs, field.Required(fldPath.Child("volumeTemplate"), "must specify volume template"))
	} else {
		allErrs = append(allErrs, validateVolumeTemplateSpecForMachine(source.VolumeTemplate, fldPath.Child("volumeTemplate"))...)
	}

	return allErrs
}

func validateVolumeTemplateSpecForMachine(template *storage.VolumeTemplateSpec, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList

	if template == nil {
		allErrs = append(allErrs, field.Required(fldPath, ""))
	} else {
		allErrs = append(allErrs, storagevalidation.ValidateVolumeTemplateSpec(template, fldPath)...)

		if template.Spec.ClaimRef != nil {
			allErrs = append(allErrs, field.Forbidden(fldPath.Child("spec", "claimRef"), "may not specify claimRef"))
		}
		if template.Spec.Unclaimable {
			allErrs = append(allErrs, field.Forbidden(fldPath.Child("spec", "unclaimable"), "may not specify unclaimable"))
		}
	}

	return allErrs
}

// validateMachineSpecUpdate validates the spec of a Machine object before an update.
func validateMachineSpecUpdate(new, old *compute.MachineSpec, deletionTimestampSet bool, fldPath *field.Path) field.ErrorList {
	var allErrs field.ErrorList

	allErrs = append(allErrs, onmetalapivalidation.ValidateImmutableField(new.Image, old.Image, fldPath.Child("image"))...)
	allErrs = append(allErrs, onmetalapivalidation.ValidateImmutableField(new.MachineClassRef, old.MachineClassRef, fldPath.Child("machineClassRef"))...)
	allErrs = append(allErrs, onmetalapivalidation.ValidateSetOnceField(new.MachinePoolRef, old.MachinePoolRef, fldPath.Child("machinePoolRef"))...)

	return allErrs
}
