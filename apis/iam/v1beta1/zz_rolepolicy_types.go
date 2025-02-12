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

type RolePolicyInitParameters struct {

	// The inline policy document. This is a JSON formatted string
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type RolePolicyObservation struct {

	// The role policy ID, in the form of role_name:role_policy_name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The inline policy document. This is a JSON formatted string
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// The name of the IAM role to attach to the policy.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`
}

type RolePolicyParameters struct {

	// The inline policy document. This is a JSON formatted string
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`

	// The name of the IAM role to attach to the policy.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Reference to a Role in iam to populate role.
	// +kubebuilder:validation:Optional
	RoleRef *v1.Reference `json:"roleRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate role.
	// +kubebuilder:validation:Optional
	RoleSelector *v1.Selector `json:"roleSelector,omitempty" tf:"-"`
}

// RolePolicySpec defines the desired state of RolePolicy
type RolePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RolePolicyParameters `json:"forProvider"`
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
	InitProvider RolePolicyInitParameters `json:"initProvider,omitempty"`
}

// RolePolicyStatus defines the observed state of RolePolicy.
type RolePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RolePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RolePolicy is the Schema for the RolePolicys API. Provides an IAM role policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RolePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.policy) || (has(self.initProvider) && has(self.initProvider.policy))",message="spec.forProvider.policy is a required parameter"
	Spec   RolePolicySpec   `json:"spec"`
	Status RolePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RolePolicyList contains a list of RolePolicys
type RolePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RolePolicy `json:"items"`
}

// Repository type metadata.
var (
	RolePolicy_Kind             = "RolePolicy"
	RolePolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RolePolicy_Kind}.String()
	RolePolicy_KindAPIVersion   = RolePolicy_Kind + "." + CRDGroupVersion.String()
	RolePolicy_GroupVersionKind = CRDGroupVersion.WithKind(RolePolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&RolePolicy{}, &RolePolicyList{})
}
