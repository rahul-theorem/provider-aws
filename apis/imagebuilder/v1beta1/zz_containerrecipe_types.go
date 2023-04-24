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

type BlockDeviceMappingObservation struct {

	// Name of the device. For example, /dev/sda or /dev/xvdb.
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Configuration block with Elastic Block Storage (EBS) block device mapping settings. Detailed below.
	EBS []EBSObservation `json:"ebs,omitempty" tf:"ebs,omitempty"`

	// Set to true to remove a mapping from the parent image.
	NoDevice *bool `json:"noDevice,omitempty" tf:"no_device,omitempty"`

	// Virtual device name. For example, ephemeral0. Instance store volumes are numbered starting from 0.
	VirtualName *string `json:"virtualName,omitempty" tf:"virtual_name,omitempty"`
}

type BlockDeviceMappingParameters struct {

	// Name of the device. For example, /dev/sda or /dev/xvdb.
	// +kubebuilder:validation:Optional
	DeviceName *string `json:"deviceName,omitempty" tf:"device_name,omitempty"`

	// Configuration block with Elastic Block Storage (EBS) block device mapping settings. Detailed below.
	// +kubebuilder:validation:Optional
	EBS []EBSParameters `json:"ebs,omitempty" tf:"ebs,omitempty"`

	// Set to true to remove a mapping from the parent image.
	// +kubebuilder:validation:Optional
	NoDevice *bool `json:"noDevice,omitempty" tf:"no_device,omitempty"`

	// Virtual device name. For example, ephemeral0. Instance store volumes are numbered starting from 0.
	// +kubebuilder:validation:Optional
	VirtualName *string `json:"virtualName,omitempty" tf:"virtual_name,omitempty"`
}

type ContainerRecipeComponentObservation struct {

	// Amazon Resource Name (ARN) of the Image Builder Component to associate.
	ComponentArn *string `json:"componentArn,omitempty" tf:"component_arn,omitempty"`

	// Configuration block(s) for parameters to configure the component. Detailed below.
	Parameter []ParameterObservation `json:"parameter,omitempty" tf:"parameter,omitempty"`
}

type ContainerRecipeComponentParameters struct {

	// Amazon Resource Name (ARN) of the Image Builder Component to associate.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/imagebuilder/v1beta1.Component
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ComponentArn *string `json:"componentArn,omitempty" tf:"component_arn,omitempty"`

	// Reference to a Component in imagebuilder to populate componentArn.
	// +kubebuilder:validation:Optional
	ComponentArnRef *v1.Reference `json:"componentArnRef,omitempty" tf:"-"`

	// Selector for a Component in imagebuilder to populate componentArn.
	// +kubebuilder:validation:Optional
	ComponentArnSelector *v1.Selector `json:"componentArnSelector,omitempty" tf:"-"`

	// Configuration block(s) for parameters to configure the component. Detailed below.
	// +kubebuilder:validation:Optional
	Parameter []ParameterParameters `json:"parameter,omitempty" tf:"parameter,omitempty"`
}

