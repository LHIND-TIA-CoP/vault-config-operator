//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package utils

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeAuthConfiguration) DeepCopyInto(out *KubeAuthConfiguration) {
	*out = *in
	if in.ServiceAccount != nil {
		in, out := &in.ServiceAccount, &out.ServiceAccount
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeAuthConfiguration.
func (in *KubeAuthConfiguration) DeepCopy() *KubeAuthConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeAuthConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootCredentialConfig) DeepCopyInto(out *RootCredentialConfig) {
	*out = *in
	if in.VaultSecret != nil {
		in, out := &in.VaultSecret, &out.VaultSecret
		*out = new(VaultSecretReference)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.RandomSecret != nil {
		in, out := &in.RandomSecret, &out.RandomSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootCredentialConfig.
func (in *RootCredentialConfig) DeepCopy() *RootCredentialConfig {
	if in == nil {
		return nil
	}
	out := new(RootCredentialConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	*out = *in
	if in.Cacert != nil {
		in, out := &in.Cacert, &out.Cacert
		*out = new(string)
		**out = **in
	}
	if in.TLSSecret != nil {
		in, out := &in.TLSSecret, &out.TLSSecret
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.TLSServerName != nil {
		in, out := &in.TLSServerName, &out.TLSServerName
		*out = new(string)
		**out = **in
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
func (in *TargetNamespaceConfig) DeepCopyInto(out *TargetNamespaceConfig) {
	*out = *in
	if in.TargetNamespaceSelector != nil {
		in, out := &in.TargetNamespaceSelector, &out.TargetNamespaceSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.TargetNamespaces != nil {
		in, out := &in.TargetNamespaces, &out.TargetNamespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetNamespaceConfig.
func (in *TargetNamespaceConfig) DeepCopy() *TargetNamespaceConfig {
	if in == nil {
		return nil
	}
	out := new(TargetNamespaceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultConnection) DeepCopyInto(out *VaultConnection) {
	*out = *in
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		*out = new(TLSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.TimeOut != nil {
		in, out := &in.TimeOut, &out.TimeOut
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.MaxRetries != nil {
		in, out := &in.MaxRetries, &out.MaxRetries
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultConnection.
func (in *VaultConnection) DeepCopy() *VaultConnection {
	if in == nil {
		return nil
	}
	out := new(VaultConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultSecretReference) DeepCopyInto(out *VaultSecretReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultSecretReference.
func (in *VaultSecretReference) DeepCopy() *VaultSecretReference {
	if in == nil {
		return nil
	}
	out := new(VaultSecretReference)
	in.DeepCopyInto(out)
	return out
}
