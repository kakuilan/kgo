package gohelper

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestGetExt(t *testing.T) {
	filename := "./file.go"
	if KFile.GetExt(filename) != "go" {
		t.Error("file extension error")
		return
	}

	KFile.GetExt("./testdata/gitkeep")
}

func BenchmarkGetExt(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i := 0; i < b.N; i++ {
		KFile.GetExt(filename)
	}
}

func TestGetContents(t *testing.T) {
	filename := "./file.go"
	cont, _ := KFile.GetContents(filename)
	if cont == "" {
		t.Error("file get contents error")
		return
	}
}

func BenchmarkGetContents(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i := 0; i < b.N; i++ {
		_, _ = KFile.GetContents(filename)
	}
}

func TestGetMime(t *testing.T) {
	filename := "./testdata/diglett.png"
	mime1 := KFile.GetMime(filename, true)
	mime2 := KFile.GetMime(filename, false)
	if mime1 != mime2 {
		t.Error("GetMime fail")
		return
	}

	KFile.GetMime("./testdata/diglett-lnk", false)
	KFile.GetMime("./", false)
}

func BenchmarkGetMimeFast(b *testing.B) {
	b.ResetTimer()
	filename := "./testdata/diglett.png"
	for i := 0; i < b.N; i++ {
		_ = KFile.GetMime(filename, true)
	}
}

func BenchmarkGetMimeReal(b *testing.B) {
	b.ResetTimer()
	filename := "./testdata/diglett.png"
	for i := 0; i < b.N; i++ {
		_ = KFile.GetMime(filename, false)
	}
}

func TestFileSize(t *testing.T) {
	filename := "./file.go"
	if KFile.FileSize(filename) <= 0 {
		t.Error("file size error")
		return
	}

	KFile.FileSize("./hello")
}

func BenchmarkFileSize(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i := 0; i < b.N; i++ {
		KFile.FileSize(filename)
	}
}

func TestDirSize(t *testing.T) {
	dirpath := "./"
	size := KFile.DirSize(dirpath)
	if size == 0 {
		t.Error("dir size error")
		return
	}
	KFile.DirSize("./hello")
}

func BenchmarkDirSize(b *testing.B) {
	b.ResetTimer()
	dirpath := "./"
	for i := 0; i < b.N; i++ {
		_ = KFile.DirSize(dirpath)
	}
}

func TestIsExist(t *testing.T) {
	filename := "./file.go"
	if !KFile.IsExist(filename) {
		t.Error("file not exist")
		return
	}
}

func BenchmarkIsExist(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i := 0; i < b.N; i++ {
		KFile.IsExist(filename)
	}
}

func TestIsWritable(t *testing.T) {
	filename := "./README.md"
	if !KFile.IsWritable(filename) {
		t.Error("file can not write")
		return
	}
	KFile.IsWritable("./hello")
}

func BenchmarkIsWritable(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i := 0; i < b.N; i++ {
		KFile.IsWritable(filename)
	}
}

func TestIsReadable(t *testing.T) {
	filename := "./README.md"
	if !KFile.IsReadable(filename) {
		t.Error("file can not read")
		return
	}
	KFile.IsReadable("./hello")
}

func BenchmarkIsReadable(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i := 0; i < b.N; i++ {
		KFile.IsReadable(filename)
	}
}

func TestIsFile(t *testing.T) {
	filename := "./file.go"
	if !KFile.IsFile(filename) {
		t.Error("isn`t a file")
		return
	}
	KFile.IsFile("./hello")
}

func BenchmarkIsFile(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i := 0; i < b.N; i++ {
		KFile.IsFile(filename)
	}
}

func TestIsLink(t *testing.T) {
	cmd := exec.Command("/bin/bash", "-c", "ln -sf ./testdata/diglett.png ./testdata/diglett-lnk")
	_ = cmd.Run()
	filename := "./testdata/diglett-lnk"
	if !KFile.IsLink(filename) {
		t.Error("isn`t a link")
		return
	}
	KFile.IsLink("./hello")
}

func BenchmarkIsLink(b *testing.B) {
	b.ResetTimer()
	filename := "./testdata/diglett-lnk"
	for i := 0; i < b.N; i++ {
		KFile.IsLink(filename)
	}
}

func TestIsDir(t *testing.T) {
	dirname := "./"
	if !KFile.IsDir(dirname) {
		t.Error("isn`t a dir")
		return
	}
	KFile.IsDir("./hello")
	KFile.IsDir("/root/.bashrc")
}

func BenchmarkIsDir(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i := 0; i < b.N; i++ {
		KFile.IsDir(filename)
	}
}

func TestFileIsBinary(t *testing.T) {
	filename := "./file.go"
	if KFile.IsBinary(filename) {
		t.Error("file isn`t binary")
		return
	}

	goroot := os.Getenv("GOROOT")
	file2 := goroot + "/bin/go"
	if !KFile.IsBinary(file2) {
		t.Error("file isn`t binary")
		return
	}

	KFile.IsBinary("./hello")
}

