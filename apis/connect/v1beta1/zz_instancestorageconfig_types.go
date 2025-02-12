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

type EncryptionConfigInitParameters struct {

	// The type of encryption. Valid Values: KMS.
	EncryptionType *string `json:"encryptionType,omitempty" tf:"encryption_type,omitempty"`
}

type EncryptionConfigObservation struct {

	// The type of encryption. Valid Values: KMS.
	EncryptionType *string `json:"encryptionType,omitempty" tf:"encryption_type,omitempty"`

	// The full ARN of the encryption key. Be sure to provide the full ARN of the encryption key, not just the ID.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`
}

type EncryptionConfigParameters struct {

	// The type of encryption. Valid Values: KMS.
	// +kubebuilder:validation:Optional
	EncryptionType *string `json:"encryptionType" tf:"encryption_type,omitempty"`

	// The full ARN of the encryption key. Be sure to provide the full ARN of the encryption key, not just the ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// Reference to a Key in kms to populate keyId.
	// +kubebuilder:validation:Optional
	KeyIDRef *v1.Reference `json:"keyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate keyId.
	// +kubebuilder:validation:Optional
	KeyIDSelector *v1.Selector `json:"keyIdSelector,omitempty" tf:"-"`
}

type InstanceStorageConfigInitParameters struct {

	// A valid resource type. Valid Values: CHAT_TRANSCRIPTS | CALL_RECORDINGS | SCHEDULED_REPORTS | MEDIA_STREAMS | CONTACT_TRACE_RECORDS | AGENT_EVENTS | REAL_TIME_CONTACT_ANALYSIS_SEGMENTS.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// Specifies the storage configuration options for the Connect Instance. Documented below.
	StorageConfig []StorageConfigInitParameters `json:"storageConfig,omitempty" tf:"storage_config,omitempty"`
}

type InstanceStorageConfigObservation struct {

	// The existing association identifier that uniquely identifies the resource type and storage config for the given instance ID.
	AssociationID *string `json:"associationId,omitempty" tf:"association_id,omitempty"`

	// The identifier of the hosting Amazon Connect Instance, association_id, and resource_type separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// A valid resource type. Valid Values: CHAT_TRANSCRIPTS | CALL_RECORDINGS | SCHEDULED_REPORTS | MEDIA_STREAMS | CONTACT_TRACE_RECORDS | AGENT_EVENTS | REAL_TIME_CONTACT_ANALYSIS_SEGMENTS.
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// Specifies the storage configuration options for the Connect Instance. Documented below.
	StorageConfig []StorageConfigObservation `json:"storageConfig,omitempty" tf:"storage_config,omitempty"`
}

type InstanceStorageConfigParameters struct {

	// Specifies the identifier of the hosting Amazon Connect Instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/connect/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// A valid resource type. Valid Values: CHAT_TRANSCRIPTS | CALL_RECORDINGS | SCHEDULED_REPORTS | MEDIA_STREAMS | CONTACT_TRACE_RECORDS | AGENT_EVENTS | REAL_TIME_CONTACT_ANALYSIS_SEGMENTS.
	// +kubebuilder:validation:Optional
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// Specifies the storage configuration options for the Connect Instance. Documented below.
	// +kubebuilder:validation:Optional
	StorageConfig []StorageConfigParameters `json:"storageConfig,omitempty" tf:"storage_config,omitempty"`
}

type KinesisFirehoseConfigInitParameters struct {
}

type KinesisFirehoseConfigObservation struct {

	// The Amazon Resource Name (ARN) of the delivery stream.
	FirehoseArn *string `json:"firehoseArn,omitempty" tf:"firehose_arn,omitempty"`
}

type KinesisFirehoseConfigParameters struct {

	// The Amazon Resource Name (ARN) of the delivery stream.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/firehose/v1beta1.DeliveryStream
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",false)
	// +kubebuilder:validation:Optional
	FirehoseArn *string `json:"firehoseArn,omitempty" tf:"firehose_arn,omitempty"`

	// Reference to a DeliveryStream in firehose to populate firehoseArn.
	// +kubebuilder:validation:Optional
	FirehoseArnRef *v1.Reference `json:"firehoseArnRef,omitempty" tf:"-"`

	// Selector for a DeliveryStream in firehose to populate firehoseArn.
	// +kubebuilder:validation:Optional
	FirehoseArnSelector *v1.Selector `json:"firehoseArnSelector,omitempty" tf:"-"`
}

type KinesisStreamConfigInitParameters struct {
}

type KinesisStreamConfigObservation struct {

	// The Amazon Resource Name (ARN) of the data stream.
	StreamArn *string `json:"streamArn,omitempty" tf:"stream_arn,omitempty"`
}

type KinesisStreamConfigParameters struct {

	// The Amazon Resource Name (ARN) of the data stream.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kinesis/v1beta1.Stream
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",false)
	// +kubebuilder:validation:Optional
	StreamArn *string `json:"streamArn,omitempty" tf:"stream_arn,omitempty"`

	// Reference to a Stream in kinesis to populate streamArn.
	// +kubebuilder:validation:Optional
	StreamArnRef *v1.Reference `json:"streamArnRef,omitempty" tf:"-"`

	// Selector for a Stream in kinesis to populate streamArn.
	// +kubebuilder:validation:Optional
	StreamArnSelector *v1.Selector `json:"streamArnSelector,omitempty" tf:"-"`
}

type KinesisVideoStreamConfigInitParameters struct {

	// The encryption configuration. Documented below.
	EncryptionConfig []EncryptionConfigInitParameters `json:"encryptionConfig,omitempty" tf:"encryption_config,omitempty"`

	// The prefix of the video stream. Minimum length of 1. Maximum length of 128. When read from the state, the value returned is <prefix>-connect-<connect_instance_alias>-contact- since the API appends additional details to the prefix.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// The number of hours data is retained in the stream. Kinesis Video Streams retains the data in a data store that is associated with the stream. Minimum value of 0. Maximum value of 87600. A value of 0, indicates that the stream does not persist data.
	RetentionPeriodHours *float64 `json:"retentionPeriodHours,omitempty" tf:"retention_period_hours,omitempty"`
}

type KinesisVideoStreamConfigObservation struct {

	// The encryption configuration. Documented below.
	EncryptionConfig []EncryptionConfigObservation `json:"encryptionConfig,omitempty" tf:"encryption_config,omitempty"`

	// The prefix of the video stream. Minimum length of 1. Maximum length of 128. When read from the state, the value returned is <prefix>-connect-<connect_instance_alias>-contact- since the API appends additional details to the prefix.
	Prefix *string `json:"prefix,omitempty" tf:"prefix,omitempty"`

	// The number of hours data is retained in the stream. Kinesis Video Streams retains the data in a data store that is associated with the stream. Minimum value of 0. Maximum value of 87600. A value of 0, indicates that the stream does not persist data.
	RetentionPeriodHours *float64 `json:"retentionPeriodHours,omitempty" tf:"retention_period_hours,omitempty"`
}

type KinesisVideoStreamConfigParameters struct {

	// The encryption configuration. Documented below.
	// +kubebuilder:validation:Optional
	EncryptionConfig []EncryptionConfigParameters `json:"encryptionConfig" tf:"encryption_config,omitempty"`

	// The prefix of the video stream. Minimum length of 1. Maximum length of 128. When read from the state, the value returned is <prefix>-connect-<connect_instance_alias>-contact- since the API appends additional details to the prefix.
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix" tf:"prefix,omitempty"`

	// The number of hours data is retained in the stream. Kinesis Video Streams retains the data in a data store that is associated with the stream. Minimum value of 0. Maximum value of 87600. A value of 0, indicates that the stream does not persist data.
	// +kubebuilder:validation:Optional
	RetentionPeriodHours *float64 `json:"retentionPeriodHours" tf:"retention_period_hours,omitempty"`
}

