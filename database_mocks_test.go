package mocks4dal

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMockDatabase(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDB := NewMockDatabase(ctrl)
	assert.NotNil(t, mockDB)
}
