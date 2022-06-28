//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Qvm) DeepCopyInto(out *Qvm) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Qvm.
func (in *Qvm) DeepCopy() *Qvm {
	if in == nil {
		return nil
	}
	out := new(Qvm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Qvm) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QvmList) DeepCopyInto(out *QvmList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Qvm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QvmList.
func (in *QvmList) DeepCopy() *QvmList {
	if in == nil {
		return nil
	}
	out := new(QvmList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QvmList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QvmSpec) DeepCopyInto(out *QvmSpec) {
	*out = *in
	in.VM.DeepCopyInto(&out.VM)
	if in.FloatingIPs != nil {
		in, out := &in.FloatingIPs, &out.FloatingIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QvmSpec.
func (in *QvmSpec) DeepCopy() *QvmSpec {
	if in == nil {
		return nil
	}
	out := new(QvmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QvmStatus) DeepCopyInto(out *QvmStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QvmStatus.
func (in *QvmStatus) DeepCopy() *QvmStatus {
	if in == nil {
		return nil
	}
	out := new(QvmStatus)
	in.DeepCopyInto(out)
	return out
}
