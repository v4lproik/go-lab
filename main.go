package main

import (
	"github.com/juju/loggo"
	"github.com/v4lproik/go-lab/client"
	"net/http"
)

var logger = loggo.GetLogger("main")

func init()  {
	loggo.ConfigureLoggers("debug")
}

func banner() {
	var banner = `
	|----------------------------------------------------------|
	|              	     Golang lab - Octo	                   |
	|----------------------------------------------------------|
	`

	logger.Infof(banner)
}

func main() {
	// display banner
	banner()

	years := generateYears()
	channels := makeChannels()

	// launch go routines
	for idx, year := range years {
		//TO COMPLETE - Hint: Find how you can call scrape("https://en.wikipedia.org/wiki/" + year + "_in_film", ch) in a new go routine
	}

	// wait for all routines to finish their work
	waitForResponse(channels)
}

func generateYears() []string {
	//TO COMPLETE - Hint: Find how to initialise an array of strings (use embedded init {})

	return nil
}

func scrape(url string, ch chan []string) {
	scraper := client.NewWikiCrawler(http.DefaultClient, url)
	doc, err := scraper.Get()
	if err != nil {
		logger.Errorf("Crawler was not able to generate a document", err)
	}

	selection := scraper.ExtractMovie(doc)

	ch <- selection
}

func makeChannels() []chan []string {
	//TO COMPLETE - Hint: Find how to initialise an array of channels of type array of strings (use embedded init {})

	return nil
}

func waitForResponse(channels []chan []string) {
	// TO COMPLETE - Hint: Find how to get the result of your goroutines (keyword select)
}

func printMovies(year string, movies []string) {
	// TO COMPLETE - Hint: Write a function that loops over an array of string and display them on stdout (use for and logger.Info)
}
