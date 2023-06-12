package helpers

import (
	"net/url"
	"strings"

	"andsp.id/andspdev/config/variables"
)

func SiteURL(uri string) string {

	if uri == "" {
		return variables.BaseURL
	}

	baseURL, err := url.Parse(variables.BaseURL)

	if err != nil {
		panic(err)
	}

	uri = strings.TrimPrefix(uri, "/")
	baseURL.Path = "/" + uri

	return baseURL.String()
}