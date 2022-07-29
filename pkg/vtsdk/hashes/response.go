package hashes

import (
	"github.com/christianrang/hackerfinder/pkg/vtsdk/responses"
)

type Response struct {
	Data  Data            `json:"data,omitempty"`
	Error responses.Error `json:"error,omitempty"`
}

type Data struct {
	Attributes Attributes      `json:"attributes"`
	Type       string          `json:"type"`
	Id         string          `json:"id"`
	Links      responses.Links `json:"links"`
}

type Attributes struct {
	TypeDescription             string                                   `json:"type_description"`
	Tlsh                        string                                   `json:"tlsh"`
	Vhash                       string                                   `json:"vhash"`
	Trid                        []Trid                                   `json:"trid"`
	SignatureInfo               SignatureInfo                            `json:"signature_info"`
	CreationDate                int64                                    `json:"creation_date"`
	Names                       []string                                 `json:"names"`
	DotNetGuid                  DotNetGuid                               `json:"dot_net_guid"`
	LastModificationDate        int                                      `json:"last_modification_date"`
	TypeTag                     string                                   `json:"type_tag"`
	TimesSubmitted              int                                      `json:"times_submitted"`
	TotalVotes                  responses.TotalVotes                     `json:"total_votes"`
	Size                        int                                      `json:"size"`
	PopularThreatClassification PopularThreatClassification              `json:"popular_threat_classification"`
	Auththentihash              string                                   `json:"auththentihash"`
	LastSubmissionDate          int                                      `json:"last_submission_date"`
	MeaningfulName              string                                   `json:"meaningful_name"`
	SandboxVerdicts             map[string]SandboxVerdicts               `json:"sandbox_verdicts"`
	Sha256                      string                                   `json:"sha256"`
	TypeExtension               string                                   `json:"type_extension"`
	Tags                        []string                                 `json:"tags"`
	LastAnalysisDate            int                                      `json:"last_analysis_date"`
	UniqueSources               int                                      `json:"unique_sources"`
	FirstSubmissionDate         int                                      `json:"first_submission_date"`
	Sha1                        string                                   `json:"sha1"`
	Ssdeep                      string                                   `json:"ssdeep"`
	Packers                     Packers                                  `json:"packers"`
	Md5                         string                                   `json:"md_5"`
	DotNetAssembly              DotNetAssembly                           `json:"dot_net_assembly"`
	PeInfo                      PeInfo                                   `json:"pe_info"`
	Magic                       string                                   `json:"magic"`
	LastAnalysisStats           responses.LastAnalysisStats              `json:"last_analysis_stats"`
	LastAnalysisResults         map[string]responses.LastAnalysisResults `json:"last_analysis_results"`
	Reputation                  int                                      `json:"reputation"`
}

type Trid struct {
	FileType    string  `json:"file_type"`
	Probability float64 `json:"probability"`
}

type SignatureInfo struct {
	Product               string                  `json:"product"`
	Verified              string                  `json:"verified"`
	Description           string                  `json:"description"`
	Copyright             string                  `json:"copyright"`
	SigningDate           string                  `json:"signing date"`
	X509                  []X509                  `json:"X509"`
	Comments              string                  `json:"comments"`
	CounterSignersDetails []CounterSignersDetails `json:"counter signers details"`
	CounterSigners        string                  `json:"counter signers"`
	Signers               string                  `json:"signers"`
	SignersDetails        []SignersDetails        `json:"signers details"`
}

type X509 struct {
	Name         string `json:"name"`
	Algorithm    string `json:"algorithm"`
	ValidFrom    string `json:"valid from"`
	ValidTo      string `json:"valid to"`
	SerialNumber string `json:"serial number"`
	CertIssuer   string `json:"cert issuer"`
	Thumbprint   string `json:"thumbprint"`
	ValidUsage   string `json:"valid_usage"`
}

type CounterSignersDetails struct {
	SignersDetails
}

type SignersDetails struct {
	Status       string `json:"status"`
	ValidUsage   string `json:"valid usage"`
	Name         string `json:"name"`
	Algorithm    string `json:"algorithm"`
	ValidTo      string `json:"valid to"`
	ValidFrom    string `json:"valid from"`
	SerialNumber string `json:"serial number"`
	CertIssuer   string `json:"cert issuer"`
	Thumbprint   string `json:"thumbprint"`
}

type DotNetGuid struct {
	Mvid string `json:"mvid"`
}

type PopularThreatClassification struct {
	SuggestedThreatLabel  string
	PopularThreatCategory []PopularThreat
	PopularThreatName     []PopularThreat
}