func BenchmarkFileIsBinary(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i := 0; i < b.N; i++ {
		KFile.IsBinary(filename)
	}
}

func TestIsImg(t *testing.T) {
	filename := "./testdata/diglett.png"
	if !KFile.IsImg(filename) {
		t.Error("file isn`t img")
		return
	}
	KFile.IsImg("./hello")
}

func BenchmarkIsImg(b *testing.B) {
	b.ResetTimer()
	filename := "./testdata/diglett.png"
	for i := 0; i < b.N; i++ {
		KFile.IsImg(filename)
	}
}

func TestAbsPath(t *testing.T) {
	filename := "./testdata/diglett.png"
	abspath := KFile.AbsPath(filename)
	if !KFile.IsExist(abspath) {
		t.Error("file not exist")
		return
	}
	KFile.AbsPath("")
}

func BenchmarkAbsPath(b *testing.B) {
	b.ResetTimer()
	filename := "./testdata/diglett.png"
	for i := 0; i < b.N; i++ {
		KFile.AbsPath(filename)
	}
}

func TestQuickFile(t *testing.T) {
	file1 := "./testdata/empty/zero"
	file2 := "./testdata/empty/2m"

	res1 := KFile.QuickFile(file1, 0)
	res2 := KFile.QuickFile(file2, 2097152)
	if !res1 || !res2 {
		t.Error("QuickFile fail")
		return
	}

	KFile.QuickFile("/root/test/empty_zero", 0)
	KFile.QuickFile("/root/empty_zero", 0)
}

func BenchmarkQuickFile(b *testing.B) {
	b.ResetTimer()
	filename := ""
	for i := 0; i < b.N; i++ {
		filename = fmt.Sprintf("./testdata/empty/zero_%d", i)
		KFile.QuickFile(filename, 0)
	}
}

func TestCopyFile(t *testing.T) {
	src := "./testdata/diglett.png"
	des := "./testdata/sub/diglett_copy.png"
	num, err := KFile.CopyFile(src, des, FILE_COVER_ALLOW)
	if err != nil || num == 0 {
		t.Error("copy file fail")
		return
	}

	_, _ = KFile.CopyFile("abc", "abc", FILE_COVER_ALLOW)
	_, _ = KFile.CopyFile("./hello", "", FILE_COVER_ALLOW)
	_, _ = KFile.CopyFile(".", "", FILE_COVER_ALLOW)
	_, _ = KFile.CopyFile("./testdata/diglett.png", "./testdata/.gitkeep", FILE_COVER_IGNORE)
	_, _ = KFile.CopyFile("./testdata/diglett.png", "./testdata/.gitkeep", FILE_COVER_DENY)

	_, _ = KFile.CopyFile("./testdata/diglett.png", "/root/test/diglett.png", FILE_COVER_ALLOW)
	_, _ = KFile.CopyFile("./testdata/diglett.png", "/root/diglett.png", FILE_COVER_ALLOW)
	_, _ = KFile.CopyFile("./testdata/empty/2m", "./testdata/empty/2m_copy", FILE_COVER_ALLOW)
	_, _ = KFile.CopyFile("./testdata/empty/2m", "./testdata/empty/2m_copy", FILE_COVER_IGNORE)

}

func BenchmarkCopyFile(b *testing.B) {
	b.ResetTimer()
	src := "./testdata/diglett.png"
	des := ""
	for i := 0; i < b.N; i++ {
		des = fmt.Sprintf("./testdata/sub/diglett_copy_%d.png", i)
		_, _ = KFile.CopyFile(src, des, FILE_COVER_ALLOW)
	}
}

func TestFastCopy(t *testing.T) {
	src := "./testdata/diglett.png"
	des := "./testdata/fast/diglett_copy.png"

	num, err := KFile.FastCopy(src, des)
	if err != nil || num == 0 {
		t.Error("fast copy file fail")
		return
	}

	_, _ = KFile.FastCopy("./hello", "")
	_, _ = KFile.FastCopy("./testdata/diglett.png", "/root/test/diglett.png")
	_, _ = KFile.FastCopy("./testdata/diglett.png", "/root/diglett.png")
}

func BenchmarkFastCopy(b *testing.B) {
	b.ResetTimer()
	src := "./testdata/diglett.png"
	des := ""
	for i := 0; i < b.N; i++ {
		des = fmt.Sprintf("./testdata/fast/diglett_fast_%d.png", i)
		_, _ = KFile.FastCopy(src, des)
	}
}

func TestCopyLink(t *testing.T) {
	src := "./testdata/diglett-lnk"
	des := "./testdata/link/diglett-lnk.copy"

	err := KFile.CopyLink(src, des)
	if err != nil {
		t.Error("copy link fail:" + err.Error())
		return
	}

	_ = KFile.CopyLink(src, des)
	_ = KFile.CopyLink("abc", "abc")
	_ = KFile.CopyLink("./helloe", "abc")
	_ = KFile.CopyLink(src, "/root/test/abc")

}

