package mocks4dalgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	assert.NotEmptyf(t, Version, "Version is empty")
}
