package kgo

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	v := Version
	assert.NotEmpty(t, v)
}