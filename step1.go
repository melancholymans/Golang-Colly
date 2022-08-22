package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

type pageInfo struct {
	StatusCode int    `json:"statusCode"`
	URL        string `json:"url"`
	Title      string `json:"title"`
}

func savePageJson(fName string, p *pageInfo) {
	// Create json file
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	// Dump json to the standard output
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	err = enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}

	// Struct to json
	b, _ := json.MarshalIndent(p, "", "  ")
	fmt.Println(string(b))
	// fmt.Println(p)
}

func main() {
	// Target URL
	url := "https://cpp-learning.com"

	p := &pageInfo{}
	c := colly.NewCollector()
	c.OnHTML("title", func(e *colly.HTMLElement) {
		p.Title = e.Text
		fmt.Println(e.Text)
	})

	// Before making a request print "Visiting URL: https://XXX"
	c.OnRequest(func(r *colly.Request) {
		p.URL = r.URL.String()
		fmt.Println("Visiting URL:", r.URL.String())
	})

	// After making a request extract status code
	c.OnResponse(func(r *colly.Response) {
		p.StatusCode = r.StatusCode
		fmt.Println("StatusCode:", r.StatusCode)
	})
	c.OnError(func(r *colly.Response, err error) {
		p.StatusCode = r.StatusCode
		log.Println("error:", r.StatusCode, err)
	})

	// Start scraping on https://XXX
	c.Visit(url)

	// Wait until threads are finished
	c.Wait()

	// Save as JSON format
	savePageJson("page.json", p)
}
