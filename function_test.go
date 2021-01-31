package kgo

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	gofakeit.Seed(0)
}

var randomName = gofakeit.Name()

func TestFun_md5Byte(t *testing.T) {
	res1 := md5Byte([]byte(""), 32)
	res2 := md5Byte([]byte(randomName), 24)
	res3 := md5Byte([]byte(randomName), 0)
	res4 := md5Byte([]byte(randomName), 64)

	assert.Equal(t, 32, len(res1))
	assert.Equal(t, 24, len(res2))
	assert.NotEqual(t, 0, len(res3))
	assert.NotEqual(t, 64, len(res4))
}

func BenchmarkFun_md5Byte(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = md5Byte([]byte(randomName), 32)
	}
}
