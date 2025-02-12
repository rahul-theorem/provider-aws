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

type ClusterEndpointInitParameters struct {
}

type ClusterEndpointObservation struct {

	// DNS hostname of the cluster configuration endpoint.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The port number on which each of the nodes accepts connections. Defaults to 6379.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type ClusterEndpointParameters struct {
}

type ClusterInitParameters struct {

	// The name of the Access Control List to associate with the cluster.
	ACLName *string `json:"aclName,omitempty" tf:"acl_name,omitempty"`

	// When set to true, the cluster will automatically receive minor engine version upgrades after launch. Defaults to true.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// Enables data tiering. This option is not supported by all instance types. For more information, see Data tiering.
	DataTiering *bool `json:"dataTiering,omitempty" tf:"data_tiering,omitempty"`

	// Description for the cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Version number of the Redis engine to be used for the cluster. Downgrades are not supported.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Name of the final cluster snapshot to be created when this resource is deleted. If omitted, no final snapshot will be made.
	FinalSnapshotName *string `json:"finalSnapshotName,omitempty" tf:"final_snapshot_name,omitempty"`

	// Specifies the weekly time range during which maintenance on the cluster is performed. Specify as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example: sun:23:00-mon:01:30.
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// The compute and memory capacity of the nodes in the cluster. See AWS documentation on supported node types as well as vertical scaling.
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// The number of replicas to apply to each shard, up to a maximum of 5. Defaults to 1 (i.e. 2 nodes per shard).
	NumReplicasPerShard *float64 `json:"numReplicasPerShard,omitempty" tf:"num_replicas_per_shard,omitempty"`

	// The number of shards in the cluster. Defaults to 1.
	NumShards *float64 `json:"numShards,omitempty" tf:"num_shards,omitempty"`

	// The name of the parameter group associated with the cluster.
	ParameterGroupName *string `json:"parameterGroupName,omitempty" tf:"parameter_group_name,omitempty"`

	// The port number on which each of the nodes accepts connections. Defaults to 6379.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// List of ARN-s that uniquely identify RDB snapshot files stored in S3. The snapshot files will be used to populate the new cluster. Object names in the ARN-s cannot contain any commas.
	SnapshotArns []*string `json:"snapshotArns,omitempty" tf:"snapshot_arns,omitempty"`

	// The name of a snapshot from which to restore data into the new cluster.
	SnapshotName *string `json:"snapshotName,omitempty" tf:"snapshot_name,omitempty"`

	// The number of days for which MemoryDB retains automatic snapshots before deleting them. When set to 0, automatic backups are disabled. Defaults to 0.
	SnapshotRetentionLimit *float64 `json:"snapshotRetentionLimit,omitempty" tf:"snapshot_retention_limit,omitempty"`

	// The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your shard. Example: 05:00-09:00.
	SnapshotWindow *string `json:"snapshotWindow,omitempty" tf:"snapshot_window,omitempty"`

	// ARN of the SNS topic to which cluster notifications are sent.
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`

	// A flag to enable in-transit encryption on the cluster. When set to false, the acl_name must be open-access. Defaults to true.
	TLSEnabled *bool `json:"tlsEnabled,omitempty" tf:"tls_enabled,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ClusterObservation struct {

	// The name of the Access Control List to associate with the cluster.
	ACLName *string `json:"aclName,omitempty" tf:"acl_name,omitempty"`

	// The ARN of the cluster.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// When set to true, the cluster will automatically receive minor engine version upgrades after launch. Defaults to true.
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	ClusterEndpoint []ClusterEndpointObservation `json:"clusterEndpoint,omitempty" tf:"cluster_endpoint,omitempty"`

	// Enables data tiering. This option is not supported by all instance types. For more information, see Data tiering.
	DataTiering *bool `json:"dataTiering,omitempty" tf:"data_tiering,omitempty"`

	// Description for the cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Patch version number of the Redis engine used by the cluster.
	EnginePatchVersion *string `json:"enginePatchVersion,omitempty" tf:"engine_patch_version,omitempty"`

	// Version number of the Redis engine to be used for the cluster. Downgrades are not supported.
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Name of the final cluster snapshot to be created when this resource is deleted. If omitted, no final snapshot will be made.
	FinalSnapshotName *string `json:"finalSnapshotName,omitempty" tf:"final_snapshot_name,omitempty"`

	// Same as name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ARN of the KMS key used to encrypt the cluster at rest.
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Specifies the weekly time range during which maintenance on the cluster is performed. Specify as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example: sun:23:00-mon:01:30.
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// The compute and memory capacity of the nodes in the cluster. See AWS documentation on supported node types as well as vertical scaling.
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// The number of replicas to apply to each shard, up to a maximum of 5. Defaults to 1 (i.e. 2 nodes per shard).
	NumReplicasPerShard *float64 `json:"numReplicasPerShard,omitempty" tf:"num_replicas_per_shard,omitempty"`

	// The number of shards in the cluster. Defaults to 1.
	NumShards *float64 `json:"numShards,omitempty" tf:"num_shards,omitempty"`

	// The name of the parameter group associated with the cluster.
	ParameterGroupName *string `json:"parameterGroupName,omitempty" tf:"parameter_group_name,omitempty"`

	// The port number on which each of the nodes accepts connections. Defaults to 6379.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Set of VPC Security Group ID-s to associate with this cluster.
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// Set of shards in this cluster.
	Shards []ShardsObservation `json:"shards,omitempty" tf:"shards,omitempty"`

	// List of ARN-s that uniquely identify RDB snapshot files stored in S3. The snapshot files will be used to populate the new cluster. Object names in the ARN-s cannot contain any commas.
	SnapshotArns []*string `json:"snapshotArns,omitempty" tf:"snapshot_arns,omitempty"`

	// The name of a snapshot from which to restore data into the new cluster.
	SnapshotName *string `json:"snapshotName,omitempty" tf:"snapshot_name,omitempty"`

	// The number of days for which MemoryDB retains automatic snapshots before deleting them. When set to 0, automatic backups are disabled. Defaults to 0.
	SnapshotRetentionLimit *float64 `json:"snapshotRetentionLimit,omitempty" tf:"snapshot_retention_limit,omitempty"`

	// The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your shard. Example: 05:00-09:00.
	SnapshotWindow *string `json:"snapshotWindow,omitempty" tf:"snapshot_window,omitempty"`

	// ARN of the SNS topic to which cluster notifications are sent.
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`

	// The name of the subnet group to be used for the cluster. Defaults to a subnet group consisting of default VPC subnets.
	SubnetGroupName *string `json:"subnetGroupName,omitempty" tf:"subnet_group_name,omitempty"`

	// A flag to enable in-transit encryption on the cluster. When set to false, the acl_name must be open-access. Defaults to true.
	TLSEnabled *bool `json:"tlsEnabled,omitempty" tf:"tls_enabled,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ClusterParameters struct {

	// The name of the Access Control List to associate with the cluster.
	// +kubebuilder:validation:Optional
	ACLName *string `json:"aclName,omitempty" tf:"acl_name,omitempty"`

	// When set to true, the cluster will automatically receive minor engine version upgrades after launch. Defaults to true.
	// +kubebuilder:validation:Optional
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade,omitempty"`

	// Enables data tiering. This option is not supported by all instance types. For more information, see Data tiering.
	// +kubebuilder:validation:Optional
	DataTiering *bool `json:"dataTiering,omitempty" tf:"data_tiering,omitempty"`

	// Description for the cluster.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Version number of the Redis engine to be used for the cluster. Downgrades are not supported.
	// +kubebuilder:validation:Optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`

	// Name of the final cluster snapshot to be created when this resource is deleted. If omitted, no final snapshot will be made.
	// +kubebuilder:validation:Optional
	FinalSnapshotName *string `json:"finalSnapshotName,omitempty" tf:"final_snapshot_name,omitempty"`

	// ARN of the KMS key used to encrypt the cluster at rest.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Reference to a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`

	// Specifies the weekly time range during which maintenance on the cluster is performed. Specify as a range in the format ddd:hh24:mi-ddd:hh24:mi (24H Clock UTC). The minimum maintenance window is a 60 minute period. Example: sun:23:00-mon:01:30.
	// +kubebuilder:validation:Optional
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// The compute and memory capacity of the nodes in the cluster. See AWS documentation on supported node types as well as vertical scaling.
	// +kubebuilder:validation:Optional
	NodeType *string `json:"nodeType,omitempty" tf:"node_type,omitempty"`

	// The number of replicas to apply to each shard, up to a maximum of 5. Defaults to 1 (i.e. 2 nodes per shard).
	// +kubebuilder:validation:Optional
	NumReplicasPerShard *float64 `json:"numReplicasPerShard,omitempty" tf:"num_replicas_per_shard,omitempty"`

	// The number of shards in the cluster. Defaults to 1.
	// +kubebuilder:validation:Optional
	NumShards *float64 `json:"numShards,omitempty" tf:"num_shards,omitempty"`

	// The name of the parameter group associated with the cluster.
	// +kubebuilder:validation:Optional
	ParameterGroupName *string `json:"parameterGroupName,omitempty" tf:"parameter_group_name,omitempty"`

	// The port number on which each of the nodes accepts connections. Defaults to 6379.
	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// References to SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDRefs []v1.Reference `json:"securityGroupIdRefs,omitempty" tf:"-"`

	// Selector for a list of SecurityGroup in ec2 to populate securityGroupIds.
	// +kubebuilder:validation:Optional
	SecurityGroupIDSelector *v1.Selector `json:"securityGroupIdSelector,omitempty" tf:"-"`

	// Set of VPC Security Group ID-s to associate with this cluster.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.SecurityGroup
	// +crossplane:generate:reference:refFieldName=SecurityGroupIDRefs
	// +crossplane:generate:reference:selectorFieldName=SecurityGroupIDSelector
	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// List of ARN-s that uniquely identify RDB snapshot files stored in S3. The snapshot files will be used to populate the new cluster. Object names in the ARN-s cannot contain any commas.
	// +kubebuilder:validation:Optional
	SnapshotArns []*string `json:"snapshotArns,omitempty" tf:"snapshot_arns,omitempty"`

	// The name of a snapshot from which to restore data into the new cluster.
	// +kubebuilder:validation:Optional
	SnapshotName *string `json:"snapshotName,omitempty" tf:"snapshot_name,omitempty"`

	// The number of days for which MemoryDB retains automatic snapshots before deleting them. When set to 0, automatic backups are disabled. Defaults to 0.
	// +kubebuilder:validation:Optional
	SnapshotRetentionLimit *float64 `json:"snapshotRetentionLimit,omitempty" tf:"snapshot_retention_limit,omitempty"`

	// The daily time range (in UTC) during which MemoryDB begins taking a daily snapshot of your shard. Example: 05:00-09:00.
	// +kubebuilder:validation:Optional
	SnapshotWindow *string `json:"snapshotWindow,omitempty" tf:"snapshot_window,omitempty"`

	// ARN of the SNS topic to which cluster notifications are sent.
	// +kubebuilder:validation:Optional
	SnsTopicArn *string `json:"snsTopicArn,omitempty" tf:"sns_topic_arn,omitempty"`

	// The name of the subnet group to be used for the cluster. Defaults to a subnet group consisting of default VPC subnets.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/memorydb/v1beta1.SubnetGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetGroupName *string `json:"subnetGroupName,omitempty" tf:"subnet_group_name,omitempty"`

	// Reference to a SubnetGroup in memorydb to populate subnetGroupName.
	// +kubebuilder:validation:Optional
	SubnetGroupNameRef *v1.Reference `json:"subnetGroupNameRef,omitempty" tf:"-"`

	// Selector for a SubnetGroup in memorydb to populate subnetGroupName.
	// +kubebuilder:validation:Optional
	SubnetGroupNameSelector *v1.Selector `json:"subnetGroupNameSelector,omitempty" tf:"-"`

	// A flag to enable in-transit encryption on the cluster. When set to false, the acl_name must be open-access. Defaults to true.
	// +kubebuilder:validation:Optional
	TLSEnabled *bool `json:"tlsEnabled,omitempty" tf:"tls_enabled,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EndpointInitParameters struct {
}

type EndpointObservation struct {

	// DNS hostname of the cluster configuration endpoint.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// The port number on which each of the nodes accepts connections. Defaults to 6379.
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`
}

