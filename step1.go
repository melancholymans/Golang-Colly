package main

import (
	"fmt"
 
	"github.com/gocolly/colly/v2"
)
 
func main() {
	// Target URL
	url := "https://cpp-learning.com"
 
	// Instantiate default collector
	c := colly.NewCollector()
 
	// Extract title element
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("Title:", e.Text)
	})
 
	// Before making a request print "Visiting URL: https://XXX"
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting URL:", r.URL.String())
	})
 
	// Start scraping on https://XXX
	c.Visit(url)
}