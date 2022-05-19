package utils

import (
	"fmt"
)

// SetResponseObjectAlienVault takes the AlienVault api response and formats it accroding to our return object struct with translations.
func SetResponseObjectAlienVault(jsonResponse AlienVaultURL, response *FrontendResponse2) {
	whitelisted := false

	for i := 0; i < len(jsonResponse.Validation); i++ {
		if jsonResponse.Validation[i].Source == "whitelist" {	//Check to see if the URL or domain is whitelisted
			whitelisted = true
		}
	}

	if whitelisted {	//If it is whitelisted set SAFE. 
		response.EN.Status = "Safe"
		response.EN.Content = "Alienvault has whitelisted this domain/URL."
		response.NO.Status = "Trygg"
		response.NO.Content = "Alienvault har hvitelistet dette domenet/URL'en."
	} else if jsonResponse.PulseInfo.Count == 0 {
		response.EN.Status = "Safe"
		response.EN.Content = "Alienvault does not have any publicly availible pulses that indicate this is malicious."
		response.NO.Status = "Trygg"
		response.NO.Content = "Alienvault har ingen aktive pulser som tyder på at dette er ondsinnet."

	} else {
		response.EN.Status = "Risk"
		response.EN.Content = "Alienvault has recorded Pulses on the URL/Domain indicating that this might be malicious"

		response.NO.Status = "Utrygg"
		response.NO.Content = "Alienvault har pulser hvor denne URLen/Domenet er nevnt, dette indikerer at dette kan være ondsinnet"
	}
	response.EN.Tags = "N/A"
	response.NO.Tags = "N/A"

	response.SourceName = "AlienVault"
}

// SetResponseObjectGoogle takes the Google Safebrowsing api response and formats it accroding to our return object struct with translations.
func SetResponeObjectGoogle(jsonResponse GoogleSafeBrowsing, response *FrontendResponse2) {
	if len(jsonResponse.Matches) != 0 {
		response.EN.Content = "This URL has been marked as malicious by Google Safebrowsing, visiting is NOT recommended"
		response.NO.Content = "Denne URLen har blitt markert som ondsinnet av Google Safebrowsing, besøk er IKKE anbefalt"
		switch jsonResponse.Matches[0].ThreatType {
		case "MALWARE":		//Contains malware, set risky. 
			response.EN.Status = "Risk"
			response.NO.Status = "Utrygg"

			response.EN.Tags = "MALWARE"
			response.NO.Tags = "SKADEVARE"

		case "SOCIAL_ENGINEERING":	//Social engineering attempt on this page, risky. 
			response.EN.Status = "Risk"
			response.NO.Status = "Utrygg"

			response.EN.Tags = "SOCIAL_ENGINEERING"
			response.NO.Tags = "SOSIAL_MANIPULERING"

		case "UNWANTED_SOFTWARE":	//Unwanted software, risky.
			response.EN.Status = "Risk"
			response.NO.Status = "Utrygg"

			response.EN.Tags = "UNWANTED_SOFTWARE"
			response.NO.Tags = "UØNSKET_PROGRAMVARE"

		default:
			response.EN.Status = "Potentially unsafe"	//Catch all potentially unsafe because of limited information. 
			response.EN.Content = "This URL has been marked as suspicious, not recommended to visit."
			response.EN.Tags = "N/A"

			response.NO.Status = "Potentielt utrygg"
			response.NO.Content = "Denne URL'EN har blitt markert som mistenkelig, besøk er IKKE anbefalt."
			response.NO.Tags = "N/A"
		}
	} else {
		response.EN.Status = "Safe"
		response.EN.Content = "Google safebrowsing has no data that indicates this is an unsafe URL/Domain"

		response.NO.Status = "Trygg"
		response.NO.Content = "Google Safebrowsing har ingen data som indikerer at dette er en utrygg URL/Domene"
	}

	response.SourceName = "Google SafeBrowsing Api"
}

