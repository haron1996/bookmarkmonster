package utils

import "strings"

func FolderLabel() string {
	charSet := "abcdefghiklmnopqrstvxyzABCDEFGHIKLMNOPQRSTVXYZ0123456789"
	str1 := randomString(charSet, 3)
	str2 := randomString(charSet, 3)
	label := strings.Join([]string{str1, str2}, "_")
	return label
}
