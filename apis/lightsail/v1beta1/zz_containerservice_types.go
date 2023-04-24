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

type ContainerServiceObservation struct {

	// The Amazon Resource Name (ARN) of the container service.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The Availability Zone. Follows the format us-east-2a (case-sensitive).
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`

	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Same as name.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A Boolean value indicating whether the container service is disabled. Defaults to false.
	IsDisabled *bool `json:"isDisabled,omitempty" tf:"is_disabled,omitempty"`

	// The power specification for the container service. The power specifies the amount of memory,
	// the number of vCPUs, and the monthly price of each node of the container service.
	// Possible values: nano, micro, small, medium, large, xlarge.
	Power *string `json:"power,omitempty" tf:"power,omitempty"`

	// The ID of the power of the container service.
	PowerID *string `json:"powerId,omitempty" tf:"power_id,omitempty"`

	// The principal ARN of the container service. The principal ARN can be used to create a trust
	// relationship between your standard AWS account and your Lightsail container service. This allows you to give your
	// service permission to access resources in your standard AWS account.
	PrincipalArn *string `json:"principalArn,omitempty" tf:"principal_arn,omitempty"`

	// The private domain name of the container service. The private domain name is accessible only
	// by other resources within the default virtual private cloud (VPC) of your Lightsail account.
	PrivateDomainName *string `json:"privateDomainName,omitempty" tf:"private_domain_name,omitempty"`

	// An object to describe the configuration for the container service to access private container image repositories, such as Amazon Elastic Container Registry (Amazon ECR) private repositories. See Private Registry Access below for more details.
	PrivateRegistryAccess []PrivateRegistryAccessObservation `json:"privateRegistryAccess,omitempty" tf:"private_registry_access,omitempty"`

	// The public domain names to use with the container service, such as example.com
	// and www.example.com. You can specify up to four public domain names for a container service. The domain names that you
	// specify are used when you create a deployment with a container configured as the public endpoint of your container
	// service. If you don't specify public domain names, then you can use the default domain of the container service.
	// Defined below.
	PublicDomainNames []PublicDomainNamesObservation `json:"publicDomainNames,omitempty" tf:"public_domain_names,omitempty"`

	// The Lightsail resource type of the container service (i.e., ContainerService).
	ResourceType *string `json:"resourceType,omitempty" tf:"resource_type,omitempty"`

	// The scale specification for the container service. The scale specifies the allocated compute
	// nodes of the container service.
	Scale *float64 `json:"scale,omitempty" tf:"scale,omitempty"`

	// The current state of the container service.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider
	// default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The publicly accessible URL of the container service. If no public endpoint is specified in the
	// currentDeployment, this URL returns a 404 response.
	URL *string `json:"url,omitempty" tf:"url,omitempty"`
}

type ContainerServiceParameters struct {

	// A Boolean value indicating whether the container service is disabled. Defaults to false.
	// +kubebuilder:validation:Optional
	IsDisabled *bool `json:"isDisabled,omitempty" tf:"is_disabled,omitempty"`

	// The power specification for the container service. The power specifies the amount of memory,
	// the number of vCPUs, and the monthly price of each node of the container service.
	// Possible values: nano, micro, small, medium, large, xlarge.
	// +kubebuilder:validation:Optional
	Power *string `json:"power,omitempty" tf:"power,omitempty"`

	// An object to describe the configuration for the container service to access private container image repositories, such as Amazon Elastic Container Registry (Amazon ECR) private repositories. See Private Registry Access below for more details.
	// +kubebuilder:validation:Optional
	PrivateRegistryAccess []PrivateRegistryAccessParameters `json:"privateRegistryAccess,omitempty" tf:"private_registry_access,omitempty"`

	// The public domain names to use with the container service, such as example.com
	// and www.example.com. You can specify up to four public domain names for a container service. The domain names that you
	// specify are used when you create a deployment with a container configured as the public endpoint of your container
	// service. If you don't specify public domain names, then you can use the default domain of the container service.
	// Defined below.
	// +kubebuilder:validation:Optional
	PublicDomainNames []PublicDomainNamesParameters `json:"publicDomainNames,omitempty" tf:"public_domain_names,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The scale specification for the container service. The scale specifies the allocated compute
	// nodes of the container service.
	// +kubebuilder:validation:Optional
	Scale *float64 `json:"scale,omitempty" tf:"scale,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EcrImagePullerRoleObservation struct {

	// A Boolean value that indicates whether to activate the role. The default is false.
	IsActive *bool `json:"isActive,omitempty" tf:"is_active,omitempty"`

	// The principal ARN of the container service. The principal ARN can be used to create a trust
	// relationship between your standard AWS account and your Lightsail container service. This allows you to give your
	// service permission to access resources in your standard AWS account.
	PrincipalArn *string `json:"principalArn,omitempty" tf:"principal_arn,omitempty"`
}

type EcrImagePullerRoleParameters struct {

	// A Boolean value that indicates whether to activate the role. The default is false.
	// +kubebuilder:validation:Optional
	IsActive *bool `json:"isActive,omitempty" tf:"is_active,omitempty"`
}

type PrivateRegistryAccessObservation struct {

	// Describes a request to configure an Amazon Lightsail container service to access private container image repositories, such as Amazon Elastic Container Registry (Amazon ECR) private repositories. See ECR Image Puller Role below for more details.
	EcrImagePullerRole []EcrImagePullerRoleObservation `json:"ecrImagePullerRole,omitempty" tf:"ecr_image_puller_role,omitempty"`
}

type PrivateRegistryAccessParameters struct {

	// Describes a request to configure an Amazon Lightsail container service to access private container image repositories, such as Amazon Elastic Container Registry (Amazon ECR) private repositories. See ECR Image Puller Role below for more details.
	// +kubebuilder:validation:Optional
	EcrImagePullerRole []EcrImagePullerRoleParameters `json:"ecrImagePullerRole,omitempty" tf:"ecr_image_puller_role,omitempty"`
}

type PublicDomainNamesCertificateObservation struct {

	// The name for the container service. Names must be of length 1 to 63, and be
	// unique within each AWS Region in your Lightsail account.
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name,omitempty"`

	DomainNames []*string `json:"domainNames,omitempty" tf:"domain_names,omitempty"`
}

