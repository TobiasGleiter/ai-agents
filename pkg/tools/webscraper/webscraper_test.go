package webscraper

import (
    "testing"
)

func TestNewScraper(t *testing.T) {
	url := "https://example.com"

	scraper := NewWebScraper(url)

	if scraper.URL != url {
		t.Fatalf(`Error creating scraper`)
	}
}

func TestExtractHeadlines(t *testing.T) {
	url := "https://example.com"

	scraper := NewWebScraper(url)

	scraper.FetchWebsiteHTML()
	headlines := scraper.ExtractHeadlines()

	if headlines[0] != "Example Domain" {
		t.Fatalf(`Error extracting headlines`)
	}
}
