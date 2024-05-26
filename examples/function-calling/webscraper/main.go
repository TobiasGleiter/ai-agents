package main

import (
	"fmt"
	"log"

	ws "github.com/TobiasGleiter/ai-agents/pkg/tools/webscraper"
)

func main() {
	url := "https://tobiasgleiter.de/en"
	webscraper := ws.NewWebScraper(url)

	if err := webscraper.FetchWebsiteHTML(); err != nil {
		log.Fatalf("Failed to fetch HTML: %v", err)
	}

	headlines := webscraper.ExtractFirstHeadline()
	for i := 0; i < len(headlines); i++ {
		fmt.Println("Headline:", headlines[i])
	}

	subtitles := webscraper.ExtractSubtitles()
	for j := 0; j < len(subtitles); j++ {
		fmt.Println("Subtitle:", subtitles[j])
	}
}
