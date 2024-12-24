package models

import (
	"log"
	"net/url"
	"strings"
)

type Website struct {
	URL     string
	HTML    string
	Assets  []string
	DirName string
}

func NewWebsite(url string) *Website {
	return &Website{
		URL:     url,
		Assets:  []string{},
		DirName: getBaseURL(url),
	}
}

func getBaseURL(websiteURL string) string {
	parsedURL, err := url.Parse(websiteURL)
	if err != nil {
		log.Fatalf("Invalid URL: %v", err)
	}
	return strings.TrimSuffix(parsedURL.Host, ".com") // просто для примера
}
