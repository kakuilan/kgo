package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncrypt_Base64Encode(t *testing.T) {
	var res []byte

	res = KEncr.Base64Encode([]byte(""))
	assert.Nil(t, res)

	res = KEncr.Base64Encode(btysHello)
	assert.Contains(t, string(res), "=")
}

func BenchmarkEncrypt_Base64Encode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KEncr.Base64Encode(btysHello)
	}
}
