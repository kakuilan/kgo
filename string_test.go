package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString_Addslashes(t *testing.T) {
	str := "Is your name O'reilly?"
	res1 := KStr.Addslashes(str)
	assert.Contains(t, res1, "\\")

	res2 := KStr.Stripslashes(res1)
	assert.NotContains(t, res2, "\\")

	KStr.Stripslashes(`Is \ your \\name O\'reilly?`)
}

func BenchmarkString_Addslashes(b *testing.B) {
	b.ResetTimer()
	str := "Is your name O'reilly?"
	for i := 0; i < b.N; i++ {
		KStr.Addslashes(str)
	}
}

func BenchmarkString_Stripslashes(b *testing.B) {
	b.ResetTimer()
	str := `Is your name O\'reilly?`
	for i := 0; i < b.N; i++ {
		KStr.Stripslashes(str)
	}
}
