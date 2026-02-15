package shortner

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"github.com/jxskiss/base62"
	b62 "github.com/jxskiss/base62"
)

func sha256of(input string) []byte {
	algo := sha256.New()

	algo.Write([]byte(input))

	return algo.Sum(nil)
}

func base62Encoded(b []byte) string {
	return b62.StdEncoding.EncodeToString(b)
}

// func GenerateShortURL(initLink, userId string) string {
// 	urlHashBytes := sha256of(initLink + userId)
// 	generatedNum := new(big.Int).SetBytes(urlHashBytes).Uint64()
// 	finalStr := base62Encoded([]byte(
// 		fmt.Sprintf("%d", generatedNum),
// 	))
// 	return finalStr
// }

func GenerateShortURL(originalURL, userID string) string {
	input := originalURL + userID
	hash := sha256.Sum256([]byte(input))

	num := binary.BigEndian.Uint64(hash[:8])

	short := base62.Encode([]byte(fmt.Sprintf("%d", num)))

	if len(short) > 10 {
		short = short[:8]
	}

	return string(short)
}
