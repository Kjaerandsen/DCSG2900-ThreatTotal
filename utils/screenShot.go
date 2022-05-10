package utils

import (
	"context"
	"log"
	"strings"

	//"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

/*
func ScreenshotURL(url string, FrondendResponse *ResultFrontendResponse){
	screenshot, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var SearchURL string

	if strings.Contains(url, "https://") {
		SearchURL = url

	} else if strings.Contains(url, "http://") {
		SearchURL = url

	} else {
		SearchURL = "https://" + url
	}


	var imageBuf []byte
	if err := chromedp.Run(screenshot, screenshotTasks(SearchURL, &imageBuf)); err != nil {
		log.Fatal(err)
	}

	FrondendResponse.Screenshot = imageBuf

	if err := ioutil.WriteFile("screenshotTest", imageBuf, 9544); err != nil {
		log.Fatal(err)
	}
}

func screenshotTasks(url string, imageBuf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.ActionFunc(func(screenshot context.Context) (err error) {
			*imageBuf, _, err = page.PrintToPDF().WithPrintBackground(false).Do(screenshot)
			*imageBuf, _, err = page.
			return err
		}),
	}
}

*/
func ScreenshotURL(url string, Response *ResultFrontendResponse) {

	var SearchURL string

	if strings.Contains(url, "https://") {
		SearchURL = url

	} else if strings.Contains(url, "http://") {
		SearchURL = url

	} else {
		SearchURL = "https://" + url
	}

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// capture screenshot of an element
	var screenshotbuf []byte

	/*
	   if err := chromedp.Run(ctx, elementScreenshot(`https://pkg.go.dev/`, `img.Homepage-logo`, &screenshotbuf)); err != nil {
	   	log.Fatal(err)
	   }
	   if err := ioutil.WriteFile("elementScreenshot.png", screenshotbuf, 0o644); err != nil {
	   	log.Fatal(err)
	   }
	*/
	//////

	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, fullScreenshot(SearchURL, 90, &screenshotbuf)); err != nil {
		log.Fatal(err)
	}
	/*
	   if err := ioutil.WriteFile("fullScreenshot.png", screenshotbuf, 0o644); err != nil {
	   	log.Fatal(err)
	   }*/

	log.Printf("wrote elementScreenshot.png and fullScreenshot.png")
	Response.Screenshot = screenshotbuf
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Note: chromedp.FullScreenshot overrides the device's emulation settings. Use
// device.Reset to reset the emulation and viewport settings.
func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}
