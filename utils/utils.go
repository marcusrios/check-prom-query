package utils

import (
	"log"
	"net/url"
	"strings"
)

//RemoveHTTPPrefix crop the HTTP prefix if is provided
func RemoveHTTPPrefix(s *string) {
	if strings.HasPrefix(strings.ToLower(*s), "http://") {
		*s = strings.Replace(*s, "http://", "", 1)
	}
}

//EncodeURL gets a full url and then return encoded
func EncodeURL(s string) string {
	URL, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}

	queryString, err := url.ParseQuery(URL.RawQuery)
	if err != nil {
		log.Fatal(err)
	}

	URL.RawQuery = queryString.Encode()
	return URL.String()
}
