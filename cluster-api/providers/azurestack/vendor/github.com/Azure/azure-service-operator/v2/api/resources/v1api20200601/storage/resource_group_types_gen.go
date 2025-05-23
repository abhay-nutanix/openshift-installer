// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=resources.azure.com,resources=resourcegroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=resources.azure.com,resources={resourcegroups/status,resourcegroups/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20200601.ResourceGroup
// Generator information:
// - Generated from: /resources/resource-manager/Microsoft.Resources/stable/2020-06-01/resources.json
// - ARM URI: /subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}
type ResourceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceGroup_Spec   `json:"spec,omitempty"`
	Status            ResourceGroup_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ResourceGroup{}

// GetConditions returns the conditions of the resource
func (group *ResourceGroup) GetConditions() conditions.Conditions {
	return group.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (group *ResourceGroup) SetConditions(conditions conditions.Conditions) {
	group.Status.Conditions = conditions
}

var _ configmaps.Exporter = &ResourceGroup{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (group *ResourceGroup) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if group.Spec.OperatorSpec == nil {
		return nil
	}
	return group.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &ResourceGroup{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (group *ResourceGroup) SecretDestinationExpressions() []*core.DestinationExpression {
	if group.Spec.OperatorSpec == nil {
		return nil
	}
	return group.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &ResourceGroup{}

// AzureName returns the Azure name of the resource
func (group *ResourceGroup) AzureName() string {
	return group.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (group ResourceGroup) GetAPIVersion() string {
	return "2020-06-01"
}

// GetResourceScope returns the scope of the resource
func (group *ResourceGroup) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeLocation
}

// GetSpec returns the specification of this resource
func (group *ResourceGroup) GetSpec() genruntime.ConvertibleSpec {
	return &group.Spec
}

// GetStatus returns the status of this resource
func (group *ResourceGroup) GetStatus() genruntime.ConvertibleStatus {
	return &group.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (group *ResourceGroup) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationHead,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Resources/resourceGroups"
func (group *ResourceGroup) GetType() string {
	return "Microsoft.Resources/resourceGroups"
}

// NewEmptyStatus returns a new empty (blank) status
func (group *ResourceGroup) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ResourceGroup_STATUS{}
}

// Owner returns nil as Location scoped resources never have an owner
func (group *ResourceGroup) Owner() *genruntime.ResourceReference {
	return nil
}

// SetStatus sets the status of this resource
func (group *ResourceGroup) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ResourceGroup_STATUS); ok {
		group.Status = *st
		return nil
	}

	// Convert status to required version
	var st ResourceGroup_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	group.Status = st
	return nil
}

var _ genruntime.LocatableResource = &ResourceGroup{}

// Location returns the location of the resource
func (group *ResourceGroup) Location() string {
	if group.Spec.Location == nil {
		return ""
	}
	return *group.Spec.Location
}

// Hub marks that this ResourceGroup is the hub type for conversion
func (group *ResourceGroup) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (group *ResourceGroup) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: group.Spec.OriginalVersion,
		Kind:    "ResourceGroup",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20200601.ResourceGroup
// Generator information:
// - Generated from: /resources/resource-manager/Microsoft.Resources/stable/2020-06-01/resources.json
// - ARM URI: /subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}
type ResourceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceGroup `json:"items"`
}

// Storage version of v1api20200601.APIVersion
// +kubebuilder:validation:Enum={"2020-06-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2020-06-01")

// Storage version of v1api20200601.ResourceGroup_Spec
type ResourceGroup_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                     `json:"azureName,omitempty"`
	Location        *string                    `json:"location,omitempty"`
	ManagedBy       *string                    `json:"managedBy,omitempty"`
	OperatorSpec    *ResourceGroupOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                     `json:"originalVersion,omitempty"`
	PropertyBag     genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	Tags            map[string]string          `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &ResourceGroup_Spec{}

// ConvertSpecFrom populates our ResourceGroup_Spec from the provided source
func (group *ResourceGroup_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == group {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(group)
}

// ConvertSpecTo populates the provided destination from our ResourceGroup_Spec
func (group *ResourceGroup_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == group {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(group)
}

// Storage version of v1api20200601.ResourceGroup_STATUS
// Resource group information.
type ResourceGroup_STATUS struct {
	Conditions  []conditions.Condition          `json:"conditions,omitempty"`
	Id          *string                         `json:"id,omitempty"`
	Location    *string                         `json:"location,omitempty"`
	ManagedBy   *string                         `json:"managedBy,omitempty"`
	Name        *string                         `json:"name,omitempty"`
	Properties  *ResourceGroupProperties_STATUS `json:"properties,omitempty"`
	PropertyBag genruntime.PropertyBag          `json:"$propertyBag,omitempty"`
	Tags        map[string]string               `json:"tags,omitempty"`
	Type        *string                         `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ResourceGroup_STATUS{}

// ConvertStatusFrom populates our ResourceGroup_STATUS from the provided source
func (group *ResourceGroup_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == group {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(group)
}

// ConvertStatusTo populates the provided destination from our ResourceGroup_STATUS
func (group *ResourceGroup_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == group {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(group)
}

// Storage version of v1api20200601.ResourceGroupOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type ResourceGroupOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// Storage version of v1api20200601.ResourceGroupProperties_STATUS
// The resource group properties.
type ResourceGroupProperties_STATUS struct {
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ResourceGroup{}, &ResourceGroupList{})
}