type EndpointParameters struct {
}

type NodesInitParameters struct {
}

type NodesObservation struct {

	// The Availability Zone in which the node resides.
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	// The date and time when the node was created. Example: 2022-01-01T21:00:00Z.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	Endpoint []EndpointObservation `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// Name of the cluster. Conflicts with name_prefix.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type NodesParameters struct {
}

type ShardsInitParameters struct {
}

type ShardsObservation struct {

	// Name of the cluster. Conflicts with name_prefix.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Set of nodes in this shard.
	Nodes []NodesObservation `json:"nodes,omitempty" tf:"nodes,omitempty"`

	// Number of individual nodes in this shard.
	NumNodes *float64 `json:"numNodes,omitempty" tf:"num_nodes,omitempty"`

	// Keyspace for this shard. Example: 0-16383.
	Slots *string `json:"slots,omitempty" tf:"slots,omitempty"`
}

type ShardsParameters struct {
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
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
	InitProvider ClusterInitParameters `json:"initProvider,omitempty"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API. Provides a MemoryDB Cluster.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.aclName) || (has(self.initProvider) && has(self.initProvider.aclName))",message="spec.forProvider.aclName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.nodeType) || (has(self.initProvider) && has(self.initProvider.nodeType))",message="spec.forProvider.nodeType is a required parameter"
	Spec   ClusterSpec   `json:"spec"`
	Status ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
