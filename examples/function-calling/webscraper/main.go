package main

import (
	"log"

	ws "github.com/TobiasGleiter/ai-agents/pkg/tools/webscraper"
)

func main() {
	url := "https://tobiasgleiter.de/en"
	webscraper := ws.NewWebScraper(url)

	if err := webscraper.FetchWebsiteHTML(); err != nil {
		log.Fatalf("Failed to fetch HTML: %v", err)
	}

	webscraper.ExtractHeadlines()
	webscraper.ExtractSubtitles()
}
