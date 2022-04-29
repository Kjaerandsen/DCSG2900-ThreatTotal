package utils

import (
	"fmt"
)

func SetResponeObjectAlienVault(jsonResponse AlienVaultURL, response *FrontendResponse2) {
	if jsonResponse.PulseInfo.Count == 0 {
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
		//VirusTotal.EN.Content ="%d/%d Sources has detected this URL/Domain as malicious", jsonResponse.Scanners[0].Positives, jsonResponse.Scanners[0].Total

		fmt.Printf("%d/%d Sources has detected this URL/Domain as malicious", jsonResponse.Scanners[0].Positives, jsonResponse.Scanners[0].Total)

		VirusTotal.NO.Status = "Utrygg"
		//VirusTotal.NO.Content = jsonResponse.Scanners[0].Positives + "/" + jsonResponse.Scanners[0].Total + " har detektert denne URLen / domenet som skadelig"
	} else if jsonResponse.Scanners[0].Status == "in-queue" {
		VirusTotal.EN.Status = "Awaiting analysis"
		VirusTotal.EN.Content = "Awaiting analysis"

		VirusTotal.NO.Status = "Venter på analyse."
		VirusTotal.NO.Content = "Venter på analyse."

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
