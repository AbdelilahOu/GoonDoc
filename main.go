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

}
