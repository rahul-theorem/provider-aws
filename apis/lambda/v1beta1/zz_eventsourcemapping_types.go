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

type AmazonManagedKafkaEventSourceConfigObservation struct {

	// A Kafka consumer group ID between 1 and 200 characters for use when creating this event source mapping. If one is not specified, this value will be automatically generated. See AmazonManagedKafkaEventSourceConfig Syntax.
	ConsumerGroupID *string `json:"consumerGroupId,omitempty" tf:"consumer_group_id,omitempty"`
}

type AmazonManagedKafkaEventSourceConfigParameters struct {

	// A Kafka consumer group ID between 1 and 200 characters for use when creating this event source mapping. If one is not specified, this value will be automatically generated. See AmazonManagedKafkaEventSourceConfig Syntax.
	// +kubebuilder:validation:Optional
	ConsumerGroupID *string `json:"consumerGroupId,omitempty" tf:"consumer_group_id,omitempty"`
}

type DestinationConfigObservation struct {

	// The destination configuration for failed invocations. Detailed below.
	OnFailure []OnFailureObservation `json:"onFailure,omitempty" tf:"on_failure,omitempty"`
}

type DestinationConfigParameters struct {

	// The destination configuration for failed invocations. Detailed below.
	// +kubebuilder:validation:Optional
	OnFailure []OnFailureParameters `json:"onFailure,omitempty" tf:"on_failure,omitempty"`
}

