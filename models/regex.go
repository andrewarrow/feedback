package models

import "regexp"

func RemoveNonAlphanumeric(s string) string {
	regex := regexp.MustCompile("[^a-z_A-Z0-9]+")
	return regex.ReplaceAllString(s, "_")
}

func RemoveMostNonAlphanumeric(s string) string {
	regex := regexp.MustCompile("[^a-z_A-Z0-9\\[\\]()'\", .:\\/]+")
	return regex.ReplaceAllString(s, "")
}

func MakeSlug(s string) string {
	regex := regexp.MustCompile("[^a-zA-Z0-9]+")
	return regex.ReplaceAllString(s, "-")
}
