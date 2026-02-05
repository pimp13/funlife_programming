package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSit_Assert(t *testing.T) {
	fixLoc1, foundNotFixLoc1 := sit("123", "142")
	assert.Equal(t, 1, fixLoc1)
	assert.Equal(t, 1, foundNotFixLoc1)

	fixLoc2, foundNotFixLoc2 := sit("123", "123")
	assert.Equal(t, 3, fixLoc2)
	assert.Equal(t, 0, foundNotFixLoc2)
}
