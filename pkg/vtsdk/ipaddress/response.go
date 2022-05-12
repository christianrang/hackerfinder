package ipaddress

type Response struct {
	Data  Data  `json:"data,omitempty"`
	Error Error `json:"error,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Data struct {
	Attributes Attributes        `json:"attributes"`
	Type       string            `json:"type"`
	Id         string            `json:"id"`
	Links      map[string]string `json:"links"`
}

type Attributes struct {
	RegionalInternetRegistery string                         `json:"regional_internet_registery"`
	Jarm                      string                         `json:"jarm"`
	Network                   string                         `json:"network"`
	LastHttpsCertificateDate  int                            `json:"last_https_certificate_date"`
	Tags                      []string                       `json:"tags"`
	CrowdsourcedContext       []CrowdsourcedContext          `json:"crowdsourced_context"`
	Country                   string                         `json:"country"`
	AsOwner                   string                         `json:"as_owner"`
	LastAnalysisStats         LastAnalysisStats              `json:"last_analysis_stats"`
	Asn                       int                            `json:"asn"`
	WhoIsDate                 int                            `json:"who_is_date"`
	LastAnalysisResults       map[string]LastAnalysisResults `json:"last_analysis_results"`
	Reputation                int                            `json:"reputation"`
	LastModificationDate      int                            `json:"last_modification_date"`
	TotalVotes                LastAnalysisStats              `json:"total_votes"`
	LastHttpsCertificate      LastHttpsCertificate           `json:"last_https_certificate"`
	Continent                 string                         `json:"continent"`
	Whois                     string                         `json:"whois"`
}

type CrowdsourcedContext struct {
	Source    string `json:"source"`
	Title     string `json:"title"`
	Detail    string `json:"detail"`
	Severity  string `json:"severity"`
	Timestamp int    `json:"timestamp"`
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

type CertSignature struct {
	Signature          string `json:"signature"`
	SignatureAlgorithm string `json:"signature_algorithm"`
}

type Validity struct {
	NotAfter  string `json:"not_after"`
	NotBefore string `json:"not_before"`
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
