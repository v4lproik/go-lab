package main

import (
	"github.com/juju/loggo"
	"github.com/jessevdk/go-flags"
	"github.com/v4lproik/go-lab/client"
	"net/http"
	"strconv"
)

var logger = loggo.GetLogger("main")

type Options struct{
	Threads string `short:"t" long:"thread" description:"Number of threads" required:"true"`
}

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
	// var
	opts := Options{}

	// display banner
	banner()

	// init parser : Pass struct pointer so the init parser can change the data inside the struct
	parser := initParser(&opts)

	// parse cli arguments
	_, err := parser.Parse()
	if err != nil {
		logger.Errorf(err.Error())
		panic(err)
	}

	// initiate channels
	nb, err := strconv.ParseInt(opts.Threads, 10, 32)
	if err != nil {
		logger.Errorf(err.Error())
		panic(err)
	}

	years := generateYears(int(nb))
	channels := makeChannels(int(nb))

	// launch go routines
	for idx, year := range years {
		ch := channels[idx]
		go scrape("https://en.wikipedia.org/wiki/" + year + "_in_film", ch)
	}

	// wait for all routines to finish their work
	waitForResponse(channels)
}

func generateYears(nb int) []string {
	years := make([]string, nb)

	for i := 0; i < nb; i++ {
		years[i] = strconv.Itoa(2015-i)
	}

	return years
}

func initParser(opts *Options) (parser *flags.Parser){
	//default behaviour is HelpFlag | PrintErrors | PassDoubleDash - we need to override the stderr output
	return flags.NewParser(opts, flags.HelpFlag)
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

func makeChannels(nb int) []chan []string {
	channels := make([]chan []string, nb)

	for i := range channels {
		channels[i] = make(chan []string)
	}

	return channels
}

func waitForResponse(channels []chan []string) {
	for i := 0; i < len(channels); i++ {
		printMovies(strconv.Itoa(2015-i), <-channels[i])
	}
}

func printMovies(year string, movies []string) {
	for _, movie := range movies {
		logger.Infof(year + "->" + movie)
	}
}
