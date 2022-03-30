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
