package utils

import (
	"context"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
)

// ScreenshotURL takes a url and returns a screenshot of the page found at the url using a headless chrome instance
func ScreenshotURL(url string, Response *ResultFrontendResponse) {

	var SearchURL string

	// Checks that http or https is included in the search url, adds http if not
	if strings.Contains(url, "https://") {
		SearchURL = url

	} else if strings.Contains(url, "http://") {
		SearchURL = url

	} else {
		SearchURL = "http://" + url
	}

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// capture screenshot of an element
	var screenshotbuf []byte

	// Take the screenshot using the screenScreenshot function
	if err := chromedp.Run(ctx, screenScreenshot(SearchURL, &screenshotbuf)); err != nil {
		log.Fatal(err)
	}

	log.Printf("Took a screenshot of a request url")
	Response.Screenshot = screenshotbuf
}

// Inspired by the emulation example at https://github.com/chromedp/examples/blob/master/emulate/main.go
// elementScreenshot takes a screenshot of a specific element.
// The image is taken with a 4 by 3 aspect ratio and a resolution of 1280*960p
func screenScreenshot(urlstr string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.EmulateViewport(1280, 960),
		chromedp.Navigate(urlstr),
		chromedp.CaptureScreenshot(res),
	}
}
