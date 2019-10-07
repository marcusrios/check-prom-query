package utils

import "strings"

//RemoveHTTPPrefix crop the HTTP prefix if is provided
func RemoveHTTPPrefix(s *string) {
	if strings.HasPrefix(strings.ToLower(*s), "http://") {
		*s = strings.Replace(*s, "http://", "", 1)
	}
}
