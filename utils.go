package fapi

import "strings"

func IfSlashPrefixString(s string) string {
	if s == "" {
		return s
	}
	s = strings.TrimSuffix(s, "/")
	// if strings.HasSuffix(s, "/") {
	// 	s = s[:len(s)-len("/")]
	// }
	if strings.HasPrefix(s, "/") {
		return ToFormat(s)
	}
	return "/" + ToFormat(s)
}

func ToFormat(s string) string {
	result := strings.ToLower(s)
	return strings.ReplaceAll(result, " ", "")
}