// SetResponseObjectVirusTotal takes the VirusTotal reponse object from HybridAnalysis and formats it accroding to our return object struct with translations.
func SetResponseObjectVirusTotal(jsonResponse HybridAnalysisURL, VirusTotal *FrontendResponse2) {
	if jsonResponse.Scanners[0].Status == "clean" {		//If clean, set safe. 

		VirusTotal.EN.Status = "Safe"
		VirusTotal.EN.Content = fmt.Sprintf("%s has no information that indicates this URL is malicious", jsonResponse.Scanners[0].Name)

		VirusTotal.NO.Status = "Trygg"
		VirusTotal.NO.Content = fmt.Sprintf("%s har ingen informasjon som tilsier at denne URL'en er skadelig.", jsonResponse.Scanners[0].Name)
	} else if jsonResponse.Scanners[0].Status == "malicious" {	//If malicious set response to risky. 
		VirusTotal.EN.Status = "Risk"
		VirusTotal.EN.Content = fmt.Sprintf("%d / %d Antivirus agents has detected this URL/Domain as malicious", jsonResponse.Scanners[0].Positives, jsonResponse.Scanners[0].Total)

		VirusTotal.NO.Status = "Utrygg"
		VirusTotal.NO.Content = fmt.Sprintf("%d / %d Antivirus agenter har detektert dette som ondsinnet", jsonResponse.Scanners[0].Positives, jsonResponse.Scanners[0].Total)
	} else if jsonResponse.Scanners[0].Status == "in-queue" {
		VirusTotal.EN.Status = "Awaiting analysis"
		VirusTotal.EN.Content = "Awaiting analysis attempt to refresh in 20 seconds."

		VirusTotal.NO.Status = "Venter på analyse."
		VirusTotal.NO.Content = "Venter på analyse forsøk å laste inn siden på nytt om 20 sekunder."

	} else if jsonResponse.Scanners[0].Status == "no-result" {	//If no result set safe. 

		VirusTotal.EN.Status = "Safe"
		VirusTotal.EN.Content = fmt.Sprintf("%s has no information that indicates this URL is malicious", jsonResponse.Scanners[0].Name)

		VirusTotal.NO.Status = "Trygg"
		VirusTotal.NO.Content = fmt.Sprintf("%s har ingen informasjon som tilsier at denne URL'en er skadelig.", jsonResponse.Scanners[0].Name)

	} else {		//If anything else unexpected set error.
		VirusTotal.EN.Status = "Error"
		VirusTotal.NO.Status = "Error"
	}
}

// SetResponseObjectAlienVault takes the UrlScanio response from HybridAnalysis and formats it accroding to our return object struct with translations.
func SetResponseObjectUrlscanio(jsonResponse HybridAnalysisURL, urlscanio *FrontendResponse2) {
	if jsonResponse.Scanners[1].Status == "clean" || jsonResponse.Scanners[1].Status == "no-classification" || jsonResponse.Scanners[1].Status == "no-result" {	//Incase of any of these outputs set to safe. 

		urlscanio.EN.Status = "Safe"
		urlscanio.EN.Content = fmt.Sprintf("%s has no information that indicates this URL is malicious", jsonResponse.Scanners[1].Name)

		urlscanio.NO.Status = "Trygg"
		urlscanio.NO.Content = fmt.Sprintf("%s har ingen informasjon som tilsier at denne URL'en er skadelig.", jsonResponse.Scanners[1].Name)
	} else if jsonResponse.Scanners[1].Status == "malicious" {		//If malicious set to risk
		urlscanio.EN.Status = "Risk"
		urlscanio.EN.Content = fmt.Sprintf("%s has detected this URL/Domain as malicious", jsonResponse.Scanners[1].Name)

		urlscanio.NO.Status = "Utrygg"
		urlscanio.NO.Content = fmt.Sprintf("%s har detektert denne URLen / domenet som skadelig", jsonResponse.Scanners[1].Name)
	} else if jsonResponse.Scanners[1].Status == "in-queue" {		//If in que, set awaiting analysis
		urlscanio.EN.Status = "Awaiting analysis"
		urlscanio.EN.Content = "Awaiting analysis attempt to refresh in 20 seconds."

		urlscanio.NO.Status = "Venter på analyse."
		urlscanio.NO.Content = "Venter på analyse forsøk å laste inn siden på nytt om 20 sekunder."

	} else {
		urlscanio.EN.Status = "Error"		//Anything else unexpected, set ERROR.
		urlscanio.NO.Status = "Error"
	}
}

