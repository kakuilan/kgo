// +build linux darwin

package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileUnix_IsReadable_Deny(t *testing.T) {
	var res bool
	res = KFile.IsReadable(rootDir)
	assert.False(t, res)
}

func TestFileUnix_IsWritable_Deny(t *testing.T) {
	var res bool
	res = KFile.IsWritable(rootDir)
	assert.False(t, res)
}

func TestFileUnix_IsExecutable_Deny(t *testing.T) {
	var res bool
	res = KFile.IsExecutable(rootDir)
	assert.False(t, res)
}
