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

type AliasObservation struct {

	// Alias ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Description of the alias.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Alias ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Name of the alias.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Specifies the fleet and/or routing type to use for the alias.
	RoutingStrategy []RoutingStrategyObservation `json:"routingStrategy,omitempty" tf:"routing_strategy,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type AliasParameters struct {

	// Description of the alias.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of the alias.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Specifies the fleet and/or routing type to use for the alias.
	// +kubebuilder:validation:Optional
	RoutingStrategy []RoutingStrategyParameters `json:"routingStrategy,omitempty" tf:"routing_strategy,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RoutingStrategyObservation struct {

	// ID of the GameLift Fleet to point the alias to.
	FleetID *string `json:"fleetId,omitempty" tf:"fleet_id,omitempty"`

	// Message text to be used with the TERMINAL routing strategy.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Type of routing strategyE.g., SIMPLE or TERMINAL
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RoutingStrategyParameters struct {

	// ID of the GameLift Fleet to point the alias to.
	// +kubebuilder:validation:Optional
	FleetID *string `json:"fleetId,omitempty" tf:"fleet_id,omitempty"`

	// Message text to be used with the TERMINAL routing strategy.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Type of routing strategyE.g., SIMPLE or TERMINAL
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

// AliasSpec defines the desired state of Alias
type AliasSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AliasParameters `json:"forProvider"`
}

// AliasStatus defines the observed state of Alias.
type AliasStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AliasObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Alias is the Schema for the Aliass API. Provides a GameLift Alias resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Alias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.routingStrategy)",message="routingStrategy is a required parameter"
	Spec   AliasSpec   `json:"spec"`
	Status AliasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AliasList contains a list of Aliass
type AliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Alias `json:"items"`
}

// Repository type metadata.
var (
	Alias_Kind             = "Alias"
	Alias_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Alias_Kind}.String()
	Alias_KindAPIVersion   = Alias_Kind + "." + CRDGroupVersion.String()
	Alias_GroupVersionKind = CRDGroupVersion.WithKind(Alias_Kind)
)

func init() {
	SchemeBuilder.Register(&Alias{}, &AliasList{})
}
