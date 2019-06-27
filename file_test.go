package gohelper

import "testing"

func TestIsExist(t *testing.T) {
	filename := "./file.go"
	if !File.IsExist(filename) {
		t.Error("file not exist")
		return
	}
}