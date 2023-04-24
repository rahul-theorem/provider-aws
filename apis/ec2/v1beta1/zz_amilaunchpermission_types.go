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

type AMILaunchPermissionObservation struct {

	// AWS account ID for the launch permission.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Name of the group for the launch permission. Valid values: "all".
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// Launch permission ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the AMI.
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// ARN of an organization for the launch permission.
	OrganizationArn *string `json:"organizationArn,omitempty" tf:"organization_arn,omitempty"`

	// ARN of an organizational unit for the launch permission.
	OrganizationalUnitArn *string `json:"organizationalUnitArn,omitempty" tf:"organizational_unit_arn,omitempty"`
}

type AMILaunchPermissionParameters struct {

	// AWS account ID for the launch permission.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Name of the group for the launch permission. Valid values: "all".
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// ID of the AMI.
	// +crossplane:generate:reference:type=AMI
	// +kubebuilder:validation:Optional
	ImageID *string `json:"imageId,omitempty" tf:"image_id,omitempty"`

	// Reference to a AMI to populate imageId.
	// +kubebuilder:validation:Optional
	ImageIDRef *v1.Reference `json:"imageIdRef,omitempty" tf:"-"`

	// Selector for a AMI to populate imageId.
	// +kubebuilder:validation:Optional
	ImageIDSelector *v1.Selector `json:"imageIdSelector,omitempty" tf:"-"`

	// ARN of an organization for the launch permission.
	// +kubebuilder:validation:Optional
	OrganizationArn *string `json:"organizationArn,omitempty" tf:"organization_arn,omitempty"`

	// ARN of an organizational unit for the launch permission.
	// +kubebuilder:validation:Optional
	OrganizationalUnitArn *string `json:"organizationalUnitArn,omitempty" tf:"organizational_unit_arn,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// AMILaunchPermissionSpec defines the desired state of AMILaunchPermission
type AMILaunchPermissionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AMILaunchPermissionParameters `json:"forProvider"`
}

// AMILaunchPermissionStatus defines the observed state of AMILaunchPermission.
type AMILaunchPermissionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AMILaunchPermissionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AMILaunchPermission is the Schema for the AMILaunchPermissions API. Adds a launch permission to an Amazon Machine Image (AMI).
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AMILaunchPermission struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AMILaunchPermissionSpec   `json:"spec"`
	Status            AMILaunchPermissionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AMILaunchPermissionList contains a list of AMILaunchPermissions
type AMILaunchPermissionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AMILaunchPermission `json:"items"`
}

// Repository type metadata.
var (
	AMILaunchPermission_Kind             = "AMILaunchPermission"
	AMILaunchPermission_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AMILaunchPermission_Kind}.String()
	AMILaunchPermission_KindAPIVersion   = AMILaunchPermission_Kind + "." + CRDGroupVersion.String()
	AMILaunchPermission_GroupVersionKind = CRDGroupVersion.WithKind(AMILaunchPermission_Kind)
)

func init() {
	SchemeBuilder.Register(&AMILaunchPermission{}, &AMILaunchPermissionList{})
}
