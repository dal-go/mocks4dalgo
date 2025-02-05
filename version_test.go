package mocks4dalgo

import (
	"strings"
	"testing"
)

func TestVersion(t *testing.T) {
	if strings.HasPrefix(Version, "v") == false {
		t.Errorf("Version should start with a 'v'")
	}
}
