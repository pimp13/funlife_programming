package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService *StoreService

func init() {
	testStoreService = InitStoreService()
}

func TestInsertaionAndRetrieve(t *testing.T) {
	initLink := "https://apophis.ir"
	userId := "550e8400-e29b-41d4-a716-446655440000"
	shortLink := "Jsuke57oXa"
	SaveUrlMapping(shortLink, initLink, userId)

	retrieveURL := RetrieveInitialUrl(shortLink)

	assert.Equal(t, initLink, retrieveURL)
}
