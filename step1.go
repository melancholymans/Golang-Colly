package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

type articleInfo struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func saveArticlesJson(fName string, a []articleInfo) {
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
	err = enc.Encode(a)
	if err != nil {
		log.Fatal(err)
	}

	// Struct to json
	b, _ := json.MarshalIndent(a, "", "  ")
	fmt.Println(string(b))
	// fmt.Println(p)
}

func main() {
	// Target URL
	url := "https://cpp-learning.com"

	articles := make([]articleInfo, 0, 4)

	// Instantiate default collector
	c := colly.NewCollector()

	i := 0
	// Extract li class="new-entry-item"
	c.OnHTML("li[class=new-entry-item]", func(e *colly.HTMLElement) {
		i++
		fmt.Println(i)
		title := e.ChildText("h3")
		fmt.Println(title)

		link, _ := e.DOM.Find("a[href]").Attr("href")
		fmt.Println(link)

		article := articleInfo{
			Title: title,
			URL:   link,
		}
		articles = append(articles, article)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting URL:", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("StatusCode:", r.StatusCode)
	})
	c.OnError(func(r *colly.Response, err error) {
		log.Println("error:", r.StatusCode, err)
	})
	c.Visit(url)
	c.Wait()
	saveArticlesJson("articles.json", articles)
}
