package models

import "regexp"

func RemoveNonAlphanumeric(s string) string {
	regex := regexp.MustCompile("[^a-z_A-Z0-9]+")
	return regex.ReplaceAllString(s, "_")
}
