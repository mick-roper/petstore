package api

import (
	"fmt"
	"net/url"
	"os"
)

const (
	defaultAPIScheme   = "https"
	defaultAPIHostname = "api.petstore.com"
)

func GetAPIRootURL() (*url.URL, error) {
	s := fmt.Sprintf("%s://%s", defaultAPIScheme, defaultAPIHostname)

	if x := os.Getenv("PETSTORE_API_URL"); x != "" {
		s = x
	}

	return url.Parse(s)
}
