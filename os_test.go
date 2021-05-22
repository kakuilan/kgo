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
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.Pwd()
	}
}

func TestOS_Getcwd_Chdir(t *testing.T) {
	var ori, res string
	var err error

	ori, err = KOS.Getcwd()
	assert.Nil(t, err)
	assert.NotEmpty(t, ori)

	//切换目录
	err = KOS.Chdir(dirTdat)
	assert.Nil(t, err)

	res, err = KOS.Getcwd()
	assert.Nil(t, err)

	//返回原来目录
	err = KOS.Chdir(ori)
	assert.Nil(t, err)
	assert.Equal(t, KFile.AbsPath(res), KFile.AbsPath(dirTdat))
}

func BenchmarkOS_Getcwd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.Getcwd()
	}
}

func BenchmarkOS_Chdir(b *testing.B) {
	b.ResetTimer()
	dir := KOS.Pwd()
	for i := 0; i < b.N; i++ {
		_ = KOS.Chdir(dir)
	}
}

func TestOS_LocalIP(t *testing.T) {
	res, err := KOS.LocalIP()
	assert.Nil(t, err)
	assert.NotEmpty(t, res)
}

func BenchmarkOS_LocalIP(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.LocalIP()
	}
}