type EventSourceMappingObservation struct {

	// Additional configuration block for Amazon Managed Kafka sources. Incompatible with "self_managed_event_source" and "self_managed_kafka_event_source_config". Detailed below.
	AmazonManagedKafkaEventSourceConfig []AmazonManagedKafkaEventSourceConfigObservation `json:"amazonManagedKafkaEventSourceConfig,omitempty" tf:"amazon_managed_kafka_event_source_config,omitempty"`

	// The largest number of records that Lambda will retrieve from your event source at the time of invocation. Defaults to 100 for DynamoDB, Kinesis, MQ and MSK, 10 for SQS.
	BatchSize *float64 `json:"batchSize,omitempty" tf:"batch_size,omitempty"`

	// If the function returns an error, split the batch in two and retry. Only available for stream sources (DynamoDB and Kinesis). Defaults to false.
	BisectBatchOnFunctionError *bool `json:"bisectBatchOnFunctionError,omitempty" tf:"bisect_batch_on_function_error,omitempty"`

	// An Amazon SQS queue or Amazon SNS topic destination for failed records. Only available for stream sources (DynamoDB and Kinesis). Detailed below.
	DestinationConfig []DestinationConfigObservation `json:"destinationConfig,omitempty" tf:"destination_config,omitempty"`

	// Determines if the mapping will be enabled on creation. Defaults to true.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The event source ARN - this is required for Kinesis stream, DynamoDB stream, SQS queue, MQ broker or MSK cluster.  It is incompatible with a Self Managed Kafka source.
	EventSourceArn *string `json:"eventSourceArn,omitempty" tf:"event_source_arn,omitempty"`

	// The criteria to use for event filtering Kinesis stream, DynamoDB stream, SQS queue event sources. Detailed below.
	FilterCriteria []FilterCriteriaObservation `json:"filterCriteria,omitempty" tf:"filter_criteria,omitempty"`

	// The the ARN of the Lambda function the event source mapping is sending events to. (Note: this is a computed value that differs from function_name above.)
	FunctionArn *string `json:"functionArn,omitempty" tf:"function_arn,omitempty"`

	// The name or the ARN of the Lambda function that will be subscribing to events.
	FunctionName *string `json:"functionName,omitempty" tf:"function_name,omitempty"`

	// A list of current response type enums applied to the event source mapping for AWS Lambda checkpointing. Only available for SQS and stream sources (DynamoDB and Kinesis). Valid values: ReportBatchItemFailures.
	FunctionResponseTypes []*string `json:"functionResponseTypes,omitempty" tf:"function_response_types,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The date this resource was last modified.
	LastModified *string `json:"lastModified,omitempty" tf:"last_modified,omitempty"`

	// The result of the last AWS Lambda invocation of your Lambda function.
	LastProcessingResult *string `json:"lastProcessingResult,omitempty" tf:"last_processing_result,omitempty"`

	// The maximum amount of time to gather records before invoking the function, in seconds (between 0 and 300). Records will continue to buffer (or accumulate in the case of an SQS queue event source) until either maximum_batching_window_in_seconds expires or batch_size has been met. For streaming event sources, defaults to as soon as records are available in the stream. If the batch it reads from the stream/queue only has one record in it, Lambda only sends one record to the function. Only available for stream sources (DynamoDB and Kinesis) and SQS standard queues.
	MaximumBatchingWindowInSeconds *float64 `json:"maximumBatchingWindowInSeconds,omitempty" tf:"maximum_batching_window_in_seconds,omitempty"`

	// The maximum age of a record that Lambda sends to a function for processing. Only available for stream sources (DynamoDB and Kinesis). Must be either -1 (forever, and the default value) or between 60 and 604800 (inclusive).
	MaximumRecordAgeInSeconds *float64 `json:"maximumRecordAgeInSeconds,omitempty" tf:"maximum_record_age_in_seconds,omitempty"`

	// The maximum number of times to retry when the function returns an error. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of -1 (forever), maximum of 10000.
	MaximumRetryAttempts *float64 `json:"maximumRetryAttempts,omitempty" tf:"maximum_retry_attempts,omitempty"`

	// The number of batches to process from each shard concurrently. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of 1, maximum of 10.
	ParallelizationFactor *float64 `json:"parallelizationFactor,omitempty" tf:"parallelization_factor,omitempty"`

	// The name of the Amazon MQ broker destination queue to consume. Only available for MQ sources. A single queue name must be specified.
	Queues []*string `json:"queues,omitempty" tf:"queues,omitempty"`

	// Scaling configuration of the event source. Only available for SQS queues. Detailed below.
	ScalingConfig []ScalingConfigObservation `json:"scalingConfig,omitempty" tf:"scaling_config,omitempty"`

	// For Self Managed Kafka sources, the location of the self managed cluster. If set, configuration must also include source_access_configuration. Detailed below.
	SelfManagedEventSource []SelfManagedEventSourceObservation `json:"selfManagedEventSource,omitempty" tf:"self_managed_event_source,omitempty"`

	// Additional configuration block for Self Managed Kafka sources. Incompatible with "event_source_arn" and "amazon_managed_kafka_event_source_config". Detailed below.
	SelfManagedKafkaEventSourceConfig []SelfManagedKafkaEventSourceConfigObservation `json:"selfManagedKafkaEventSourceConfig,omitempty" tf:"self_managed_kafka_event_source_config,omitempty"`

	// :  For Self Managed Kafka sources, the access configuration for the source. If set, configuration must also include self_managed_event_source. Detailed below.
	SourceAccessConfiguration []SourceAccessConfigurationObservation `json:"sourceAccessConfiguration,omitempty" tf:"source_access_configuration,omitempty"`

	// The position in the stream where AWS Lambda should start reading. Must be one of AT_TIMESTAMP (Kinesis only), LATEST or TRIM_HORIZON if getting events from Kinesis, DynamoDB, MSK or Self Managed Apache Kafka. Must not be provided if getting events from SQS. More information about these positions can be found in the AWS DynamoDB Streams API Reference and AWS Kinesis API Reference.
	StartingPosition *string `json:"startingPosition,omitempty" tf:"starting_position,omitempty"`

	// A timestamp in RFC3339 format of the data record which to start reading when using starting_position set to AT_TIMESTAMP. If a record with this exact timestamp does not exist, the next later record is chosen. If the timestamp is older than the current trim horizon, the oldest available record is chosen.
	StartingPositionTimestamp *string `json:"startingPositionTimestamp,omitempty" tf:"starting_position_timestamp,omitempty"`

	// The state of the event source mapping.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The reason the event source mapping is in its current state.
	StateTransitionReason *string `json:"stateTransitionReason,omitempty" tf:"state_transition_reason,omitempty"`

	// The name of the Kafka topics. Only available for MSK sources. A single topic name must be specified.
	Topics []*string `json:"topics,omitempty" tf:"topics,omitempty"`

	// The duration in seconds of a processing window for AWS Lambda streaming analytics. The range is between 1 second up to 900 seconds. Only available for stream sources (DynamoDB and Kinesis).
	TumblingWindowInSeconds *float64 `json:"tumblingWindowInSeconds,omitempty" tf:"tumbling_window_in_seconds,omitempty"`

	// The UUID of the created event source mapping.
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`
}

