<p>Packages:</p>
<ul>
<li>
<a href="#common.api.onmetal.de%2fv1alpha1">common.api.onmetal.de/v1alpha1</a>
</li>
</ul>
<h2 id="common.api.onmetal.de/v1alpha1">common.api.onmetal.de/v1alpha1</h2>
<div>
<p>Package v1alpha1 is the v1alpha1 version of the API.</p>
</div>
Resource Types:
<ul></ul>
<h3 id="common.api.onmetal.de/v1alpha1.ConfigMapKeySelector">ConfigMapKeySelector
</h3>
<div>
<p>ConfigMapKeySelector is a reference to a specific &lsquo;key&rsquo; within a ConfigMap resource.
In some instances, <code>key</code> is a required field.</p>
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
<code>LocalObjectReference</code><br/>
<em>
<a href="https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#localobjectreference-v1-core">
Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
<p>
(Members of <code>LocalObjectReference</code> are embedded into this type.)
</p>
<p>The name of the ConfigMap resource being referred to.</p>
</td>
</tr>
<tr>
<td>
<code>key</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The key of the entry in the ConfigMap resource&rsquo;s <code>data</code> field to be used.
Some instances of this field may be defaulted, in others it may be
required.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="common.api.onmetal.de/v1alpha1.IP">IP
</h3>
<p>
(<em>Appears on:</em><a href="#common.api.onmetal.de/v1alpha1.IPRange">IPRange</a>)
</p>
<div>
<p>IP is an IP address.</p>
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
<code>-</code><br/>
<em>
net/netip.Addr
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="common.api.onmetal.de/v1alpha1.IPPrefix">IPPrefix
</h3>
<div>
<p>IPPrefix represents a network prefix.</p>
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
<code>-</code><br/>
<em>
net/netip.Prefix
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="common.api.onmetal.de/v1alpha1.IPRange">IPRange
</h3>
<div>
<p>IPRange is an IP range.</p>
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
<code>from</code><br/>
<em>
<a href="#common.api.onmetal.de/v1alpha1.IP">
IP
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>to</code><br/>
<em>
<a href="#common.api.onmetal.de/v1alpha1.IP">
IP
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="common.api.onmetal.de/v1alpha1.LocalUIDReference">LocalUIDReference
</h3>
<div>
<p>LocalUIDReference is a reference to another entity including its UID</p>
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
<p>Name is the name of the referenced entity.</p>
</td>
</tr>
<tr>
<td>
<code>uid</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apimachinery/pkg/types#UID">
k8s.io/apimachinery/pkg/types.UID
</a>
</em>
</td>
<td>
<p>UID is the UID of the referenced entity.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="common.api.onmetal.de/v1alpha1.SecretKeySelector">SecretKeySelector
</h3>
<div>
<p>SecretKeySelector is a reference to a specific &lsquo;key&rsquo; within a Secret resource.
In some instances, <code>key</code> is a required field.</p>
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
<code>LocalObjectReference</code><br/>
<em>
<a href="https://v1-23.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.23/#localobjectreference-v1-core">
Kubernetes core/v1.LocalObjectReference
</a>
</em>
</td>
<td>
<p>
(Members of <code>LocalObjectReference</code> are embedded into this type.)
</p>
<p>The name of the Secret resource being referred to.</p>
</td>
</tr>
<tr>
<td>
<code>key</code><br/>
<em>
string
</em>
</td>
<td>
<em>(Optional)</em>
<p>The key of the entry in the Secret resource&rsquo;s <code>data</code> field to be used.
Some instances of this field may be defaulted, in others it may be
required.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="common.api.onmetal.de/v1alpha1.Taint">Taint
</h3>
<div>
<p>The resource pool this Taint is attached to has the &ldquo;effect&rdquo; on
any resource that does not tolerate the Taint.</p>
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
<code>key</code><br/>
<em>
string
</em>
</td>
<td>
<p>The taint key to be applied to a resource pool.</p>
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
<p>The taint value corresponding to the taint key.</p>
</td>
</tr>
<tr>
<td>
<code>effect</code><br/>
<em>
<a href="#common.api.onmetal.de/v1alpha1.TaintEffect">
TaintEffect
</a>
</em>
</td>
<td>
<p>The effect of the taint on resources
that do not tolerate the taint.
Valid effects are NoSchedule, PreferNoSchedule and NoExecute.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="common.api.onmetal.de/v1alpha1.TaintEffect">TaintEffect
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#common.api.onmetal.de/v1alpha1.Taint">Taint</a>, <a href="#common.api.onmetal.de/v1alpha1.Toleration">Toleration</a>)
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
<tbody><tr><td><p>&#34;NoSchedule&#34;</p></td>
<td><p>Do not allow new resources to schedule onto the resource pool unless they tolerate the taint,
but allow all already-running resources to continue running.
Enforced by the scheduler.</p>
</td>
</tr></tbody>
</table>
<h3 id="common.api.onmetal.de/v1alpha1.Toleration">Toleration
</h3>
<div>
<p>The resource this Toleration is attached to tolerates any taint that matches
the triple <key,value,effect> using the matching operator <operator>.</p>
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
<code>key</code><br/>
<em>
string
</em>
</td>
<td>
<p>Key is the taint key that the toleration applies to. Empty means match all taint keys.
If the key is empty, operator must be Exists; this combination means to match all values and all keys.</p>
</td>
</tr>
<tr>
<td>
<code>operator</code><br/>
<em>
<a href="#common.api.onmetal.de/v1alpha1.TolerationOperator">
TolerationOperator
</a>
</em>
</td>
<td>
<p>Operator represents a key&rsquo;s relationship to the value.
Valid operators are Exists and Equal. Defaults to Equal.
Exists is equivalent to wildcard for value, so that a resource can
tolerate all taints of a particular category.</p>
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
<p>Value is the taint value the toleration matches to.
If the operator is Exists, the value should be empty, otherwise just a regular string.</p>
</td>
</tr>
<tr>
<td>
<code>effect</code><br/>
<em>
<a href="#common.api.onmetal.de/v1alpha1.TaintEffect">
TaintEffect
</a>
</em>
</td>
<td>
<p>Effect indicates the taint effect to match. Empty means match all taint effects.
When specified, allowed values are NoSchedule.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="common.api.onmetal.de/v1alpha1.TolerationOperator">TolerationOperator
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#common.api.onmetal.de/v1alpha1.Toleration">Toleration</a>)
</p>
<div>
<p>A toleration operator is the set of operators that can be used in a toleration.</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Equal&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Exists&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<hr/>
<p><em>
Generated with <code>gen-crd-api-reference-docs</code>
</em></p>
