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

func (ws *WebScraper) fetchWebsiteHTML() error {
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

func (ws *WebScraper) ExtractHeadline() {
	var extractHeadline func(*html.Node)
	extractHeadline = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "h1" {
			for c := n.FirstChild; c != nil; c = n.NextSibling {
				if c.Type == html.TextNode {
					log.Printf("Headline found.")
					return
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				extractHeadline(c)
			}
		}
	}

	if ws.Doc == nil {
		log.Printf("Document not found")
		return
	}

	extractHeadline(ws.Doc)
}