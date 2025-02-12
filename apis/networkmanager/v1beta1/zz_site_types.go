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

type SiteInitParameters struct {

	// Description of the Site.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The site location as documented below.
	Location []SiteLocationInitParameters `json:"location,omitempty" tf:"location,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SiteLocationInitParameters struct {

	// Address of the location.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Latitude of the location.
	Latitude *string `json:"latitude,omitempty" tf:"latitude,omitempty"`

	// Longitude of the location.
	Longitude *string `json:"longitude,omitempty" tf:"longitude,omitempty"`
}

type SiteLocationObservation struct {

	// Address of the location.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Latitude of the location.
	Latitude *string `json:"latitude,omitempty" tf:"latitude,omitempty"`

	// Longitude of the location.
	Longitude *string `json:"longitude,omitempty" tf:"longitude,omitempty"`
}

type SiteLocationParameters struct {

	// Address of the location.
	// +kubebuilder:validation:Optional
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// Latitude of the location.
	// +kubebuilder:validation:Optional
	Latitude *string `json:"latitude,omitempty" tf:"latitude,omitempty"`

	// Longitude of the location.
	// +kubebuilder:validation:Optional
	Longitude *string `json:"longitude,omitempty" tf:"longitude,omitempty"`
}

type SiteObservation struct {

	// Site Amazon Resource Name (ARN)
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Description of the Site.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Global Network to create the site in.
	GlobalNetworkID *string `json:"globalNetworkId,omitempty" tf:"global_network_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The site location as documented below.
	Location []SiteLocationObservation `json:"location,omitempty" tf:"location,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type SiteParameters struct {

	// Description of the Site.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Global Network to create the site in.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkmanager/v1beta1.GlobalNetwork
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GlobalNetworkID *string `json:"globalNetworkId,omitempty" tf:"global_network_id,omitempty"`

	// Reference to a GlobalNetwork in networkmanager to populate globalNetworkId.
	// +kubebuilder:validation:Optional
	GlobalNetworkIDRef *v1.Reference `json:"globalNetworkIdRef,omitempty" tf:"-"`

	// Selector for a GlobalNetwork in networkmanager to populate globalNetworkId.
	// +kubebuilder:validation:Optional
	GlobalNetworkIDSelector *v1.Selector `json:"globalNetworkIdSelector,omitempty" tf:"-"`

	// The site location as documented below.
	// +kubebuilder:validation:Optional
	Location []SiteLocationParameters `json:"location,omitempty" tf:"location,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SiteSpec defines the desired state of Site
type SiteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SiteParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SiteInitParameters `json:"initProvider,omitempty"`
}

// SiteStatus defines the observed state of Site.
type SiteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SiteObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Site is the Schema for the Sites API. Creates a site in a global network.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Site struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SiteSpec   `json:"spec"`
	Status            SiteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SiteList contains a list of Sites
type SiteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Site `json:"items"`
}

// Repository type metadata.
var (
	Site_Kind             = "Site"
	Site_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Site_Kind}.String()
	Site_KindAPIVersion   = Site_Kind + "." + CRDGroupVersion.String()
	Site_GroupVersionKind = CRDGroupVersion.WithKind(Site_Kind)
)

func init() {
	SchemeBuilder.Register(&Site{}, &SiteList{})
}
