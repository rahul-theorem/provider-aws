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

type ExtensionAssociationInitParameters struct {

	// The parameter names and values defined for the association.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type ExtensionAssociationObservation struct {

	// ARN of the AppConfig Extension Association.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ARN of the extension defined in the association.
	ExtensionArn *string `json:"extensionArn,omitempty" tf:"extension_arn,omitempty"`

	// The version number for the extension defined in the association.
	ExtensionVersion *float64 `json:"extensionVersion,omitempty" tf:"extension_version,omitempty"`

	// AppConfig Extension Association ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The parameter names and values defined for the association.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The ARN of the application, configuration profile, or environment to associate with the extension.
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`
}

type ExtensionAssociationParameters struct {

	// The ARN of the extension defined in the association.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appconfig/v1beta1.Extension
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ExtensionArn *string `json:"extensionArn,omitempty" tf:"extension_arn,omitempty"`

	// Reference to a Extension in appconfig to populate extensionArn.
	// +kubebuilder:validation:Optional
	ExtensionArnRef *v1.Reference `json:"extensionArnRef,omitempty" tf:"-"`

	// Selector for a Extension in appconfig to populate extensionArn.
	// +kubebuilder:validation:Optional
	ExtensionArnSelector *v1.Selector `json:"extensionArnSelector,omitempty" tf:"-"`

	// The parameter names and values defined for the association.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The ARN of the application, configuration profile, or environment to associate with the extension.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/appconfig/v1beta1.Application
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// Reference to a Application in appconfig to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnRef *v1.Reference `json:"resourceArnRef,omitempty" tf:"-"`

	// Selector for a Application in appconfig to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnSelector *v1.Selector `json:"resourceArnSelector,omitempty" tf:"-"`
}

// ExtensionAssociationSpec defines the desired state of ExtensionAssociation
type ExtensionAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExtensionAssociationParameters `json:"forProvider"`
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
	InitProvider ExtensionAssociationInitParameters `json:"initProvider,omitempty"`
}

// ExtensionAssociationStatus defines the observed state of ExtensionAssociation.
type ExtensionAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExtensionAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ExtensionAssociation is the Schema for the ExtensionAssociations API. Associates an AppConfig Extension with a Resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ExtensionAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExtensionAssociationSpec   `json:"spec"`
	Status            ExtensionAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExtensionAssociationList contains a list of ExtensionAssociations
type ExtensionAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExtensionAssociation `json:"items"`
}

// Repository type metadata.
var (
	ExtensionAssociation_Kind             = "ExtensionAssociation"
	ExtensionAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ExtensionAssociation_Kind}.String()
	ExtensionAssociation_KindAPIVersion   = ExtensionAssociation_Kind + "." + CRDGroupVersion.String()
	ExtensionAssociation_GroupVersionKind = CRDGroupVersion.WithKind(ExtensionAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&ExtensionAssociation{}, &ExtensionAssociationList{})
}
