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
		ch := channels[idx]
		go scrape("https://en.wikipedia.org/wiki/" + year + "_in_film", ch)
	}

	// wait for all routines to finish their work
	waitForResponse(channels)
}

func generateYears() []string {
	return []string{"2011", "2012", "2013", "2014", "2015"}
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
	channels := []chan []string{
		make(chan []string),
		make(chan []string),
		make(chan []string),
		make(chan []string),
		make(chan []string),
	}
	return channels
}

func waitForResponse(channels []chan []string) {
	for i := 0; i < 5; i++ {
		select {
		case movies2011 := <-channels[0]:
			printMovies("2011", movies2011)
		case movies2012 := <-channels[1]:
			printMovies("2012", movies2012)
		case movies2013 := <-channels[2]:
			printMovies("2013", movies2013)
		case movies2014 := <-channels[3]:
			printMovies("2014", movies2014)
		case movies2015 := <-channels[4]:
			printMovies("2015", movies2015)
		}
	}
}

func printMovies(year string, movies []string) {
	for _, movie := range movies {
		logger.Infof(year + "->" + movie)
	}
}
