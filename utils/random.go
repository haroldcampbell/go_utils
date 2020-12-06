package utils

/*
Based on http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
*/
import (
	crypto_rand "crypto/rand"
	"encoding/base64"
	"math/rand"
	"time"
)

// SecureRandBytes returns a byte array that in n bytes long based on crypto/rand
func SecureRandBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := crypto_rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// SecureToken returns a string token of that contains n bytes based on crypto/rand
func SecureToken(n int) (string, error) {
	rawToken, err := SecureRandBytes(n)
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(rawToken), nil
}

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// RandString is a fast RandomString generator. The function was originally called RandStringBytesMaskImprSrc
func RandString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

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

// GenerateUUID generates a random uuid
func GenerateUUID() string {
	//E621E1F8-C36C-495A-93FC-0C247A3E6E5F
	return RandCharacters(1) + RandDigits(3) + RandCharNumPair() + RandCharNumPair() +
		"-" + RandCharNumPair() + RandCharNumPair() +
		"-" + RandDigits(3) + RandCharacters(1) +
		"-" + RandCharNumPair() + RandCharNumPair() +
		"-" + RandCharNumPair() + RandDigits(3) + RandCharNumPair() + RandCharNumPair() + RandCharNumPair() + RandCharacters(1)
}

// GenerateSLUG generates the first 8 characters of random uuid
func GenerateSLUG() string {
	//E621E1F8
	return RandCharacters(1) + RandDigits(3) + RandCharNumPair() + RandCharNumPair()
}
