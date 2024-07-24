package api

import "strings"

func IfSlashPrefixString(s string) string {
	if strings.HasSuffix(s, "/") {
		s = s[:len(s)-len("/")]
	}
	if strings.HasPrefix(s, "/") {
		return ToFormat(s)
	}
	return "/" + ToFormat(s)
}

func ToFormat(s string) string {
	result := strings.ToLower(s)
	return strings.ReplaceAll(result, " ", "")
}