type EventSourceMappingParameters struct {

	// Additional configuration block for Amazon Managed Kafka sources. Incompatible with "self_managed_event_source" and "self_managed_kafka_event_source_config". Detailed below.
	// +kubebuilder:validation:Optional
	AmazonManagedKafkaEventSourceConfig []AmazonManagedKafkaEventSourceConfigParameters `json:"amazonManagedKafkaEventSourceConfig,omitempty" tf:"amazon_managed_kafka_event_source_config,omitempty"`

	// The largest number of records that Lambda will retrieve from your event source at the time of invocation. Defaults to 100 for DynamoDB, Kinesis, MQ and MSK, 10 for SQS.
	// +kubebuilder:validation:Optional
	BatchSize *float64 `json:"batchSize,omitempty" tf:"batch_size,omitempty"`

	// If the function returns an error, split the batch in two and retry. Only available for stream sources (DynamoDB and Kinesis). Defaults to false.
	// +kubebuilder:validation:Optional
	BisectBatchOnFunctionError *bool `json:"bisectBatchOnFunctionError,omitempty" tf:"bisect_batch_on_function_error,omitempty"`

	// An Amazon SQS queue or Amazon SNS topic destination for failed records. Only available for stream sources (DynamoDB and Kinesis). Detailed below.
	// +kubebuilder:validation:Optional
	DestinationConfig []DestinationConfigParameters `json:"destinationConfig,omitempty" tf:"destination_config,omitempty"`

	// Determines if the mapping will be enabled on creation. Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The event source ARN - this is required for Kinesis stream, DynamoDB stream, SQS queue, MQ broker or MSK cluster.  It is incompatible with a Self Managed Kafka source.
	// +kubebuilder:validation:Optional
	EventSourceArn *string `json:"eventSourceArn,omitempty" tf:"event_source_arn,omitempty"`

	// The criteria to use for event filtering Kinesis stream, DynamoDB stream, SQS queue event sources. Detailed below.
	// +kubebuilder:validation:Optional
	FilterCriteria []FilterCriteriaParameters `json:"filterCriteria,omitempty" tf:"filter_criteria,omitempty"`

	// The name or the ARN of the Lambda function that will be subscribing to events.
	// +crossplane:generate:reference:type=Function
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	FunctionName *string `json:"functionName,omitempty" tf:"function_name,omitempty"`

	// Reference to a Function to populate functionName.
	// +kubebuilder:validation:Optional
	FunctionNameRef *v1.Reference `json:"functionNameRef,omitempty" tf:"-"`

	// Selector for a Function to populate functionName.
	// +kubebuilder:validation:Optional
	FunctionNameSelector *v1.Selector `json:"functionNameSelector,omitempty" tf:"-"`

	// A list of current response type enums applied to the event source mapping for AWS Lambda checkpointing. Only available for SQS and stream sources (DynamoDB and Kinesis). Valid values: ReportBatchItemFailures.
	// +kubebuilder:validation:Optional
	FunctionResponseTypes []*string `json:"functionResponseTypes,omitempty" tf:"function_response_types,omitempty"`

	// The maximum amount of time to gather records before invoking the function, in seconds (between 0 and 300). Records will continue to buffer (or accumulate in the case of an SQS queue event source) until either maximum_batching_window_in_seconds expires or batch_size has been met. For streaming event sources, defaults to as soon as records are available in the stream. If the batch it reads from the stream/queue only has one record in it, Lambda only sends one record to the function. Only available for stream sources (DynamoDB and Kinesis) and SQS standard queues.
	// +kubebuilder:validation:Optional
	MaximumBatchingWindowInSeconds *float64 `json:"maximumBatchingWindowInSeconds,omitempty" tf:"maximum_batching_window_in_seconds,omitempty"`

	// The maximum age of a record that Lambda sends to a function for processing. Only available for stream sources (DynamoDB and Kinesis). Must be either -1 (forever, and the default value) or between 60 and 604800 (inclusive).
	// +kubebuilder:validation:Optional
	MaximumRecordAgeInSeconds *float64 `json:"maximumRecordAgeInSeconds,omitempty" tf:"maximum_record_age_in_seconds,omitempty"`

	// The maximum number of times to retry when the function returns an error. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of -1 (forever), maximum of 10000.
	// +kubebuilder:validation:Optional
	MaximumRetryAttempts *float64 `json:"maximumRetryAttempts,omitempty" tf:"maximum_retry_attempts,omitempty"`

	// The number of batches to process from each shard concurrently. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of 1, maximum of 10.
	// +kubebuilder:validation:Optional
	ParallelizationFactor *float64 `json:"parallelizationFactor,omitempty" tf:"parallelization_factor,omitempty"`

	// The name of the Amazon MQ broker destination queue to consume. Only available for MQ sources. A single queue name must be specified.
	// +kubebuilder:validation:Optional
	Queues []*string `json:"queues,omitempty" tf:"queues,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Scaling configuration of the event source. Only available for SQS queues. Detailed below.
	// +kubebuilder:validation:Optional
	ScalingConfig []ScalingConfigParameters `json:"scalingConfig,omitempty" tf:"scaling_config,omitempty"`

	// For Self Managed Kafka sources, the location of the self managed cluster. If set, configuration must also include source_access_configuration. Detailed below.
	// +kubebuilder:validation:Optional
	SelfManagedEventSource []SelfManagedEventSourceParameters `json:"selfManagedEventSource,omitempty" tf:"self_managed_event_source,omitempty"`

	// Additional configuration block for Self Managed Kafka sources. Incompatible with "event_source_arn" and "amazon_managed_kafka_event_source_config". Detailed below.
	// +kubebuilder:validation:Optional
	SelfManagedKafkaEventSourceConfig []SelfManagedKafkaEventSourceConfigParameters `json:"selfManagedKafkaEventSourceConfig,omitempty" tf:"self_managed_kafka_event_source_config,omitempty"`

	// :  For Self Managed Kafka sources, the access configuration for the source. If set, configuration must also include self_managed_event_source. Detailed below.
	// +kubebuilder:validation:Optional
	SourceAccessConfiguration []SourceAccessConfigurationParameters `json:"sourceAccessConfiguration,omitempty" tf:"source_access_configuration,omitempty"`

	// The position in the stream where AWS Lambda should start reading. Must be one of AT_TIMESTAMP (Kinesis only), LATEST or TRIM_HORIZON if getting events from Kinesis, DynamoDB, MSK or Self Managed Apache Kafka. Must not be provided if getting events from SQS. More information about these positions can be found in the AWS DynamoDB Streams API Reference and AWS Kinesis API Reference.
	// +kubebuilder:validation:Optional
	StartingPosition *string `json:"startingPosition,omitempty" tf:"starting_position,omitempty"`

	// A timestamp in RFC3339 format of the data record which to start reading when using starting_position set to AT_TIMESTAMP. If a record with this exact timestamp does not exist, the next later record is chosen. If the timestamp is older than the current trim horizon, the oldest available record is chosen.
	// +kubebuilder:validation:Optional
	StartingPositionTimestamp *string `json:"startingPositionTimestamp,omitempty" tf:"starting_position_timestamp,omitempty"`

	// The name of the Kafka topics. Only available for MSK sources. A single topic name must be specified.
	// +kubebuilder:validation:Optional
	Topics []*string `json:"topics,omitempty" tf:"topics,omitempty"`

	// The duration in seconds of a processing window for AWS Lambda streaming analytics. The range is between 1 second up to 900 seconds. Only available for stream sources (DynamoDB and Kinesis).
	// +kubebuilder:validation:Optional
	TumblingWindowInSeconds *float64 `json:"tumblingWindowInSeconds,omitempty" tf:"tumbling_window_in_seconds,omitempty"`
}

type FilterCriteriaObservation struct {

	// A set of up to 5 filter. If an event satisfies at least one, Lambda sends the event to the function or adds it to the next batch. Detailed below.
	Filter []FilterObservation `json:"filter,omitempty" tf:"filter,omitempty"`
}

type FilterCriteriaParameters struct {

	// A set of up to 5 filter. If an event satisfies at least one, Lambda sends the event to the function or adds it to the next batch. Detailed below.
	// +kubebuilder:validation:Optional
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter,omitempty"`
}

