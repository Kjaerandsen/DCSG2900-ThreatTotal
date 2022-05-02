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

type APIresponseResult struct {
	Result       string
	ResponseData []FrontendResponse2
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

type FrontendResponse2 struct {
	ID         int    `json:"id"`
	SourceName string `json:"sourceName"`
	EN         struct {
		Status      string `json:"status"`
		Content     string `json:"content"`
		Description string `json:"description"`
		Tags        string `json:"tags"` //fjerner denne fra å være []string for now.
		Result      string
	} `json:"en"`
	NO struct {
		Status      string `json:"status"`
		Content     string `json:"content"`
		Description string `json:"description"`
		Tags        string `json:"tags"`
		Result      string
	} `json:"no"`
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
	AnalysisStartTime       string        `json:"analysis_start_time"`
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
		CreatedAt    string      `json:"created_at"`
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

type AlienVaultURL struct {
	Sections      []string      `json:"sections"`
	Indicator     string        `json:"indicator"`
	Type          string        `json:"type"`
	TypeTitle     string        `json:"type_title"`
	Validation    []interface{} `json:"validation"`
	BaseIndicator struct {
	} `json:"base_indicator"`
	PulseInfo struct {
		Count      int           `json:"count"`
		Pulses     []interface{} `json:"pulses"`
		References []interface{} `json:"references"`
		Related    struct {
			Alienvault struct {
				Adversary        []interface{} `json:"adversary"`
				MalwareFamilies  []interface{} `json:"malware_families"`
				Industries       []interface{} `json:"industries"`
				UniqueIndicators int           `json:"unique_indicators"`
			} `json:"alienvault"`
			Other struct {
				Adversary        []interface{} `json:"adversary"`
				MalwareFamilies  []interface{} `json:"malware_families"`
				Industries       []interface{} `json:"industries"`
				UniqueIndicators int           `json:"unique_indicators"`
			} `json:"other"`
		} `json:"related"`
	} `json:"pulse_info"`
	FalsePositive []interface{} `json:"false_positive"`
	Alexa         string        `json:"alexa"`
	Whois         string        `json:"whois"`
	Domain        string        `json:"domain"`
	Hostname      string        `json:"hostname"`
}

type AlienVaultHash struct {
	Sections      []string      `json:"sections"`
	Type          string        `json:"type"`
	TypeTitle     string        `json:"type_title"`
	Indicator     string        `json:"indicator"`
	Validation    []interface{} `json:"validation"`
	BaseIndicator struct {
		ID           int64  `json:"id"`
		Indicator    string `json:"indicator"`
		Type         string `json:"type"`
		Title        string `json:"title"`
		Description  string `json:"description"`
		Content      string `json:"content"`
		AccessType   string `json:"access_type"`
		AccessReason string `json:"access_reason"`
	} `json:"base_indicator"`
	PulseInfo struct {
		Count  int `json:"count"`
		Pulses []struct {
			ID                string        `json:"id"`
			Name              string        `json:"name"`
			Description       string        `json:"description"`
			Modified          string        `json:"modified"`
			Created           string        `json:"created"`
			Tags              []interface{} `json:"tags"`
			References        []string      `json:"references"`
			Public            int           `json:"public"`
			Adversary         string        `json:"adversary"`
			TargetedCountries []interface{} `json:"targeted_countries"`
			MalwareFamilies   []struct {
				ID          string      `json:"id"`
				DisplayName string      `json:"display_name"`
				Target      interface{} `json:"target"`
			} `json:"malware_families"`
			AttackIds []struct {
				ID          string `json:"id"`
				Name        string `json:"name"`
				DisplayName string `json:"display_name"`
			} `json:"attack_ids"`
			Industries     []interface{} `json:"industries"`
			Tlp            string        `json:"TLP"`
			ClonedFrom     interface{}   `json:"cloned_from"`
			ExportCount    int           `json:"export_count"`
			UpvotesCount   int           `json:"upvotes_count"`
			DownvotesCount int           `json:"downvotes_count"`
			VotesCount     int           `json:"votes_count"`
			Locked         bool          `json:"locked"`
			PulseSource    string        `json:"pulse_source"`
			ValidatorCount int           `json:"validator_count"`
			CommentCount   int           `json:"comment_count"`
			FollowerCount  int           `json:"follower_count"`
			Vote           int           `json:"vote"`
			Author         struct {
				Username     string `json:"username"`
				ID           string `json:"id"`
				AvatarURL    string `json:"avatar_url"`
				IsSubscribed bool   `json:"is_subscribed"`
				IsFollowing  bool   `json:"is_following"`
			} `json:"author"`
			IndicatorTypeCounts struct {
				FileHashSHA256 int `json:"FileHash-SHA256"`
			} `json:"indicator_type_counts"`
			IndicatorCount           int           `json:"indicator_count"`
			IsAuthor                 bool          `json:"is_author"`
			IsSubscribing            interface{}   `json:"is_subscribing"`
			SubscriberCount          int           `json:"subscriber_count"`
			ModifiedText             string        `json:"modified_text"`
			IsModified               bool          `json:"is_modified"`
			Groups                   []interface{} `json:"groups"`
			InGroup                  bool          `json:"in_group"`
			ThreatHunterScannable    bool          `json:"threat_hunter_scannable"`
			ThreatHunterHasAgents    int           `json:"threat_hunter_has_agents"`
			RelatedIndicatorType     string        `json:"related_indicator_type"`
			RelatedIndicatorIsActive int           `json:"related_indicator_is_active"`
		} `json:"pulses"`
		References []string `json:"references"`
		Related    struct {
			Alienvault struct {
				Adversary       []interface{} `json:"adversary"`
				MalwareFamilies []interface{} `json:"malware_families"`
				Industries      []interface{} `json:"industries"`
			} `json:"alienvault"`
			Other struct {
				Adversary       []interface{} `json:"adversary"`
				MalwareFamilies []string      `json:"malware_families"`
				Industries      []interface{} `json:"industries"`
			} `json:"other"`
		} `json:"related"`
	} `json:"pulse_info"`
	FalsePositive []interface{} `json:"false_positive"`
}

// feideToken response to oauth2token request
type FeideToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	IDToken     string `json:"id_token"`
}

type FeideJWT struct {
	Header struct {
		Typ string `json:"typ"`
		Alg string `json:"alg"`
		Kid string `json:"kid"`
	} `json:"header"`
	Payload struct {
		Iss      string `json:"iss"`
		Jti      string `json:"jti"`
		Aud      string `json:"aud"`
		Sub      string `json:"sub"`
		Iat      int    `json:"iat"`
		Exp      int    `json:"exp"`
		AuthTime int    `json:"auth_time"`
		Email    string `json:"email"`
	} `json:"payload"`
}
