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

func TestFileUnix_TarGzUnTarGz(t *testing.T) {
	var res1, res2 bool
	var err1, err2 error

	//打包无权限的目录
	res1, err1 = KFile.TarGz(rootDir, targzfile1)
	assert.False(t, res1)
	assert.NotNil(t, err1)

	//解压到无权限的目录
	res2, err2 = KFile.UnTarGz(targzfile1, rootDir)
	assert.False(t, res2)
	assert.NotNil(t, err2)
}
