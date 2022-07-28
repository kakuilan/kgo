//go:build linux || darwin
// +build linux darwin

package kgo

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFileUnix_IsReadable_Deny(t *testing.T) {
	res := KFile.IsReadable(rootDir)
	assert.False(t, res)
}

func TestFileUnix_IsWritable_Deny(t *testing.T) {
	res := KFile.IsWritable(rootDir)
	assert.False(t, res)
}

func TestFileUnix_IsExecutable_Deny(t *testing.T) {
	var res bool
	res = KFile.IsExecutable(rootDir)
	assert.False(t, res)
}

func TestFileUnix_CopyFile_Deny(t *testing.T) {
	var err error
	//目标路径无权限
	_, err = KFile.CopyFile(imgPng, rootFile1, FILE_COVER_ALLOW)
	assert.NotNil(t, err)
}

func TestFileUnix_FastCopy_Deny(t *testing.T) {
	var err error
	//目录无权限
	_, err = KFile.FastCopy(imgJpg, rootFile1)
	assert.NotNil(t, err)
}

func TestFileUnix_CopyLink_Deny(t *testing.T) {
	var err error

	//创建链接文件
	if !KFile.IsExist(fileLink) {
		_ = os.Symlink(filePubPem, fileLink)
	}

	//目标路径无权限
	err = KFile.CopyLink(fileLink, rootFile1, FILE_COVER_ALLOW)
	assert.NotNil(t, err)
}

func TestFileUnix_CopyDir_Deny(t *testing.T) {
	var err error
	//目标路径无权限
	_, err = KFile.CopyDir(dirVendor, rootDir2, FILE_COVER_ALLOW)
	assert.NotNil(t, err)

	//源路径无权限
	_, err = KFile.CopyDir(rootDir, dirTdat, FILE_COVER_ALLOW)
	assert.NotNil(t, err)
}

func TestFileUnix_DelDir_Deny(t *testing.T) {
	var err error
	//目录无权限
	err = KFile.DelDir(rootDir, false)
	assert.NotNil(t, err)
}

func TestFileUnix_TarGzUnTarGz(t *testing.T) {
	var res bool
	var err error

	//打包-源目录无权限
	res, err = KFile.TarGz(rootDir, targzfile2)
	assert.False(t, res)
	assert.NotNil(t, err)

	//打包-目标目录无权限
	res, err = KFile.TarGz(dirVendor, rootFile3)
	assert.False(t, res)
	assert.NotNil(t, err)

	//解压到无权限的目录
	if !KFile.IsExist(targzfile1) {
		_, _ = KFile.TarGz(dirVendor, targzfile1)
	}
	res, err = KFile.UnTarGz(targzfile1, rootDir)
	assert.False(t, res)
	assert.NotNil(t, err)
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