type PopularThreat struct {
	Count int
	Value string
}

type SandboxVerdicts struct {
	Category              string   `json:"category"`
	SandboxName           string   `json:"sandbox_name"`
	MalwareClassification []string `json:"malware_classification"`
	Confidence            int      `json:"confidence"`
}

type Packers struct {
	PEiD string `json:"PEiD"`
}

type DotNetAssembly struct {
	StrongnameVa        int                           `json:"strongname_va"`
	ExternalAssemblies  map[string]ExternalAssemblies `json:"external_assemblies"`
	TablesRowsMap       string                        `json:"tables_rows_map"`
	TablesPresentMap    string                        `json:"tables_present_map"`
	ManifestResource    []string                      `json:"manifest_resource"`
	TablesRowsMapLog    string                        `json:"tables_rows_map_log"`
	TypeDefinitionList  []TypeDefinitionList          `json:"type_definition_list"`
	UnmanagedMethodList []UnmanagedMethodList         `json:"unmanaged_method_list"`
	MetadataHeaderRVA   int                           `json:"metadata_header_rva"`
	ExternalModules     []string                      `json:"external_modules"`
	AssemblyFlags       int                           `json:"assembly_flags"`
	AssemblyFlagsTxt    string                        `json:"assembly_flags_txt"`
	EntryPointToken     int                           `json:"entry_point_token"`
	EntryPointRVA       int                           `json:"entry_point_rva"`
	AssemblyName        string                        `json:"assembly_name"`
	ResourcesVA         int                           `json:"resources_va"`
	AssemblyData        AssemblyData                  `json:"assembly_data"`
	Streams             map[string]Stream             `json:"streams"`
	TablesPresent       int                           `json:"tables_present"`
	ClrVersion          string                        `json:"clr_version"`
	UnmanagedMethods    map[string][]string           `json:"unmanaged_methods"`
	ClrMetaVersion      string                        `json:"clr_meta_version"`
}

type ExternalAssemblies struct {
	Version string `json:"version"`
}

type TypeDefinitionList struct {
	TypeDefinitions []string `json:"type_definitions"`
	Namespace       string   `json:"namespace"`
}

type UnmanagedMethodList struct {
	Name    string   `json:"name"`
	Methods []string `json:"methods"`
}

type AssemblyData struct {
	MajorVersion   int    `json:"major_version"`
	MinorVersion   int    `json:"minor_version"`
	Hashalgid      int    `json:"hashalgid"`
	FlagsText      string `json:"flags_text"`
	BuildNumber    int    `json:"build_number"`
	Flags          int    `json:"flags"`
	RevisionNumber int    `json:"revision_number"`
	Name           string `json:"name"`
}

type Stream struct {
	Chi2    float64 `json:"chi2"`
	Size    int     `json:"size"`
	Entropy float64 `json:"entropy"`
	Md5     string  `json:"md5"`
}

type PeInfo struct {
	ResourceDetails []ResourceDetails `json:"resource_details"`
	ResourceTypes   ResourceTypes     `json:"resource_types"`
	ImpHash         string            `json:"imphash"`
	Overlay         ResourceDetails   `json:"overlay"`
	ResourceLangs   map[string]int    `json:"resource_langs"`
	MachineType     int               `json:"machine_type"`
	Timestamp       int               `json:"timestamp"`
	EntryPoint      int               `json:"entry_point"`
	Sections        []Sections        `json:"sections"`
	ImportList      []ImportList      `json:"import_list"`
}

type ResourceDetails struct {
	Lang     string  `json:"lang"`
	Entropy  float64 `json:"entropy"`
	Chi2     float64 `json:"chi2"`
	Filetype string  `json:"filetype"`
	Sha256   string  `json:"sha256"`
	Type     string  `json:"type"`
}

type ResourceTypes struct {
	RtIcon      int `json:"RT_ICON"`
	RtVersion   int `json:"RT_VERSION"`
	RtGroupIcon int `json:"RT_GROUP_ICON"`
}

type Overlay struct {
	ResourceDetails `json:"resource_details"`
}

type Sections struct {
	Name          string  `json:"name"`
	Chi2          float64 `json:"chi2"`
	VitualAddress int     `json:"vitual_address"`
	Entropy       float64 `json:"entropy"`
	RawSize       int     `json:"raw_size"`
	Flags         string  `json:"flags"`
	VirtualSize   int     `json:"virtual_size"`
	Md5           string  `json:"md5"`
}

type ImportList struct {
	LibraryName       string   `json:"library_name"`
	ImportedFunctions []string `json:"imported_functions"`
}
