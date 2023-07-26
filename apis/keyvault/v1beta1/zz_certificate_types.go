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

type ActionInitParameters struct {

	// The Type of action to be performed when the lifetime trigger is triggerec. Possible values include AutoRenew and EmailContacts. Changing this forces a new resource to be created.
	ActionType *string `json:"actionType,omitempty" tf:"action_type,omitempty"`
}

type ActionObservation struct {

	// The Type of action to be performed when the lifetime trigger is triggerec. Possible values include AutoRenew and EmailContacts. Changing this forces a new resource to be created.
	ActionType *string `json:"actionType,omitempty" tf:"action_type,omitempty"`
}

type ActionParameters struct {

	// The Type of action to be performed when the lifetime trigger is triggerec. Possible values include AutoRenew and EmailContacts. Changing this forces a new resource to be created.
	ActionType *string `json:"actionType,omitempty" tf:"action_type,omitempty"`
}

type CertificateAttributeInitParameters struct {
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

type CertificateCertificateInitParameters struct {
}

type CertificateCertificateObservation struct {
}

type CertificateCertificateParameters struct {

	// The base64-encoded certificate contents.
	// +kubebuilder:validation:Required
	ContentsSecretRef v1.SecretKeySelector `json:"contentsSecretRef" tf:"-"`

	// The password associated with the certificate.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`
}

type CertificateInitParameters struct {

	// A certificate block as defined below, used to Import an existing certificate.
	Certificate []CertificateCertificateInitParameters `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// A certificate_policy block as defined below. Changing this forces a new resource to be created.
	CertificatePolicy []CertificatePolicyInitParameters `json:"certificatePolicy,omitempty" tf:"certificate_policy,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CertificateObservation struct {

	// A certificate block as defined below, used to Import an existing certificate.
	Certificate []CertificateCertificateParameters `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// A certificate_attribute block as defined below.
	CertificateAttribute []CertificateAttributeObservation `json:"certificateAttribute,omitempty" tf:"certificate_attribute,omitempty"`

	// The raw Key Vault Certificate data represented as a hexadecimal string.
	CertificateData *string `json:"certificateData,omitempty" tf:"certificate_data,omitempty"`

	// The Base64 encoded Key Vault Certificate data.
	CertificateDataBase64 *string `json:"certificateDataBase64,omitempty" tf:"certificate_data_base64,omitempty"`

	// A certificate_policy block as defined below. Changing this forces a new resource to be created.
	CertificatePolicy []CertificatePolicyObservation `json:"certificatePolicy,omitempty" tf:"certificate_policy,omitempty"`

	// The Key Vault Certificate ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Key Vault where the Certificate should be created. Changing this forces a new resource to be created.
	KeyVaultID *string `json:"keyVaultId,omitempty" tf:"key_vault_id,omitempty"`

	// The (Versioned) ID for this Key Vault Certificate. This property points to a specific version of a Key Vault Certificate, as such using this won't auto-rotate values if used in other Azure Services.
	ResourceManagerID *string `json:"resourceManagerId,omitempty" tf:"resource_manager_id,omitempty"`

	// The Versionless ID of the Key Vault Certificate. This property allows other Azure Services (that support it) to auto-rotate their value when the Key Vault Certificate is updated.
	ResourceManagerVersionlessID *string `json:"resourceManagerVersionlessId,omitempty" tf:"resource_manager_versionless_id,omitempty"`

	// The ID of the associated Key Vault Secret.
	SecretID *string `json:"secretId,omitempty" tf:"secret_id,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

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
	Certificate []CertificateCertificateParameters `json:"certificate,omitempty" tf:"certificate,omitempty"`

	// A certificate_policy block as defined below. Changing this forces a new resource to be created.
	CertificatePolicy []CertificatePolicyParameters `json:"certificatePolicy,omitempty" tf:"certificate_policy,omitempty"`

	// The ID of the Key Vault where the Certificate should be created. Changing this forces a new resource to be created.
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
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CertificatePolicyInitParameters struct {

	// A issuer_parameters block as defined below.
	IssuerParameters []IssuerParametersInitParameters `json:"issuerParameters,omitempty" tf:"issuer_parameters,omitempty"`

	// A key_properties block as defined below.
	KeyProperties []KeyPropertiesInitParameters `json:"keyProperties,omitempty" tf:"key_properties,omitempty"`

	// A lifetime_action block as defined below.
	LifetimeAction []LifetimeActionInitParameters `json:"lifetimeAction,omitempty" tf:"lifetime_action,omitempty"`

	// A secret_properties block as defined below.
	SecretProperties []SecretPropertiesInitParameters `json:"secretProperties,omitempty" tf:"secret_properties,omitempty"`

	// A x509_certificate_properties block as defined below. Required when certificate block is not specified.
	X509CertificateProperties []X509CertificatePropertiesInitParameters `json:"x509CertificateProperties,omitempty" tf:"x509_certificate_properties,omitempty"`
}

type CertificatePolicyObservation struct {

	// A issuer_parameters block as defined below.
	IssuerParameters []IssuerParametersObservation `json:"issuerParameters,omitempty" tf:"issuer_parameters,omitempty"`

	// A key_properties block as defined below.
	KeyProperties []KeyPropertiesObservation `json:"keyProperties,omitempty" tf:"key_properties,omitempty"`

	// A lifetime_action block as defined below.
	LifetimeAction []LifetimeActionObservation `json:"lifetimeAction,omitempty" tf:"lifetime_action,omitempty"`

	// A secret_properties block as defined below.
	SecretProperties []SecretPropertiesObservation `json:"secretProperties,omitempty" tf:"secret_properties,omitempty"`

	// A x509_certificate_properties block as defined below. Required when certificate block is not specified.
	X509CertificateProperties []X509CertificatePropertiesObservation `json:"x509CertificateProperties,omitempty" tf:"x509_certificate_properties,omitempty"`
}

type CertificatePolicyParameters struct {

	// A issuer_parameters block as defined below.
	IssuerParameters []IssuerParametersParameters `json:"issuerParameters,omitempty" tf:"issuer_parameters,omitempty"`

	// A key_properties block as defined below.
	KeyProperties []KeyPropertiesParameters `json:"keyProperties,omitempty" tf:"key_properties,omitempty"`

	// A lifetime_action block as defined below.
	LifetimeAction []LifetimeActionParameters `json:"lifetimeAction,omitempty" tf:"lifetime_action,omitempty"`

	// A secret_properties block as defined below.
	SecretProperties []SecretPropertiesParameters `json:"secretProperties,omitempty" tf:"secret_properties,omitempty"`

	// A x509_certificate_properties block as defined below. Required when certificate block is not specified.
	X509CertificateProperties []X509CertificatePropertiesParameters `json:"x509CertificateProperties,omitempty" tf:"x509_certificate_properties,omitempty"`
}

type IssuerParametersInitParameters struct {

	// The name of the Certificate Issuer. Possible values include Self (for self-signed certificate), or Unknown (for a certificate issuing authority like Let's Encrypt and Azure direct supported ones). Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type IssuerParametersObservation struct {

	// The name of the Certificate Issuer. Possible values include Self (for self-signed certificate), or Unknown (for a certificate issuing authority like Let's Encrypt and Azure direct supported ones). Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type IssuerParametersParameters struct {

	// The name of the Certificate Issuer. Possible values include Self (for self-signed certificate), or Unknown (for a certificate issuing authority like Let's Encrypt and Azure direct supported ones). Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type KeyPropertiesInitParameters struct {

	// Specifies the curve to use when creating an EC key. Possible values are P-256, P-256K, P-384, and P-521. This field will be required in a future release if key_type is EC or EC-HSM. Changing this forces a new resource to be created.
	Curve *string `json:"curve,omitempty" tf:"curve,omitempty"`

	// Is this certificate exportable? Changing this forces a new resource to be created.
	Exportable *bool `json:"exportable,omitempty" tf:"exportable,omitempty"`

	// The size of the key used in the certificate. Possible values include 2048, 3072, and 4096 for RSA keys, or 256, 384, and 521 for EC keys. This property is required when using RSA keys. Changing this forces a new resource to be created.
	KeySize *float64 `json:"keySize,omitempty" tf:"key_size,omitempty"`

	// Specifies the type of key. Possible values are EC, EC-HSM, RSA, RSA-HSM and oct. Changing this forces a new resource to be created.
	KeyType *string `json:"keyType,omitempty" tf:"key_type,omitempty"`

	// Is the key reusable? Changing this forces a new resource to be created.
	ReuseKey *bool `json:"reuseKey,omitempty" tf:"reuse_key,omitempty"`
}

type KeyPropertiesObservation struct {

	// Specifies the curve to use when creating an EC key. Possible values are P-256, P-256K, P-384, and P-521. This field will be required in a future release if key_type is EC or EC-HSM. Changing this forces a new resource to be created.
	Curve *string `json:"curve,omitempty" tf:"curve,omitempty"`

	// Is this certificate exportable? Changing this forces a new resource to be created.
	Exportable *bool `json:"exportable,omitempty" tf:"exportable,omitempty"`

	// The size of the key used in the certificate. Possible values include 2048, 3072, and 4096 for RSA keys, or 256, 384, and 521 for EC keys. This property is required when using RSA keys. Changing this forces a new resource to be created.
	KeySize *float64 `json:"keySize,omitempty" tf:"key_size,omitempty"`

	// Specifies the type of key. Possible values are EC, EC-HSM, RSA, RSA-HSM and oct. Changing this forces a new resource to be created.
	KeyType *string `json:"keyType,omitempty" tf:"key_type,omitempty"`

	// Is the key reusable? Changing this forces a new resource to be created.
	ReuseKey *bool `json:"reuseKey,omitempty" tf:"reuse_key,omitempty"`
}

type KeyPropertiesParameters struct {

	// Specifies the curve to use when creating an EC key. Possible values are P-256, P-256K, P-384, and P-521. This field will be required in a future release if key_type is EC or EC-HSM. Changing this forces a new resource to be created.
	Curve *string `json:"curve,omitempty" tf:"curve,omitempty"`

	// Is this certificate exportable? Changing this forces a new resource to be created.
	Exportable *bool `json:"exportable,omitempty" tf:"exportable,omitempty"`

	// The size of the key used in the certificate. Possible values include 2048, 3072, and 4096 for RSA keys, or 256, 384, and 521 for EC keys. This property is required when using RSA keys. Changing this forces a new resource to be created.
	KeySize *float64 `json:"keySize,omitempty" tf:"key_size,omitempty"`

	// Specifies the type of key. Possible values are EC, EC-HSM, RSA, RSA-HSM and oct. Changing this forces a new resource to be created.
	KeyType *string `json:"keyType,omitempty" tf:"key_type,omitempty"`

	// Is the key reusable? Changing this forces a new resource to be created.
	ReuseKey *bool `json:"reuseKey,omitempty" tf:"reuse_key,omitempty"`
}

type LifetimeActionInitParameters struct {

	// A action block as defined below.
	Action []ActionInitParameters `json:"action,omitempty" tf:"action,omitempty"`

	// A trigger block as defined below.
	Trigger []TriggerInitParameters `json:"trigger,omitempty" tf:"trigger,omitempty"`
}

type LifetimeActionObservation struct {

	// A action block as defined below.
	Action []ActionObservation `json:"action,omitempty" tf:"action,omitempty"`

	// A trigger block as defined below.
	Trigger []TriggerObservation `json:"trigger,omitempty" tf:"trigger,omitempty"`
}

type LifetimeActionParameters struct {

	// A action block as defined below.
	Action []ActionParameters `json:"action,omitempty" tf:"action,omitempty"`

	// A trigger block as defined below.
	Trigger []TriggerParameters `json:"trigger,omitempty" tf:"trigger,omitempty"`
}

type SecretPropertiesInitParameters struct {

	// The Content-Type of the Certificate, such as application/x-pkcs12 for a PFX or application/x-pem-file for a PEM. Changing this forces a new resource to be created.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`
}

type SecretPropertiesObservation struct {

	// The Content-Type of the Certificate, such as application/x-pkcs12 for a PFX or application/x-pem-file for a PEM. Changing this forces a new resource to be created.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`
}

type SecretPropertiesParameters struct {

	// The Content-Type of the Certificate, such as application/x-pkcs12 for a PFX or application/x-pem-file for a PEM. Changing this forces a new resource to be created.
	ContentType *string `json:"contentType,omitempty" tf:"content_type,omitempty"`
}

type SubjectAlternativeNamesInitParameters struct {

	// A list of alternative DNS names (FQDNs) identified by the Certificate. Changing this forces a new resource to be created.
	DNSNames []*string `json:"dnsNames,omitempty" tf:"dns_names,omitempty"`

	// A list of email addresses identified by this Certificate. Changing this forces a new resource to be created.
	Emails []*string `json:"emails,omitempty" tf:"emails,omitempty"`

	// A list of User Principal Names identified by the Certificate. Changing this forces a new resource to be created.
	Upns []*string `json:"upns,omitempty" tf:"upns,omitempty"`
}

type SubjectAlternativeNamesObservation struct {

	// A list of alternative DNS names (FQDNs) identified by the Certificate. Changing this forces a new resource to be created.
	DNSNames []*string `json:"dnsNames,omitempty" tf:"dns_names,omitempty"`

	// A list of email addresses identified by this Certificate. Changing this forces a new resource to be created.
	Emails []*string `json:"emails,omitempty" tf:"emails,omitempty"`

	// A list of User Principal Names identified by the Certificate. Changing this forces a new resource to be created.
	Upns []*string `json:"upns,omitempty" tf:"upns,omitempty"`
}

type SubjectAlternativeNamesParameters struct {

	// A list of alternative DNS names (FQDNs) identified by the Certificate. Changing this forces a new resource to be created.
	DNSNames []*string `json:"dnsNames,omitempty" tf:"dns_names,omitempty"`

	// A list of email addresses identified by this Certificate. Changing this forces a new resource to be created.
	Emails []*string `json:"emails,omitempty" tf:"emails,omitempty"`

	// A list of User Principal Names identified by the Certificate. Changing this forces a new resource to be created.
	Upns []*string `json:"upns,omitempty" tf:"upns,omitempty"`
}

type TriggerInitParameters struct {

	// The number of days before the Certificate expires that the action associated with this Trigger should run. Changing this forces a new resource to be created. Conflicts with lifetime_percentage.
	DaysBeforeExpiry *float64 `json:"daysBeforeExpiry,omitempty" tf:"days_before_expiry,omitempty"`

	// The percentage at which during the Certificates Lifetime the action associated with this Trigger should run. Changing this forces a new resource to be created. Conflicts with days_before_expiry.
	LifetimePercentage *float64 `json:"lifetimePercentage,omitempty" tf:"lifetime_percentage,omitempty"`
}

type TriggerObservation struct {

	// The number of days before the Certificate expires that the action associated with this Trigger should run. Changing this forces a new resource to be created. Conflicts with lifetime_percentage.
	DaysBeforeExpiry *float64 `json:"daysBeforeExpiry,omitempty" tf:"days_before_expiry,omitempty"`

	// The percentage at which during the Certificates Lifetime the action associated with this Trigger should run. Changing this forces a new resource to be created. Conflicts with days_before_expiry.
	LifetimePercentage *float64 `json:"lifetimePercentage,omitempty" tf:"lifetime_percentage,omitempty"`
}

type TriggerParameters struct {

	// The number of days before the Certificate expires that the action associated with this Trigger should run. Changing this forces a new resource to be created. Conflicts with lifetime_percentage.
	DaysBeforeExpiry *float64 `json:"daysBeforeExpiry,omitempty" tf:"days_before_expiry,omitempty"`

	// The percentage at which during the Certificates Lifetime the action associated with this Trigger should run. Changing this forces a new resource to be created. Conflicts with days_before_expiry.
	LifetimePercentage *float64 `json:"lifetimePercentage,omitempty" tf:"lifetime_percentage,omitempty"`
}

type X509CertificatePropertiesInitParameters struct {

	// A list of Extended/Enhanced Key Usages. Changing this forces a new resource to be created.
	ExtendedKeyUsage []*string `json:"extendedKeyUsage,omitempty" tf:"extended_key_usage,omitempty"`

	// A list of uses associated with this Key. Possible values include cRLSign, dataEncipherment, decipherOnly, digitalSignature, encipherOnly, keyAgreement, keyCertSign, keyEncipherment and nonRepudiation and are case-sensitive. Changing this forces a new resource to be created.
	KeyUsage []*string `json:"keyUsage,omitempty" tf:"key_usage,omitempty"`

	// The Certificate's Subject. Changing this forces a new resource to be created.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// A subject_alternative_names block as defined below. Changing this forces a new resource to be created.
	SubjectAlternativeNames []SubjectAlternativeNamesInitParameters `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`

	// The Certificates Validity Period in Months. Changing this forces a new resource to be created.
	ValidityInMonths *float64 `json:"validityInMonths,omitempty" tf:"validity_in_months,omitempty"`
}

type X509CertificatePropertiesObservation struct {

	// A list of Extended/Enhanced Key Usages. Changing this forces a new resource to be created.
	ExtendedKeyUsage []*string `json:"extendedKeyUsage,omitempty" tf:"extended_key_usage,omitempty"`

	// A list of uses associated with this Key. Possible values include cRLSign, dataEncipherment, decipherOnly, digitalSignature, encipherOnly, keyAgreement, keyCertSign, keyEncipherment and nonRepudiation and are case-sensitive. Changing this forces a new resource to be created.
	KeyUsage []*string `json:"keyUsage,omitempty" tf:"key_usage,omitempty"`

	// The Certificate's Subject. Changing this forces a new resource to be created.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// A subject_alternative_names block as defined below. Changing this forces a new resource to be created.
	SubjectAlternativeNames []SubjectAlternativeNamesObservation `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`

	// The Certificates Validity Period in Months. Changing this forces a new resource to be created.
	ValidityInMonths *float64 `json:"validityInMonths,omitempty" tf:"validity_in_months,omitempty"`
}

type X509CertificatePropertiesParameters struct {

	// A list of Extended/Enhanced Key Usages. Changing this forces a new resource to be created.
	ExtendedKeyUsage []*string `json:"extendedKeyUsage,omitempty" tf:"extended_key_usage,omitempty"`

	// A list of uses associated with this Key. Possible values include cRLSign, dataEncipherment, decipherOnly, digitalSignature, encipherOnly, keyAgreement, keyCertSign, keyEncipherment and nonRepudiation and are case-sensitive. Changing this forces a new resource to be created.
	KeyUsage []*string `json:"keyUsage,omitempty" tf:"key_usage,omitempty"`

	// The Certificate's Subject. Changing this forces a new resource to be created.
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// A subject_alternative_names block as defined below. Changing this forces a new resource to be created.
	SubjectAlternativeNames []SubjectAlternativeNamesParameters `json:"subjectAlternativeNames,omitempty" tf:"subject_alternative_names,omitempty"`

	// The Certificates Validity Period in Months. Changing this forces a new resource to be created.
	ValidityInMonths *float64 `json:"validityInMonths,omitempty" tf:"validity_in_months,omitempty"`
}

// CertificateSpec defines the desired state of Certificate
type CertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	InitProvider CertificateInitParameters `json:"initProvider,omitempty"`
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
