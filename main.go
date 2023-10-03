package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

func main() {
	// apparently the content isnt generated on the server
	// its dynamically generated so we have to use other packages
	// to process the page with javascript
	start("https://www.youtube.com/")
}

func start(URL string) {
	// Load the HTML document
	doc, err := ParseWebApp(URL)
	if err != nil {
		fmt.Println("error parsing the webapp")
	}
	// select all vedio cards
	if err != nil {
		fmt.Println("error reading the page")
	}
	// select all vedio cards
	vedioCards := doc.Find("ytd-rich-grid-row >  #contents > ytd-rich-item-renderer > #content > ytd-rich-grid-media > #dismissible")
	//
	vedioCards.Each(func(i int, card *goquery.Selection) {
		// get the vedio image and the title and the channel name also the time it was published
		cardImages := card.Find("img")
		imageSrc, exists := cardImages.First().Attr("src")
		if exists {
			fmt.Println("image src :", imageSrc)
		}
	})
}

func ParseWebApp(url string) (*goquery.Document, error) {
	// where to store generated html
	var outterHTML string
	// create ctx
	ctx, cancel := chromedp.NewContext(context.Background())
	// cancel whene we done
	defer cancel()
	//
	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		// js rendering happens asynchronously and this call seems to be enough to account for that
		chromedp.WaitReady(":root"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			// get html
			outterHTML, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			return err
		}),
	}); err != nil {
		return nil, fmt.Errorf("ParseWebApp(): ActionFunc(): %w", err)
	}
	// parse html
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(outterHTML))
	if err != nil {
		return nil, fmt.Errorf("ParseWebApp(): goquery.NewDocumentFromReader(): %w", err)
	}

	return doc, nil
}