func BenchmarkCopyLink(b *testing.B) {
	b.ResetTimer()
	src := "./testdata/diglett-lnk"
	des := ""
	for i := 0; i < b.N; i++ {
		des = fmt.Sprintf("./testdata/link/diglett-lnk_%d.copy", i)
		_ = KFile.CopyLink(src, des)
	}
}

func TestCopyDir(t *testing.T) {
	src := "./testdata"
	des := "./test/copy"
	des2 := "./test/copy2"

	num, err := KFile.CopyDir(src, des, FILE_COVER_ALLOW)
	if err != nil || num == 0 {
		t.Error("copy directory fail")
		return
	}

	_, _ = KFile.CopyDir("./hello", des, FILE_COVER_ALLOW)
	_, _ = KFile.CopyDir("./file.go", des, FILE_COVER_ALLOW)
	_, _ = KFile.CopyDir(src, "/root/test/tdir", FILE_COVER_ALLOW)
	_, _ = KFile.CopyDir("/root/", des, FILE_COVER_ALLOW)
	_, _ = KFile.CopyDir(src, des2, FILE_COVER_ALLOW)
	_, _ = KFile.CopyDir(des, des2, FILE_COVER_IGNORE)
	_, _ = KFile.CopyDir(des, des2, FILE_COVER_ALLOW)
}

func BenchmarkCopyDir(b *testing.B) {
	b.ResetTimer()
	src := "./testdata"
	des := ""
	for i := 0; i < b.N; i++ {
		des = fmt.Sprintf("./test/copy_%d", i)
		_, _ = KFile.CopyDir(src, des, FILE_COVER_ALLOW)
	}
}

func TestImg2Base64(t *testing.T) {
	img := "./testdata/diglett.png"
	str, err := KFile.Img2Base64(img)
	if err != nil || str == "" {
		t.Error("Img2Base64 fail")
		return
	}

	_, _ = KFile.Img2Base64("./testdata/.gitkeep")
	_, _ = KFile.Img2Base64("./testdata/hello.png")

}

func BenchmarkImg2Base64(b *testing.B) {
	b.ResetTimer()
	img := "./testdata/diglett.png"
	for i := 0; i < b.N; i++ {
		_, _ = KFile.Img2Base64(img)
	}
}

func TestDelDir(t *testing.T) {
	dir := "./test"
	err := KFile.DelDir(dir, true)
	if err != nil || KFile.IsDir(dir) {
		t.Error("DelDir fail")
		return
	}

	_ = KFile.DelDir("./hello", true)
	_ = KFile.DelDir("/root", true)

}

func BenchmarkDelDir(b *testing.B) {
	b.ResetTimer()
	dir := "./test"
	for i := 0; i < b.N; i++ {
		_ = KFile.DelDir(dir, true)
	}
}

func TestFileTree(t *testing.T) {
	dir := "./"
	tree := KFile.FileTree(dir, FILE_TREE_ALL, true)
	//fmt.Printf("%v", tree)
	if len(tree) == 0 {
		t.Error("FileTree fail")
		return
	}

	KFile.FileTree("", FILE_TREE_ALL, true)
	KFile.FileTree("./README.md", FILE_TREE_ALL, true)
	KFile.FileTree("/root", FILE_TREE_ALL, true)
}

func BenchmarkFileTree(b *testing.B) {
	b.ResetTimer()
	dir := "./"
	for i := 0; i < b.N; i++ {
		_ = KFile.FileTree(dir, FILE_TREE_ALL, true)
	}
}

func TestFormatDir(t *testing.T) {
	dir := `/usr\bin\\golang//fmt`
	res := KFile.FormatDir(dir)
	if strings.Contains(res, `\`) {
		t.Error("FormatDir fail")
		return
	}

	KFile.FormatDir("")
}

func BenchmarkFormatDir(b *testing.B) {
	b.ResetTimer()
	dir := `/usr\bin\\golang//fmt`
	for i := 0; i < b.N; i++ {
		_ = KFile.FormatDir(dir)
	}
}

func TestFileMd5(t *testing.T) {
	file := `./file.go`
	res1, _ := KFile.Md5(file, 0)
	res2, _ := KFile.Md5(file, 16)
	if len(res1) != 32 || !strings.Contains(res1, res2) {
		t.Error("File Md5 fail")
		return
	}
	_, _ = KFile.Md5("./hello", 32)
	_, _ = KFile.Md5("/tmp", 32)
}

func BenchmarkFileMd5(b *testing.B) {
	b.ResetTimer()
	file := `./file.go`
	for i := 0; i < b.N; i++ {
		_, _ = KFile.Md5(file, 32)
	}
}
