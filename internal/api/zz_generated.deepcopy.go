//go:build !ignore_autogenerated

/*
Copyright 2025.

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

package api

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespacedSelector) DeepCopyInto(out *NamespacedSelector) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespacedSelector.
func (in *NamespacedSelector) DeepCopy() *NamespacedSelector {
	if in == nil {
		return nil
	}
	out := new(NamespacedSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SopsMetadata) DeepCopyInto(out *SopsMetadata) {
	*out = *in
	if in.AwsKms != nil {
		in, out := &in.AwsKms, &out.AwsKms
		*out = make([]KmsDataItem, len(*in))
		copy(*out, *in)
	}
	if in.Pgp != nil {
		in, out := &in.Pgp, &out.Pgp
		*out = make([]PgpDataItem, len(*in))
		copy(*out, *in)
	}
	if in.AzureKms != nil {
		in, out := &in.AzureKms, &out.AzureKms
		*out = make([]AzureKmsItem, len(*in))
		copy(*out, *in)
	}
	if in.HcVault != nil {
		in, out := &in.HcVault, &out.HcVault
		*out = make([]HcVaultItem, len(*in))
		copy(*out, *in)
	}
	if in.GcpKms != nil {
		in, out := &in.GcpKms, &out.GcpKms
		*out = make([]GcpKmsDataItem, len(*in))
		copy(*out, *in)
	}
	if in.Age != nil {
		in, out := &in.Age, &out.Age
		*out = make([]AgeItem, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SopsMetadata.
func (in *SopsMetadata) DeepCopy() *SopsMetadata {
	if in == nil {
		return nil
	}
	out := new(SopsMetadata)
	in.DeepCopyInto(out)
	return out
}
