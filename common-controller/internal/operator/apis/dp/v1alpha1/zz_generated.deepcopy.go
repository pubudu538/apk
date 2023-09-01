//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
 *  Copyright (c) 2022, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *API) DeepCopyInto(out *API) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new API.
func (in *API) DeepCopy() *API {
	if in == nil {
		return nil
	}
	out := new(API)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *API) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIList) DeepCopyInto(out *APIList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]API, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIList.
func (in *APIList) DeepCopy() *APIList {
	if in == nil {
		return nil
	}
	out := new(APIList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIPolicy) DeepCopyInto(out *APIPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIPolicy.
func (in *APIPolicy) DeepCopy() *APIPolicy {
	if in == nil {
		return nil
	}
	out := new(APIPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIPolicyList) DeepCopyInto(out *APIPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]APIPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIPolicyList.
func (in *APIPolicyList) DeepCopy() *APIPolicyList {
	if in == nil {
		return nil
	}
	out := new(APIPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIPolicySpec) DeepCopyInto(out *APIPolicySpec) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(PolicySpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Override != nil {
		in, out := &in.Override, &out.Override
		*out = new(PolicySpec)
		(*in).DeepCopyInto(*out)
	}
	in.TargetRef.DeepCopyInto(&out.TargetRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIPolicySpec.
func (in *APIPolicySpec) DeepCopy() *APIPolicySpec {
	if in == nil {
		return nil
	}
	out := new(APIPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIPolicyStatus) DeepCopyInto(out *APIPolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIPolicyStatus.
func (in *APIPolicyStatus) DeepCopy() *APIPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(APIPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIRateLimitPolicy) DeepCopyInto(out *APIRateLimitPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIRateLimitPolicy.
func (in *APIRateLimitPolicy) DeepCopy() *APIRateLimitPolicy {
	if in == nil {
		return nil
	}
	out := new(APIRateLimitPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APISpec) DeepCopyInto(out *APISpec) {
	*out = *in
	if in.Production != nil {
		in, out := &in.Production, &out.Production
		*out = make([]EnvConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Sandbox != nil {
		in, out := &in.Sandbox, &out.Sandbox
		*out = make([]EnvConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.APIProperties != nil {
		in, out := &in.APIProperties, &out.APIProperties
		*out = make([]Property, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APISpec.
func (in *APISpec) DeepCopy() *APISpec {
	if in == nil {
		return nil
	}
	out := new(APISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIStatus) DeepCopyInto(out *APIStatus) {
	*out = *in
	in.DeploymentStatus.DeepCopyInto(&out.DeploymentStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIStatus.
func (in *APIStatus) DeepCopy() *APIStatus {
	if in == nil {
		return nil
	}
	out := new(APIStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backend) DeepCopyInto(out *Backend) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backend.
func (in *Backend) DeepCopy() *Backend {
	if in == nil {
		return nil
	}
	out := new(Backend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Backend) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendJWT) DeepCopyInto(out *BackendJWT) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendJWT.
func (in *BackendJWT) DeepCopy() *BackendJWT {
	if in == nil {
		return nil
	}
	out := new(BackendJWT)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackendJWT) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendJWTList) DeepCopyInto(out *BackendJWTList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackendJWT, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendJWTList.
func (in *BackendJWTList) DeepCopy() *BackendJWTList {
	if in == nil {
		return nil
	}
	out := new(BackendJWTList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackendJWTList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendJWTSpec) DeepCopyInto(out *BackendJWTSpec) {
	*out = *in
	if in.CustomClaims != nil {
		in, out := &in.CustomClaims, &out.CustomClaims
		*out = make([]CustomClaim, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendJWTSpec.
func (in *BackendJWTSpec) DeepCopy() *BackendJWTSpec {
	if in == nil {
		return nil
	}
	out := new(BackendJWTSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendJWTStatus) DeepCopyInto(out *BackendJWTStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendJWTStatus.
func (in *BackendJWTStatus) DeepCopy() *BackendJWTStatus {
	if in == nil {
		return nil
	}
	out := new(BackendJWTStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendJWTToken) DeepCopyInto(out *BackendJWTToken) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendJWTToken.
func (in *BackendJWTToken) DeepCopy() *BackendJWTToken {
	if in == nil {
		return nil
	}
	out := new(BackendJWTToken)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendList) DeepCopyInto(out *BackendList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Backend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendList.
func (in *BackendList) DeepCopy() *BackendList {
	if in == nil {
		return nil
	}
	out := new(BackendList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackendList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendReference) DeepCopyInto(out *BackendReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendReference.
func (in *BackendReference) DeepCopy() *BackendReference {
	if in == nil {
		return nil
	}
	out := new(BackendReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendSpec) DeepCopyInto(out *BackendSpec) {
	*out = *in
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]Service, len(*in))
		copy(*out, *in)
	}
	if in.TLS != nil {
		in, out := &in.TLS, &out.TLS
		*out = new(TLSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = new(SecurityConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.CircuitBreaker != nil {
		in, out := &in.CircuitBreaker, &out.CircuitBreaker
		*out = new(CircuitBreaker)
		**out = **in
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(Timeout)
		**out = **in
	}
	if in.Retry != nil {
		in, out := &in.Retry, &out.Retry
		*out = new(RetryConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.HealthCheck != nil {
		in, out := &in.HealthCheck, &out.HealthCheck
		*out = new(HealthCheck)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendSpec.
func (in *BackendSpec) DeepCopy() *BackendSpec {
	if in == nil {
		return nil
	}
	out := new(BackendSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendStatus) DeepCopyInto(out *BackendStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendStatus.
func (in *BackendStatus) DeepCopy() *BackendStatus {
	if in == nil {
		return nil
	}
	out := new(BackendStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BasicSecurityConfig) DeepCopyInto(out *BasicSecurityConfig) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BasicSecurityConfig.
func (in *BasicSecurityConfig) DeepCopy() *BasicSecurityConfig {
	if in == nil {
		return nil
	}
	out := new(BasicSecurityConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CORSPolicy) DeepCopyInto(out *CORSPolicy) {
	*out = *in
	if in.AccessControlAllowHeaders != nil {
		in, out := &in.AccessControlAllowHeaders, &out.AccessControlAllowHeaders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AccessControlAllowMethods != nil {
		in, out := &in.AccessControlAllowMethods, &out.AccessControlAllowMethods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AccessControlAllowOrigins != nil {
		in, out := &in.AccessControlAllowOrigins, &out.AccessControlAllowOrigins
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AccessControlExposeHeaders != nil {
		in, out := &in.AccessControlExposeHeaders, &out.AccessControlExposeHeaders
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AccessControlMaxAge != nil {
		in, out := &in.AccessControlMaxAge, &out.AccessControlMaxAge
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CORSPolicy.
func (in *CORSPolicy) DeepCopy() *CORSPolicy {
	if in == nil {
		return nil
	}
	out := new(CORSPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CircuitBreaker) DeepCopyInto(out *CircuitBreaker) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CircuitBreaker.
func (in *CircuitBreaker) DeepCopy() *CircuitBreaker {
	if in == nil {
		return nil
	}
	out := new(CircuitBreaker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomClaim) DeepCopyInto(out *CustomClaim) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomClaim.
func (in *CustomClaim) DeepCopy() *CustomClaim {
	if in == nil {
		return nil
	}
	out := new(CustomClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRateLimitPolicy) DeepCopyInto(out *CustomRateLimitPolicy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRateLimitPolicy.
func (in *CustomRateLimitPolicy) DeepCopy() *CustomRateLimitPolicy {
	if in == nil {
		return nil
	}
	out := new(CustomRateLimitPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomRateLimitPolicyDef) DeepCopyInto(out *CustomRateLimitPolicyDef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomRateLimitPolicyDef.
func (in *CustomRateLimitPolicyDef) DeepCopy() *CustomRateLimitPolicyDef {
	if in == nil {
		return nil
	}
	out := new(CustomRateLimitPolicyDef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentStatus) DeepCopyInto(out *DeploymentStatus) {
	*out = *in
	if in.TransitionTime != nil {
		in, out := &in.TransitionTime, &out.TransitionTime
		*out = (*in).DeepCopy()
	}
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentStatus.
func (in *DeploymentStatus) DeepCopy() *DeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(DeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvConfig) DeepCopyInto(out *EnvConfig) {
	*out = *in
	if in.HTTPRouteRefs != nil {
		in, out := &in.HTTPRouteRefs, &out.HTTPRouteRefs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvConfig.
func (in *EnvConfig) DeepCopy() *EnvConfig {
	if in == nil {
		return nil
	}
	out := new(EnvConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthCheck) DeepCopyInto(out *HealthCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthCheck.
func (in *HealthCheck) DeepCopy() *HealthCheck {
	if in == nil {
		return nil
	}
	out := new(HealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterceptorReference) DeepCopyInto(out *InterceptorReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterceptorReference.
func (in *InterceptorReference) DeepCopy() *InterceptorReference {
	if in == nil {
		return nil
	}
	out := new(InterceptorReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterceptorService) DeepCopyInto(out *InterceptorService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterceptorService.
func (in *InterceptorService) DeepCopy() *InterceptorService {
	if in == nil {
		return nil
	}
	out := new(InterceptorService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InterceptorService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterceptorServiceList) DeepCopyInto(out *InterceptorServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InterceptorService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterceptorServiceList.
func (in *InterceptorServiceList) DeepCopy() *InterceptorServiceList {
	if in == nil {
		return nil
	}
	out := new(InterceptorServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InterceptorServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterceptorServiceSpec) DeepCopyInto(out *InterceptorServiceSpec) {
	*out = *in
	out.BackendRef = in.BackendRef
	if in.Includes != nil {
		in, out := &in.Includes, &out.Includes
		*out = make([]InterceptorInclusion, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterceptorServiceSpec.
func (in *InterceptorServiceSpec) DeepCopy() *InterceptorServiceSpec {
	if in == nil {
		return nil
	}
	out := new(InterceptorServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterceptorServiceStatus) DeepCopyInto(out *InterceptorServiceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterceptorServiceStatus.
func (in *InterceptorServiceStatus) DeepCopy() *InterceptorServiceStatus {
	if in == nil {
		return nil
	}
	out := new(InterceptorServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicySpec) DeepCopyInto(out *PolicySpec) {
	*out = *in
	if in.RequestInterceptors != nil {
		in, out := &in.RequestInterceptors, &out.RequestInterceptors
		*out = make([]InterceptorReference, len(*in))
		copy(*out, *in)
	}
	if in.ResponseInterceptors != nil {
		in, out := &in.ResponseInterceptors, &out.ResponseInterceptors
		*out = make([]InterceptorReference, len(*in))
		copy(*out, *in)
	}
	if in.BackendJWTPolicy != nil {
		in, out := &in.BackendJWTPolicy, &out.BackendJWTPolicy
		*out = new(BackendJWTToken)
		**out = **in
	}
	if in.CORSPolicy != nil {
		in, out := &in.CORSPolicy, &out.CORSPolicy
		*out = new(CORSPolicy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicySpec.
func (in *PolicySpec) DeepCopy() *PolicySpec {
	if in == nil {
		return nil
	}
	out := new(PolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Property) DeepCopyInto(out *Property) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Property.
func (in *Property) DeepCopy() *Property {
	if in == nil {
		return nil
	}
	out := new(Property)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitAPIPolicy) DeepCopyInto(out *RateLimitAPIPolicy) {
	*out = *in
	if in.API != nil {
		in, out := &in.API, &out.API
		*out = new(APIRateLimitPolicy)
		**out = **in
	}
	if in.Custom != nil {
		in, out := &in.Custom, &out.Custom
		*out = new(CustomRateLimitPolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitAPIPolicy.
func (in *RateLimitAPIPolicy) DeepCopy() *RateLimitAPIPolicy {
	if in == nil {
		return nil
	}
	out := new(RateLimitAPIPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitPolicy) DeepCopyInto(out *RateLimitPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitPolicy.
func (in *RateLimitPolicy) DeepCopy() *RateLimitPolicy {
	if in == nil {
		return nil
	}
	out := new(RateLimitPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RateLimitPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitPolicyList) DeepCopyInto(out *RateLimitPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RateLimitPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitPolicyList.
func (in *RateLimitPolicyList) DeepCopy() *RateLimitPolicyList {
	if in == nil {
		return nil
	}
	out := new(RateLimitPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RateLimitPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitPolicySpec) DeepCopyInto(out *RateLimitPolicySpec) {
	*out = *in
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(RateLimitAPIPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Override != nil {
		in, out := &in.Override, &out.Override
		*out = new(RateLimitAPIPolicy)
		(*in).DeepCopyInto(*out)
	}
	in.TargetRef.DeepCopyInto(&out.TargetRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitPolicySpec.
func (in *RateLimitPolicySpec) DeepCopy() *RateLimitPolicySpec {
	if in == nil {
		return nil
	}
	out := new(RateLimitPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitPolicyStatus) DeepCopyInto(out *RateLimitPolicyStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitPolicyStatus.
func (in *RateLimitPolicyStatus) DeepCopy() *RateLimitPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(RateLimitPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RefConfig) DeepCopyInto(out *RefConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RefConfig.
func (in *RefConfig) DeepCopy() *RefConfig {
	if in == nil {
		return nil
	}
	out := new(RefConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolveRateLimit) DeepCopyInto(out *ResolveRateLimit) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolveRateLimit.
func (in *ResolveRateLimit) DeepCopy() *ResolveRateLimit {
	if in == nil {
		return nil
	}
	out := new(ResolveRateLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolveRateLimitAPIPolicy) DeepCopyInto(out *ResolveRateLimitAPIPolicy) {
	*out = *in
	out.API = in.API
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResolveResource, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolveRateLimitAPIPolicy.
func (in *ResolveRateLimitAPIPolicy) DeepCopy() *ResolveRateLimitAPIPolicy {
	if in == nil {
		return nil
	}
	out := new(ResolveRateLimitAPIPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolveResource) DeepCopyInto(out *ResolveResource) {
	*out = *in
	out.ResourceRatelimit = in.ResourceRatelimit
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolveResource.
func (in *ResolveResource) DeepCopy() *ResolveResource {
	if in == nil {
		return nil
	}
	out := new(ResolveResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RetryConfig) DeepCopyInto(out *RetryConfig) {
	*out = *in
	if in.StatusCodes != nil {
		in, out := &in.StatusCodes, &out.StatusCodes
		*out = make([]uint32, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RetryConfig.
func (in *RetryConfig) DeepCopy() *RetryConfig {
	if in == nil {
		return nil
	}
	out := new(RetryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityConfig) DeepCopyInto(out *SecurityConfig) {
	*out = *in
	if in.Basic != nil {
		in, out := &in.Basic, &out.Basic
		*out = new(BasicSecurityConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityConfig.
func (in *SecurityConfig) DeepCopy() *SecurityConfig {
	if in == nil {
		return nil
	}
	out := new(SecurityConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	*out = *in
	if in.CertificateInline != nil {
		in, out := &in.CertificateInline, &out.CertificateInline
		*out = new(string)
		**out = **in
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(RefConfig)
		**out = **in
	}
	if in.ConfigMapRef != nil {
		in, out := &in.ConfigMapRef, &out.ConfigMapRef
		*out = new(RefConfig)
		**out = **in
	}
	if in.AllowedSANs != nil {
		in, out := &in.AllowedSANs, &out.AllowedSANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfig.
func (in *TLSConfig) DeepCopy() *TLSConfig {
	if in == nil {
		return nil
	}
	out := new(TLSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Timeout) DeepCopyInto(out *Timeout) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Timeout.
func (in *Timeout) DeepCopy() *Timeout {
	if in == nil {
		return nil
	}
	out := new(Timeout)
	in.DeepCopyInto(out)
	return out
}