type S3ConfigEncryptionConfigInitParameters struct {

	// The type of encryption. Valid Values: KMS.
	EncryptionType *string `json:"encryptionType,omitempty" tf:"encryption_type,omitempty"`
}

type S3ConfigEncryptionConfigObservation struct {

	// The type of encryption. Valid Values: KMS.
	EncryptionType *string `json:"encryptionType,omitempty" tf:"encryption_type,omitempty"`

	// The full ARN of the encryption key. Be sure to provide the full ARN of the encryption key, not just the ID.
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`
}

type S3ConfigEncryptionConfigParameters struct {

	// The type of encryption. Valid Values: KMS.
	// +kubebuilder:validation:Optional
	EncryptionType *string `json:"encryptionType" tf:"encryption_type,omitempty"`

	// The full ARN of the encryption key. Be sure to provide the full ARN of the encryption key, not just the ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	KeyID *string `json:"keyId,omitempty" tf:"key_id,omitempty"`

	// Reference to a Key in kms to populate keyId.
	// +kubebuilder:validation:Optional
	KeyIDRef *v1.Reference `json:"keyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate keyId.
	// +kubebuilder:validation:Optional
	KeyIDSelector *v1.Selector `json:"keyIdSelector,omitempty" tf:"-"`
}

type S3ConfigInitParameters struct {

	// The S3 bucket prefix.
	BucketPrefix *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`

	// The encryption configuration. Documented below.
	EncryptionConfig []S3ConfigEncryptionConfigInitParameters `json:"encryptionConfig,omitempty" tf:"encryption_config,omitempty"`
}

type S3ConfigObservation struct {

	// The S3 bucket name.
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// The S3 bucket prefix.
	BucketPrefix *string `json:"bucketPrefix,omitempty" tf:"bucket_prefix,omitempty"`

	// The encryption configuration. Documented below.
	EncryptionConfig []S3ConfigEncryptionConfigObservation `json:"encryptionConfig,omitempty" tf:"encryption_config,omitempty"`
}

type S3ConfigParameters struct {

	// The S3 bucket name.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Reference to a Bucket in s3 to populate bucketName.
	// +kubebuilder:validation:Optional
	BucketNameRef *v1.Reference `json:"bucketNameRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucketName.
	// +kubebuilder:validation:Optional
	BucketNameSelector *v1.Selector `json:"bucketNameSelector,omitempty" tf:"-"`

	// The S3 bucket prefix.
	// +kubebuilder:validation:Optional
	BucketPrefix *string `json:"bucketPrefix" tf:"bucket_prefix,omitempty"`

	// The encryption configuration. Documented below.
	// +kubebuilder:validation:Optional
	EncryptionConfig []S3ConfigEncryptionConfigParameters `json:"encryptionConfig,omitempty" tf:"encryption_config,omitempty"`
}

