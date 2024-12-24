package utils

import (
	"net/url"
	"strings"
)

func GetFileName(parsedURL *url.URL) string {
	if strings.HasSuffix(parsedURL.Path, "/") {
		return "index.html"
	}
	segments := strings.Split(parsedURL.Path, "/")
	return segments[len(segments)-1]
}
