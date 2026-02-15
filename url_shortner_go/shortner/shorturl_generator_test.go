package shortner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "4f76f9df-840d-4e46-8e91-233611657c4f"

func TestGenerateShortURL(t *testing.T) {
	initLink_1 := "https://apophis.ir"
	initLink_2 := "https://github.com/pimp13"

	shortLink_1 := GenerateShortURL(initLink_1, UserId)
	shortLink_2 := GenerateShortURL(initLink_2, UserId)
	// fmt.Println(shortLink_1, shortLink_2)
	// zYDM4gjM4M 4ADN2MzNwM

	assert.Equal(t, "zYDM4gjM4M", shortLink_1)
	assert.Equal(t, "4ADN2MzNwM", shortLink_2)

}
