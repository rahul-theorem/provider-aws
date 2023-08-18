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

type AccessPointInitParameters struct {

	// Operating system user and group applied to all file system requests made using the access point. Detailed below.
	PosixUser []PosixUserInitParameters `json:"posixUser,omitempty" tf:"posix_user,omitempty"`

	// Directory on the Amazon EFS file system that the access point provides access to. Detailed below.
	RootDirectory []RootDirectoryInitParameters `json:"rootDirectory,omitempty" tf:"root_directory,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AccessPointObservation struct {

	// ARN of the access point.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// ARN of the file system.
	FileSystemArn *string `json:"fileSystemArn,omitempty" tf:"file_system_arn,omitempty"`

	// ID of the file system for which the access point is intended.
	FileSystemID *string `json:"fileSystemId,omitempty" tf:"file_system_id,omitempty"`

	// ID of the access point.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the access point.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// Operating system user and group applied to all file system requests made using the access point. Detailed below.
	PosixUser []PosixUserObservation `json:"posixUser,omitempty" tf:"posix_user,omitempty"`

	// Directory on the Amazon EFS file system that the access point provides access to. Detailed below.
	RootDirectory []RootDirectoryObservation `json:"rootDirectory,omitempty" tf:"root_directory,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type AccessPointParameters struct {

	// ID of the file system for which the access point is intended.
	// +crossplane:generate:reference:type=FileSystem
	// +kubebuilder:validation:Optional
	FileSystemID *string `json:"fileSystemId,omitempty" tf:"file_system_id,omitempty"`

	// Reference to a FileSystem to populate fileSystemId.
	// +kubebuilder:validation:Optional
	FileSystemIDRef *v1.Reference `json:"fileSystemIdRef,omitempty" tf:"-"`

	// Selector for a FileSystem to populate fileSystemId.
	// +kubebuilder:validation:Optional
	FileSystemIDSelector *v1.Selector `json:"fileSystemIdSelector,omitempty" tf:"-"`

	// Operating system user and group applied to all file system requests made using the access point. Detailed below.
	// +kubebuilder:validation:Optional
	PosixUser []PosixUserParameters `json:"posixUser,omitempty" tf:"posix_user,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Directory on the Amazon EFS file system that the access point provides access to. Detailed below.
	// +kubebuilder:validation:Optional
	RootDirectory []RootDirectoryParameters `json:"rootDirectory,omitempty" tf:"root_directory,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CreationInfoInitParameters struct {

	// POSIX group ID to apply to the root_directory.
	OwnerGID *float64 `json:"ownerGid,omitempty" tf:"owner_gid,omitempty"`

	// POSIX user ID to apply to the root_directory.
	OwnerUID *float64 `json:"ownerUid,omitempty" tf:"owner_uid,omitempty"`

	// POSIX permissions to apply to the RootDirectory, in the format of an octal number representing the file's mode bits.
	Permissions *string `json:"permissions,omitempty" tf:"permissions,omitempty"`
}

type CreationInfoObservation struct {

	// POSIX group ID to apply to the root_directory.
	OwnerGID *float64 `json:"ownerGid,omitempty" tf:"owner_gid,omitempty"`

	// POSIX user ID to apply to the root_directory.
	OwnerUID *float64 `json:"ownerUid,omitempty" tf:"owner_uid,omitempty"`

	// POSIX permissions to apply to the RootDirectory, in the format of an octal number representing the file's mode bits.
	Permissions *string `json:"permissions,omitempty" tf:"permissions,omitempty"`
}

type CreationInfoParameters struct {

	// POSIX group ID to apply to the root_directory.
	// +kubebuilder:validation:Optional
	OwnerGID *float64 `json:"ownerGid" tf:"owner_gid,omitempty"`

	// POSIX user ID to apply to the root_directory.
	// +kubebuilder:validation:Optional
	OwnerUID *float64 `json:"ownerUid" tf:"owner_uid,omitempty"`

	// POSIX permissions to apply to the RootDirectory, in the format of an octal number representing the file's mode bits.
	// +kubebuilder:validation:Optional
	Permissions *string `json:"permissions" tf:"permissions,omitempty"`
}

