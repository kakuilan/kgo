package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTime_UnixTime(t *testing.T) {
	var res int64

	res = KTime.UnixTime()
	assert.Equal(t, 10, len(toStr(res)))
}

func BenchmarkTime_UnixTime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.UnixTime()
	}
}