type StorageConfigInitParameters struct {

	// A block that specifies the configuration of the Kinesis Firehose delivery stream. Documented below.
	KinesisFirehoseConfig []KinesisFirehoseConfigInitParameters `json:"kinesisFirehoseConfig,omitempty" tf:"kinesis_firehose_config,omitempty"`

	// A block that specifies the configuration of the Kinesis data stream. Documented below.
	KinesisStreamConfig []KinesisStreamConfigInitParameters `json:"kinesisStreamConfig,omitempty" tf:"kinesis_stream_config,omitempty"`

	// A block that specifies the configuration of the Kinesis video stream. Documented below.
	KinesisVideoStreamConfig []KinesisVideoStreamConfigInitParameters `json:"kinesisVideoStreamConfig,omitempty" tf:"kinesis_video_stream_config,omitempty"`

	// A block that specifies the configuration of S3 Bucket. Documented below.
	S3Config []S3ConfigInitParameters `json:"s3Config,omitempty" tf:"s3_config,omitempty"`

	// A valid storage type. Valid Values: S3 | KINESIS_VIDEO_STREAM | KINESIS_STREAM | KINESIS_FIREHOSE.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`
}

type StorageConfigObservation struct {

	// A block that specifies the configuration of the Kinesis Firehose delivery stream. Documented below.
	KinesisFirehoseConfig []KinesisFirehoseConfigObservation `json:"kinesisFirehoseConfig,omitempty" tf:"kinesis_firehose_config,omitempty"`

	// A block that specifies the configuration of the Kinesis data stream. Documented below.
	KinesisStreamConfig []KinesisStreamConfigObservation `json:"kinesisStreamConfig,omitempty" tf:"kinesis_stream_config,omitempty"`

	// A block that specifies the configuration of the Kinesis video stream. Documented below.
	KinesisVideoStreamConfig []KinesisVideoStreamConfigObservation `json:"kinesisVideoStreamConfig,omitempty" tf:"kinesis_video_stream_config,omitempty"`

	// A block that specifies the configuration of S3 Bucket. Documented below.
	S3Config []S3ConfigObservation `json:"s3Config,omitempty" tf:"s3_config,omitempty"`

	// A valid storage type. Valid Values: S3 | KINESIS_VIDEO_STREAM | KINESIS_STREAM | KINESIS_FIREHOSE.
	StorageType *string `json:"storageType,omitempty" tf:"storage_type,omitempty"`
}

