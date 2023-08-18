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

type PredicatesInitParameters struct {

	// Set this to false if you want to allow, block, or count requests
	// based on the settings in the specified ByteMatchSet, IPSet, SqlInjectionMatchSet, XssMatchSet, or SizeConstraintSet.
	// For example, if an IPSet includes the IP address 192.0.2.44, AWS WAF will allow or block requests based on that IP address.
	// If set to true, AWS WAF will allow, block, or count requests based on all IP addresses except 192.0.2.44.
	Negated *bool `json:"negated,omitempty" tf:"negated,omitempty"`

	// The type of predicate in a rule. Valid values: ByteMatch, GeoMatch, IPMatch, RegexMatch, SizeConstraint, SqlInjectionMatch, or XssMatch.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PredicatesObservation struct {

	// A unique identifier for a predicate in the rule, such as Byte Match Set ID or IPSet ID.
	DataID *string `json:"dataId,omitempty" tf:"data_id,omitempty"`

	// Set this to false if you want to allow, block, or count requests
	// based on the settings in the specified ByteMatchSet, IPSet, SqlInjectionMatchSet, XssMatchSet, or SizeConstraintSet.
	// For example, if an IPSet includes the IP address 192.0.2.44, AWS WAF will allow or block requests based on that IP address.
	// If set to true, AWS WAF will allow, block, or count requests based on all IP addresses except 192.0.2.44.
	Negated *bool `json:"negated,omitempty" tf:"negated,omitempty"`

	// The type of predicate in a rule. Valid values: ByteMatch, GeoMatch, IPMatch, RegexMatch, SizeConstraint, SqlInjectionMatch, or XssMatch.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type PredicatesParameters struct {

	// A unique identifier for a predicate in the rule, such as Byte Match Set ID or IPSet ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/waf/v1beta1.IPSet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataID *string `json:"dataId,omitempty" tf:"data_id,omitempty"`

	// Reference to a IPSet in waf to populate dataId.
	// +kubebuilder:validation:Optional
	DataIDRef *v1.Reference `json:"dataIdRef,omitempty" tf:"-"`

	// Selector for a IPSet in waf to populate dataId.
	// +kubebuilder:validation:Optional
	DataIDSelector *v1.Selector `json:"dataIdSelector,omitempty" tf:"-"`

	// Set this to false if you want to allow, block, or count requests
	// based on the settings in the specified ByteMatchSet, IPSet, SqlInjectionMatchSet, XssMatchSet, or SizeConstraintSet.
	// For example, if an IPSet includes the IP address 192.0.2.44, AWS WAF will allow or block requests based on that IP address.
	// If set to true, AWS WAF will allow, block, or count requests based on all IP addresses except 192.0.2.44.
	// +kubebuilder:validation:Optional
	Negated *bool `json:"negated" tf:"negated,omitempty"`

	// The type of predicate in a rule. Valid values: ByteMatch, GeoMatch, IPMatch, RegexMatch, SizeConstraint, SqlInjectionMatch, or XssMatch.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type RateBasedRuleInitParameters struct {

	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// The name or description of the rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The objects to include in a rule (documented below).
	Predicates []PredicatesInitParameters `json:"predicates,omitempty" tf:"predicates,omitempty"`

	// Valid value is IP.
	RateKey *string `json:"rateKey,omitempty" tf:"rate_key,omitempty"`

	// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
	RateLimit *float64 `json:"rateLimit,omitempty" tf:"rate_limit,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RateBasedRuleObservation struct {

	// Amazon Resource Name (ARN)
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the WAF rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name or description for the Amazon CloudWatch metric of this rule.
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// The name or description of the rule.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The objects to include in a rule (documented below).
	Predicates []PredicatesObservation `json:"predicates,omitempty" tf:"predicates,omitempty"`

	// Valid value is IP.
	RateKey *string `json:"rateKey,omitempty" tf:"rate_key,omitempty"`

	// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
	RateLimit *float64 `json:"rateLimit,omitempty" tf:"rate_limit,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type RateBasedRuleParameters struct {

	// The name or description for the Amazon CloudWatch metric of this rule.
	// +kubebuilder:validation:Optional
	MetricName *string `json:"metricName,omitempty" tf:"metric_name,omitempty"`

	// The name or description of the rule.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The objects to include in a rule (documented below).
	// +kubebuilder:validation:Optional
	Predicates []PredicatesParameters `json:"predicates,omitempty" tf:"predicates,omitempty"`

	// Valid value is IP.
	// +kubebuilder:validation:Optional
	RateKey *string `json:"rateKey,omitempty" tf:"rate_key,omitempty"`

	// The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period. Minimum value is 100.
	// +kubebuilder:validation:Optional
	RateLimit *float64 `json:"rateLimit,omitempty" tf:"rate_limit,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// RateBasedRuleSpec defines the desired state of RateBasedRule
type RateBasedRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RateBasedRuleParameters `json:"forProvider"`
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
	InitProvider RateBasedRuleInitParameters `json:"initProvider,omitempty"`
}

// RateBasedRuleStatus defines the observed state of RateBasedRule.
type RateBasedRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RateBasedRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RateBasedRule is the Schema for the RateBasedRules API. Provides a AWS WAF rule resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RateBasedRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.metricName) || has(self.initProvider.metricName)",message="metricName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rateKey) || has(self.initProvider.rateKey)",message="rateKey is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rateLimit) || has(self.initProvider.rateLimit)",message="rateLimit is a required parameter"
	Spec   RateBasedRuleSpec   `json:"spec"`
	Status RateBasedRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RateBasedRuleList contains a list of RateBasedRules
type RateBasedRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RateBasedRule `json:"items"`
}

// Repository type metadata.
var (
	RateBasedRule_Kind             = "RateBasedRule"
	RateBasedRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RateBasedRule_Kind}.String()
	RateBasedRule_KindAPIVersion   = RateBasedRule_Kind + "." + CRDGroupVersion.String()
	RateBasedRule_GroupVersionKind = CRDGroupVersion.WithKind(RateBasedRule_Kind)
)

func init() {
	SchemeBuilder.Register(&RateBasedRule{}, &RateBasedRuleList{})
}
