package utils

import (
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

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

type FrontendResponse2 struct {
	ID         int    `json:"id"`
	SourceName string `json:"sourceName"`
	EN         struct {
		Status      string `json:"status"`
		Content     string `json:"content"`
		Description string `json:"description"`
		Tags        string `json:"tags"` //fjerner denne fra å være []string for now.
	} `json:"en"`
	NO struct {
		Status      string `json:"status"`
		Content     string `json:"content"`
		Description string `json:"description"`
		Tags        string `json:"tags"`
	} `json:"no"`
}

/*
Lionic struct {

				} `json:"Lionic"`
*/

type FrontendResponse4 struct {
	Category      string      `json:"category"`
	EngineName    string      `json:"engine_name"`
	EngineVersion string      `json:"engine_version"`
	Result        interface{} `json:"result"`
	Method        string      `json:"method"`
	EngineUpdate  string      `json:"engine_update"`
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
	Validation    []ValidationAlienVault `json:"validation"`
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

type ResultFrontendResponse struct {
	EN struct {
		Result string
	}

	NO struct {
		Result string
	}
	Screenshot []byte
	FrontendResponse []FrontendResponse2
}

type VirusTotalUploadID struct {
	Data struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"data"`
}

type FileUploadData2 struct {
	Data struct {
		Attributes struct {
			TypeDescription      string   `json:"type_description"`
			Tlsh                 string   `json:"tlsh"`
			Vhash                string   `json:"vhash"`
			Names                []string `json:"names"`
			LastModificationDate int      `json:"last_modification_date"`
			TypeTag              string   `json:"type_tag"`
			TimesSubmitted       int      `json:"times_submitted"`
			TotalVotes           struct {
				Harmless  int `json:"harmless"`
				Malicious int `json:"malicious"`
			} `json:"total_votes"`
			Size                int                          `json:"size"`
			TypeExtension       string                       `json:"type_extension"`
			LastSubmissionDate  int                          `json:"last_submission_date"`
			LastAnalysisResults map[string]FrontendResponse4 `json:"last_analysis_results"`
			Sha256              string                       `json:"sha256"`
			Tags                []string                     `json:"tags"`
			LastAnalysisDate    int                          `json:"last_analysis_date"`
			UniqueSources       int                          `json:"unique_sources"`
			FirstSubmissionDate int                          `json:"first_submission_date"`
			Ssdeep              string                       `json:"ssdeep"`
			Md5                 string                       `json:"md5"`
			Sha1                string                       `json:"sha1"`
			Magic               string                       `json:"magic"`
			LastAnalysisStats   struct {
				Harmless         int `json:"harmless"`
				TypeUnsupported  int `json:"type-unsupported"`
				Suspicious       int `json:"suspicious"`
				ConfirmedTimeout int `json:"confirmed-timeout"`
				Timeout          int `json:"timeout"`
				Failure          int `json:"failure"`
				Malicious        int `json:"malicious"`
				Undetected       int `json:"undetected"`
			} `json:"last_analysis_stats"`
			MeaningfulName   string `json:"meaningful_name"`
			Reputation       int    `json:"reputation"`
			FirstSeenItwDate int    `json:"first_seen_itw_date"`
		} `json:"attributes"`
		Type  string `json:"type"`
		ID    string `json:"id"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
}

// Struct for storing the oauth2 token, the jwt and the jwt data
type IdAndJwt struct {
	Oauth2Token oauth2.Token           `json:"oauth2Token"`
	Jwt         oidc.IDToken           `json:"jwt"`
	Claims      map[string]interface{} `json:"email"`
}

type ValidationAlienVault struct {
	Source  string `json:"source"`
	Message string `json:"message"`
	Name    string `json:"name"`
}