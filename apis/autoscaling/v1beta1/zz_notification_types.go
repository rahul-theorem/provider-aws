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

type NotificationObservation struct {

	// List of AutoScaling Group Names
	GroupNames []*string `json:"groupNames,omitempty" tf:"group_names,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// List of Notification Types that trigger
	// notifications. Acceptable values are documented in the AWS documentation here
	Notifications []*string `json:"notifications,omitempty" tf:"notifications,omitempty"`

	// Topic ARN for notifications to be sent through
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`
}

type NotificationParameters struct {

	// List of AutoScaling Group Names
	// +kubebuilder:validation:Optional
	GroupNames []*string `json:"groupNames,omitempty" tf:"group_names,omitempty"`

	// List of Notification Types that trigger
	// notifications. Acceptable values are documented in the AWS documentation here
	// +kubebuilder:validation:Optional
	Notifications []*string `json:"notifications,omitempty" tf:"notifications,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Topic ARN for notifications to be sent through
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/sns/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	TopicArn *string `json:"topicArn,omitempty" tf:"topic_arn,omitempty"`

	// Reference to a Topic in sns to populate topicArn.
	// +kubebuilder:validation:Optional
	TopicArnRef *v1.Reference `json:"topicArnRef,omitempty" tf:"-"`

	// Selector for a Topic in sns to populate topicArn.
	// +kubebuilder:validation:Optional
	TopicArnSelector *v1.Selector `json:"topicArnSelector,omitempty" tf:"-"`
}

// NotificationSpec defines the desired state of Notification
type NotificationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NotificationParameters `json:"forProvider"`
}

// NotificationStatus defines the observed state of Notification.
type NotificationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NotificationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Notification is the Schema for the Notifications API. Provides an AutoScaling Group with Notification support
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Notification struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.groupNames)",message="groupNames is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.notifications)",message="notifications is a required parameter"
	Spec   NotificationSpec   `json:"spec"`
	Status NotificationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NotificationList contains a list of Notifications
type NotificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Notification `json:"items"`
}

// Repository type metadata.
var (
	Notification_Kind             = "Notification"
	Notification_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Notification_Kind}.String()
	Notification_KindAPIVersion   = Notification_Kind + "." + CRDGroupVersion.String()
	Notification_GroupVersionKind = CRDGroupVersion.WithKind(Notification_Kind)
)

func init() {
	SchemeBuilder.Register(&Notification{}, &NotificationList{})
}
