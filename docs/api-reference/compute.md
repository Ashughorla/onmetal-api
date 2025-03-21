<p>Packages:</p>
<ul>
<li>
<a href="#compute.api.onmetal.de%2fv1alpha1">compute.api.onmetal.de/v1alpha1</a>
</li>
</ul>
<h2 id="compute.api.onmetal.de/v1alpha1">compute.api.onmetal.de/v1alpha1</h2>
<div>
<p>Package v1alpha1 is the v1alpha1 version of the API.</p>
</div>
Resource Types:
<ul><li>
<a href="#compute.api.onmetal.de/v1alpha1.Machine">Machine</a>
</li><li>
<a href="#compute.api.onmetal.de/v1alpha1.MachineClass">MachineClass</a>
</li><li>
<a href="#compute.api.onmetal.de/v1alpha1.MachinePool">MachinePool</a>
</li></ul>
<h3 id="compute.api.onmetal.de/v1alpha1.Machine">Machine
</h3>
<div>
<p>Machine is the Schema for the machines API</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
compute.api.onmetal.de/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>Machine</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.MachineSpec">
MachineSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>machineClassRef</code><br/>
<em>
<a href="https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#localobjectreference-v1-core">
Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
<p>MachineClassRef is a reference to the machine class/flavor of the machine.</p>
</td>
</tr>
<tr>
<td>
<code>machinePoolSelector</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>MachinePoolSelector selects a suitable MachinePoolRef by the given labels.</p>
</td>
</tr>
<tr>
<td>
<code>machinePoolRef</code><br/>
<em>
<a href="https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#localobjectreference-v1-core">
Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
<p>MachinePoolRef defines machine pool to run the machine in.
If empty, a scheduler will figure out an appropriate pool to run the machine in.</p>
</td>
</tr>
<tr>
<td>
<code>image</code><br/>
<em>
string
</em>
</td>
<td>
<p>Image is the URL providing the operating system image of the machine.</p>
</td>
</tr>
<tr>
<td>
<code>imagePullSecret</code><br/>
<em>
<a href="https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#localobjectreference-v1-core">
Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
<p>ImagePullSecretRef is an optional secret for pulling the image of a machine.</p>
</td>
</tr>
<tr>
<td>
<code>networkInterfaces</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.NetworkInterface">
[]NetworkInterface
</a>
</em>
</td>
<td>
<p>NetworkInterfaces define a list of network interfaces present on the machine</p>
</td>
</tr>
<tr>
<td>
<code>volumes</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.Volume">
[]Volume
</a>
</em>
</td>
<td>
<p>Volumes are volumes attached to this machine.</p>
</td>
</tr>
<tr>
<td>
<code>ignitionRef</code><br/>
<em>
<a href="/api-reference/common/#common.onmetal.de/v1alpha1.SecretKeySelector">
github.com/onmetal/onmetal-api/apis/common/v1alpha1.SecretKeySelector
</a>
</em>
</td>
<td>
<p>IgnitionRef is a reference to a config map containing the ignition YAML for the machine to boot up.
If key is empty, DefaultIgnitionKey will be used as fallback.</p>
</td>
</tr>
<tr>
<td>
<code>efiVars</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.EFIVar">
[]EFIVar
</a>
</em>
</td>
<td>
<p>EFIVars are variables to pass to EFI while booting up.</p>
</td>
</tr>
<tr>
<td>
<code>tolerations</code><br/>
<em>
<a href="/api-reference/common/#common.onmetal.de/v1alpha1.Toleration">
[]github.com/onmetal/onmetal-api/apis/common/v1alpha1.Toleration
</a>
</em>
</td>
<td>
<p>Tolerations define tolerations the Machine has. Only MachinePools whose taints
covered by Tolerations will be considered to run the Machine.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.MachineStatus">
MachineStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.MachineClass">MachineClass
</h3>
<div>
<p>MachineClass is the Schema for the machineclasses API</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
compute.api.onmetal.de/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>MachineClass</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>capabilities</code><br/>
<em>
<a href="https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#resourcelist-v1-core">
Kubernetes core/v1.ResourceList
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.MachinePool">MachinePool
</h3>
<div>
<p>MachinePool is the Schema for the machinepools API</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>apiVersion</code><br/>
string</td>
<td>
<code>
compute.api.onmetal.de/v1alpha1
</code>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
string
</td>
<td><code>MachinePool</code></td>
</tr>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.MachinePoolSpec">
MachinePoolSpec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>providerID</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProviderID identifies the MachinePool on provider side.</p>
</td>
</tr>
<tr>
<td>
<code>taints</code><br/>
<em>
<a href="/api-reference/common/#common.onmetal.de/v1alpha1.Taint">
[]github.com/onmetal/onmetal-api/apis/common/v1alpha1.Taint
</a>
</em>
</td>
<td>
<p>Taints of the MachinePool. Only Machines who tolerate all the taints
will land in the MachinePool.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.MachinePoolStatus">
MachinePoolStatus
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.DaemonEndpoint">DaemonEndpoint
</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.MachinePoolDaemonEndpoints">MachinePoolDaemonEndpoints</a>)
</p>
<div>
<p>DaemonEndpoint contains information about a single Daemon endpoint.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>port</code><br/>
<em>
int32
</em>
</td>
<td>
<p>Port number of the given endpoint.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.EFIVar">EFIVar
</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.MachineSpec">MachineSpec</a>)
</p>
<div>
<p>EFIVar is a variable to pass to EFI while booting up.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name is the name of the EFIVar.</p>
</td>
</tr>
<tr>
<td>
<code>uuid</code><br/>
<em>
string
</em>
</td>
<td>
<p>UUID is the uuid of the EFIVar.</p>
</td>
</tr>
<tr>
<td>
<code>value</code><br/>
<em>
string
</em>
</td>
<td>
<p>Value is the value of the EFIVar.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.EmptyDiskVolumeSource">EmptyDiskVolumeSource
</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.VolumeSource">VolumeSource</a>)
</p>
<div>
<p>EmptyDiskVolumeSource is a volume that&rsquo;s offered by the machine pool provider.
Usually ephemeral (i.e. deleted when the surrounding entity is deleted), with
varying performance characteristics. Potentially not recoverable.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>sizeLimit</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource#Duration">
k8s.io/apimachinery/pkg/api/resource.Quantity
</a>
</em>
</td>
<td>
<p>SizeLimit is the total amount of local storage required for this EmptyDisk volume.
The default is nil which means that the limit is undefined.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.EphemeralNetworkInterfaceSource">EphemeralNetworkInterfaceSource
</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.NetworkInterfaceSource">NetworkInterfaceSource</a>)
</p>
<div>
<p>EphemeralNetworkInterfaceSource is a definition for an ephemeral (i.e. coupled to the lifetime of the surrounding
object) networking.NetworkInterface.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>networkInterfaceTemplate</code><br/>
<em>
github.com/onmetal/onmetal-api/apis/networking/v1alpha1.NetworkInterfaceTemplateSpec
</em>
</td>
<td>
<p>NetworkInterfaceTemplate is the template definition of the networking.NetworkInterface.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.EphemeralVolumeSource">EphemeralVolumeSource
</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.VolumeSource">VolumeSource</a>)
</p>
<div>
<p>EphemeralVolumeSource is a definition for an ephemeral (i.e. coupled to the lifetime of the surrounding object)
storage.Volume.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>volumeTemplate</code><br/>
<em>
github.com/onmetal/onmetal-api/apis/storage/v1alpha1.VolumeTemplateSpec
</em>
</td>
<td>
<p>VolumeTemplate is the template definition of the storage.Volume.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.MachineCondition">MachineCondition
</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.MachineStatus">MachineStatus</a>)
</p>
<div>
<p>MachineCondition is one of the conditions of a volume.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.MachineConditionType">
MachineConditionType
</a>
</em>
</td>
<td>
<p>Type is the type of the condition.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#conditionstatus-v1-core">
Kubernetes core/v1.ConditionStatus
</a>
</em>
</td>
<td>
<p>Status is the status of the condition.</p>
</td>
</tr>
<tr>
<td>
<code>reason</code><br/>
<em>
string
</em>
</td>
<td>
<p>Reason is a machine-readable indication of why the condition is in a certain state.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<p>Message is a human-readable explanation of why the condition has a certain reason / state.</p>
</td>
</tr>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<p>ObservedGeneration represents the .metadata.generation that the condition was set based upon.</p>
</td>
</tr>
<tr>
<td>
<code>lastTransitionTime</code><br/>
<em>
<a href="https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#time-v1-meta">
Kubernetes meta/v1.Time
</a>
</em>
</td>
<td>
<p>LastTransitionTime is the last time the status of a condition has transitioned from one state to another.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.MachineConditionType">MachineConditionType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.MachineCondition">MachineCondition</a>)
</p>
<div>
<p>MachineConditionType is a type a MachineCondition can have.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Synced&#34;</p></td>
<td><p>MachineSynced represents the condition of a machine being synced with its backing resources</p>
</td>
</tr></tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.MachineExecOptions">MachineExecOptions
</h3>
<div>
<p>MachineExecOptions is the query options to a Machine&rsquo;s remote exec call</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>insecureSkipTLSVerifyBackend</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.MachinePoolAddress">MachinePoolAddress
</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.MachinePoolStatus">MachinePoolStatus</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.MachinePoolAddressType">
MachinePoolAddressType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>address</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.MachinePoolAddressType">MachinePoolAddressType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.MachinePoolAddress">MachinePoolAddress</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;ExternalDNS&#34;</p></td>
<td><p>MachinePoolExternalDNS identifies a DNS name which resolves to an IP address which has the characteristics
of MachinePoolExternalIP. The IP it resolves to may or may not be a listed MachineExternalIP address.</p>
</td>
</tr><tr><td><p>&#34;ExternalIP&#34;</p></td>
<td><p>MachinePoolExternalIP identifies an IP address which is, in some way, intended to be more usable from outside
the cluster than an internal IP, though no specific semantics are defined.</p>
</td>
</tr><tr><td><p>&#34;Hostname&#34;</p></td>
<td><p>MachinePoolHostName identifies a name of the machine pool. Although every machine pool can be assumed
to have a MachinePoolAddress of this type, its exact syntax and semantics are not
defined, and are not consistent between different clusters.</p>
</td>
</tr><tr><td><p>&#34;InternalDNS&#34;</p></td>
<td><p>MachinePoolInternalDNS identifies a DNS name which resolves to an IP address which has
the characteristics of a MachinePoolInternalIP. The IP it resolves to may or may not
be a listed MachinePoolInternalIP address.</p>
</td>
</tr><tr><td><p>&#34;InternalIP&#34;</p></td>
<td><p>MachinePoolInternalIP identifies an IP address which may not be visible to hosts outside the cluster.
By default, it is assumed that onmetal-api-apiserver can reach machine pool internal IPs, though it is possible
to configure clusters where this is not the case.</p>
<p>MachinePoolInternalIP is the default type of machine pool IP, and does not necessarily imply
that the IP is ONLY reachable internally. If a machine pool has multiple internal IPs,
no specific semantics are assigned to the additional IPs.</p>
</td>
</tr></tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.MachinePoolCondition">MachinePoolCondition
</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.MachinePoolStatus">MachinePoolStatus</a>)
</p>
<div>
<p>MachinePoolCondition is one of the conditions of a volume.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.MachinePoolConditionType">
MachinePoolConditionType
</a>
</em>
</td>
<td>
<p>Type is the type of the condition.</p>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#conditionstatus-v1-core">
Kubernetes core/v1.ConditionStatus
</a>
</em>
</td>
<td>
<p>Status is the status of the condition.</p>
</td>
</tr>
<tr>
<td>
<code>reason</code><br/>
<em>
string
</em>
</td>
<td>
<p>Reason is a machine-readable indication of why the condition is in a certain state.</p>
</td>
</tr>
<tr>
<td>
<code>message</code><br/>
<em>
string
</em>
</td>
<td>
<p>Message is a human-readable explanation of why the condition has a certain reason / state.</p>
</td>
</tr>
<tr>
<td>
<code>observedGeneration</code><br/>
<em>
int64
</em>
</td>
<td>
<p>ObservedGeneration represents the .metadata.generation that the condition was set based upon.</p>
</td>
</tr>
<tr>
<td>
<code>lastTransitionTime</code><br/>
<em>
<a href="https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#time-v1-meta">
Kubernetes meta/v1.Time
</a>
</em>
</td>
<td>
<p>LastTransitionTime is the last time the status of a condition has transitioned from one state to another.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.MachinePoolConditionType">MachinePoolConditionType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.MachinePoolCondition">MachinePoolCondition</a>)
</p>
<div>
<p>MachinePoolConditionType is a type a MachinePoolCondition can have.</p>
</div>
<h3 id="compute.api.onmetal.de/v1alpha1.MachinePoolDaemonEndpoints">MachinePoolDaemonEndpoints
</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.MachinePoolStatus">MachinePoolStatus</a>)
</p>
<div>
<p>MachinePoolDaemonEndpoints lists ports opened by daemons running on the MachinePool.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>machinepoolletEndpoint</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.DaemonEndpoint">
DaemonEndpoint
</a>
</em>
</td>
<td>
<em>(Optional)</em>
<p>Endpoint on which machinepoollet is listening.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.MachinePoolSpec">MachinePoolSpec
</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.MachinePool">MachinePool</a>)
</p>
<div>
<p>MachinePoolSpec defines the desired state of MachinePool</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>providerID</code><br/>
<em>
string
</em>
</td>
<td>
<p>ProviderID identifies the MachinePool on provider side.</p>
</td>
</tr>
<tr>
<td>
<code>taints</code><br/>
<em>
<a href="/api-reference/common/#common.onmetal.de/v1alpha1.Taint">
[]github.com/onmetal/onmetal-api/apis/common/v1alpha1.Taint
</a>
</em>
</td>
<td>
<p>Taints of the MachinePool. Only Machines who tolerate all the taints
will land in the MachinePool.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.MachinePoolState">MachinePoolState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.MachinePoolStatus">MachinePoolStatus</a>)
</p>
<div>
<p>MachinePoolState is a state a MachinePool can be in.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Error&#34;</p></td>
<td><p>MachinePoolStateError marks a MachinePool in an error state.</p>
</td>
</tr><tr><td><p>&#34;Offline&#34;</p></td>
<td><p>MachinePoolStateOffline marks a MachinePool as offline.</p>
</td>
</tr><tr><td><p>&#34;Pending&#34;</p></td>
<td><p>MachinePoolStatePending marks a MachinePool as pending readiness.</p>
</td>
</tr><tr><td><p>&#34;Ready&#34;</p></td>
<td><p>MachinePoolStateReady marks a MachinePool as ready for accepting a Machine.</p>
</td>
</tr></tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.MachinePoolStatus">MachinePoolStatus
</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.MachinePool">MachinePool</a>)
</p>
<div>
<p>MachinePoolStatus defines the observed state of MachinePool</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.MachinePoolState">
MachinePoolState
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.MachinePoolCondition">
[]MachinePoolCondition
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>availableMachineClasses</code><br/>
<em>
<a href="https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#localobjectreference-v1-core">
[]Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>addresses</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.MachinePoolAddress">
[]MachinePoolAddress
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>daemonEndpoints</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.MachinePoolDaemonEndpoints">
MachinePoolDaemonEndpoints
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.MachineSpec">MachineSpec
</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.Machine">Machine</a>)
</p>
<div>
<p>MachineSpec defines the desired state of Machine</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>machineClassRef</code><br/>
<em>
<a href="https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#localobjectreference-v1-core">
Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
<p>MachineClassRef is a reference to the machine class/flavor of the machine.</p>
</td>
</tr>
<tr>
<td>
<code>machinePoolSelector</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>MachinePoolSelector selects a suitable MachinePoolRef by the given labels.</p>
</td>
</tr>
<tr>
<td>
<code>machinePoolRef</code><br/>
<em>
<a href="https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#localobjectreference-v1-core">
Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
<p>MachinePoolRef defines machine pool to run the machine in.
If empty, a scheduler will figure out an appropriate pool to run the machine in.</p>
</td>
</tr>
<tr>
<td>
<code>image</code><br/>
<em>
string
</em>
</td>
<td>
<p>Image is the URL providing the operating system image of the machine.</p>
</td>
</tr>
<tr>
<td>
<code>imagePullSecret</code><br/>
<em>
<a href="https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#localobjectreference-v1-core">
Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
<p>ImagePullSecretRef is an optional secret for pulling the image of a machine.</p>
</td>
</tr>
<tr>
<td>
<code>networkInterfaces</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.NetworkInterface">
[]NetworkInterface
</a>
</em>
</td>
<td>
<p>NetworkInterfaces define a list of network interfaces present on the machine</p>
</td>
</tr>
<tr>
<td>
<code>volumes</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.Volume">
[]Volume
</a>
</em>
</td>
<td>
<p>Volumes are volumes attached to this machine.</p>
</td>
</tr>
<tr>
<td>
<code>ignitionRef</code><br/>
<em>
<a href="/api-reference/common/#common.onmetal.de/v1alpha1.SecretKeySelector">
github.com/onmetal/onmetal-api/apis/common/v1alpha1.SecretKeySelector
</a>
</em>
</td>
<td>
<p>IgnitionRef is a reference to a config map containing the ignition YAML for the machine to boot up.
If key is empty, DefaultIgnitionKey will be used as fallback.</p>
</td>
</tr>
<tr>
<td>
<code>efiVars</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.EFIVar">
[]EFIVar
</a>
</em>
</td>
<td>
<p>EFIVars are variables to pass to EFI while booting up.</p>
</td>
</tr>
<tr>
<td>
<code>tolerations</code><br/>
<em>
<a href="/api-reference/common/#common.onmetal.de/v1alpha1.Toleration">
[]github.com/onmetal/onmetal-api/apis/common/v1alpha1.Toleration
</a>
</em>
</td>
<td>
<p>Tolerations define tolerations the Machine has. Only MachinePools whose taints
covered by Tolerations will be considered to run the Machine.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.MachineState">MachineState
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.MachineStatus">MachineStatus</a>)
</p>
<div>
<p>MachineState is the state of a machine.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Error&#34;</p></td>
<td><p>MachineStateError means the machine is in an error state.</p>
</td>
</tr><tr><td><p>&#34;Pending&#34;</p></td>
<td><p>MachineStatePending means the Machine has been accepted by the system, but not yet completely started.
This includes time before being bound to a MachinePool, as well as time spent setting up the Machine on that
MachinePool.</p>
</td>
</tr><tr><td><p>&#34;Running&#34;</p></td>
<td><p>MachineStateRunning means the machine is running on a MachinePool.</p>
</td>
</tr><tr><td><p>&#34;Shutdown&#34;</p></td>
<td><p>MachineStateShutdown means the machine is shut down.</p>
</td>
</tr></tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.MachineStatus">MachineStatus
</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.Machine">Machine</a>)
</p>
<div>
<p>MachineStatus defines the observed state of Machine</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>state</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.MachineState">
MachineState
</a>
</em>
</td>
<td>
<p>State is the state of the machine.</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.MachineCondition">
[]MachineCondition
</a>
</em>
</td>
<td>
<p>Conditions are the conditions of the machines.</p>
</td>
</tr>
<tr>
<td>
<code>networkInterfaces</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.NetworkInterfaceStatus">
[]NetworkInterfaceStatus
</a>
</em>
</td>
<td>
<p>NetworkInterfaces is the list of network interface states for the machine.</p>
</td>
</tr>
<tr>
<td>
<code>volumes</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.VolumeStatus">
[]VolumeStatus
</a>
</em>
</td>
<td>
<p>Volumes is the list of volume states for the machine.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.NetworkInterface">NetworkInterface
</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.MachineSpec">MachineSpec</a>)
</p>
<div>
<p>NetworkInterface is the definition of a single interface</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name is the name of the network interface.</p>
</td>
</tr>
<tr>
<td>
<code>NetworkInterfaceSource</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.NetworkInterfaceSource">
NetworkInterfaceSource
</a>
</em>
</td>
<td>
<p>
(Members of <code>NetworkInterfaceSource</code> are embedded into this type.)
</p>
<p>NetworkInterfaceSource is where to obtain the interface from.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.NetworkInterfacePhase">NetworkInterfacePhase
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.NetworkInterfaceStatus">NetworkInterfaceStatus</a>)
</p>
<div>
<p>NetworkInterfacePhase represents the binding phase a NetworkInterface can be in.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Bound&#34;</p></td>
<td><p>NetworkInterfacePhaseBound is used for a NetworkInterface that is bound.</p>
</td>
</tr><tr><td><p>&#34;Pending&#34;</p></td>
<td><p>NetworkInterfacePhasePending is used for a NetworkInterface that is not yet bound.</p>
</td>
</tr></tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.NetworkInterfaceSource">NetworkInterfaceSource
</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.NetworkInterface">NetworkInterface</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>networkInterfaceRef</code><br/>
<em>
<a href="https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#localobjectreference-v1-core">
Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
<p>NetworkInterfaceRef instructs to use the NetworkInterface at the target reference.</p>
</td>
</tr>
<tr>
<td>
<code>ephemeral</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.EphemeralNetworkInterfaceSource">
EphemeralNetworkInterfaceSource
</a>
</em>
</td>
<td>
<p>Ephemeral instructs to create an ephemeral (i.e. coupled to the lifetime of the surrounding object)
NetworkInterface to use.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.NetworkInterfaceStatus">NetworkInterfaceStatus
</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.MachineStatus">MachineStatus</a>)
</p>
<div>
<p>NetworkInterfaceStatus reports the status of an NetworkInterfaceSource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name is the name of the NetworkInterface to whom the status belongs to.</p>
</td>
</tr>
<tr>
<td>
<code>phase</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.NetworkInterfacePhase">
NetworkInterfacePhase
</a>
</em>
</td>
<td>
<p>Phase is the NetworkInterface binding phase of the NetworkInterface.</p>
</td>
</tr>
<tr>
<td>
<code>lastPhaseTransitionTime</code><br/>
<em>
<a href="https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#time-v1-meta">
Kubernetes meta/v1.Time
</a>
</em>
</td>
<td>
<p>LastPhaseTransitionTime is the last time the Phase transitioned.</p>
</td>
</tr>
<tr>
<td>
<code>ips</code><br/>
<em>
<a href="/api-reference/common/#common.onmetal.de/v1alpha1.IP">
[]github.com/onmetal/onmetal-api/apis/common/v1alpha1.IP
</a>
</em>
</td>
<td>
<p>IPs are the ips allocated for the network interface.</p>
</td>
</tr>
<tr>
<td>
<code>virtualIP</code><br/>
<em>
<a href="/api-reference/common/#common.onmetal.de/v1alpha1.IP">
github.com/onmetal/onmetal-api/apis/common/v1alpha1.IP
</a>
</em>
</td>
<td>
<p>VirtualIP is the virtual ip allocated for the network interface.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.Volume">Volume
</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.MachineSpec">MachineSpec</a>)
</p>
<div>
<p>Volume defines a volume attachment of a machine</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name is the name of the Volume</p>
</td>
</tr>
<tr>
<td>
<code>device</code><br/>
<em>
string
</em>
</td>
<td>
<p>Device is the device name where the volume should be attached. If empty,
an unused device name will be determined if possible.</p>
</td>
</tr>
<tr>
<td>
<code>VolumeSource</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.VolumeSource">
VolumeSource
</a>
</em>
</td>
<td>
<p>
(Members of <code>VolumeSource</code> are embedded into this type.)
</p>
<p>VolumeSource is the source where the storage for the Volume resides at.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.VolumePhase">VolumePhase
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.VolumeStatus">VolumeStatus</a>)
</p>
<div>
<p>VolumePhase represents the binding phase a Volume can be in.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Bound&#34;</p></td>
<td><p>VolumePhaseBound is used for a Volume that is bound.</p>
</td>
</tr><tr><td><p>&#34;Pending&#34;</p></td>
<td><p>VolumePhasePending is used for a Volume that is not yet bound.</p>
</td>
</tr></tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.VolumeSource">VolumeSource
</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.Volume">Volume</a>)
</p>
<div>
<p>VolumeSource specifies the source to use for a Volume.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>volumeRef</code><br/>
<em>
<a href="https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#localobjectreference-v1-core">
Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
<p>VolumeRef instructs to use the specified Volume as source for the attachment.</p>
</td>
</tr>
<tr>
<td>
<code>emptyDisk</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.EmptyDiskVolumeSource">
EmptyDiskVolumeSource
</a>
</em>
</td>
<td>
<p>EmptyDisk instructs to use a Volume offered by the machine pool provider.</p>
</td>
</tr>
<tr>
<td>
<code>ephemeral</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.EphemeralVolumeSource">
EphemeralVolumeSource
</a>
</em>
</td>
<td>
<p>Ephemeral instructs to create an ephemeral (i.e. coupled to the lifetime of the surrounding object)
Volume to use.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="compute.api.onmetal.de/v1alpha1.VolumeStatus">VolumeStatus
</h3>
<p>
(<em>Appears on:</em><a href="#compute.api.onmetal.de/v1alpha1.MachineStatus">MachineStatus</a>)
</p>
<div>
<p>VolumeStatus is the status of a Volume.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name is the name of a volume attachment.</p>
</td>
</tr>
<tr>
<td>
<code>phase</code><br/>
<em>
<a href="#compute.api.onmetal.de/v1alpha1.VolumePhase">
VolumePhase
</a>
</em>
</td>
<td>
<p>Phase represents the binding phase of a Volume.</p>
</td>
</tr>
<tr>
<td>
<code>lastPhaseTransitionTime</code><br/>
<em>
<a href="https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#time-v1-meta">
Kubernetes meta/v1.Time
</a>
</em>
</td>
<td>
<p>LastPhaseTransitionTime is the last time the Phase transitioned.</p>
</td>
</tr>
<tr>
<td>
<code>deviceID</code><br/>
<em>
string
</em>
</td>
<td>
<p>DeviceID is the disk device ID on the host.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
<p><em>
Generated with <code>gen-crd-api-reference-docs</code>
</em></p>
