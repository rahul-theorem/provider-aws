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

type ConnectAttachmentObservation struct {

	// The ARN of the attachment.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the attachment.
	AttachmentID *string `json:"attachmentId,omitempty" tf:"attachment_id,omitempty"`

	// The policy rule number associated with the attachment.
	AttachmentPolicyRuleNumber *float64 `json:"attachmentPolicyRuleNumber,omitempty" tf:"attachment_policy_rule_number,omitempty"`

	// The type of attachment.
	AttachmentType *string `json:"attachmentType,omitempty" tf:"attachment_type,omitempty"`

	// The ARN of a core network.
	CoreNetworkArn *string `json:"coreNetworkArn,omitempty" tf:"core_network_arn,omitempty"`

	// The ID of a core network where you want to create the attachment.
	CoreNetworkID *string `json:"coreNetworkId,omitempty" tf:"core_network_id,omitempty"`

	// The Region where the edge is located.
	EdgeLocation *string `json:"edgeLocation,omitempty" tf:"edge_location,omitempty"`

	// The ID of the attachment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Options for creating an attachment.
	Options []OptionsObservation `json:"options,omitempty" tf:"options,omitempty"`

	// The ID of the attachment account owner.
	OwnerAccountID *string `json:"ownerAccountId,omitempty" tf:"owner_account_id,omitempty"`

	// The attachment resource ARN.
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// The name of the segment attachment.
	SegmentName *string `json:"segmentName,omitempty" tf:"segment_name,omitempty"`

	// The state of the attachment.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The ID of the attachment between the two connections.
	TransportAttachmentID *string `json:"transportAttachmentId,omitempty" tf:"transport_attachment_id,omitempty"`
}

type ConnectAttachmentParameters struct {

	// The ID of a core network where you want to create the attachment.
	// +crossplane:generate:reference:type=CoreNetwork
	// +kubebuilder:validation:Optional
	CoreNetworkID *string `json:"coreNetworkId,omitempty" tf:"core_network_id,omitempty"`

	// Reference to a CoreNetwork to populate coreNetworkId.
	// +kubebuilder:validation:Optional
	CoreNetworkIDRef *v1.Reference `json:"coreNetworkIdRef,omitempty" tf:"-"`

	// Selector for a CoreNetwork to populate coreNetworkId.
	// +kubebuilder:validation:Optional
	CoreNetworkIDSelector *v1.Selector `json:"coreNetworkIdSelector,omitempty" tf:"-"`

	// The Region where the edge is located.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkmanager/v1beta1.VPCAttachment
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("edge_location",true)
	// +kubebuilder:validation:Optional
	EdgeLocation *string `json:"edgeLocation,omitempty" tf:"edge_location,omitempty"`

	// Reference to a VPCAttachment in networkmanager to populate edgeLocation.
	// +kubebuilder:validation:Optional
	EdgeLocationRef *v1.Reference `json:"edgeLocationRef,omitempty" tf:"-"`

	// Selector for a VPCAttachment in networkmanager to populate edgeLocation.
	// +kubebuilder:validation:Optional
	EdgeLocationSelector *v1.Selector `json:"edgeLocationSelector,omitempty" tf:"-"`

	// Options for creating an attachment.
	// +kubebuilder:validation:Optional
	Options []OptionsParameters `json:"options,omitempty" tf:"options,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the attachment between the two connections.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkmanager/v1beta1.VPCAttachment
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TransportAttachmentID *string `json:"transportAttachmentId,omitempty" tf:"transport_attachment_id,omitempty"`

	// Reference to a VPCAttachment in networkmanager to populate transportAttachmentId.
	// +kubebuilder:validation:Optional
	TransportAttachmentIDRef *v1.Reference `json:"transportAttachmentIdRef,omitempty" tf:"-"`

	// Selector for a VPCAttachment in networkmanager to populate transportAttachmentId.
	// +kubebuilder:validation:Optional
	TransportAttachmentIDSelector *v1.Selector `json:"transportAttachmentIdSelector,omitempty" tf:"-"`
}

type OptionsObservation struct {
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

type OptionsParameters struct {

	// +kubebuilder:validation:Optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol,omitempty"`
}

// ConnectAttachmentSpec defines the desired state of ConnectAttachment
type ConnectAttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ConnectAttachmentParameters `json:"forProvider"`
}

// ConnectAttachmentStatus defines the observed state of ConnectAttachment.
type ConnectAttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ConnectAttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectAttachment is the Schema for the ConnectAttachments API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ConnectAttachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.options)",message="options is a required parameter"
	Spec   ConnectAttachmentSpec   `json:"spec"`
	Status ConnectAttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectAttachmentList contains a list of ConnectAttachments
type ConnectAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConnectAttachment `json:"items"`
}

// Repository type metadata.
var (
	ConnectAttachment_Kind             = "ConnectAttachment"
	ConnectAttachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConnectAttachment_Kind}.String()
	ConnectAttachment_KindAPIVersion   = ConnectAttachment_Kind + "." + CRDGroupVersion.String()
	ConnectAttachment_GroupVersionKind = CRDGroupVersion.WithKind(ConnectAttachment_Kind)
)

func init() {
	SchemeBuilder.Register(&ConnectAttachment{}, &ConnectAttachmentList{})
}