// SetResponseObjectVirusTotal takes the Alienvault api response and formats it accroding to our return object struct with translations.
func SetResponseObjectAlienVaultHash(jsonResponse AlienVaultHash, response *FrontendResponse2) {
	if jsonResponse.PulseInfo.Count == 0 || len(jsonResponse.PulseInfo.Related.Other.MalwareFamilies) == 0 {	//Set safe if this is correct
		response.EN.Status = "Safe"
		response.EN.Content = "We have no information indicating that this file is malicious."

		response.NO.Status = "Trygg"
		response.NO.Content = "Vi har ingen informasjon som tyder på at dette er en ondsinnet fil."
	} else {		//Else set malicious 
		response.EN.Status = "Risk"
		response.EN.Tags = "Malicious"
		response.EN.Content = jsonResponse.PulseInfo.Related.Other.MalwareFamilies[0]

		response.NO.Status = "Risk"
		response.NO.Tags = "Ondsinnet"
		response.NO.Content = jsonResponse.PulseInfo.Related.Other.MalwareFamilies[0]
	}
}

// SetResponseObjectVirusTotal takes the HybridAnalysis api response and formats it accroding to our return object struct with translations.
func SetResponseObjectHybridAnalysisHash(jsonResponse HybridAnalysishash, response *FrontendResponse2) {
	response.SourceName = "Hybrid Analysis"

	if len(jsonResponse) >= 1 {		//Check to see if response is not empty. 

		if jsonResponse[0].Verdict == "malicious" {		//Filter data based on different inputs
			response.EN.Status = "Risk"
			response.EN.Content = "This file is recognized as malicious."

			response.NO.Status = "Utrygg"
			response.NO.Content = "Denne filen er gjenkjent som ondsinnet."
		
		} else if jsonResponse[0].Verdict == "whitelisted." {
			response.EN.Status = "Safe"
			response.EN.Content = "This file is known to be good - whitelisted."

			response.NO.Status = "Trygg"
			response.NO.Content = "Denne filen er hvitelistet av HybridAnalysis - Ikke ondsinnet."
			
		} else if jsonResponse[0].Verdict == "no specific threat" {
			response.EN.Status = "Safe"
			response.EN.Content = "According to HybridAnalysis does this file not pose any specific threat."

			response.NO.Status = "Trygg"
			response.NO.Content = "I henhold til informasjon gitt av HybridAnalysis tilsier ikke denne filen noen trussel."
		} else {
			response.EN.Status = "Unknown"
			response.EN.Content = "This file hash is not known to Hybrid Analysis."

			response.NO.Status = "Ukjent"
			response.NO.Content = "Denne filhashen er ukjent for Hybrid Analysis."
		}
		//fmt.Println(jsonResponse[0].Verdict)
		// Set the filename field if known
		if jsonResponse[0].Submissions != nil {
			if jsonResponse[0].Submissions[0].Filename != "" {
				response.EN.Content = response.EN.Content + " filename: " + jsonResponse[0].Submissions[0].Filename
				response.NO.Content = response.NO.Content + " filnavn: " + jsonResponse[0].Submissions[0].Filename

				response.EN.Tags = "Known filename: " + jsonResponse[0].Submissions[0].Filename
				response.NO.Tags = "Kjent filnavn: " + jsonResponse[0].Submissions[0].Filename
			}
		}
	} else {
		response.EN.Status = "Unknown" 
		response.EN.Content = "This file hash is not known to Hybrid Analysis."

		response.NO.Status = "Ukjent"
		response.NO.Content = "Denne filhashen er ukjent for Hybrid Analysis."
	}

}

