// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: extensions/v1alpha1/wasm.proto

// WasmPlugins provides a mechanism to extend the functionality provided by
// the Istio proxy through WebAssembly filters.
//
// Order of execution (as part of Envoy's filter chain) is determined by
// phase and priority settings, allowing the configuration of complex
// interactions between user-supplied WasmPlugins and Istio's internal
// filters.
//
// Examples:
//
// AuthN Filter deployed to ingress-gateway that implements an OpenID flow
// and populates the `Authorization` header with a JWT to be consumed by
// Istio AuthN.
//
// ```yaml
// apiVersion: extensions.istio.io/v1alpha1
// kind: WasmPlugin
// metadata:
//   name: openid-connect
//   namespace: istio-ingress
// spec:
//   selector:
//     matchLabels:
//       istio: ingressgateway
//   url: file:///opt/filters/openid.wasm
//   sha256: 1ef0c9a92b0420cf25f7fe5d481b231464bc88f486ca3b9c83ed5cc21d2f6210
//   phase: AUTHN
//   pluginConfig:
//     openid_server: authn
//     openid_realm: ingress
// ```
//
// This is the same as the last example, but using an OCI image.
//
// ```yaml
// apiVersion: extensions.istio.io/v1alpha1
// kind: WasmPlugin
// metadata:
//   name: openid-connect
//   namespace: istio-ingress
// spec:
//   selector:
//     labels:
//       istio: ingressgateway
//   url: oci://private-registry:5000/openid-connect/openid:latest
//   imagePullPolicy: IfNotPresent
//   imagePullSecret: private-registry-pull-secret
//   phase: AUTHN
//   pluginConfig:
//     openid_server: authn
//     openid_realm: ingress
// ```
//
// This is the same as the last example, but using VmConfig to configure environment variables in the VM.
//
// ```yaml
// apiVersion: extensions.istio.io/v1alpha1
// kind: WasmPlugin
// metadata:
//   name: openid-connect
//   namespace: istio-ingress
// spec:
//   selector:
//     labels:
//       istio: ingressgateway
//   url: oci://private-registry:5000/openid-connect/openid:latest
//   imagePullPolicy: IfNotPresent
//   imagePullSecret: private-registry-pull-secret
//   phase: AUTHN
//   pluginConfig:
//     openid_server: authn
//     openid_realm: ingress
//   vmConfig:
//     env:
//     - name: POD_NAME
//       valueFrom: HOST
//     - name: TRUST_DOMAIN
//       value: "cluster.local"
// ```
//
// And a more complex example that deploys three WasmPlugins and orders them
// using `phase` and `priority`. The (hypothetical) setup is that the
// `openid-connect` filter performs an OpenID Connect flow to authenticate the
// user, writing a signed JWT into the Authorization header of the request,
// which can be verified by the Istio authn plugin. Then, the `acl-check` plugin
// kicks in, passing the JWT to a policy server, which in turn responds with a
// signed token that contains information about which files and functions of the
// system are available to the user that was previously authenticated. The
// `acl-check` filter writes this token to a header. Finally, the `check-header`
// filter verifies the token in that header and makes sure that the token's
// contents (the permitted 'function') matches its plugin configuration.
//
// The resulting filter chain looks like this:
// -> openid-connect -> istio.authn -> acl-check -> check-header -> router
//
// ```yaml
// apiVersion: extensions.istio.io/v1alpha1
// kind: WasmPlugin
// metadata:
//   name: openid-connect
//   namespace: istio-ingress
// spec:
//   selector:
//     matchLabels:
//       istio: ingressgateway
//   url: oci://private-registry:5000/openid-connect/openid:latest
//   imagePullPolicy: IfNotPresent
//   imagePullSecret: private-registry-pull-secret
//   phase: AUTHN
//   pluginConfig:
//     openid_server: authn
//     openid_realm: ingress
// ```
//
// ```yaml
// apiVersion: extensions.istio.io/v1alpha1
// kind: WasmPlugin
// metadata:
//   name: acl-check
//   namespace: istio-ingress
// spec:
//   selector:
//     matchLabels:
//       istio: ingressgateway
//   url: oci://private-registry:5000/acl-check/acl:latest
//   imagePullPolicy: Always
//   imagePullSecret: private-registry-pull-secret
//   phase: AUTHZ
//   priority: 1000
//   pluginConfig:
//     acl_server: some_server
//     set_header: authz_complete
// ```
//
// ```yaml
// apiVersion: extensions.istio.io/v1alpha1
// kind: WasmPlugin
// metadata:
//   name: check-header
//   namespace: istio-ingress
// spec:
//   selector:
//     matchLabels:
//       istio: ingressgateway
//   url: oci://private-registry:5000/check-header:latest
//   imagePullPolicy: IfNotPresent
//   imagePullSecret: private-registry-pull-secret
//   phase: AUTHZ
//   priority: 10
//   pluginConfig:
//     read_header: authz_complete
//     verification_key: a89gAzxvls0JKAKIJSBnnvvvkIO
//     function: read_data
// ```
//

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "istio.io/api/type/v1beta1"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// DeepCopyInto supports using WasmPlugin within kubernetes types, where deepcopy-gen is used.
func (in *WasmPlugin) DeepCopyInto(out *WasmPlugin) {
	p := proto.Clone(in).(*WasmPlugin)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WasmPlugin. Required by controller-gen.
func (in *WasmPlugin) DeepCopy() *WasmPlugin {
	if in == nil {
		return nil
	}
	out := new(WasmPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new WasmPlugin. Required by controller-gen.
func (in *WasmPlugin) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using VmConfig within kubernetes types, where deepcopy-gen is used.
func (in *VmConfig) DeepCopyInto(out *VmConfig) {
	p := proto.Clone(in).(*VmConfig)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VmConfig. Required by controller-gen.
func (in *VmConfig) DeepCopy() *VmConfig {
	if in == nil {
		return nil
	}
	out := new(VmConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new VmConfig. Required by controller-gen.
func (in *VmConfig) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using EnvVar within kubernetes types, where deepcopy-gen is used.
func (in *EnvVar) DeepCopyInto(out *EnvVar) {
	p := proto.Clone(in).(*EnvVar)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvVar. Required by controller-gen.
func (in *EnvVar) DeepCopy() *EnvVar {
	if in == nil {
		return nil
	}
	out := new(EnvVar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new EnvVar. Required by controller-gen.
func (in *EnvVar) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
