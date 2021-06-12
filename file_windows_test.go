// +build windows

package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileWins_IsReadable_Deny(t *testing.T) {
	var res bool
	res = KFile.IsReadable(admDir)
	assert.False(t, res)
}

func TestFileWins_IsWritable_Deny(t *testing.T) {
	var res bool
	res = KFile.IsWritable(admDir)
	assert.False(t, res)
}

func TestFileWins_IsExecutable_Deny(t *testing.T) {
	var res bool
	res = KFile.IsExecutable(admDir)
	assert.False(t, res)
}
