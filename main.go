package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/AbdelilahOu/GoonDoc/model"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

func main() {
	// apparently the content isnt generated on the server
	// its dynamically generated so we have to use other packages
	// to process the page with javascript
	start("https://www.youtube.co.uk/")
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
	var videos = []model.YtVideo{}
	//
	vedioCards.Each(func(i int, card *goquery.Selection) {
		// get the vedio image and the title and the channel name also the time it was published
		thumbnailImg := card.Find("#thumbnail yt-image img")
		imageSrc, _ := thumbnailImg.Attr("src")
		channel := card.Find("#details a #avatar img")
		channelImg, _ := channel.Attr("src")
		title := card.Find("#details #meta h3 a yt-formatted-string").Text()
		channelName := card.Find("#details #meta ytd-video-meta-block #metadata #byline-container ytd-channel-name a").Text()
		views := card.Find("#details #meta ytd-video-meta-block #metadata span").First().Text()
		releaseTime := card.Find("#details #meta ytd-video-meta-block #metadata-line span").Last().Text()

		data := model.YtVideo{
			Title:     title,
			Thumbnail: imageSrc,
			ChannelData: model.Channel{
				Name:  channelName,
				Image: channelImg,
			},
			DetailsData: model.Details{
				Views:       views,
				ReleaseTime: releaseTime,
			},
		}

		videos = append(videos, data)
	})

	saveDataToJson(videos)
}

func saveDataToJson(data []model.YtVideo) {
	file, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Create a JSON encoder
	encoder := json.NewEncoder(file)

	// Encode the struct into JSON and write it to the file
	if err := encoder.Encode(model.Result{
		Videos: data,
	}); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println("Data has been written to output.json")
}

// parse dynamic webapp
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
