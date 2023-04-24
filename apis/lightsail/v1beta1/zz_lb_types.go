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

type LBObservation struct {

	// The ARN of the Lightsail load balancer.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The timestamp when the load balancer was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The DNS name of the load balancer.
	DNSName *string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`

	// The health check path of the load balancer. Default value "/".
	HealthCheckPath *string `json:"healthCheckPath,omitempty" tf:"health_check_path,omitempty"`

	// The name used for this load balancer (matches name).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`

	// The instance port the load balancer will connect.
	InstancePort *float64 `json:"instancePort,omitempty" tf:"instance_port,omitempty"`

	// The protocol of the load balancer.
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`

	// The public ports of the load balancer.
	PublicPorts []*float64 `json:"publicPorts,omitempty" tf:"public_ports,omitempty"`

	// The support code for the database. Include this code in your email to support when you have questions about a database in Lightsail. This code enables our support team to look up your Lightsail information more easily.
	SupportCode *string `json:"supportCode,omitempty" tf:"support_code,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type LBParameters struct {

	// The health check path of the load balancer. Default value "/".
	// +kubebuilder:validation:Optional
	HealthCheckPath *string `json:"healthCheckPath,omitempty" tf:"health_check_path,omitempty"`

	// +kubebuilder:validation:Optional
	IPAddressType *string `json:"ipAddressType,omitempty" tf:"ip_address_type,omitempty"`

	// The instance port the load balancer will connect.
	// +kubebuilder:validation:Optional
	InstancePort *float64 `json:"instancePort,omitempty" tf:"instance_port,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// LBSpec defines the desired state of LB
type LBSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBParameters `json:"forProvider"`
}

// LBStatus defines the observed state of LB.
type LBStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LB is the Schema for the LBs API. Provides a Lightsail Load Balancer
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.instancePort)",message="instancePort is a required parameter"
	Spec   LBSpec   `json:"spec"`
	Status LBStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBList contains a list of LBs
type LBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LB `json:"items"`
}

// Repository type metadata.
var (
	LB_Kind             = "LB"
	LB_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LB_Kind}.String()
	LB_KindAPIVersion   = LB_Kind + "." + CRDGroupVersion.String()
	LB_GroupVersionKind = CRDGroupVersion.WithKind(LB_Kind)
)

func init() {
	SchemeBuilder.Register(&LB{}, &LBList{})
}
