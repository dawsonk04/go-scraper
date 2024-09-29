package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	// Create a new collector
	c := colly.NewCollector()

	//Define what happens after visitng webpage
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("Page title:", e.Text)
	})

	// Error Handeling
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	// Start scraping
	err := c.Visit("http://examplewebsite.com")
	if err != nil {
		log.Fatal(err)
	}

	// Example of targeting certain elements on a webpage
	c.OnHTML("a", func(e *colly.HTMLElement) {
		fmt.Println("Link found: ", e.Attr("href"))
	})

}
