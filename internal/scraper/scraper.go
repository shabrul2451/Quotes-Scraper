package scraper

import (
	"strings"
	"time"
	"web_Scraping/internal/storage"

	"github.com/PuerkitoBio/goquery"
	"go.mongodb.org/mongo-driver/mongo"
)

func ScrapeQuotes(url string) ([]storage.Quote, error) {
	var quotes []storage.Quote

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}

	doc.Find(".quote").Each(func(i int, s *goquery.Selection) {
		quoteText := strings.TrimSpace(s.Find(".quoteText").Contents().First().Text())
		author := strings.TrimSpace(s.Find(".authorOrTitle").Text())

		quote := storage.Quote{
			Text:      quoteText,
			Author:    author,
			CreatedAt: time.Now(),
		}

		quotes = append(quotes, quote)
	})

	return quotes, nil
}

func ScrapeAndStoreQuotes(url string, client *mongo.Client) error {
	quotes, err := ScrapeQuotes(url)
	if err != nil {
		return err
	}

	for _, quote := range quotes {
		if !storage.QuoteExists(client, quote) {
			err := storage.StoreQuote(client, quote)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