type FilterObservation struct {

	// A filter pattern up to 4096 characters. See Filter Rule Syntax.
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`
}

type FilterParameters struct {

	// A filter pattern up to 4096 characters. See Filter Rule Syntax.
	// +kubebuilder:validation:Optional
	Pattern *string `json:"pattern,omitempty" tf:"pattern,omitempty"`
}

type OnFailureObservation struct {

	// The Amazon Resource Name (ARN) of the destination resource.
	DestinationArn *string `json:"destinationArn,omitempty" tf:"destination_arn,omitempty"`
}

type OnFailureParameters struct {

	// The Amazon Resource Name (ARN) of the destination resource.
	// +kubebuilder:validation:Required
	DestinationArn *string `json:"destinationArn" tf:"destination_arn,omitempty"`
}

type ScalingConfigObservation struct {

	// Limits the number of concurrent instances that the Amazon SQS event source can invoke. Must be between 2 and 1000. See Configuring maximum concurrency for Amazon SQS event sources.
	MaximumConcurrency *float64 `json:"maximumConcurrency,omitempty" tf:"maximum_concurrency,omitempty"`
}

type ScalingConfigParameters struct {

	// Limits the number of concurrent instances that the Amazon SQS event source can invoke. Must be between 2 and 1000. See Configuring maximum concurrency for Amazon SQS event sources.
	// +kubebuilder:validation:Optional
	MaximumConcurrency *float64 `json:"maximumConcurrency,omitempty" tf:"maximum_concurrency,omitempty"`
}

type SelfManagedEventSourceObservation struct {

	// A map of endpoints for the self managed source.  For Kafka self-managed sources, the key should be KAFKA_BOOTSTRAP_SERVERS and the value should be a string with a comma separated list of broker endpoints.
	Endpoints map[string]*string `json:"endpoints,omitempty" tf:"endpoints,omitempty"`
}

type SelfManagedEventSourceParameters struct {

	// A map of endpoints for the self managed source.  For Kafka self-managed sources, the key should be KAFKA_BOOTSTRAP_SERVERS and the value should be a string with a comma separated list of broker endpoints.
	// +kubebuilder:validation:Required
	Endpoints map[string]*string `json:"endpoints" tf:"endpoints,omitempty"`
}

type SelfManagedKafkaEventSourceConfigObservation struct {

	// A Kafka consumer group ID between 1 and 200 characters for use when creating this event source mapping. If one is not specified, this value will be automatically generated. See SelfManagedKafkaEventSourceConfig Syntax.
	ConsumerGroupID *string `json:"consumerGroupId,omitempty" tf:"consumer_group_id,omitempty"`
}

type SelfManagedKafkaEventSourceConfigParameters struct {

	// A Kafka consumer group ID between 1 and 200 characters for use when creating this event source mapping. If one is not specified, this value will be automatically generated. See SelfManagedKafkaEventSourceConfig Syntax.
	// +kubebuilder:validation:Optional
	ConsumerGroupID *string `json:"consumerGroupId,omitempty" tf:"consumer_group_id,omitempty"`
}

type SourceAccessConfigurationObservation struct {

	// The type of this configuration.  For Self Managed Kafka you will need to supply blocks for type VPC_SUBNET and VPC_SECURITY_GROUP.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// The URI for this configuration.  For type VPC_SUBNET the value should be subnet:subnet_id where subnet_id is the value you would find in an aws_subnet resource's id attribute.  For type VPC_SECURITY_GROUP the value should be security_group:security_group_id where security_group_id is the value you would find in an aws_security_group resource's id attribute.
	URI *string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type SourceAccessConfigurationParameters struct {

	// The type of this configuration.  For Self Managed Kafka you will need to supply blocks for type VPC_SUBNET and VPC_SECURITY_GROUP.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// The URI for this configuration.  For type VPC_SUBNET the value should be subnet:subnet_id where subnet_id is the value you would find in an aws_subnet resource's id attribute.  For type VPC_SECURITY_GROUP the value should be security_group:security_group_id where security_group_id is the value you would find in an aws_security_group resource's id attribute.
	// +kubebuilder:validation:Required
	URI *string `json:"uri" tf:"uri,omitempty"`
}

// EventSourceMappingSpec defines the desired state of EventSourceMapping
type EventSourceMappingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EventSourceMappingParameters `json:"forProvider"`
}

// EventSourceMappingStatus defines the observed state of EventSourceMapping.
type EventSourceMappingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EventSourceMappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EventSourceMapping is the Schema for the EventSourceMappings API. Provides a Lambda event source mapping. This allows Lambda functions to get events from Kinesis, DynamoDB, SQS, Amazon MQ and Managed Streaming for Apache Kafka (MSK).
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type EventSourceMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventSourceMappingSpec   `json:"spec"`
	Status            EventSourceMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventSourceMappingList contains a list of EventSourceMappings
type EventSourceMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventSourceMapping `json:"items"`
}

// Repository type metadata.
var (
	EventSourceMapping_Kind             = "EventSourceMapping"
	EventSourceMapping_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EventSourceMapping_Kind}.String()
	EventSourceMapping_KindAPIVersion   = EventSourceMapping_Kind + "." + CRDGroupVersion.String()
	EventSourceMapping_GroupVersionKind = CRDGroupVersion.WithKind(EventSourceMapping_Kind)
)

func init() {
	SchemeBuilder.Register(&EventSourceMapping{}, &EventSourceMappingList{})
}
