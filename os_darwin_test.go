// +build darwin

package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOS_IsMac(t *testing.T) {
	res := KOS.IsMac()
	assert.True(t, res)
}

func BenchmarkOS_IsMac(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.IsMac()
	}
}
