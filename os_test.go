package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOS_Pwd(t *testing.T) {
	res := KOS.Pwd()
	assert.NotEmpty(t, res)
}

func BenchmarkOS_Pwd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KOS.Pwd()
	}
}