// Helper function which creates a description of the intelligence result for a url / domain search
func SetResultURL(Responses *ResultFrontendResponse, size int) {

	for i := 0; i <= size-1; i++ {
		if Responses.FrontendResponse[i].EN.Status == "Risk" {		//If any are marked as risk set default risk string. 
			Responses.EN.Result = "This URL/Domain has been marked as malicious by atleast one of our threat intelligence sources visiting is not reccomended."
			Responses.NO.Result = "Denne URLen/Domenet har blitt markert som ondsinnet av minst en av våre trusseletteretningskilder, besøk er ikke anbefalt."
		}
	}
	if Responses.EN.Result == "" { //If the for loop does not assign a value it means that no agent found this as risky.
		Responses.EN.Result = "We do not have any intelligence indicating that this URL/Domain is malicious."
		Responses.NO.Result = "Vi har ingen informasjon som tilsier at denne URLen/Domenet er ondsinnet"
	}
}

// Helper function which creates a description of the intelligence result for a file hash
func SetResultHash(Responses *ResultFrontendResponse, size int) {

	for i := 0; i <= size-1; i++ {
		if Responses.FrontendResponse[i].EN.Status == "Risk" {	//Set default risk string if malicious
			Responses.EN.Result = "This file hash has been marked as malicious by atleast one of our threat intelligence sources, if this file is on the machine we reccomend to delete it and run a full antivirus scan of the machine."
			Responses.NO.Result = "Denne filhashen har blitt markert som ondsinnet av minst en av våre trusseletteretningskilder, hvis du har denne filen på datamaskinen anbefaler vi å slette filen og kjøre en full antivirus skann av maskinen."
		}
	}
	if Responses.EN.Result == "" {				//Set default safe string if for loop has not set it as malicious
		Responses.EN.Result = "We do not have any intelligence indicating that this file is malicious."
		Responses.NO.Result = "Vi har ingen informasjon som tilsier at denne filen er ondsinnet"
	}
}

// // Helper function which creates a generic error response
func SetGenericError(Response *FrontendResponse2) {

	Response.EN.Status = "ERROR"
	Response.NO.Status = "ERROR"

	Response.EN.Content = "We have encountered an error"
	Response.NO.Content = "Vi har støtt på en error"
}

// Function which handles translations for a frontendResponse struct of a parameter length
func SetResultFile(Response *ResultFrontendResponse, size int) {
	// tell the input to be translated, use standardized output.
	// probably more edits to be done here, figure out which fields are actually printed out
	// type-unsupported, timeout == not relevant to show
	for i := 0; i < size; i++ {
		switch Response.FrontendResponse[i].EN.Status {
		case "undetected":
			Response.FrontendResponse[i].EN.Status = "Safe"
			Response.FrontendResponse[i].EN.Content = "This file has not been marked as malicious, and can be considered safe"

			Response.FrontendResponse[i].NO.Status = "Trygg"
			Response.FrontendResponse[i].NO.Content = "Denne filen har ikke blitt merket som mistenksom, og kan vurderes som trygg"
		case "malicious", "suspicious":
			Response.FrontendResponse[i].EN.Status = "Risk"
			Response.FrontendResponse[i].EN.Content = "This file has been marked as malicious by known sources, it is advised not to interact with this file."

			Response.FrontendResponse[i].NO.Status = "Utrygg"
			Response.FrontendResponse[i].NO.Content = "Denne filen er markert som utrygg basert på kjente kilder, det anbefales å ikke videre behandle denne filen"
		case "harmless":
			Response.FrontendResponse[i].EN.Status = "Confirmed safe"
			Response.FrontendResponse[i].EN.Content = "This file has been marked as benign, based on known sources. Further handling of this file is considered safe"

			Response.FrontendResponse[i].NO.Status = "Bekreftet trygg"
			Response.FrontendResponse[i].NO.Content = "Denne filen har blitt bekreftet som godartet, basert på kjente kilder. Håndtering av denne filen er trygt."
		}

	}
}
