package utils

import (
	"fmt"
)

func SetResponeObjectAlienVault(jsonResponse AlienVaultURL, response *FrontendResponse2) {
	whitelisted := false
	
	for i := 0; i<len(jsonResponse.Validation); i++{
		if jsonResponse.Validation[i].Source == "whitelist"{
		fmt.Println("This is whitelisted")
		whitelisted = true
		}
	} 

	if whitelisted == true{
		response.EN.Status = "Safe"
		response.EN.Content = "Alienvault has whitelisted this domain/URL."
		response.NO.Status = "Trygg"
		response.NO.Content = "Alienvault har hvitelistet dette domenet/URL'en."
	}else if jsonResponse.PulseInfo.Count == 0 {
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

func SetResponeObjectGoogle(jsonResponse GoogleSafeBrowsing, response *FrontendResponse2) {
	if len(jsonResponse.Matches) != 0 {
		response.EN.Content = "This URL has been marked as malicious by Google Safebrowsing, visiting is NOT recommended"
		response.NO.Content = "Denne URLen har blitt markert som ondsinnet av Google Safebrowsing, besøk er IKKE anbefalt"
		switch jsonResponse.Matches[0].ThreatType {
		case "MALWARE":
			response.EN.Status = "Risk"
			response.NO.Status = "Utrygg"

			response.EN.Tags = "MALWARE"
			response.NO.Tags = "SKADEVARE"

		case "SOCIAL_ENGINEERING":
			response.EN.Status = "Risk"
			response.NO.Status = "Utrygg"

			response.EN.Tags = "SOCIAL_ENGINEERING"
			response.NO.Tags = "SOSIAL_MANIPULERING"

		case "UNWANTED_SOFTWARE":
			response.EN.Status = "Risk"
			response.NO.Status = "Utrygg"

			response.EN.Tags = "UNWANTED_SOFTWARE"
			response.NO.Tags = "UØNSKET_PROGRAMVARE"

		default:
			response.EN.Status = "Potentially unsafe"
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

func SetResponeObjectVirusTotal(jsonResponse HybridAnalysisURL, VirusTotal *FrontendResponse2) {
	if jsonResponse.Scanners[0].Status == "clean" {

		VirusTotal.EN.Status = "Safe"
		VirusTotal.EN.Content = "VirusTotal has no information that indicates this URL is malicious"

		VirusTotal.NO.Status = "Trygg"
		VirusTotal.NO.Content = "VirusTotal har ingen informasjon som tilsier at denne URL'en er skadelig."
	} else if jsonResponse.Scanners[0].Status == "malicious" {
		VirusTotal.EN.Status = "Risk"
		VirusTotal.EN.Content = fmt.Sprintf("%d / %d Antivirus agents has detected this URL/Domain as malicious", jsonResponse.Scanners[0].Positives, jsonResponse.Scanners[0].Total)

		VirusTotal.NO.Status = "Utrygg"
		VirusTotal.NO.Content = fmt.Sprintf("%d / %d Antivirus agenter har detektert dette som ondsinnet", jsonResponse.Scanners[0].Positives, jsonResponse.Scanners[0].Total)
	} else if jsonResponse.Scanners[0].Status == "in-queue" {
		VirusTotal.EN.Status = "Awaiting analysis"
		VirusTotal.EN.Content = "Awaiting analysis"

		VirusTotal.NO.Status = "Venter på analyse."
		VirusTotal.NO.Content = "Venter på analyse."

	} else if jsonResponse.Scanners[0].Status == "no-result" {

		VirusTotal.EN.Status = "Safe"
		VirusTotal.EN.Content = "VirusTotal has no information that indicates this URL is malicious"

		VirusTotal.NO.Status = "Trygg"
		VirusTotal.NO.Content = "VirusTotal har ingen informasjon som tilsier at denne URL'en er skadelig."

	} else {
		VirusTotal.EN.Status = "Error"
		VirusTotal.NO.Status = "Error"
	}
}

func SetResponeObjectUrlscanio(jsonResponse HybridAnalysisURL, urlscanio *FrontendResponse2) {
	if jsonResponse.Scanners[1].Status == "clean" || jsonResponse.Scanners[1].Status == "no-classification" {

		urlscanio.EN.Status = "Safe"
		urlscanio.EN.Content = "Urlscan has no information that indicates this URL is malicious"

		urlscanio.NO.Status = "Trygg"
		urlscanio.NO.Content = "Urlscan har ingen informasjon som tilsier at denne URL'en er skadelig."
	} else if jsonResponse.Scanners[1].Status == "malicious" {
		urlscanio.EN.Status = "Risk"
		urlscanio.EN.Content = "Urlscan has detected this URL/Domain as malicious"

		urlscanio.NO.Status = "Utrygg"
		urlscanio.NO.Content = "Urlscan har detektert denne URLen / domenet som skadelig"
	} else if jsonResponse.Scanners[1].Status == "in-queue" {
		urlscanio.EN.Status = "Awaiting analysis"
		urlscanio.EN.Content = "Awaiting analysis"

		urlscanio.NO.Status = "Venter på analyse."
		urlscanio.NO.Content = "Venter på analyse."

	} else {
		urlscanio.EN.Status = "Error"
		urlscanio.NO.Status = "Error"
	}
}

func SetResponseObjectAlienVaultHash(jsonResponse AlienVaultHash, response *FrontendResponse2) {
	if jsonResponse.PulseInfo.Count == 0 || len(jsonResponse.PulseInfo.Related.Other.MalwareFamilies) == 0 {
		response.EN.Status = "Safe"
		response.EN.Content = "We have no information indicating that this is malicious."

		response.NO.Content = "Trygg"
		response.NO.Content = "Vi har ingen informasjon som tyder på at dette er ondsinnet."
	} else {
		response.EN.Status = "Risk"
		response.EN.Content = jsonResponse.PulseInfo.Related.Other.MalwareFamilies[0]

		response.NO.Status = "Risk"
		response.NO.Content = jsonResponse.PulseInfo.Related.Other.MalwareFamilies[0]
	}
}

func SetResponseObjectHybridAnalysisHash(jsonResponse HybridAnalysishash, response *FrontendResponse2) {
	response.SourceName = "Hybrid Analysis"

	if jsonResponse.Verdict == "malicious" {
		response.EN.Status = "Risk"
		response.EN.Content = "This file is malicious"

		response.NO.Status = "Utrygg"
		response.EN.Content = "Denne filen er gjenkjent som ondsinnet"
		//response.SourceName = jsonResponse.Submissions[0].Filename
	} else if jsonResponse.Verdict == "whitelisted" {
		response.EN.Status = "Safe"
		response.EN.Content = "This file is known to be good"

		response.NO.Status = "Trygg"
		response.EN.Content = "Denne filen er hvitelistet av HybridAnalysis - Ikke ondsinnet."
		//response.SourceName = jsonResponse.Submissions[0].Filename
	} else {
		response.EN.Status = "Unknown" //Denne må byttes til at den er ukjent // grå farge elns på frontend.
		response.EN.Content = "This filehash is not known to Hybrid Analysis"

		response.NO.Status = "Ukjent"
		response.NO.Status = "Denne filhashen er ukjent for Hybrid Analysis"
	}

	// Set the filename field if known
	if jsonResponse.Submissions != nil {
		if jsonResponse.Submissions[0].Filename != "" {
			response.EN.Content = response.EN.Content + " " + jsonResponse.Submissions[0].Filename
			response.NO.Content = response.NO.Content + " " + jsonResponse.Submissions[0].Filename
		}
	}
}

func SetResultURL(Responses *ResultFrontendResponse, size int) {

	for i := 0; i <= size-1; i++ {
		if Responses.FrontendResponse[i].EN.Status == "Risk" {
			Responses.EN.Result = "This URL/Domain has been marked as malicious by atleast one of our threat intelligence sources visiting is not reccomended."
			Responses.NO.Result = "Denne URLen/Domenet har blitt markert som ondsinnet av minst en av våre trusseletteretningskilder, besøk er ikke anbefalt."
		}
	}
	if Responses.EN.Result == "" { //If the for loop does not assign a value it means that no agent found this as risky.
		Responses.EN.Result = "We do not have any intelligence indicating that this URL/Domain is malicious."
		Responses.NO.Result = "Vi har ingen informasjon som tilsier at denne URLen/Domenet er ondsinnet"
	}
}

func SetResultHash(Responses *ResultFrontendResponse, size int) {

	for i := 0; i <= size-1; i++ {
		if Responses.FrontendResponse[i].EN.Status == "Risk" {
			Responses.EN.Result = "This filehash has been marked as malicious by atleast one of our threat intelligence sources visiting is not reccomended."
			Responses.NO.Result = "Denne filhashen har blitt markert som ondsinnet av minst en av våre trusseletteretningskilder, besøk er ikke anbefalt."
		}
	}
	if Responses.EN.Result == "" {
		Responses.EN.Result = "We do not have any intelligence indicating that this filehash is malicious."
		Responses.NO.Result = "Vi har ingen informasjon som tilsier at denne filhashen er ondsinnet"
	}
}

func SetGenericError(Response *FrontendResponse2) {

	Response.EN.Status = "ERROR"
	Response.NO.Status = "ERROR"

	Response.EN.Content = "We have encountered an error"
	Response.NO.Content = "Vi har støtt på en error"
}

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
			Response.FrontendResponse[i].EN.Status = "Unsafe"
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
