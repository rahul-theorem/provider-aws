/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type HostObservation struct {

	// The ARN of the Dedicated Host.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the allocated Dedicated Host. This is used to launch an instance onto a specific host.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the AWS account that owns the Dedicated Host.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type HostParameters struct {

	// Indicates whether the host accepts any untargeted instance launches that match its instance type configuration, or if it only accepts Host tenancy instance launches that specify its unique host ID. Valid values: on, off. Default: on.
	// +kubebuilder:validation:Optional
	AutoPlacement *string `json:"autoPlacement,omitempty" tf:"auto_placement,omitempty"`

	// The Availability Zone in which to allocate the Dedicated Host.
	// +kubebuilder:validation:Required
	AvailabilityZone *string `json:"availabilityZone" tf:"availability_zone,omitempty"`

	// Indicates whether to enable or disable host recovery for the Dedicated Host. Valid values: on, off. Default: off.
	// +kubebuilder:validation:Optional
	HostRecovery *string `json:"hostRecovery,omitempty" tf:"host_recovery,omitempty"`

	// Specifies the instance family to be supported by the Dedicated Hosts. If you specify an instance family, the Dedicated Hosts support multiple instance types within that instance family. Exactly one of instance_family or instance_type must be specified.
	// +kubebuilder:validation:Optional
	InstanceFamily *string `json:"instanceFamily,omitempty" tf:"instance_family,omitempty"`

	// Specifies the instance type to be supported by the Dedicated Hosts. If you specify an instance type, the Dedicated Hosts support instances of the specified instance type only.  Exactly one of instance_family or instance_type must be specified.
	// +kubebuilder:validation:Optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// HostSpec defines the desired state of Host
type HostSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostParameters `json:"forProvider"`
}

// HostStatus defines the observed state of Host.
type HostStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Host is the Schema for the Hosts API. Provides an EC2 Host resource. This allows Dedicated Hosts to be allocated, modified, and released.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Host struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostSpec   `json:"spec"`
	Status            HostStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostList contains a list of Hosts
type HostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Host `json:"items"`
}

// Repository type metadata.
var (
	Host_Kind             = "Host"
	Host_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Host_Kind}.String()
	Host_KindAPIVersion   = Host_Kind + "." + CRDGroupVersion.String()
	Host_GroupVersionKind = CRDGroupVersion.WithKind(Host_Kind)
)

func init() {
	SchemeBuilder.Register(&Host{}, &HostList{})
}