type PublicDomainNamesCertificateParameters struct {

	// The name for the container service. Names must be of length 1 to 63, and be
	// unique within each AWS Region in your Lightsail account.
	// +kubebuilder:validation:Required
	CertificateName *string `json:"certificateName" tf:"certificate_name,omitempty"`

	// +kubebuilder:validation:Required
	DomainNames []*string `json:"domainNames" tf:"domain_names,omitempty"`
}

type PublicDomainNamesObservation struct {
	Certificate []PublicDomainNamesCertificateObservation `json:"certificate,omitempty" tf:"certificate,omitempty"`
}

type PublicDomainNamesParameters struct {

	// +kubebuilder:validation:Required
	Certificate []PublicDomainNamesCertificateParameters `json:"certificate" tf:"certificate,omitempty"`
}

// ContainerServiceSpec defines the desired state of ContainerService
type ContainerServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContainerServiceParameters `json:"forProvider"`
}

// ContainerServiceStatus defines the observed state of ContainerService.
type ContainerServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContainerServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerService is the Schema for the ContainerServices API. Provides a resource to manage Lightsail container service
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ContainerService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.power)",message="power is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.scale)",message="scale is a required parameter"
	Spec   ContainerServiceSpec   `json:"spec"`
	Status ContainerServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerServiceList contains a list of ContainerServices
type ContainerServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerService `json:"items"`
}

// Repository type metadata.
var (
	ContainerService_Kind             = "ContainerService"
	ContainerService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ContainerService_Kind}.String()
	ContainerService_KindAPIVersion   = ContainerService_Kind + "." + CRDGroupVersion.String()
	ContainerService_GroupVersionKind = CRDGroupVersion.WithKind(ContainerService_Kind)
)

func init() {
	SchemeBuilder.Register(&ContainerService{}, &ContainerServiceList{})
}
