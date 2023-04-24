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

type CatalogDatabaseObservation struct {

	// ARN of the Glue Catalog Database.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// ID of the Glue Catalog to create the database in. If omitted, this defaults to the AWS Account ID.
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// Creates a set of default permissions on the table for principals. See create_table_default_permission below.
	CreateTableDefaultPermission []CreateTableDefaultPermissionObservation `json:"createTableDefaultPermission,omitempty" tf:"create_table_default_permission,omitempty"`

	// Description of the database.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Catalog ID and name of the database
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Location of the database (for example, an HDFS path).
	LocationURI *string `json:"locationUri,omitempty" tf:"location_uri,omitempty"`

	// List of key-value pairs that define parameters and properties of the database.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Configuration block for a target database for resource linking. See target_database below.
	TargetDatabase []TargetDatabaseObservation `json:"targetDatabase,omitempty" tf:"target_database,omitempty"`
}

type CatalogDatabaseParameters struct {

	// ID of the Glue Catalog to create the database in. If omitted, this defaults to the AWS Account ID.
	// +kubebuilder:validation:Required
	CatalogID *string `json:"catalogId" tf:"catalog_id,omitempty"`

	// Creates a set of default permissions on the table for principals. See create_table_default_permission below.
	// +kubebuilder:validation:Optional
	CreateTableDefaultPermission []CreateTableDefaultPermissionParameters `json:"createTableDefaultPermission,omitempty" tf:"create_table_default_permission,omitempty"`

	// Description of the database.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Location of the database (for example, an HDFS path).
	// +kubebuilder:validation:Optional
	LocationURI *string `json:"locationUri,omitempty" tf:"location_uri,omitempty"`

	// List of key-value pairs that define parameters and properties of the database.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Configuration block for a target database for resource linking. See target_database below.
	// +kubebuilder:validation:Optional
	TargetDatabase []TargetDatabaseParameters `json:"targetDatabase,omitempty" tf:"target_database,omitempty"`
}

type CreateTableDefaultPermissionObservation struct {

	// The permissions that are granted to the principal.
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// The principal who is granted permissions.. See principal below.
	Principal []PrincipalObservation `json:"principal,omitempty" tf:"principal,omitempty"`
}

type CreateTableDefaultPermissionParameters struct {

	// The permissions that are granted to the principal.
	// +kubebuilder:validation:Optional
	Permissions []*string `json:"permissions,omitempty" tf:"permissions,omitempty"`

	// The principal who is granted permissions.. See principal below.
	// +kubebuilder:validation:Optional
	Principal []PrincipalParameters `json:"principal,omitempty" tf:"principal,omitempty"`
}

type PrincipalObservation struct {

	// An identifier for the Lake Formation principal.
	DataLakePrincipalIdentifier *string `json:"dataLakePrincipalIdentifier,omitempty" tf:"data_lake_principal_identifier,omitempty"`
}

type PrincipalParameters struct {

	// An identifier for the Lake Formation principal.
	// +kubebuilder:validation:Optional
	DataLakePrincipalIdentifier *string `json:"dataLakePrincipalIdentifier,omitempty" tf:"data_lake_principal_identifier,omitempty"`
}

type TargetDatabaseObservation struct {

	// ID of the Data Catalog in which the database resides.
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id,omitempty"`

	// Name of the catalog database.
	DatabaseName *string `json:"databaseName,omitempty" tf:"database_name,omitempty"`
}

type TargetDatabaseParameters struct {

	// ID of the Data Catalog in which the database resides.
	// +kubebuilder:validation:Required
	CatalogID *string `json:"catalogId" tf:"catalog_id,omitempty"`

	// Name of the catalog database.
	// +kubebuilder:validation:Required
	DatabaseName *string `json:"databaseName" tf:"database_name,omitempty"`
}

// CatalogDatabaseSpec defines the desired state of CatalogDatabase
type CatalogDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CatalogDatabaseParameters `json:"forProvider"`
}

// CatalogDatabaseStatus defines the observed state of CatalogDatabase.
type CatalogDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CatalogDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CatalogDatabase is the Schema for the CatalogDatabases API. Provides a Glue Catalog Database.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type CatalogDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CatalogDatabaseSpec   `json:"spec"`
	Status            CatalogDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CatalogDatabaseList contains a list of CatalogDatabases
type CatalogDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CatalogDatabase `json:"items"`
}

// Repository type metadata.
var (
	CatalogDatabase_Kind             = "CatalogDatabase"
	CatalogDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CatalogDatabase_Kind}.String()
	CatalogDatabase_KindAPIVersion   = CatalogDatabase_Kind + "." + CRDGroupVersion.String()
	CatalogDatabase_GroupVersionKind = CRDGroupVersion.WithKind(CatalogDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&CatalogDatabase{}, &CatalogDatabaseList{})
}
