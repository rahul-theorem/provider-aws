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

type AMIInitParameters struct {

	// Machine architecture for created instances. Defaults to "x86_64".
	Architecture *string `json:"architecture,omitempty" tf:"architecture,omitempty"`

	// Boot mode of the AMI. For more information, see Boot modes in the Amazon Elastic Compute Cloud User Guide.
	BootMode *string `json:"bootMode,omitempty" tf:"boot_mode,omitempty"`

	// Date and time to deprecate the AMI. If you specified a value for seconds, Amazon EC2 rounds the seconds to the nearest minute. Valid values: RFC3339 time string (YYYY-MM-DDTHH:MM:SSZ)
	DeprecationTime *string `json:"deprecationTime,omitempty" tf:"deprecation_time,omitempty"`

	// Longer, human-readable description for the AMI.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EBSBlockDevice []EBSBlockDeviceInitParameters `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`

	// Whether enhanced networking with ENA is enabled. Defaults to false.
	EnaSupport *bool `json:"enaSupport,omitempty" tf:"ena_support,omitempty"`

	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevice []EphemeralBlockDeviceInitParameters `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`

	// Path to an S3 object containing an image manifest, e.g., created
	// by the ec2-upload-bundle command in the EC2 command line tools.
	ImageLocation *string `json:"imageLocation,omitempty" tf:"image_location,omitempty"`

	// If EC2 instances started from this image should require the use of the Instance Metadata Service V2 (IMDSv2), set this argument to v2.0. For more information, see Configure instance metadata options for new instances.
	ImdsSupport *string `json:"imdsSupport,omitempty" tf:"imds_support,omitempty"`

	// ID of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelID *string `json:"kernelId,omitempty" tf:"kernel_id,omitempty"`

	// Region-unique name for the AMI.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of an initrd image (ARI) that will be used when booting the
	// created instances.
	RamdiskID *string `json:"ramdiskId,omitempty" tf:"ramdisk_id,omitempty"`

	// Name of the root device (for example, /dev/sda1, or /dev/xvda).
	RootDeviceName *string `json:"rootDeviceName,omitempty" tf:"root_device_name,omitempty"`

	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	SriovNetSupport *string `json:"sriovNetSupport,omitempty" tf:"sriov_net_support,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// If the image is configured for NitroTPM support, the value is v2.0. For more information, see NitroTPM in the Amazon Elastic Compute Cloud User Guide.
	TpmSupport *string `json:"tpmSupport,omitempty" tf:"tpm_support,omitempty"`

	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType *string `json:"virtualizationType,omitempty" tf:"virtualization_type,omitempty"`
}

type AMIObservation struct {

	// Machine architecture for created instances. Defaults to "x86_64".
	Architecture *string `json:"architecture,omitempty" tf:"architecture,omitempty"`

	// ARN of the AMI.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Boot mode of the AMI. For more information, see Boot modes in the Amazon Elastic Compute Cloud User Guide.
	BootMode *string `json:"bootMode,omitempty" tf:"boot_mode,omitempty"`

	// Date and time to deprecate the AMI. If you specified a value for seconds, Amazon EC2 rounds the seconds to the nearest minute. Valid values: RFC3339 time string (YYYY-MM-DDTHH:MM:SSZ)
	DeprecationTime *string `json:"deprecationTime,omitempty" tf:"deprecation_time,omitempty"`

	// Longer, human-readable description for the AMI.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	EBSBlockDevice []EBSBlockDeviceObservation `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`

	// Whether enhanced networking with ENA is enabled. Defaults to false.
	EnaSupport *bool `json:"enaSupport,omitempty" tf:"ena_support,omitempty"`

	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	EphemeralBlockDevice []EphemeralBlockDeviceObservation `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`

	// Hypervisor type of the image.
	Hypervisor *string `json:"hypervisor,omitempty" tf:"hypervisor,omitempty"`

	// ID of the created AMI.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Path to an S3 object containing an image manifest, e.g., created
	// by the ec2-upload-bundle command in the EC2 command line tools.
	ImageLocation *string `json:"imageLocation,omitempty" tf:"image_location,omitempty"`

	// AWS account alias (for example, amazon, self) or the AWS account ID of the AMI owner.
	ImageOwnerAlias *string `json:"imageOwnerAlias,omitempty" tf:"image_owner_alias,omitempty"`

	// Type of image.
	ImageType *string `json:"imageType,omitempty" tf:"image_type,omitempty"`

	// If EC2 instances started from this image should require the use of the Instance Metadata Service V2 (IMDSv2), set this argument to v2.0. For more information, see Configure instance metadata options for new instances.
	ImdsSupport *string `json:"imdsSupport,omitempty" tf:"imds_support,omitempty"`

	// ID of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	KernelID *string `json:"kernelId,omitempty" tf:"kernel_id,omitempty"`

	ManageEBSSnapshots *bool `json:"manageEbsSnapshots,omitempty" tf:"manage_ebs_snapshots,omitempty"`

	// Region-unique name for the AMI.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// AWS account ID of the image owner.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// This value is set to windows for Windows AMIs; otherwise, it is blank.
	Platform *string `json:"platform,omitempty" tf:"platform,omitempty"`

	// Platform details associated with the billing code of the AMI.
	PlatformDetails *string `json:"platformDetails,omitempty" tf:"platform_details,omitempty"`

	// Whether the image has public launch permissions.
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	// ID of an initrd image (ARI) that will be used when booting the
	// created instances.
	RamdiskID *string `json:"ramdiskId,omitempty" tf:"ramdisk_id,omitempty"`

	// Name of the root device (for example, /dev/sda1, or /dev/xvda).
	RootDeviceName *string `json:"rootDeviceName,omitempty" tf:"root_device_name,omitempty"`

	// Snapshot ID for the root volume (for EBS-backed AMIs)
	RootSnapshotID *string `json:"rootSnapshotId,omitempty" tf:"root_snapshot_id,omitempty"`

	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	SriovNetSupport *string `json:"sriovNetSupport,omitempty" tf:"sriov_net_support,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// If the image is configured for NitroTPM support, the value is v2.0. For more information, see NitroTPM in the Amazon Elastic Compute Cloud User Guide.
	TpmSupport *string `json:"tpmSupport,omitempty" tf:"tpm_support,omitempty"`

	// Operation of the Amazon EC2 instance and the billing code that is associated with the AMI.
	UsageOperation *string `json:"usageOperation,omitempty" tf:"usage_operation,omitempty"`

	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	VirtualizationType *string `json:"virtualizationType,omitempty" tf:"virtualization_type,omitempty"`
}

