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

type RDSDBInstanceObservation struct {

	// A db username
	DBUser *string `json:"dbUser,omitempty" tf:"db_user,omitempty"`

	// The computed id. Please note that this is only used internally to identify the stack <-> instance relation. This value is not used in aws.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The db instance to register for this stack. Changing this will force a new resource.
	RDSDBInstanceArn *string `json:"rdsDbInstanceArn,omitempty" tf:"rds_db_instance_arn,omitempty"`

	// The stack to register a db instance for. Changing this will force a new resource.
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`
}

type RDSDBInstanceParameters struct {

	// A db password
	// +kubebuilder:validation:Optional
	DBPasswordSecretRef v1.SecretKeySelector `json:"dbPasswordSecretRef" tf:"-"`

	// A db username
	// +kubebuilder:validation:Optional
	DBUser *string `json:"dbUser,omitempty" tf:"db_user,omitempty"`

	// The db instance to register for this stack. Changing this will force a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/rds/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	RDSDBInstanceArn *string `json:"rdsDbInstanceArn,omitempty" tf:"rds_db_instance_arn,omitempty"`

	// Reference to a Instance in rds to populate rdsDbInstanceArn.
	// +kubebuilder:validation:Optional
	RDSDBInstanceArnRef *v1.Reference `json:"rdsDbInstanceArnRef,omitempty" tf:"-"`

	// Selector for a Instance in rds to populate rdsDbInstanceArn.
	// +kubebuilder:validation:Optional
	RDSDBInstanceArnSelector *v1.Selector `json:"rdsDbInstanceArnSelector,omitempty" tf:"-"`

	// The stack to register a db instance for. Changing this will force a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/opsworks/v1beta1.Stack
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// Reference to a Stack in opsworks to populate stackId.
	// +kubebuilder:validation:Optional
	StackIDRef *v1.Reference `json:"stackIdRef,omitempty" tf:"-"`

	// Selector for a Stack in opsworks to populate stackId.
	// +kubebuilder:validation:Optional
	StackIDSelector *v1.Selector `json:"stackIdSelector,omitempty" tf:"-"`
}

// RDSDBInstanceSpec defines the desired state of RDSDBInstance
type RDSDBInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RDSDBInstanceParameters `json:"forProvider"`
}

// RDSDBInstanceStatus defines the observed state of RDSDBInstance.
type RDSDBInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RDSDBInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RDSDBInstance is the Schema for the RDSDBInstances API. Provides an OpsWorks RDS DB Instance resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type RDSDBInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.dbPasswordSecretRef)",message="dbPasswordSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.dbUser)",message="dbUser is a required parameter"
	Spec   RDSDBInstanceSpec   `json:"spec"`
	Status RDSDBInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RDSDBInstanceList contains a list of RDSDBInstances
type RDSDBInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RDSDBInstance `json:"items"`
}

// Repository type metadata.
var (
	RDSDBInstance_Kind             = "RDSDBInstance"
	RDSDBInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RDSDBInstance_Kind}.String()
	RDSDBInstance_KindAPIVersion   = RDSDBInstance_Kind + "." + CRDGroupVersion.String()
	RDSDBInstance_GroupVersionKind = CRDGroupVersion.WithKind(RDSDBInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&RDSDBInstance{}, &RDSDBInstanceList{})
}