type PosixUserInitParameters struct {

	// POSIX group ID used for all file system operations using this access point.
	GID *float64 `json:"gid,omitempty" tf:"gid,omitempty"`

	// Secondary POSIX group IDs used for all file system operations using this access point.
	SecondaryGids []*float64 `json:"secondaryGids,omitempty" tf:"secondary_gids,omitempty"`

	// POSIX user ID used for all file system operations using this access point.
	UID *float64 `json:"uid,omitempty" tf:"uid,omitempty"`
}

type PosixUserObservation struct {

	// POSIX group ID used for all file system operations using this access point.
	GID *float64 `json:"gid,omitempty" tf:"gid,omitempty"`

	// Secondary POSIX group IDs used for all file system operations using this access point.
	SecondaryGids []*float64 `json:"secondaryGids,omitempty" tf:"secondary_gids,omitempty"`

	// POSIX user ID used for all file system operations using this access point.
	UID *float64 `json:"uid,omitempty" tf:"uid,omitempty"`
}

type PosixUserParameters struct {

	// POSIX group ID used for all file system operations using this access point.
	// +kubebuilder:validation:Optional
	GID *float64 `json:"gid" tf:"gid,omitempty"`

	// Secondary POSIX group IDs used for all file system operations using this access point.
	// +kubebuilder:validation:Optional
	SecondaryGids []*float64 `json:"secondaryGids,omitempty" tf:"secondary_gids,omitempty"`

	// POSIX user ID used for all file system operations using this access point.
	// +kubebuilder:validation:Optional
	UID *float64 `json:"uid" tf:"uid,omitempty"`
}

type RootDirectoryInitParameters struct {

	// POSIX IDs and permissions to apply to the access point's Root Directory. See Creation Info below.
	CreationInfo []CreationInfoInitParameters `json:"creationInfo,omitempty" tf:"creation_info,omitempty"`

	// Path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system. A path can have up to four subdirectories. If the specified path does not exist, you are required to provide creation_info.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type RootDirectoryObservation struct {

	// POSIX IDs and permissions to apply to the access point's Root Directory. See Creation Info below.
	CreationInfo []CreationInfoObservation `json:"creationInfo,omitempty" tf:"creation_info,omitempty"`

	// Path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system. A path can have up to four subdirectories. If the specified path does not exist, you are required to provide creation_info.
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type RootDirectoryParameters struct {

	// POSIX IDs and permissions to apply to the access point's Root Directory. See Creation Info below.
	// +kubebuilder:validation:Optional
	CreationInfo []CreationInfoParameters `json:"creationInfo,omitempty" tf:"creation_info,omitempty"`

	// Path on the EFS file system to expose as the root directory to NFS clients using the access point to access the EFS file system. A path can have up to four subdirectories. If the specified path does not exist, you are required to provide creation_info.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

// AccessPointSpec defines the desired state of AccessPoint
type AccessPointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AccessPointParameters `json:"forProvider"`
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
	InitProvider AccessPointInitParameters `json:"initProvider,omitempty"`
}

// AccessPointStatus defines the observed state of AccessPoint.
type AccessPointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AccessPointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AccessPoint is the Schema for the AccessPoints API. Provides an Elastic File System (EFS) access point.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AccessPoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AccessPointSpec   `json:"spec"`
	Status            AccessPointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AccessPointList contains a list of AccessPoints
type AccessPointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessPoint `json:"items"`
}

// Repository type metadata.
var (
	AccessPoint_Kind             = "AccessPoint"
	AccessPoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AccessPoint_Kind}.String()
	AccessPoint_KindAPIVersion   = AccessPoint_Kind + "." + CRDGroupVersion.String()
	AccessPoint_GroupVersionKind = CRDGroupVersion.WithKind(AccessPoint_Kind)
)

func init() {
	SchemeBuilder.Register(&AccessPoint{}, &AccessPointList{})
}
