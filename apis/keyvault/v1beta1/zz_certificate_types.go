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

	// The Type of action to be performed when the lifetime trigger is triggerec. Possible values include AutoRenew and EmailContacts. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	ActionType *string `json:"actionType" tf:"action_type,omitempty"`
}

type CertificateAttributeObservation struct {

	// The create time of the Key Vault Certificate.
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	// whether the Key Vault Certificate is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The expires time of the Key Vault Certificate.
	Expires *string `json:"expires,omitempty" tf:"expires,omitempty"`

	// The not before valid time of the Key Vault Certificate.
	NotBefore *string `json:"notBefore,omitempty" tf:"not_before,omitempty"`

	// The deletion recovery level of the Key Vault Certificate.
	RecoveryLevel *string `json:"recoveryLevel,omitempty" tf:"recovery_level,omitempty"`

	// The recent update time of the Key Vault Certificate.
	Updated *string `json:"updated,omitempty" tf:"updated,omitempty"`
}

type CertificateAttributeParameters struct {
}

type CertificateCertificateObservation struct {
}

type CertificateCertificateParameters struct {

	// The base64-encoded certificate contents. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	ContentsSecretRef v1.SecretKeySelector `json:"contentsSecretRef" tf:"-"`

	// The password associated with the certificate. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`
}

type CertificateObservation struct {

	// A certificate_attribute block as defined below.
	CertificateAttribute []CertificateAttributeObservation `json:"certificateAttribute,omitempty" tf:"certificate_attribute,omitempty"`

	// The raw Key Vault Certificate data represented as a hexadecimal string.
	CertificateData *string `json:"certificateData,omitempty" tf:"certificate_data,omitempty"`

	// The Base64 encoded Key Vault Certificate data.
	CertificateDataBase64 *string `json:"certificateDataBase64,omitempty" tf:"certificate_data_base64,omitempty"`

	// The Key Vault Certificate ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the associated Key Vault Secret.
	SecretID *string `json:"secretId,omitempty" tf:"secret_id,omitempty"`

	// The X509 Thumbprint of the Key Vault Certificate represented as a hexadecimal string.
	Thumbprint *string `json:"thumbprint,omitempty" tf:"thumbprint,omitempty"`

	// The current version of the Key Vault Certificate.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// The Base ID of the Key Vault Certificate.
	VersionlessID *string `json:"versionlessId,omitempty" tf:"versionless_id,omitempty"`

	// The Base ID of the Key Vault Secret.
	VersionlessSecretID *string `json:"versionlessSecretId,omitempty" tf:"versionless_secret_id,omitempty"`
}

type CertificateParameters struct {

	// A certificate block as defined below, used to Import an existing certificate.
	// +kubebuilder:validation:Optional
	Certificate []CertificateCertificateParameters `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// A certificate_policy block as defined below.
	// +kubebuilder:validation:Optional
	CertificatePolicy []CertificatePolicyParameters `json:"certificatePolicy,omitempty" tf:"certificate_policy,omitempty"`

	// The ID of the Key Vault where the Certificate should be created.
	// +crossplane:generate:reference:type=Vault
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// Reference to a Vault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDRef *v1.Reference `json:"keyVaultIdRef,omitempty" tf:"-"`

	// Selector for a Vault to populate keyVaultId.
	// +kubebuilder:validation:Optional
	KeyVaultIDSelector *v1.Selector `json:"keyVaultIdSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CertificatePolicyObservation struct {
}

type CertificatePolicyParameters struct {

	// A issuer_parameters block as defined below.
	// +kubebuilder:validation:Required
	IssuerParameters []IssuerParametersParameters `json:"issuerParameters" tf:"issuer_parameters,omitempty"`

	// A key_properties block as defined below.
	// +kubebuilder:validation:Required
	KeyProperties []KeyPropertiesParameters `json:"keyProperties" tf:"key_properties,omitempty"`

	// A lifetime_action block as defined below.
	// +kubebuilder:validation:Optional
	LifetimeAction []LifetimeActionParameters `json:"lifetimeAction,omitempty" tf:"lifetime_action,omitempty"`

	// A secret_properties block as defined below.
	// +kubebuilder:validation:Required
	SecretProperties []SecretPropertiesParameters `json:"secretProperties" tf:"secret_properties,omitempty"`

	// A x509_certificate_properties block as defined below. Required when certificate block is not specified.
	// +kubebuilder:validation:Optional
	X509CertificateProperties []X509CertificatePropertiesParameters `json:"x509CertificateProperties,omitempty" tf:"x509_certificate_properties,omitempty"`
}

type IssuerParametersObservation struct {
}

type IssuerParametersParameters struct {

	// Specifies the name of the Key Vault Certificate. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type KeyPropertiesObservation struct {
}