type AMIParameters struct {

	// Machine architecture for created instances. Defaults to "x86_64".
	// +kubebuilder:validation:Optional
	Architecture *string `json:"architecture,omitempty" tf:"architecture,omitempty"`

	// Boot mode of the AMI. For more information, see Boot modes in the Amazon Elastic Compute Cloud User Guide.
	// +kubebuilder:validation:Optional
	BootMode *string `json:"bootMode,omitempty" tf:"boot_mode,omitempty"`

	// Date and time to deprecate the AMI. If you specified a value for seconds, Amazon EC2 rounds the seconds to the nearest minute. Valid values: RFC3339 time string (YYYY-MM-DDTHH:MM:SSZ)
	// +kubebuilder:validation:Optional
	DeprecationTime *string `json:"deprecationTime,omitempty" tf:"deprecation_time,omitempty"`

	// Longer, human-readable description for the AMI.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Nested block describing an EBS block device that should be
	// attached to created instances. The structure of this block is described below.
	// +kubebuilder:validation:Optional
	EBSBlockDevice []EBSBlockDeviceParameters `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`

	// Whether enhanced networking with ENA is enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	EnaSupport *bool `json:"enaSupport,omitempty" tf:"ena_support,omitempty"`

	// Nested block describing an ephemeral block device that
	// should be attached to created instances. The structure of this block is described below.
	// +kubebuilder:validation:Optional
	EphemeralBlockDevice []EphemeralBlockDeviceParameters `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`

	// Path to an S3 object containing an image manifest, e.g., created
	// by the ec2-upload-bundle command in the EC2 command line tools.
	// +kubebuilder:validation:Optional
	ImageLocation *string `json:"imageLocation,omitempty" tf:"image_location,omitempty"`

	// If EC2 instances started from this image should require the use of the Instance Metadata Service V2 (IMDSv2), set this argument to v2.0. For more information, see Configure instance metadata options for new instances.
	// +kubebuilder:validation:Optional
	ImdsSupport *string `json:"imdsSupport,omitempty" tf:"imds_support,omitempty"`

	// ID of the kernel image (AKI) that will be used as the paravirtual
	// kernel in created instances.
	// +kubebuilder:validation:Optional
	KernelID *string `json:"kernelId,omitempty" tf:"kernel_id,omitempty"`

	// Region-unique name for the AMI.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of an initrd image (ARI) that will be used when booting the
	// created instances.
	// +kubebuilder:validation:Optional
	RamdiskID *string `json:"ramdiskId,omitempty" tf:"ramdisk_id,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Name of the root device (for example, /dev/sda1, or /dev/xvda).
	// +kubebuilder:validation:Optional
	RootDeviceName *string `json:"rootDeviceName,omitempty" tf:"root_device_name,omitempty"`

	// When set to "simple" (the default), enables enhanced networking
	// for created instances. No other value is supported at this time.
	// +kubebuilder:validation:Optional
	SriovNetSupport *string `json:"sriovNetSupport,omitempty" tf:"sriov_net_support,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// If the image is configured for NitroTPM support, the value is v2.0. For more information, see NitroTPM in the Amazon Elastic Compute Cloud User Guide.
	// +kubebuilder:validation:Optional
	TpmSupport *string `json:"tpmSupport,omitempty" tf:"tpm_support,omitempty"`

	// Keyword to choose what virtualization mode created instances
	// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
	// changes the set of further arguments that are required, as described below.
	// +kubebuilder:validation:Optional
	VirtualizationType *string `json:"virtualizationType,omitempty" tf:"virtualization_type,omitempty"`
}

