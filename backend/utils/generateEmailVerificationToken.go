package utils

func GenerateEmailVerificationToken() string {
	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	return randomString(charSet, 100)
}
