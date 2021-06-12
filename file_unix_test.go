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
	res1, err1 = KFile.TarGz(rootDir, targzfile2)
	assert.False(t, res1)
	assert.NotNil(t, err1)

	//解压到无权限的目录
	res2, err2 = KFile.UnTarGz(targzfile1, rootDir)
	assert.False(t, res2)
	assert.NotNil(t, err2)
}

func TestFileUnix_ChmodBatch(t *testing.T) {
	var res bool

	//无权限的目录
	res = KFile.ChmodBatch(rootDir, 0777, 0777)
	assert.False(t, res)
}

func TestFileUnix_ZipIszipUnzip(t *testing.T) {
	var res1, res2 bool
	var err1, err2 error

	//打包无权限的目录
	res1, err1 = KFile.Zip(zipfile2, rootDir)
	assert.False(t, res1)
	assert.NotNil(t, err1)

	//解压到无权限的目录
	res2, err2 = KFile.UnZip(zipfile1, rootDir)
	assert.False(t, res2)
	assert.NotNil(t, err2)
}
