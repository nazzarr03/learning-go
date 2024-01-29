package main

import (
	"fmt"
	"time"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		// Visit only domains: www.reddit.com and reddit.com
		colly.AllowedDomains("www.reddit.com", "reddit.com"),
	)

	// Callback for when a scraped page contains an article element
	c.OnHTML(".article-container", func(e *colly.HTMLElement) {
		fmt.Println("Article found: ", e.DOM.Find("h2").Text())
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// AbsoluteURL converts link to an absolute URL
		//Request is a pointer to the current request 
		// Visit link found on page on a new thread
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Limit(&colly.LimitRule{
		DomainGlob: "*",
		RandomDelay: 1 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.reddit.com/r/golang/")
}