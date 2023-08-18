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

type ConfigInitParameters struct {

	// Specifies the day that the hours of operation applies to.
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// A end time block specifies the time that your contact center closes. The end_time is documented below.
	EndTime []EndTimeInitParameters `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// A start time block specifies the time that your contact center opens. The start_time is documented below.
	StartTime []StartTimeInitParameters `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type ConfigObservation struct {

	// Specifies the day that the hours of operation applies to.
	Day *string `json:"day,omitempty" tf:"day,omitempty"`

	// A end time block specifies the time that your contact center closes. The end_time is documented below.
	EndTime []EndTimeObservation `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// A start time block specifies the time that your contact center opens. The start_time is documented below.
	StartTime []StartTimeObservation `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type ConfigParameters struct {

	// Specifies the day that the hours of operation applies to.
	// +kubebuilder:validation:Optional
	Day *string `json:"day" tf:"day,omitempty"`

	// A end time block specifies the time that your contact center closes. The end_time is documented below.
	// +kubebuilder:validation:Optional
	EndTime []EndTimeParameters `json:"endTime" tf:"end_time,omitempty"`

	// A start time block specifies the time that your contact center opens. The start_time is documented below.
	// +kubebuilder:validation:Optional
	StartTime []StartTimeParameters `json:"startTime" tf:"start_time,omitempty"`
}

type EndTimeInitParameters struct {

	// Specifies the hour of closing.
	Hours *float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// Specifies the minute of closing.
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`
}

type EndTimeObservation struct {

	// Specifies the hour of closing.
	Hours *float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// Specifies the minute of closing.
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`
}

type EndTimeParameters struct {

	// Specifies the hour of closing.
	// +kubebuilder:validation:Optional
	Hours *float64 `json:"hours" tf:"hours,omitempty"`

	// Specifies the minute of closing.
	// +kubebuilder:validation:Optional
	Minutes *float64 `json:"minutes" tf:"minutes,omitempty"`
}

type HoursOfOperationInitParameters struct {

	// One or more config blocks which define the configuration information for the hours of operation: day, start time, and end time . Config blocks are documented below.
	Config []ConfigInitParameters `json:"config,omitempty" tf:"config,omitempty"`

	// Specifies the description of the Hours of Operation.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the name of the Hours of Operation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the time zone of the Hours of Operation.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type HoursOfOperationObservation struct {

	// The Amazon Resource Name (ARN) of the Hours of Operation.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// One or more config blocks which define the configuration information for the hours of operation: day, start time, and end time . Config blocks are documented below.
	Config []ConfigObservation `json:"config,omitempty" tf:"config,omitempty"`

	// Specifies the description of the Hours of Operation.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Deprecated) The Amazon Resource Name (ARN) of the Hours of Operation.
	HoursOfOperationArn *string `json:"hoursOfOperationArn,omitempty" tf:"hours_of_operation_arn,omitempty"`

	// The identifier for the hours of operation.
	HoursOfOperationID *string `json:"hoursOfOperationId,omitempty" tf:"hours_of_operation_id,omitempty"`

	// The identifier of the hosting Amazon Connect Instance and identifier of the Hours of Operation separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Specifies the name of the Hours of Operation.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Specifies the time zone of the Hours of Operation.
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type HoursOfOperationParameters struct {

	// One or more config blocks which define the configuration information for the hours of operation: day, start time, and end time . Config blocks are documented below.
	// +kubebuilder:validation:Optional
	Config []ConfigParameters `json:"config,omitempty" tf:"config,omitempty"`

	// Specifies the description of the Hours of Operation.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

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

	// Specifies the name of the Hours of Operation.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the time zone of the Hours of Operation.
	// +kubebuilder:validation:Optional
	TimeZone *string `json:"timeZone,omitempty" tf:"time_zone,omitempty"`
}

type StartTimeInitParameters struct {

	// Specifies the hour of opening.
	Hours *float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// Specifies the minute of opening.
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`
}

type StartTimeObservation struct {

	// Specifies the hour of opening.
	Hours *float64 `json:"hours,omitempty" tf:"hours,omitempty"`

	// Specifies the minute of opening.
	Minutes *float64 `json:"minutes,omitempty" tf:"minutes,omitempty"`
}

type StartTimeParameters struct {

	// Specifies the hour of opening.
	// +kubebuilder:validation:Optional
	Hours *float64 `json:"hours" tf:"hours,omitempty"`

	// Specifies the minute of opening.
	// +kubebuilder:validation:Optional
	Minutes *float64 `json:"minutes" tf:"minutes,omitempty"`
}

// HoursOfOperationSpec defines the desired state of HoursOfOperation
type HoursOfOperationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HoursOfOperationParameters `json:"forProvider"`
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
	InitProvider HoursOfOperationInitParameters `json:"initProvider,omitempty"`
}

// HoursOfOperationStatus defines the observed state of HoursOfOperation.
type HoursOfOperationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HoursOfOperationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HoursOfOperation is the Schema for the HoursOfOperations API. Provides details about a specific Amazon Connect Hours of Operation.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type HoursOfOperation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.config) || has(self.initProvider.config)",message="config is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || has(self.initProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.timeZone) || has(self.initProvider.timeZone)",message="timeZone is a required parameter"
	Spec   HoursOfOperationSpec   `json:"spec"`
	Status HoursOfOperationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HoursOfOperationList contains a list of HoursOfOperations
type HoursOfOperationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HoursOfOperation `json:"items"`
}

// Repository type metadata.
var (
	HoursOfOperation_Kind             = "HoursOfOperation"
	HoursOfOperation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HoursOfOperation_Kind}.String()
	HoursOfOperation_KindAPIVersion   = HoursOfOperation_Kind + "." + CRDGroupVersion.String()
	HoursOfOperation_GroupVersionKind = CRDGroupVersion.WithKind(HoursOfOperation_Kind)
)

func init() {
	SchemeBuilder.Register(&HoursOfOperation{}, &HoursOfOperationList{})
}
