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

type ContainerObservation struct {

	// The ARN of the container.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The DNS endpoint of the container.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ContainerParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ContainerSpec defines the desired state of Container
type ContainerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContainerParameters `json:"forProvider"`
}

// ContainerStatus defines the observed state of Container.
type ContainerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContainerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Container is the Schema for the Containers API. Provides a MediaStore Container.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Container struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContainerSpec   `json:"spec"`
	Status            ContainerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerList contains a list of Containers
type ContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Container `json:"items"`
}

// Repository type metadata.
var (
	Container_Kind             = "Container"
	Container_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Container_Kind}.String()
	Container_KindAPIVersion   = Container_Kind + "." + CRDGroupVersion.String()
	Container_GroupVersionKind = CRDGroupVersion.WithKind(Container_Kind)
)

func init() {
	SchemeBuilder.Register(&Container{}, &ContainerList{})
}
