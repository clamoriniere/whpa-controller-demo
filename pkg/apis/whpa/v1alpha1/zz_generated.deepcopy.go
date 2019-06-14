// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatermarkHorizontalPodAutoscaler) DeepCopyInto(out *WatermarkHorizontalPodAutoscaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatermarkHorizontalPodAutoscaler.
func (in *WatermarkHorizontalPodAutoscaler) DeepCopy() *WatermarkHorizontalPodAutoscaler {
	if in == nil {
		return nil
	}
	out := new(WatermarkHorizontalPodAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WatermarkHorizontalPodAutoscaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatermarkHorizontalPodAutoscalerList) DeepCopyInto(out *WatermarkHorizontalPodAutoscalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WatermarkHorizontalPodAutoscaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatermarkHorizontalPodAutoscalerList.
func (in *WatermarkHorizontalPodAutoscalerList) DeepCopy() *WatermarkHorizontalPodAutoscalerList {
	if in == nil {
		return nil
	}
	out := new(WatermarkHorizontalPodAutoscalerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WatermarkHorizontalPodAutoscalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatermarkHorizontalPodAutoscalerSpec) DeepCopyInto(out *WatermarkHorizontalPodAutoscalerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatermarkHorizontalPodAutoscalerSpec.
func (in *WatermarkHorizontalPodAutoscalerSpec) DeepCopy() *WatermarkHorizontalPodAutoscalerSpec {
	if in == nil {
		return nil
	}
	out := new(WatermarkHorizontalPodAutoscalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WatermarkHorizontalPodAutoscalerStatus) DeepCopyInto(out *WatermarkHorizontalPodAutoscalerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WatermarkHorizontalPodAutoscalerStatus.
func (in *WatermarkHorizontalPodAutoscalerStatus) DeepCopy() *WatermarkHorizontalPodAutoscalerStatus {
	if in == nil {
		return nil
	}
	out := new(WatermarkHorizontalPodAutoscalerStatus)
	in.DeepCopyInto(out)
	return out
}