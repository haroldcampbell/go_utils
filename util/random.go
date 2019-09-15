package util

import (
	"math/rand"
	"time"
)

// Src: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func randBytes(n int, stringSrc string) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, n)
	for i := range b {
		b[i] = stringSrc[rand.Intn(len(stringSrc))]
	}
	return string(b)
}

const Letters = "ABCDEFGHJKMNPQRSTUVWXYZ"
const Numbers = "123456789"

// RandCharacters generates a string containing random letters of length n
func RandCharacters(n int) string {
	// Intentionally removed the letters ILO so that these aren't confused with 1 or 0
	return randBytes(n, Letters)
}

// RandDigits generates a string containing random numbers of length n
func RandDigits(n int) string {
	// Intentionally remove 0 digit
	return randBytes(n, Numbers)
}

// RandChar ..
func RandChar() string {
	return RandCharacters(1)
}

// RandCharNumPair ..
func RandCharNumPair() string {
	return RandCharacters(1) + RandDigits(1)
}
