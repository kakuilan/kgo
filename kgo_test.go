package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	v := Version
	assert.NotEmpty(t, v)
}