type EBSBlockDeviceInitParameters struct {

	// Boolean controlling whether the EBS volumes created to
	// support each created instance will be deleted once that instance is terminated.
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// Path at which the device is exposed to created instances.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Boolean controlling whether the created EBS volumes will be encrypted. Can't be used with snapshot_id.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// Number of I/O operations per second the
	// created volumes will support.
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// ARN of the Outpost on which the snapshot is stored.
	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn,omitempty"`

	// Throughput that the EBS volume supports, in MiB/s. Only valid for volume_type of gp3.
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// Size of created volumes in GiB.
	// If snapshot_id is set and volume_size is omitted then the volume will have the same size
	// as the selected snapshot.
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// Type of EBS volume to create. Can be standard, gp2, gp3, io1, io2, sc1 or st1 (Default: standard).
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type EBSBlockDeviceObservation struct {

	// Boolean controlling whether the EBS volumes created to
	// support each created instance will be deleted once that instance is terminated.
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// Path at which the device is exposed to created instances.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Boolean controlling whether the created EBS volumes will be encrypted. Can't be used with snapshot_id.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// Number of I/O operations per second the
	// created volumes will support.
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// ARN of the Outpost on which the snapshot is stored.
	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn,omitempty"`

	// ID of an EBS snapshot that will be used to initialize the created
	// EBS volumes. If set, the volume_size attribute must be at least as large as the referenced
	// snapshot.
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// Throughput that the EBS volume supports, in MiB/s. Only valid for volume_type of gp3.
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// Size of created volumes in GiB.
	// If snapshot_id is set and volume_size is omitted then the volume will have the same size
	// as the selected snapshot.
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// Type of EBS volume to create. Can be standard, gp2, gp3, io1, io2, sc1 or st1 (Default: standard).
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type EBSBlockDeviceParameters struct {

	// Boolean controlling whether the EBS volumes created to
	// support each created instance will be deleted once that instance is terminated.
	// +kubebuilder:validation:Optional
	DeleteOnTermination *bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// Path at which the device is exposed to created instances.
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName" tf:"device_name,omitempty"`

	// Boolean controlling whether the created EBS volumes will be encrypted. Can't be used with snapshot_id.
	// +kubebuilder:validation:Optional
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// Number of I/O operations per second the
	// created volumes will support.
	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// ARN of the Outpost on which the snapshot is stored.
	// +kubebuilder:validation:Optional
	OutpostArn *string `json:"outpostArn,omitempty" tf:"outpost_arn,omitempty"`

	// ID of an EBS snapshot that will be used to initialize the created
	// EBS volumes. If set, the volume_size attribute must be at least as large as the referenced
	// snapshot.
	// +crossplane:generate:reference:type=EBSSnapshot
	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// Reference to a EBSSnapshot to populate snapshotId.
	// +kubebuilder:validation:Optional
	SnapshotIDRef *v1.Reference `json:"snapshotIdRef,omitempty" tf:"-"`

	// Selector for a EBSSnapshot to populate snapshotId.
	// +kubebuilder:validation:Optional
	SnapshotIDSelector *v1.Selector `json:"snapshotIdSelector,omitempty" tf:"-"`

	// Throughput that the EBS volume supports, in MiB/s. Only valid for volume_type of gp3.
	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// Size of created volumes in GiB.
	// If snapshot_id is set and volume_size is omitted then the volume will have the same size
	// as the selected snapshot.
	// +kubebuilder:validation:Optional
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// Type of EBS volume to create. Can be standard, gp2, gp3, io1, io2, sc1 or st1 (Default: standard).
	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type EphemeralBlockDeviceInitParameters struct {

	// Path at which the device is exposed to created instances.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Name for the ephemeral device, of the form "ephemeralN" where
	// N is a volume number starting from zero.
	VirtualName *string `json:"virtualName,omitempty" tf:"virtual_name,omitempty"`
}

type EphemeralBlockDeviceObservation struct {

	// Path at which the device is exposed to created instances.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Name for the ephemeral device, of the form "ephemeralN" where
	// N is a volume number starting from zero.
	VirtualName *string `json:"virtualName,omitempty" tf:"virtual_name,omitempty"`
}

type EphemeralBlockDeviceParameters struct {

	// Path at which the device is exposed to created instances.
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName" tf:"device_name,omitempty"`

	// Name for the ephemeral device, of the form "ephemeralN" where
	// N is a volume number starting from zero.
	// +kubebuilder:validation:Optional
	VirtualName *string `json:"virtualName" tf:"virtual_name,omitempty"`
}

// AMISpec defines the desired state of AMI
type AMISpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AMIParameters `json:"forProvider"`
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
	InitProvider AMIInitParameters `json:"initProvider,omitempty"`
}

// AMIStatus defines the observed state of AMI.
type AMIStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AMIObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AMI is the Schema for the AMIs API. Creates and manages a custom Amazon Machine Image (AMI).
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type AMI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   AMISpec   `json:"spec"`
	Status AMIStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AMIList contains a list of AMIs
type AMIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AMI `json:"items"`
}

// Repository type metadata.
var (
	AMI_Kind             = "AMI"
	AMI_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AMI_Kind}.String()
	AMI_KindAPIVersion   = AMI_Kind + "." + CRDGroupVersion.String()
	AMI_GroupVersionKind = CRDGroupVersion.WithKind(AMI_Kind)
)

func init() {
	SchemeBuilder.Register(&AMI{}, &AMIList{})
}
