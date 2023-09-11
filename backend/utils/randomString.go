package utils

import (
	"math/rand"
	"time"
)

func randomNumber(min, max int32) int32 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + int32(rand.Intn(int(max-min)))
}

func randomString(charSet string, codeLenghth int32) string {
	code := ""

	charSetLenghth := int32(len(charSet))

	for i := int32(0); i < codeLenghth; i++ {
		index := randomNumber(0, charSetLenghth)
		code += string(charSet[index])
	}

	return code
}

func RandomString() string {
	charSet := "abcdefghijklmnopqrstuvwxyz-_"
	str := randomString(charSet, 33)
	return str
}
