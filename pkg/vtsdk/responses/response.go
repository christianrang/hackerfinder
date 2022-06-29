package responses

// This file contains generic structs for objects that are used in many
// different responses from VirusTotal

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type LastAnalysisStats struct {
	Harmless   int `json:"harmless"`
	Malicious  int `json:"malicious"`
	Suspicious int `json:"suspicious"`
	Undetected int `json:"undetected"`
	Timeout    int `json:"timeout"`
}

type LastAnalysisResults struct {
	Category   string `json:"category"`
	Result     string `json:"result"`
	Method     string `json:"method"`
	EngineName string `json:"engine_name"`
}

type LastHttpsCertificate struct {
	Size               int               `json:"size"`
	PublicKey          PublicKey         `json:"public_key"`
	ThumbPrintSha256   string            `json:"thumb_print_sha_256"`
	Tags               []string          `json:"tags"`
	CertSignature      CertSignature     `json:"cert_signature"`
	Validity           Validity          `json:"validity"`
	Version            string            `json:"version"`
	Extensions         Extensions        `json:"extensions"`
	SignatureAlgorithm string            `json:"signature_algorithm"`
	SerialNumber       string            `json:"serial_number"`
	Thumbprint         string            `json:"thumbprint"`
	Issuer             map[string]string `json:"issuer"`
	Subject            map[string]string `json:"subject"`
}

type PublicKey struct {
	Keys      map[string]KeyType `json:"keys"`
	Algorithm string             `json:"algorithm"`
}

type KeyType struct {
	KeySize  int    `json:"key_size"`
	Modulus  string `json:"modulus"`
	Exponent string `json:"exponent"`
}

type CertSignature struct {
	Signature          string `json:"signature"`
	SignatureAlgorithm string `json:"signature_algorithm"`
}

type Extensions struct {
	CertificatePolicies    []string          `json:"certificate_policies"`
	ExtendedKeyUsage       []string          `json:"extended_key_usage"`
	AuthorityKeyIdentifier map[string]string `json:"authority_key_identifier"`
	SubjectAlternativeName []string          `json:"subject_alternative_name"`
	Tags                   []string          `json:"tags"`
	SubjectKeyIdentifier   string            `json:"subject_key_identifier"`
	CrlDistributionPoint   []string          `json:"crl_distribution_point"`
	KeyUsage               []string          `json:"key_usage"`
	Ca                     bool              `json:"ca"`
	CaInformationAccess    map[string]string `json:"ca_information_access"`
}

type Validity struct {
	NotAfter  string `json:"not_after"`
	NotBefore string `json:"not_before"`
}