type StorageConfigParameters struct {

	// A block that specifies the configuration of the Kinesis Firehose delivery stream. Documented below.
	// +kubebuilder:validation:Optional
	KinesisFirehoseConfig []KinesisFirehoseConfigParameters `json:"kinesisFirehoseConfig,omitempty" tf:"kinesis_firehose_config,omitempty"`

	// A block that specifies the configuration of the Kinesis data stream. Documented below.
	// +kubebuilder:validation:Optional
	KinesisStreamConfig []KinesisStreamConfigParameters `json:"kinesisStreamConfig,omitempty" tf:"kinesis_stream_config,omitempty"`

	// A block that specifies the configuration of the Kinesis video stream. Documented below.
	// +kubebuilder:validation:Optional
	KinesisVideoStreamConfig []KinesisVideoStreamConfigParameters `json:"kinesisVideoStreamConfig,omitempty" tf:"kinesis_video_stream_config,omitempty"`

	// A block that specifies the configuration of S3 Bucket. Documented below.
	// +kubebuilder:validation:Optional
	S3Config []S3ConfigParameters `json:"s3Config,omitempty" tf:"s3_config,omitempty"`

	// A valid storage type. Valid Values: S3 | KINESIS_VIDEO_STREAM | KINESIS_STREAM | KINESIS_FIREHOSE.
	// +kubebuilder:validation:Optional
	StorageType *string `json:"storageType" tf:"storage_type,omitempty"`
}

// InstanceStorageConfigSpec defines the desired state of InstanceStorageConfig
type InstanceStorageConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceStorageConfigParameters `json:"forProvider"`
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
	InitProvider InstanceStorageConfigInitParameters `json:"initProvider,omitempty"`
}

// InstanceStorageConfigStatus defines the observed state of InstanceStorageConfig.
type InstanceStorageConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceStorageConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceStorageConfig is the Schema for the InstanceStorageConfigs API. Provides details about a specific Amazon Connect Instance Storage Config.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type InstanceStorageConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourceType) || (has(self.initProvider) && has(self.initProvider.resourceType))",message="spec.forProvider.resourceType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.storageConfig) || (has(self.initProvider) && has(self.initProvider.storageConfig))",message="spec.forProvider.storageConfig is a required parameter"
	Spec   InstanceStorageConfigSpec   `json:"spec"`
	Status InstanceStorageConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceStorageConfigList contains a list of InstanceStorageConfigs
type InstanceStorageConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceStorageConfig `json:"items"`
}

// Repository type metadata.
var (
	InstanceStorageConfig_Kind             = "InstanceStorageConfig"
	InstanceStorageConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceStorageConfig_Kind}.String()
	InstanceStorageConfig_KindAPIVersion   = InstanceStorageConfig_Kind + "." + CRDGroupVersion.String()
	InstanceStorageConfig_GroupVersionKind = CRDGroupVersion.WithKind(InstanceStorageConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceStorageConfig{}, &InstanceStorageConfigList{})
}
