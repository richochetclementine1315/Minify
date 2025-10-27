package helpers

import (
	"os"
	"strings"
)

func EnforceHTTP(url string) string {
	// this function will enforce http:// in the URL if not present
	// This says that if the first 4 characters are not http then add http:// at the beginning
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}
func RemoveDomainError(url string) bool {
	// this function will check if the user is trying to shorten our own domain link
	if url == os.Getenv("APP_DOMAIN") {
		return false
	}
	// removing https:// , http:// , www. from the URL
	// because the user can try to shorten our domain link in different ways
	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	if newURL == os.Getenv("APP_DOMAIN") {
		return false
	}
	return true
}
