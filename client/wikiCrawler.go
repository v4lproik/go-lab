package client

import (
	"net/http"
	"net/url"
	"github.com/PuerkitoBio/goquery"
)

type WikiCrawler struct {
	clientWeb *http.Client
	url string

	selector string
}


func NewWikiCrawler(http *http.Client, urlToHit string) *WikiCrawler {

	_, err := url.Parse(urlToHit)
	if err != nil {
		panic(err)
	}

	return &WikiCrawler{http, urlToHit, "table.wikitable i a"}
}

func (s *WikiCrawler) Get() (*goquery.Document, error) {
	res, err := s.clientWeb.Get(s.url)
	if err != nil {
		return nil, err
	}

	return s.getDocument(res), nil
}

func (s *WikiCrawler) getDocument(res *http.Response) *goquery.Document {
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		panic(err)
	}

	return doc
}

func (s *WikiCrawler) ExtractMovie(doc *goquery.Document) []string {
	selection := make([]string, 10)
	doc.Find(s.selector).Each(func(i int, s *goquery.Selection) {
		selection = append(selection, s.Text())
	})
	return selection
	return nil
}
