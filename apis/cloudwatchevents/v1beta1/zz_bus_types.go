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

type BusObservation struct {

	// The Amazon Resource Name (ARN) of the event bus.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The partner event source that the new event bus will be matched with. Must match name.
	EventSourceName *string `json:"eventSourceName,omitempty" tf:"event_source_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type BusParameters struct {

	// The partner event source that the new event bus will be matched with. Must match name.
	// +kubebuilder:validation:Optional
	EventSourceName *string `json:"eventSourceName,omitempty" tf:"event_source_name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// BusSpec defines the desired state of Bus
type BusSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BusParameters `json:"forProvider"`
}

// BusStatus defines the observed state of Bus.
type BusStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BusObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Bus is the Schema for the Buss API. Provides an EventBridge event bus resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Bus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BusSpec   `json:"spec"`
	Status            BusStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BusList contains a list of Buss
type BusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bus `json:"items"`
}

// Repository type metadata.
var (
	Bus_Kind             = "Bus"
	Bus_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Bus_Kind}.String()
	Bus_KindAPIVersion   = Bus_Kind + "." + CRDGroupVersion.String()
	Bus_GroupVersionKind = CRDGroupVersion.WithKind(Bus_Kind)
)

func init() {
	SchemeBuilder.Register(&Bus{}, &BusList{})
}
