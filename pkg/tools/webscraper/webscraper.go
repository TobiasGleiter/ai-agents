package webscraper

import (
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

func (ws *WebScraper) ExtractFirstHeadline() []string {
	var headlines []string
	var extractFirstHeadline func(*html.Node) bool
	extractFirstHeadline = func(n *html.Node) bool {
		if n.Type == html.ElementNode && n.Data == "h1" {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.TextNode {
					headlines = append(headlines, c.Data)
					return true
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if extractFirstHeadline(c) {
				return true
			}
		}
		return false
	}

	if ws.Doc != nil {
		if extractFirstHeadline(ws.Doc) {
			return headlines
		}
	}

	return headlines
}


func (ws *WebScraper) ExtractSubtitles() []string {
	var subtitles []string
	var extractSubtitles func(*html.Node) bool
	extractSubtitles = func(n *html.Node) bool {	
		if n.Type == html.ElementNode && n.Data == "h2" {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.TextNode {
					subtitles = append(subtitles, c.Data)
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
		return subtitles
	}
	return subtitles
}