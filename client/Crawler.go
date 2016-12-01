package client

import (
	"github.com/PuerkitoBio/goquery"
)

type Crawler interface {
	Get() (*goquery.Document, error)
	ExtractMovie(*goquery.Document) (error)
}