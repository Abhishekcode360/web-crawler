package service

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

func ExtractURLs(body io.Reader, domain string) []string {
	tokens := html.NewTokenizer(body) // convert the html to different token(tags)
	var urls []string

	for {
		tokenType := tokens.Next() // Get all the tags
		switch tokenType {
		case html.ErrorToken:
			return urls
		case html.StartTagToken, html.SelfClosingTagToken: // iterating over two types of tages inlcuding div, img, href etc.
			t := tokens.Token()
			for _, attr := range t.Attr { // check for all the attributes
				if attr.Key == "href" {
					if strings.HasPrefix(attr.Val, "/") {
						fullURL := fmt.Sprintf("https://%s%s", domain, attr.Val)
						if IsValidURL(fullURL) {
							urls = append(urls, fullURL)
						}
					} else if IsValidURL(attr.Val) {
						urls = append(urls, attr.Val)
					}
				}
			}
		}
	}
}

func IsValidURL(url string) bool {
	invalidExtensions := []string{
		".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp",
		".css", ".js", ".svg", ".woff", ".woff2", ".ttf", ".eot",
		".ico", ".pdf", ".zip", ".rar", ".tar", ".gz", ".7z",
		".mp3", ".mp4", ".avi", ".mkv", ".flv", ".mov",
	} // remove all the irrelvant files

	for _, ext := range invalidExtensions {
		if strings.HasSuffix(strings.ToLower(url), ext) {
			return false
		}
	}

	return true
}
