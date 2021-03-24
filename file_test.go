package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFile_GetExt(t *testing.T) {
	var ext string

	ext = KFile.GetExt("./file.go")
	assert.Equal(t, "go", ext)

	ext = KFile.GetExt("./testdata/gitkeep")
	assert.Empty(t, ext)
}