type ContainerRecipeObservation struct {

	// Amazon Resource Name (ARN) of the container recipe.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// Ordered configuration block(s) with components for the container recipe. Detailed below.
	Component []ContainerRecipeComponentObservation `json:"component,omitempty" tf:"component,omitempty"`

	// The type of the container to create. Valid values: DOCKER.
	ContainerType *string `json:"containerType,omitempty" tf:"container_type,omitempty"`

	// Date the container recipe was created.
	DateCreated *string `json:"dateCreated,omitempty" tf:"date_created,omitempty"`

	// The description of the container recipe.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Dockerfile template used to build the image as an inline data blob.
	DockerfileTemplateData *string `json:"dockerfileTemplateData,omitempty" tf:"dockerfile_template_data,omitempty"`

	// The Amazon S3 URI for the Dockerfile that will be used to build the container image.
	DockerfileTemplateURI *string `json:"dockerfileTemplateUri,omitempty" tf:"dockerfile_template_uri,omitempty"`

	// A flag that indicates if the target container is encrypted.
	Encrypted *bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Configuration block used to configure an instance for building and testing container images. Detailed below.
	InstanceConfiguration []InstanceConfigurationObservation `json:"instanceConfiguration,omitempty" tf:"instance_configuration,omitempty"`

	// The KMS key used to encrypt the container image.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// The name of the container recipe.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Owner of the container recipe.
	Owner *string `json:"owner,omitempty" tf:"owner,omitempty"`

	// The base image for the container recipe.
	ParentImage *string `json:"parentImage,omitempty" tf:"parent_image,omitempty"`

	// Platform of the container recipe.
	Platform *string `json:"platform,omitempty" tf:"platform,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The destination repository for the container image. Detailed below.
	TargetRepository []TargetRepositoryObservation `json:"targetRepository,omitempty" tf:"target_repository,omitempty"`

	// Version of the container recipe.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// The working directory to be used during build and test workflows.
	WorkingDirectory *string `json:"workingDirectory,omitempty" tf:"working_directory,omitempty"`
}

type ContainerRecipeParameters struct {

	// Ordered configuration block(s) with components for the container recipe. Detailed below.
	// +kubebuilder:validation:Optional
	Component []ContainerRecipeComponentParameters `json:"component,omitempty" tf:"component,omitempty"`

	// The type of the container to create. Valid values: DOCKER.
	// +kubebuilder:validation:Optional
	ContainerType *string `json:"containerType,omitempty" tf:"container_type,omitempty"`

	// The description of the container recipe.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The Dockerfile template used to build the image as an inline data blob.
	// +kubebuilder:validation:Optional
	DockerfileTemplateData *string `json:"dockerfileTemplateData,omitempty" tf:"dockerfile_template_data,omitempty"`

	// The Amazon S3 URI for the Dockerfile that will be used to build the container image.
	// +kubebuilder:validation:Optional
	DockerfileTemplateURI *string `json:"dockerfileTemplateUri,omitempty" tf:"dockerfile_template_uri,omitempty"`

	// Configuration block used to configure an instance for building and testing container images. Detailed below.
	// +kubebuilder:validation:Optional
	InstanceConfiguration []InstanceConfigurationParameters `json:"instanceConfiguration,omitempty" tf:"instance_configuration,omitempty"`

	// The KMS key used to encrypt the container image.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Reference to a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDRef *v1.Reference `json:"kmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyId.
	// +kubebuilder:validation:Optional
	KMSKeyIDSelector *v1.Selector `json:"kmsKeyIdSelector,omitempty" tf:"-"`

	// The name of the container recipe.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The base image for the container recipe.
	// +kubebuilder:validation:Optional
	ParentImage *string `json:"parentImage,omitempty" tf:"parent_image,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The destination repository for the container image. Detailed below.
	// +kubebuilder:validation:Optional
	TargetRepository []TargetRepositoryParameters `json:"targetRepository,omitempty" tf:"target_repository,omitempty"`

	// Version of the container recipe.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// The working directory to be used during build and test workflows.
	// +kubebuilder:validation:Optional
	WorkingDirectory *string `json:"workingDirectory,omitempty" tf:"working_directory,omitempty"`
}

type EBSObservation struct {

	// Whether to delete the volume on termination. Defaults to unset, which is the value inherited from the parent image.
	DeleteOnTermination *string `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// Whether to encrypt the volume. Defaults to unset, which is the value inherited from the parent image.
	Encrypted *string `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// Number of Input/Output (I/O) operations per second to provision for an io1 or io2 volume.
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// Amazon Resource Name (ARN) of the Key Management Service (KMS) Key for encryption.
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Identifier of the EC2 Volume Snapshot.
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// For GP3 volumes only. The throughput in MiB/s that the volume supports.
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// Size of the volume, in GiB.
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// Type of the volume. For example, gp2 or io2.
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type EBSParameters struct {

	// Whether to delete the volume on termination. Defaults to unset, which is the value inherited from the parent image.
	// +kubebuilder:validation:Optional
	DeleteOnTermination *string `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`

	// Whether to encrypt the volume. Defaults to unset, which is the value inherited from the parent image.
	// +kubebuilder:validation:Optional
	Encrypted *string `json:"encrypted,omitempty" tf:"encrypted,omitempty"`

	// Number of Input/Output (I/O) operations per second to provision for an io1 or io2 volume.
	// +kubebuilder:validation:Optional
	Iops *float64 `json:"iops,omitempty" tf:"iops,omitempty"`

	// Amazon Resource Name (ARN) of the Key Management Service (KMS) Key for encryption.
	// +kubebuilder:validation:Optional
	KMSKeyID *string `json:"kmsKeyId,omitempty" tf:"kms_key_id,omitempty"`

	// Identifier of the EC2 Volume Snapshot.
	// +kubebuilder:validation:Optional
	SnapshotID *string `json:"snapshotId,omitempty" tf:"snapshot_id,omitempty"`

	// For GP3 volumes only. The throughput in MiB/s that the volume supports.
	// +kubebuilder:validation:Optional
	Throughput *float64 `json:"throughput,omitempty" tf:"throughput,omitempty"`

	// Size of the volume, in GiB.
	// +kubebuilder:validation:Optional
	VolumeSize *float64 `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`

	// Type of the volume. For example, gp2 or io2.
	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type InstanceConfigurationObservation struct {

	// Configuration block(s) with block device mappings for the container recipe. Detailed below.
	BlockDeviceMapping []BlockDeviceMappingObservation `json:"blockDeviceMapping,omitempty" tf:"block_device_mapping,omitempty"`

	// The AMI ID to use as the base image for a container build and test instance. If not specified, Image Builder will use the appropriate ECS-optimized AMI as a base image.
	Image *string `json:"image,omitempty" tf:"image,omitempty"`
}

type InstanceConfigurationParameters struct {

	// Configuration block(s) with block device mappings for the container recipe. Detailed below.
	// +kubebuilder:validation:Optional
	BlockDeviceMapping []BlockDeviceMappingParameters `json:"blockDeviceMapping,omitempty" tf:"block_device_mapping,omitempty"`

	// The AMI ID to use as the base image for a container build and test instance. If not specified, Image Builder will use the appropriate ECS-optimized AMI as a base image.
	// +kubebuilder:validation:Optional
	Image *string `json:"image,omitempty" tf:"image,omitempty"`
}

type ParameterObservation struct {

	// The name of the component parameter.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value for the named component parameter.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ParameterParameters struct {

	// The name of the component parameter.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The value for the named component parameter.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type TargetRepositoryObservation struct {

	// The name of the container repository where the output container image is stored. This name is prefixed by the repository location.
	RepositoryName *string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`

	// The service in which this image is registered. Valid values: ECR.
	Service *string `json:"service,omitempty" tf:"service,omitempty"`
}