type KeyPropertiesParameters struct {

	// Specifies the curve to use when creating an EC key. Possible values are P-256, P-256K, P-384, and P-521. This field will be required in a future release if key_type is EC or EC-HSM. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Curve *string `json:"curve,omitempty" tf:"curve,omitempty"`

	// Is this certificate exportable? Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Exportable *bool `json:"exportable" tf:"exportable,omitempty"`

	// The size of the key used in the certificate. Possible values include 2048, 3072, and 4096 for RSA keys, or 256, 384, and 521 for EC keys. This property is required when using RSA keys. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	KeySize *float64 `json:"keySize,omitempty" tf:"key_size,omitempty"`

	// Specifies the type of key, such as RSA or EC. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	KeyType *string `json:"keyType" tf:"key_type,omitempty"`

	// Is the key reusable? Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	ReuseKey *bool `json:"reuseKey" tf:"reuse_key,omitempty"`
}

type LifetimeActionObservation struct {
}

type LifetimeActionParameters struct {

	// A action block as defined below.
	// +kubebuilder:validation:Required
	Action []ActionParameters `json:"action" tf:"action,omitempty"`

	// A trigger block as defined below.
	// +kubebuilder:validation:Required
	Trigger []TriggerParameters `json:"trigger" tf:"trigger,omitempty"`
}

type SecretPropertiesObservation struct {
}

type SecretPropertiesParameters struct {

	// The Content-Type of the Certificate, such as application/x-pkcs12 for a PFX or application/x-pem-file for a PEM. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	ContentType *string `json:"contentType" tf:"content_type,omitempty"`
}

type SubjectAlternativeNamesObservation struct {
}

type SubjectAlternativeNamesParameters struct {

	// A list of alternative DNS names (FQDNs) identified by the Certificate. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DNSNames []*string `json:"dnsNames,omitempty" tf:"dns_names,omitempty"`

	// A list of email addresses identified by this Certificate. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Emails []*string `json:"emails,omitempty" tf:"emails,omitempty"`

	// A list of User Principal Names identified by the Certificate. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Upns []*string `json:"upns,omitempty" tf:"upns,omitempty"`
}

type TriggerObservation struct {
}

type TriggerParameters struct {

	// The number of days before the Certificate expires that the action associated with this Trigger should run. Changing this forces a new resource to be created. Conflicts with lifetime_percentage.
	// +kubebuilder:validation:Optional
	DaysBeforeExpiry *float64 `json:"daysBeforeExpiry,omitempty" tf:"days_before_expiry,omitempty"`

	// The percentage at which during the Certificates Lifetime the action associated with this Trigger should run. Changing this forces a new resource to be created. Conflicts with days_before_expiry.
	// +kubebuilder:validation:Optional
	LifetimePercentage *float64 `json:"lifetimePercentage,omitempty" tf:"lifetime_percentage,omitempty"`
}

type X509CertificatePropertiesObservation struct {
}

type X509CertificatePropertiesParameters struct {

	// A list of Extended/Enhanced Key Usages. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	ExtendedKeyUsage []*string `json:"extendedKeyUsage,omitempty" tf:"extended_key_usage,omitempty"`

	// A list of uses associated with this Key. Possible values include cRLSign, dataEncipherment, decipherOnly, digitalSignature, encipherOnly, keyAgreement, keyCertSign, keyEncipherment and nonRepudiation and are case-sensitive. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	KeyUsage []*string `json:"keyUsage" tf:"key_usage,omitempty"`

	// The Certificate's Subject. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Subject *string `json:"subject" tf:"subject,omitempty"`

	// A subject_alternative_names block as defined below.
	// +kubebuilder:validation:Optional
	SubjectAlternativeNames []SubjectAlternativeNamesParameters `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`

	// The Certificates Validity Period in Months. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	ValidityInMonths *float64 `json:"validityInMonths" tf:"validity_in_months,omitempty"`
}

// CertificateSpec defines the desired state of Certificate
type CertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateParameters `json:"forProvider"`
}

// CertificateStatus defines the observed state of Certificate.
type CertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Certificate is the Schema for the Certificates API. Manages a Key Vault Certificate.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CertificateSpec   `json:"spec"`
	Status            CertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateList contains a list of Certificates
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}

// Repository type metadata.
var (
	Certificate_Kind             = "Certificate"
	Certificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Certificate_Kind}.String()
	Certificate_KindAPIVersion   = Certificate_Kind + "." + CRDGroupVersion.String()
	Certificate_GroupVersionKind = CRDGroupVersion.WithKind(Certificate_Kind)
)

func init() {
	SchemeBuilder.Register(&Certificate{}, &CertificateList{})
}
