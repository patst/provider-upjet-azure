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

type ActionObservation struct {
}

type ActionParameters struct {

	// A request_header block as defined below.
	// +kubebuilder:validation:Optional
	RequestHeader []RequestHeaderParameters `json:"requestHeader,omitempty" tf:"request_header,omitempty"`

	// A response_header block as defined below.
	// +kubebuilder:validation:Optional
	ResponseHeader []ResponseHeaderParameters `json:"responseHeader,omitempty" tf:"response_header,omitempty"`
}

type FrontdoorRulesEngineObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Location *string `json:"location,omitempty" tf:"location,omitempty"`
}

type FrontdoorRulesEngineParameters struct {

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The name of the Front Door instance. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=FrontDoor
	// +kubebuilder:validation:Optional
	FrontdoorName *string `json:"frontdoorName,omitempty" tf:"frontdoor_name,omitempty"`

	// Reference to a FrontDoor to populate frontdoorName.
	// +kubebuilder:validation:Optional
	FrontdoorNameRef *v1.Reference `json:"frontdoorNameRef,omitempty" tf:"-"`

	// Selector for a FrontDoor to populate frontdoorName.
	// +kubebuilder:validation:Optional
	FrontdoorNameSelector *v1.Selector `json:"frontdoorNameSelector,omitempty" tf:"-"`

	// The name of the resource group. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A rule block as defined below.
	// +kubebuilder:validation:Optional
	Rule []FrontdoorRulesEngineRuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type FrontdoorRulesEngineRuleObservation struct {
}

type FrontdoorRulesEngineRuleParameters struct {

	// A rule_action block as defined below.
	// +kubebuilder:validation:Optional
	Action []ActionParameters `json:"action,omitempty" tf:"action,omitempty"`

	// One or more match_condition block as defined below.
	// +kubebuilder:validation:Optional
	MatchCondition []RuleMatchConditionParameters `json:"matchCondition,omitempty" tf:"match_condition,omitempty"`

	// The name of the rule.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Priority of the rule, must be unique per rules engine definition.
	// +kubebuilder:validation:Required
	Priority *float64 `json:"priority" tf:"priority,omitempty"`
}

type RequestHeaderObservation struct {
}

type RequestHeaderParameters struct {

	// can be set to Overwrite, Append or Delete.
	// +kubebuilder:validation:Optional
	HeaderActionType *string `json:"headerActionType,omitempty" tf:"header_action_type,omitempty"`

	// header name (string).
	// +kubebuilder:validation:Optional
	HeaderName *string `json:"headerName,omitempty" tf:"header_name,omitempty"`

	// value name (string).
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ResponseHeaderObservation struct {
}

type ResponseHeaderParameters struct {

	// can be set to Overwrite, Append or Delete.
	// +kubebuilder:validation:Optional
	HeaderActionType *string `json:"headerActionType,omitempty" tf:"header_action_type,omitempty"`

	// header name (string).
	// +kubebuilder:validation:Optional
	HeaderName *string `json:"headerName,omitempty" tf:"header_name,omitempty"`

	// value name (string).
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type RuleMatchConditionObservation struct {
}

type RuleMatchConditionParameters struct {

	// can be set to true or false to negate the given condition. Defaults to true.
	// +kubebuilder:validation:Optional
	NegateCondition *bool `json:"negateCondition,omitempty" tf:"negate_condition,omitempty"`

	// can be set to Any, IPMatch, GeoMatch, Equal, Contains, LessThan, GreaterThan, LessThanOrEqual, GreaterThanOrEqual, BeginsWith or EndsWith
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// match against a specific key when variable is set to PostArgs or RequestHeader. It cannot be used with QueryString and RequestMethod. Defaults to null.
	// +kubebuilder:validation:Optional
	Selector *string `json:"selector,omitempty" tf:"selector,omitempty"`

	// can be set to one or more values out of Lowercase, RemoveNulls, Trim, Uppercase, UrlDecode and UrlEncode
	// +kubebuilder:validation:Optional
	Transform []*string `json:"transform,omitempty" tf:"transform,omitempty"`

	// value name (string).
	// +kubebuilder:validation:Optional
	Value []*string `json:"value,omitempty" tf:"value,omitempty"`

	// can be set to IsMobile, RemoteAddr, RequestMethod, QueryString, PostArgs, RequestURI, RequestPath, RequestFilename, RequestFilenameExtension,RequestHeader,RequestBody or RequestScheme.
	// +kubebuilder:validation:Optional
	Variable *string `json:"variable,omitempty" tf:"variable,omitempty"`
}

// FrontdoorRulesEngineSpec defines the desired state of FrontdoorRulesEngine
type FrontdoorRulesEngineSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FrontdoorRulesEngineParameters `json:"forProvider"`
}

// FrontdoorRulesEngineStatus defines the observed state of FrontdoorRulesEngine.
type FrontdoorRulesEngineStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FrontdoorRulesEngineObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorRulesEngine is the Schema for the FrontdoorRulesEngines API. Manages an Azure Front Door Rules Engine configuration and rules.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FrontdoorRulesEngine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FrontdoorRulesEngineSpec   `json:"spec"`
	Status            FrontdoorRulesEngineStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FrontdoorRulesEngineList contains a list of FrontdoorRulesEngines
type FrontdoorRulesEngineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FrontdoorRulesEngine `json:"items"`
}

// Repository type metadata.
var (
	FrontdoorRulesEngine_Kind             = "FrontdoorRulesEngine"
	FrontdoorRulesEngine_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FrontdoorRulesEngine_Kind}.String()
	FrontdoorRulesEngine_KindAPIVersion   = FrontdoorRulesEngine_Kind + "." + CRDGroupVersion.String()
	FrontdoorRulesEngine_GroupVersionKind = CRDGroupVersion.WithKind(FrontdoorRulesEngine_Kind)
)

func init() {
	SchemeBuilder.Register(&FrontdoorRulesEngine{}, &FrontdoorRulesEngineList{})
}
