package customgenerate

import (
	"math/rand"
	"time"
)

func GenerateCode(length int) string {
	rand.Seed(time.Now().UnixNano())

	const availableChars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	code := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(availableChars))
		code[i] = availableChars[randomIndex]
	}

	return string(code)
}