type TargetRepositoryParameters struct {

	// The name of the container repository where the output container image is stored. This name is prefixed by the repository location.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ecr/v1beta1.Repository
	// +kubebuilder:validation:Optional
	RepositoryName *string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`

	// Reference to a Repository in ecr to populate repositoryName.
	// +kubebuilder:validation:Optional
	RepositoryNameRef *v1.Reference `json:"repositoryNameRef,omitempty" tf:"-"`

	// Selector for a Repository in ecr to populate repositoryName.
	// +kubebuilder:validation:Optional
	RepositoryNameSelector *v1.Selector `json:"repositoryNameSelector,omitempty" tf:"-"`

	// The service in which this image is registered. Valid values: ECR.
	// +kubebuilder:validation:Required
	Service *string `json:"service" tf:"service,omitempty"`
}

// ContainerRecipeSpec defines the desired state of ContainerRecipe
type ContainerRecipeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContainerRecipeParameters `json:"forProvider"`
}

// ContainerRecipeStatus defines the observed state of ContainerRecipe.
type ContainerRecipeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContainerRecipeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerRecipe is the Schema for the ContainerRecipes API. Manage an Image Builder Container Recipe
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ContainerRecipe struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.component)",message="component is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.containerType)",message="containerType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.parentImage)",message="parentImage is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.targetRepository)",message="targetRepository is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.version)",message="version is a required parameter"
	Spec   ContainerRecipeSpec   `json:"spec"`
	Status ContainerRecipeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerRecipeList contains a list of ContainerRecipes
type ContainerRecipeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerRecipe `json:"items"`
}

// Repository type metadata.
var (
	ContainerRecipe_Kind             = "ContainerRecipe"
	ContainerRecipe_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ContainerRecipe_Kind}.String()
	ContainerRecipe_KindAPIVersion   = ContainerRecipe_Kind + "." + CRDGroupVersion.String()
	ContainerRecipe_GroupVersionKind = CRDGroupVersion.WithKind(ContainerRecipe_Kind)
)

func init() {
	SchemeBuilder.Register(&ContainerRecipe{}, &ContainerRecipeList{})
}
