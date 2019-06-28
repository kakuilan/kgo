package gohelper

import "testing"

func TestIsExist(t *testing.T) {
	filename := "./file.go"
	if !File.IsExist(filename) {
		t.Error("file not exist")
		return
	}
}

func TestGetExt(t *testing.T) {
	filename := "./file.go"
	if File.GetExt(filename) !="go" {
		t.Error("file extension error")
		return
	}
}

func TestGetSize(t *testing.T)  {
	filename := "./file.go"
	if File.GetSize(filename) <=0 {
		t.Error("file size error")
		return
	}
}