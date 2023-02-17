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

type ApplicationInsightsStandardWebTestObservation struct {

	// The ID of the Application Insights Standard WebTest.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Unique ID of this WebTest. This is typically the same value as the Name field.
	SyntheticMonitorID *string `json:"syntheticMonitorId,omitempty" tf:"synthetic_monitor_id,omitempty"`
}

type ApplicationInsightsStandardWebTestParameters struct {

	// The ID of the Application Insights instance on which the WebTest operates. Changing this forces a new Application Insights Standard WebTest to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/insights/v1beta1.ApplicationInsights
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ApplicationInsightsID *string `json:"applicationInsightsId,omitempty" tf:"application_insights_id,omitempty"`

	// Reference to a ApplicationInsights in insights to populate applicationInsightsId.
	// +kubebuilder:validation:Optional
	ApplicationInsightsIDRef *v1.Reference `json:"applicationInsightsIdRef,omitempty" tf:"-"`

	// Selector for a ApplicationInsights in insights to populate applicationInsightsId.
	// +kubebuilder:validation:Optional
	ApplicationInsightsIDSelector *v1.Selector `json:"applicationInsightsIdSelector,omitempty" tf:"-"`

	// Purpose/user defined descriptive test for this WebTest.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Should the WebTest be enabled?
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Interval in seconds between test runs for this WebTest. Valid options are 300, 600 and 900. Defaults to 300.
	// +kubebuilder:validation:Optional
	Frequency *float64 `json:"frequency,omitempty" tf:"frequency,omitempty"`

	// Specifies a list of where to physically run the tests from to give global coverage for accessibility of your application.
	// +kubebuilder:validation:Required
	GeoLocations []*string `json:"geoLocations" tf:"geo_locations,omitempty"`

	// The Azure Region where the Application Insights Standard WebTest should exist. Changing this forces a new Application Insights Standard WebTest to be created. It needs to correlate with location of the parent resource (azurerm_application_insights)
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// A request block as defined below.
	// +kubebuilder:validation:Required
	Request []RequestParameters `json:"request" tf:"request,omitempty"`

	// The name of the Resource Group where the Application Insights Standard WebTest should exist. Changing this forces a new Application Insights Standard WebTest to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Should the retry on WebTest failure be enabled?
	// +kubebuilder:validation:Optional
	RetryEnabled *bool `json:"retryEnabled,omitempty" tf:"retry_enabled,omitempty"`

	// A mapping of tags which should be assigned to the Application Insights Standard WebTest.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Seconds until this WebTest will timeout and fail. Default is 30.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// A validation_rules block as defined below.
	// +kubebuilder:validation:Optional
	ValidationRules []ValidationRulesParameters `json:"validationRules,omitempty" tf:"validation_rules,omitempty"`
}

type ContentObservation struct {
}

type ContentParameters struct {

	// A string value containing the content to match on.
	// +kubebuilder:validation:Required
	ContentMatch *string `json:"contentMatch" tf:"content_match,omitempty"`

	// Ignore the casing in the content_match value.
	// +kubebuilder:validation:Optional
	IgnoreCase *bool `json:"ignoreCase,omitempty" tf:"ignore_case,omitempty"`

	// If the content of content_match is found, pass the test. If set to false, the WebTest is failing if the content of content_match is found.
	// +kubebuilder:validation:Optional
	PassIfTextFound *bool `json:"passIfTextFound,omitempty" tf:"pass_if_text_found,omitempty"`
}

type HeaderObservation struct {
}

type HeaderParameters struct {

	// The name which should be used for this Application Insights Standard WebTest. Changing this forces a new Application Insights Standard WebTest to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The value which should be used for a header in the request.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type RequestObservation struct {
}

type RequestParameters struct {

	// The WebTest request body.
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// Should the following of redirects be enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	FollowRedirectsEnabled *bool `json:"followRedirectsEnabled,omitempty" tf:"follow_redirects_enabled,omitempty"`

	// Which HTTP verb to use for the call. Options are 'GET', 'POST', 'PUT', 'PATCH', and 'DELETE'.
	// +kubebuilder:validation:Optional
	HTTPVerb *string `json:"httpVerb,omitempty" tf:"http_verb,omitempty"`

	// One or more header blocks as defined above.
	// +kubebuilder:validation:Optional
	Header []HeaderParameters `json:"header,omitempty" tf:"header,omitempty"`

	// Should the parsing of dependend requests be enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	ParseDependentRequestsEnabled *bool `json:"parseDependentRequestsEnabled,omitempty" tf:"parse_dependent_requests_enabled,omitempty"`

	// The WebTest request URL.
	// +kubebuilder:validation:Required
	URL *string `json:"url" tf:"url,omitempty"`
}

type ValidationRulesObservation struct {
}

type ValidationRulesParameters struct {

	// A content block as defined above.
	// +kubebuilder:validation:Optional
	Content []ContentParameters `json:"content,omitempty" tf:"content,omitempty"`

	// The expected status code of the response. Default is '200', '0' means 'response code < 400'
	// +kubebuilder:validation:Optional
	ExpectedStatusCode *float64 `json:"expectedStatusCode,omitempty" tf:"expected_status_code,omitempty"`

	// The number of days of SSL certificate validity remaining for the checked endpoint. If the certificate has a shorter remaining lifetime left, the test will fail. This number should be between 1 and 365.
	// +kubebuilder:validation:Optional
	SSLCertRemainingLifetime *float64 `json:"sslCertRemainingLifetime,omitempty" tf:"ssl_cert_remaining_lifetime,omitempty"`

	// Should the SSL check be enabled?
	// +kubebuilder:validation:Optional
	SSLCheckEnabled *bool `json:"sslCheckEnabled,omitempty" tf:"ssl_check_enabled,omitempty"`
}

// ApplicationInsightsStandardWebTestSpec defines the desired state of ApplicationInsightsStandardWebTest
type ApplicationInsightsStandardWebTestSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationInsightsStandardWebTestParameters `json:"forProvider"`
}

// ApplicationInsightsStandardWebTestStatus defines the observed state of ApplicationInsightsStandardWebTest.
type ApplicationInsightsStandardWebTestStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationInsightsStandardWebTestObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationInsightsStandardWebTest is the Schema for the ApplicationInsightsStandardWebTests API. Manages a Application Insights Standard WebTest.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ApplicationInsightsStandardWebTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationInsightsStandardWebTestSpec   `json:"spec"`
	Status            ApplicationInsightsStandardWebTestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationInsightsStandardWebTestList contains a list of ApplicationInsightsStandardWebTests
type ApplicationInsightsStandardWebTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationInsightsStandardWebTest `json:"items"`
}

// Repository type metadata.
var (
	ApplicationInsightsStandardWebTest_Kind             = "ApplicationInsightsStandardWebTest"
	ApplicationInsightsStandardWebTest_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationInsightsStandardWebTest_Kind}.String()
	ApplicationInsightsStandardWebTest_KindAPIVersion   = ApplicationInsightsStandardWebTest_Kind + "." + CRDGroupVersion.String()
	ApplicationInsightsStandardWebTest_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationInsightsStandardWebTest_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationInsightsStandardWebTest{}, &ApplicationInsightsStandardWebTestList{})
}
