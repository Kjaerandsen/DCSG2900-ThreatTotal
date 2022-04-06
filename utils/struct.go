package utils

// HybridAnalysisUrl struct for hybrid analysis api url request
type HybridAnalysisUrl struct {
	SubmissionType string                   `json:"submission_type"`
	Id             string                   `json:"id"`
	Sha256         string                   `json:"sha256"`
	Scanners       []Scanners               `json:"scanners"`
	Whitelist      []map[string]interface{} `json:"whitelist"`
	Reports        []string                 `json:"reports"`
	Finished       bool                     `json:"finished"`
}

// Scanners struct for the scanners in the hybrid analysis api url request
type Scanners struct {
	Name             string                   `json:"name"`
	Status           string                   `json:"status"`
	ErrorMessage     string                   `json:"error_message"`
	Progress         int                      `json:"progress"`
	Total            int                      `json:"total"`
	Positives        int                      `json:"positives"`
	Percent          int                      `json:"percent"`
	AntiVirusResults []map[string]interface{} `json:"anti_virus_results"`
}

// FrontendResponse struct for the response sent to the frontend to be displayed as cards
type FrontendResponse struct {
	ID          int      `json:"id"`
	SourceName  string   `json:"sourceName"`
	Status      string   `json:"status"`
	Content     string   `json:"content"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

type GoogleSafeBrowsing struct {
	Matches []struct {
		ThreatType   string `json:"threatType"`
		PlatformType string `json:"platformType"`
		Threat       struct {
			URL string `json:"url"`
		} `json:"threat"`
		CacheDuration   string `json:"cacheDuration"`
		ThreatEntryType string `json:"threatEntryType"`
	} `json:"matches"`
}

type HybridAnalysishash struct {
	JobID                   string        `json:"job_id"`
	EnvironmentID           int           `json:"environment_id"`
	EnvironmentDescription  string        `json:"environment_description"`
	Size                    int           `json:"size"`
	Type                    string        `json:"type"`
	TypeShort               []string      `json:"type_short"`
	TargetURL               interface{}   `json:"target_url"`
	State                   string        `json:"state"`
	ErrorType               interface{}   `json:"error_type"`
	ErrorOrigin             interface{}   `json:"error_origin"`
	SubmitName              string        `json:"submit_name"`
	Md5                     string        `json:"md5"`
	Sha1                    string        `json:"sha1"`
	Sha256                  string        `json:"sha256"`
	Sha512                  string        `json:"sha512"`
	Ssdeep                  string        `json:"ssdeep"`
	Imphash                 string        `json:"imphash"`
	AvDetect                int           `json:"av_detect"`
	VxFamily                string        `json:"vx_family"`
	URLAnalysis             bool          `json:"url_analysis"`
	AnalysisStartTime       string     `json:"analysis_start_time"`
	ThreatScore             int           `json:"threat_score"`
	Interesting             bool          `json:"interesting"`
	ThreatLevel             int           `json:"threat_level"`
	Verdict                 string        `json:"verdict"`
	Certificates            []interface{} `json:"certificates"`
	Domains                 []interface{} `json:"domains"`
	ClassificationTags      []string      `json:"classification_tags"`
	CompromisedHosts        []interface{} `json:"compromised_hosts"`
	Hosts                   []interface{} `json:"hosts"`
	TotalNetworkConnections int           `json:"total_network_connections"`
	TotalProcesses          int           `json:"total_processes"`
	TotalSignatures         int           `json:"total_signatures"`
	ExtractedFiles          []interface{} `json:"extracted_files"`
	FileMetadata            interface{}   `json:"file_metadata"`
	Processes               []interface{} `json:"processes"`
	Tags                    []string      `json:"tags"`
	MitreAttcks             []struct {
		Tactic                      string        `json:"tactic"`
		Technique                   string        `json:"technique"`
		AttckID                     string        `json:"attck_id"`
		AttckIDWiki                 string        `json:"attck_id_wiki"`
		MaliciousIdentifiersCount   int           `json:"malicious_identifiers_count"`
		MaliciousIdentifiers        []interface{} `json:"malicious_identifiers"`
		SuspiciousIdentifiersCount  int           `json:"suspicious_identifiers_count"`
		SuspiciousIdentifiers       []interface{} `json:"suspicious_identifiers"`
		InformativeIdentifiersCount int           `json:"informative_identifiers_count"`
		InformativeIdentifiers      []interface{} `json:"informative_identifiers"`
		Parent                      struct {
			Technique   string `json:"technique"`
			AttckID     string `json:"attck_id"`
			AttckIDWiki string `json:"attck_id_wiki"`
		} `json:"parent"`
	} `json:"mitre_attcks"`
	Submissions []struct {
		SubmissionID string      `json:"submission_id"`
		Filename     string      `json:"filename"`
		URL          interface{} `json:"url"`
		CreatedAt    string   `json:"created_at"`
	} `json:"submissions"`
	NetworkMode           string        `json:"network_mode"`
	MachineLearningModels []interface{} `json:"machine_learning_models"`
}

type HybridAnalysisURL struct {
	SubmissionType string `json:"submission_type"`
	ID             string `json:"id"`
	Sha256         string `json:"sha256"`
	Scanners       []struct {
		Name             string        `json:"name"`
		Status           string        `json:"status"`
		ErrorMessage     interface{}   `json:"error_message"`
		Progress         int           `json:"progress"`
		Total            int           `json:"total"`
		Positives        int           `json:"positives"`
		Percent          int           `json:"percent"`
		AntiVirusResults []interface{} `json:"anti_virus_results"`
	} `json:"scanners"`
	Whitelist []interface{} `json:"whitelist"`
	Reports   []string      `json:"reports"`
	Finished  bool          `json:"finished"`
}