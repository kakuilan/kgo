// +build linux darwin

package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOS_NotWindows(t *testing.T) {
	res := KOS.IsWindows()
	assert.False(t, res)
}
