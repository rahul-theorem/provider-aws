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

type MappingRuleInitParameters struct {

	// The claim name that must be present in the token, for example, "isAdmin" or "paid".
	Claim *string `json:"claim,omitempty" tf:"claim,omitempty"`

	// The match condition that specifies how closely the claim value in the IdP token must match Value.
	MatchType *string `json:"matchType,omitempty" tf:"match_type,omitempty"`

	// A brief string that the claim must match, for example, "paid" or "yes".
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type MappingRuleObservation struct {

	// The claim name that must be present in the token, for example, "isAdmin" or "paid".
	Claim *string `json:"claim,omitempty" tf:"claim,omitempty"`

	// The match condition that specifies how closely the claim value in the IdP token must match Value.
	MatchType *string `json:"matchType,omitempty" tf:"match_type,omitempty"`

	// The role ARN.
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// A brief string that the claim must match, for example, "paid" or "yes".
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type MappingRuleParameters struct {

	// The claim name that must be present in the token, for example, "isAdmin" or "paid".
	// +kubebuilder:validation:Optional
	Claim *string `json:"claim" tf:"claim,omitempty"`

	// The match condition that specifies how closely the claim value in the IdP token must match Value.
	// +kubebuilder:validation:Optional
	MatchType *string `json:"matchType" tf:"match_type,omitempty"`

	// The role ARN.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`

	// Reference to a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnRef *v1.Reference `json:"roleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate roleArn.
	// +kubebuilder:validation:Optional
	RoleArnSelector *v1.Selector `json:"roleArnSelector,omitempty" tf:"-"`

	// A brief string that the claim must match, for example, "paid" or "yes".
	// +kubebuilder:validation:Optional
	Value *string `json:"value" tf:"value,omitempty"`
}

type PoolRolesAttachmentInitParameters struct {

	// A List of Role Mapping.
	RoleMapping []RoleMappingInitParameters `json:"roleMapping,omitempty" tf:"role_mapping,omitempty"`

	// The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
	Roles map[string]*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

type PoolRolesAttachmentObservation struct {

	// The identity pool ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity pool ID in the format REGION_GUID.
	IdentityPoolID *string `json:"identityPoolId,omitempty" tf:"identity_pool_id,omitempty"`

	// A List of Role Mapping.
	RoleMapping []RoleMappingObservation `json:"roleMapping,omitempty" tf:"role_mapping,omitempty"`

	// The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
	Roles map[string]*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

type PoolRolesAttachmentParameters struct {

	// An identity pool ID in the format REGION_GUID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cognitoidentity/v1beta1.Pool
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	IdentityPoolID *string `json:"identityPoolId,omitempty" tf:"identity_pool_id,omitempty"`

	// Reference to a Pool in cognitoidentity to populate identityPoolId.
	// +kubebuilder:validation:Optional
	IdentityPoolIDRef *v1.Reference `json:"identityPoolIdRef,omitempty" tf:"-"`

	// Selector for a Pool in cognitoidentity to populate identityPoolId.
	// +kubebuilder:validation:Optional
	IdentityPoolIDSelector *v1.Selector `json:"identityPoolIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A List of Role Mapping.
	// +kubebuilder:validation:Optional
	RoleMapping []RoleMappingParameters `json:"roleMapping,omitempty" tf:"role_mapping,omitempty"`

	// The map of roles associated with this pool. For a given role, the key will be either "authenticated" or "unauthenticated" and the value will be the Role ARN.
	// +kubebuilder:validation:Optional
	Roles map[string]*string `json:"roles,omitempty" tf:"roles,omitempty"`
}

type RoleMappingInitParameters struct {

	// Specifies the action to be taken if either no rules match the claim value for the Rules type, or there is no cognito:preferred_role claim and there are multiple cognito:roles matches for the Token type. Required if you specify Token or Rules as the Type.
	AmbiguousRoleResolution *string `json:"ambiguousRoleResolution,omitempty" tf:"ambiguous_role_resolution,omitempty"`

	// A string identifying the identity provider, for example, "graph.facebook.com" or "cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id". Depends on cognito_identity_providers set on aws_cognito_identity_pool resource or a aws_cognito_identity_provider resource.
	IdentityProvider *string `json:"identityProvider,omitempty" tf:"identity_provider,omitempty"`

	// The Rules Configuration to be used for mapping users to roles. You can specify up to 25 rules per identity provider. Rules are evaluated in order. The first one to match specifies the role.
	MappingRule []MappingRuleInitParameters `json:"mappingRule,omitempty" tf:"mapping_rule,omitempty"`

	// The role mapping type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RoleMappingObservation struct {

	// Specifies the action to be taken if either no rules match the claim value for the Rules type, or there is no cognito:preferred_role claim and there are multiple cognito:roles matches for the Token type. Required if you specify Token or Rules as the Type.
	AmbiguousRoleResolution *string `json:"ambiguousRoleResolution,omitempty" tf:"ambiguous_role_resolution,omitempty"`

	// A string identifying the identity provider, for example, "graph.facebook.com" or "cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id". Depends on cognito_identity_providers set on aws_cognito_identity_pool resource or a aws_cognito_identity_provider resource.
	IdentityProvider *string `json:"identityProvider,omitempty" tf:"identity_provider,omitempty"`

	// The Rules Configuration to be used for mapping users to roles. You can specify up to 25 rules per identity provider. Rules are evaluated in order. The first one to match specifies the role.
	MappingRule []MappingRuleObservation `json:"mappingRule,omitempty" tf:"mapping_rule,omitempty"`

	// The role mapping type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RoleMappingParameters struct {

	// Specifies the action to be taken if either no rules match the claim value for the Rules type, or there is no cognito:preferred_role claim and there are multiple cognito:roles matches for the Token type. Required if you specify Token or Rules as the Type.
	// +kubebuilder:validation:Optional
	AmbiguousRoleResolution *string `json:"ambiguousRoleResolution,omitempty" tf:"ambiguous_role_resolution,omitempty"`

	// A string identifying the identity provider, for example, "graph.facebook.com" or "cognito-idp.us-east-1.amazonaws.com/us-east-1_abcdefghi:app_client_id". Depends on cognito_identity_providers set on aws_cognito_identity_pool resource or a aws_cognito_identity_provider resource.
	// +kubebuilder:validation:Optional
	IdentityProvider *string `json:"identityProvider" tf:"identity_provider,omitempty"`

	// The Rules Configuration to be used for mapping users to roles. You can specify up to 25 rules per identity provider. Rules are evaluated in order. The first one to match specifies the role.
	// +kubebuilder:validation:Optional
	MappingRule []MappingRuleParameters `json:"mappingRule,omitempty" tf:"mapping_rule,omitempty"`

	// The role mapping type.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

// PoolRolesAttachmentSpec defines the desired state of PoolRolesAttachment
type PoolRolesAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PoolRolesAttachmentParameters `json:"forProvider"`
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
	InitProvider PoolRolesAttachmentInitParameters `json:"initProvider,omitempty"`
}

// PoolRolesAttachmentStatus defines the observed state of PoolRolesAttachment.
type PoolRolesAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PoolRolesAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PoolRolesAttachment is the Schema for the PoolRolesAttachments API. Provides an AWS Cognito Identity Pool Roles Attachment.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type PoolRolesAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.roles) || (has(self.initProvider) && has(self.initProvider.roles))",message="spec.forProvider.roles is a required parameter"
	Spec   PoolRolesAttachmentSpec   `json:"spec"`
	Status PoolRolesAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PoolRolesAttachmentList contains a list of PoolRolesAttachments
type PoolRolesAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PoolRolesAttachment `json:"items"`
}

// Repository type metadata.
var (
	PoolRolesAttachment_Kind             = "PoolRolesAttachment"
	PoolRolesAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PoolRolesAttachment_Kind}.String()
	PoolRolesAttachment_KindAPIVersion   = PoolRolesAttachment_Kind + "." + CRDGroupVersion.String()
	PoolRolesAttachment_GroupVersionKind = CRDGroupVersion.WithKind(PoolRolesAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&PoolRolesAttachment{}, &PoolRolesAttachmentList{})
}
