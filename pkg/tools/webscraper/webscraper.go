package webscraper

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {

}

type WebScraper struct {
	URL string
	Doc *html.Node
}

func NewWebScraper(url string) *WebScraper {
	return &WebScraper{URL: url}
}

func (ws *WebScraper) FetchWebsiteHTML() error {
	resp, err := http.Get(ws.URL)
	if err != nil {
		log.Printf("error fetching URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("error: revieved status code: %d", resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body) 
	if err != nil {
		log.Printf("error: parse error: %v", err)
	}

	ws.Doc = doc
	return nil
} 

func (ws *WebScraper) ExtractHeadlines() {
	var extractHeadline func(*html.Node) bool
	extractHeadline = func(n *html.Node) bool {	
		if n.Type == html.ElementNode && n.Data == "h1" {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.TextNode {
					fmt.Println("Headline found:", c.Data)
					return false
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if extractHeadline(c) {
				return true
			}
		}
		return false
	}

	if ws.Doc != nil {
		extractHeadline(ws.Doc)
	}
}

func (ws *WebScraper) ExtractSubtitles() {
	var extractSubtitles func(*html.Node) bool
	extractSubtitles = func(n *html.Node) bool {	
		if n.Type == html.ElementNode && n.Data == "h2" {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.TextNode {
					fmt.Println("Subtitles found:", c.Data)
					return false
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if extractSubtitles(c) {
				return true
			}
		}
		return false
	}

	if ws.Doc != nil {
		extractSubtitles(ws.Doc)
	}
}