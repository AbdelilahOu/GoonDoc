package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	start("https://www.youtube.com/")
}

func start(URL string) {
	// request the page
	res, err := http.Get(URL)
	// check for errors
	if err != nil {
		fmt.Println("error getting the page")
	}
	// close res process is done
	defer res.Body.Close()
	// check if we got a valid res
	if res.StatusCode != 200 {
		fmt.Println("status code error")
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("error reading the page")
	}
	// select all vedio cards
	vedioCards := doc.Find("ytd-rich-grid-row >  #contents > ytd-rich-item-renderer > #content > ytd-rich-grid-media > #dismissible")
	//
	vedioCards.Each(func(i int, card *goquery.Selection) {
		// get the vedio image and the title and the channel name also the time it was published
		cardImageUrl, exists := card.Find("img").Attr("src")
		if exists {
			fmt.Println("image src :", cardImageUrl)
		}
	})
}
