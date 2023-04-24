//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnomalyMonitor) DeepCopyInto(out *AnomalyMonitor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnomalyMonitor.
func (in *AnomalyMonitor) DeepCopy() *AnomalyMonitor {
	if in == nil {
		return nil
	}
	out := new(AnomalyMonitor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnomalyMonitor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnomalyMonitorList) DeepCopyInto(out *AnomalyMonitorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AnomalyMonitor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnomalyMonitorList.
func (in *AnomalyMonitorList) DeepCopy() *AnomalyMonitorList {
	if in == nil {
		return nil
	}
	out := new(AnomalyMonitorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnomalyMonitorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnomalyMonitorObservation) DeepCopyInto(out *AnomalyMonitorObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MonitorDimension != nil {
		in, out := &in.MonitorDimension, &out.MonitorDimension
		*out = new(string)
		**out = **in
	}
	if in.MonitorSpecification != nil {
		in, out := &in.MonitorSpecification, &out.MonitorSpecification
		*out = new(string)
		**out = **in
	}
	if in.MonitorType != nil {
		in, out := &in.MonitorType, &out.MonitorType
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnomalyMonitorObservation.
func (in *AnomalyMonitorObservation) DeepCopy() *AnomalyMonitorObservation {
	if in == nil {
		return nil
	}
	out := new(AnomalyMonitorObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnomalyMonitorParameters) DeepCopyInto(out *AnomalyMonitorParameters) {
	*out = *in
	if in.MonitorDimension != nil {
		in, out := &in.MonitorDimension, &out.MonitorDimension
		*out = new(string)
		**out = **in
	}
	if in.MonitorSpecification != nil {
		in, out := &in.MonitorSpecification, &out.MonitorSpecification
		*out = new(string)
		**out = **in
	}
	if in.MonitorType != nil {
		in, out := &in.MonitorType, &out.MonitorType
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnomalyMonitorParameters.
func (in *AnomalyMonitorParameters) DeepCopy() *AnomalyMonitorParameters {
	if in == nil {
		return nil
	}
	out := new(AnomalyMonitorParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnomalyMonitorSpec) DeepCopyInto(out *AnomalyMonitorSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnomalyMonitorSpec.
func (in *AnomalyMonitorSpec) DeepCopy() *AnomalyMonitorSpec {
	if in == nil {
		return nil
	}
	out := new(AnomalyMonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnomalyMonitorStatus) DeepCopyInto(out *AnomalyMonitorStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnomalyMonitorStatus.
func (in *AnomalyMonitorStatus) DeepCopy() *AnomalyMonitorStatus {
	if in == nil {
		return nil
	}
	out := new(AnomalyMonitorStatus)
	in.DeepCopyInto(out)
	return out
}
