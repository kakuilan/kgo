package kgo

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

func TestReadFile(t *testing.T) {
	filename := "./file.go"
	cont, _ := KFile.ReadFile(filename)
	if string(cont) == "" {
		t.Error("file get contents error")
		return
	}
	_, _ = KFile.ReadFile("")
}

func BenchmarkReadFile(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i := 0; i < b.N; i++ {
		_, _ = KFile.ReadFile(filename)
	}
}

func TestWriteFile(t *testing.T) {
	str := []byte("Hello World!")
	err := KFile.WriteFile("./testdata/putfile", str)
	if err != nil {
		t.Error("file get contents error")
		return
	}
	_ = KFile.WriteFile("./testdata/putfile", str, 0777)
	_ = KFile.WriteFile("/root/hello/world", str)
	_ = KFile.WriteFile("/root/how/are", str)
}

func BenchmarkWriteFile(b *testing.B) {
	b.ResetTimer()
	str := []byte("Hello World!")
	for i := 0; i < b.N; i++ {
		filename := fmt.Sprintf("./testdata/file/putfile_%d", i)
		_ = KFile.WriteFile(filename, str)
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

func TestIsExecutable(t *testing.T) {
	filename := "./hello"
	res := KFile.IsExecutable(filename)
	if res {
		t.Error("file can not execute")
		return
	}
}

func BenchmarkIsExecutable(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i := 0; i < b.N; i++ {
		KFile.IsExecutable(filename)
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

func TestIsFile(t *testing.T) {
	tests := []struct {
		f        string
		t        LkkFileType
		expected bool
	}{
		{"", FILE_TYPE_ANY, false},
		{"./hello.go", FILE_TYPE_ANY, false},
		{"./file.go", FILE_TYPE_ANY, true},
		{"./file.go", FILE_TYPE_LINK, false},
		{"./file.go", FILE_TYPE_REGULAR, true},
		{"./file.go", FILE_TYPE_COMMON, true},
		{"./testdata/diglett-lnk", FILE_TYPE_LINK, true},
		{"./", FILE_TYPE_ANY, false},
	}

	for _, test := range tests {
		actual := KFile.IsFile(test.f, test.t)
		if actual != test.expected {
			t.Errorf("Expected IsFile(%q, %q) to be %v, got %v", test.f, test.t, test.expected, actual)
			return
		}
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

func TestMkdir(t *testing.T) {
	dir := "./testdata/hello/world"
	err := KFile.Mkdir(dir, 0777)

	if err != nil {
		t.Error("Mkdir fail")
		return
	}
}

func BenchmarkMkdir(b *testing.B) {
	b.ResetTimer()
	dir := "./testdata/hello/world"
	for i := 0; i < b.N; i++ {
		_ = KFile.Mkdir(dir, 0777)
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
	KFile.AbsPath("file:///c:/test.go")

	//手工引发 filepath.Abs 错误
	//创建目录
	testpath := "./testdata/testhehetcl/abspath/123"
	err := os.MkdirAll(testpath, 0755)
	if err == nil {
		//当前目录
		testDir, _ := os.Getwd()

		filename = "../../test.jpg"
		//进入目录
		_ = os.Chdir(testpath)
		pwdir, _ := os.Getwd()

		//删除目录
		_ = os.Remove(pwdir)

		//再获取路径
		res := KFile.AbsPath(filename)
		if res != "/test.jpg" {
			t.Error("KFile.AbsPath fail")
		}

		//回到旧目录
		_ = os.Chdir(testDir)
	}
}

func BenchmarkAbsPath(b *testing.B) {
	b.ResetTimer()
	filename := "./testdata/diglett.png"
	for i := 0; i < b.N; i++ {
		KFile.AbsPath(filename)
	}
}

func TestRealPath(t *testing.T) {
	pwd, _ := KOS.Getcwd()
	path1 := "testdata/diglett.png"
	path2 := "./testdata/diglett.png"
	path3 := pwd + `/` + path1

	res1 := KFile.RealPath("./hello/nothing")
	res2 := KFile.RealPath(path3)
	res3 := KFile.RealPath(path2)
	if res1 != "" || res2 != res3 {
		t.Error("RealPath fail")
		return
	}

	//手工引发 os.Getwd 错误
	//创建目录
	testpath := "./testdata/testhehetcl/abspath/456"
	err := os.MkdirAll(testpath, 0755)
	if err == nil {
		//当前目录
		testDir, _ := os.Getwd()
		filename := "../../test.jpg"

		//进入目录
		_ = os.Chdir(testpath)
		pwdir, _ := os.Getwd()

		//删除目录
		_ = os.Remove(pwdir)

		//再获取路径
		res := KFile.RealPath(filename)
		if res != "" {
			t.Error("KFile.RealPath fail")
		}

		//回到旧目录
		_ = os.Chdir(testDir)
	}
}

func BenchmarkRealPath(b *testing.B) {
	b.ResetTimer()
	path := "testdata/diglett.png"
	for i := 0; i < b.N; i++ {
		KFile.RealPath(path)
	}
}

func TestTouchRenameUnlink(t *testing.T) {
	file1 := "./testdata/empty/zero"
	file2 := "./testdata/empty/2m"
	file3 := "/root/test/empty_zero"
	file4 := "/root/empty_zero"

	//创建文件
	res1 := KFile.Touch(file1, 0)
	res2 := KFile.Touch(file2, 2097152)
	if !res1 || !res2 {
		t.Error("Touch fail")
		return
	}

	//重命名
	file5 := "./testdata/empty/zero_re"
	file6 := "./testdata/empty/2m_re"
	err1 := KFile.Rename(file1, file5)
	err2 := KFile.Rename(file2, file6)
	if err1 != nil || err2 != nil {
		t.Error("Unlink fail")
		return
	}

	//删除文件
	err3 := KFile.Unlink(file5)
	err4 := KFile.Unlink(file6)
	if err3 != nil || err4 != nil {
		t.Error("Unlink fail")
		return
	}

	KFile.Touch(file3, 0)
	KFile.Touch(file4, 0)
}

func BenchmarkTouch(b *testing.B) {
	b.ResetTimer()
	filename := ""
	for i := 0; i < b.N; i++ {
		filename = fmt.Sprintf("./testdata/empty/zero_%d", i)
		KFile.Touch(filename, 0)
	}
}

func BenchmarkRename(b *testing.B) {
	b.ResetTimer()
	filename1 := ""
	filename2 := ""
	for i := 0; i < b.N; i++ {
		filename1 = fmt.Sprintf("./testdata/empty/zero_%d", i)
		filename2 = fmt.Sprintf("./testdata/empty/zero_re%d", i)
		_ = KFile.Rename(filename1, filename2)
	}
}

func BenchmarkUnlink(b *testing.B) {
	b.ResetTimer()
	filename := ""
	for i := 0; i < b.N; i++ {
		filename = fmt.Sprintf("./testdata/empty/zero_re%d", i)
		_ = KFile.Unlink(filename)
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

	//拷贝大文件
	src = "./testdata/2mfile"
	des = "./testdata/2mfile_copy"
	KFile.Touch(src, 2097152)
	_, _ = KFile.CopyFile(src, des, FILE_COVER_ALLOW)

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

func BenchmarkCopyFileErrorRead(b *testing.B) {
	b.ResetTimer()
	src := "./testdata/diglett.png"
	des := ""
	for i := 0; i < b.N; i++ {
		des = fmt.Sprintf("./testdata/sub/diglett_copy_%d.png", i)
		go func(src, des string) {
			_, _ = KFile.CopyFile(src, des, FILE_COVER_ALLOW)
			_ = KFile.Unlink(src)
		}(src, des)
	}
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

func BenchmarkFastCopyErrorRead(b *testing.B) {
	b.ResetTimer()
	src := "./testdata/diglett.png"
	des := ""
	for i := 0; i < b.N; i++ {
		des = fmt.Sprintf("./testdata/fast/diglett_fast_%d.png", i)
		go func(src, des string) {
			_, _ = KFile.FastCopy(src, des)
			_ = KFile.Unlink(src)
		}(src, des)
	}
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

func TestFileImg2Base64(t *testing.T) {
	img := "./testdata/diglett.png"
	str, err := KFile.Img2Base64(img)
	if err != nil || str == "" {
		t.Error("Img2Base64 fail")
		return
	}

	_, _ = KFile.Img2Base64("./testdata/.gitkeep")
	_, _ = KFile.Img2Base64("./testdata/hello.png")

}

func BenchmarkFileImg2Base64(b *testing.B) {
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
	tree := KFile.FileTree("./", FILE_TREE_ALL, true)
	if len(tree) == 0 {
		t.Error("FileTree fail")
		return
	}

	KFile.FileTree("", FILE_TREE_ALL, true)
	KFile.FileTree("./README.md", FILE_TREE_ALL, true)
	KFile.FileTree("/root", FILE_TREE_ALL, true)

	home, _ := KOS.HomeDir()
	KFile.FileTree(home, FILE_TREE_ALL, true)
}

func BenchmarkFileTree(b *testing.B) {
	b.ResetTimer()
	dir := "./"
	for i := 0; i < b.N; i++ {
		_ = KFile.FileTree(dir, FILE_TREE_ALL, true)
	}
}

func TestFormatDir(t *testing.T) {
	var res string
	d1 := `/usr\bin\\golang//fmt/\test\/hehe`
	d2 := `/usr|///tmp:\\\123/\abc<|\hello>\/%world?\\how$\\are`

	res = KFile.FormatDir(d1)
	if strings.Contains(res, `\`) || res != `/usr/bin/golang/fmt/test/hehe/` {
		t.Error("FormatDir fail")
		return
	}

	res = KFile.FormatDir(d2)
	if strings.Contains(res, `\`) || res != `/usr/tmp/123/abc/hello/%world/how$/are/` {
		t.Error("FormatDir fail")
		return
	}

	r1 := KFile.FormatDir(".")
	r2 := KFile.FormatDir("./")
	if r1 != r2 || r2 != "./" {
		t.Error("FormatDir fail")
		return
	}

	KFile.FormatDir("")
}

func TestFormatPath(t *testing.T) {
	p1 := `/usr\bin\\golang//fmt/\test\/hehe`
	p2 := `/usr|///tmp:\\\123/\abc<|\hello>\/%world?\\how$\\are\@#test.png`
	p3 := `test.log`
	p4 := `./test.log`

	res1 := KFile.FormatPath(p1)
	res2 := KFile.FormatPath(p2)
	res3 := KFile.FormatPath(p3)
	res4 := KFile.FormatPath(p4)

	if res1 != `/usr/bin/golang/fmt/test/hehe` {
		t.Error("FormatPath fail")
		return
	}
	if res2 != `/usr/tmp/123/abc/hello/%world/how$/are/@#test.png` {
		t.Error("FormatPath fail")
		return
	}
	if res3 != `test.log` {
		t.Error("FormatPath fail")
		return
	}
	if res4 != `./test.log` {
		t.Error("FormatPath fail")
		return
	}

	KFile.FormatPath("")
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

func TestFileShaX(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	file := "./testdata/diglett.png"

	_, err := KFile.ShaX(file, 1)
	if err != nil {
		t.Error("File ShaX[1] fail")
		return
	}

	_, err = KFile.ShaX(file, 256)
	if err != nil {
		t.Error("File ShaX[256] fail")
		return
	}

	_, err = KFile.ShaX(file, 512)
	if err != nil {
		t.Error("File ShaX[512] fail")
		return
	}

	_, _ = KFile.ShaX("./testdata/hello", 256)
	_, _ = KFile.ShaX(file, 32)
}

func BenchmarkFileShaX(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KFile.ShaX("./testdata/diglett.png", 256)
	}
}

func TestPathinfo(t *testing.T) {
	filename := "./testdata/diglett.png"
	res1 := KFile.Pathinfo(filename, -1)
	res2 := KFile.Pathinfo(filename, 1)
	res3 := KFile.Pathinfo(filename, 2)
	res4 := KFile.Pathinfo(filename, 4)
	res5 := KFile.Pathinfo(filename, 8)
	res6 := KFile.Pathinfo("./testdata/.gitkeep", -1)
	res7 := KFile.Pathinfo("./testdata/hello", -1)

	if len(res1) != 4 {
		t.Error("Pathinfo[all] fail")
		return
	} else if _, ok := res2["dirname"]; !ok {
		t.Error("Pathinfo[dirname] fail")
		return
	} else if _, ok := res3["basename"]; !ok {
		t.Error("Pathinfo[basename] fail")
		return
	} else if _, ok := res4["extension"]; !ok {
		t.Error("Pathinfo[extension] fail")
		return
	} else if _, ok := res5["filename"]; !ok {
		t.Error("Pathinfo[filename] fail")
		return
	} else if ext, _ := res6["extension"]; ext != "gitkeep" {
		t.Error("Pathinfo fail")
		return
	} else if ext, _ := res7["extension"]; ext != "" {
		t.Error("Pathinfo fail")
		return
	}
}

func BenchmarkPathinfo(b *testing.B) {
	b.ResetTimer()
	filename := "./testdata/diglett.png"
	for i := 0; i < b.N; i++ {
		KFile.Pathinfo(filename, -1)
	}
}

func TestBasename(t *testing.T) {
	path := "./testdata/diglett.png"
	res := KFile.Basename(path)
	if res != "diglett.png" {
		t.Error("Basename fail")
		return
	}
}

func BenchmarkBasename(b *testing.B) {
	b.ResetTimer()
	path := "./testdata/diglett.png"
	for i := 0; i < b.N; i++ {
		KFile.Basename(path)
	}
}

func TestDirname(t *testing.T) {
	path1 := "/home/arnie/amelia.jpg"
	path2 := "/mnt/photos/"
	path3 := "rabbit.jpg"
	path4 := "/usr/local//go"
	path5 := ""

	res1 := KFile.Dirname(path1)
	res2 := KFile.Dirname(path2)
	res3 := KFile.Dirname(path3) //返回"."
	res4 := KFile.Dirname(path4)
	res5 := KFile.Dirname(path5) //返回"."

	if res1 == "" || res2 == "" || res4 == "" || res3 != res5 || res5 != "." {
		t.Error("Dirname fail")
		return
	}
}

func BenchmarkDirname(b *testing.B) {
	b.ResetTimer()
	path := "/home/arnie/amelia.jpg"
	for i := 0; i < b.N; i++ {
		KFile.Dirname(path)
	}
}

func TestGetModTime(t *testing.T) {
	path := "./testdata/diglett.png"
	res := KFile.GetModTime(path)
	if res == 0 {
		t.Error("GetModTime fail")
		return
	}

	KFile.GetModTime("./hello")
}

func BenchmarkGetModTime(b *testing.B) {
	b.ResetTimer()
	path := "./testdata/diglett.png"
	for i := 0; i < b.N; i++ {
		KFile.GetModTime(path)
	}
}

func TestGlob(t *testing.T) {
	pattern := "*test.go"
	res, err := KFile.Glob(pattern)
	if err != nil || len(res) == 0 {
		t.Error("Glob fail")
		return
	}
}

func BenchmarkGlob(b *testing.B) {
	b.ResetTimer()
	pattern := "*test.go"
	for i := 0; i < b.N; i++ {
		_, _ = KFile.Glob(pattern)
	}
}

func TestTarGzUnTarGz(t *testing.T) {
	//打包
	patterns := []string{".*_test.go", ".*.yml", "*_test"}
	_, err := KFile.TarGz("./", "./targz/test.tar.gz", patterns...)
	if err != nil {
		t.Error("TarGz fail")
		return
	}

	//解压
	_, err = KFile.UnTarGz("./targz/test.tar.gz", "/tmp/targz/tmp")
	if err != nil {
		t.Error("UnTarGz fail")
		return
	}

	_, _ = KFile.TarGz("", "./targz/test.tar.gz")
	_, _ = KFile.TarGz("/root/hello", "./targz/test.tar.gz", patterns...)
	_, _ = KFile.TarGz("./", "/root/test.tar.gz", patterns...)
	_, _ = KFile.UnTarGz("./targz/hello.tar.gz", "/root/targz/tmp")
	_, _ = KFile.UnTarGz("./targz/test.tar.gz", "/root/targz/tmp")
}

func TestTarGzUnTarGzError(t *testing.T) {
	tarDir := "/tmp/targz/limit"
	go func(tarDir string) {
		_, _ = KFile.TarGz("/tmp/targz/tmp", tarDir+"/test.tar.gz")
	}(tarDir)

	go func(tarDir string) {
		tarDir = "/tmp/targz/tmp"
		KOS.Chmod(tarDir, 0111)
	}(tarDir)

	go func(tarDir string) {
		_, _ = KFile.TarGz("/tmp/targz/tmp", tarDir+"/test.tar.gz")
	}(tarDir)
}

func BenchmarkTarGz(b *testing.B) {
	b.ResetTimer()
	src := "./README.md"
	for i := 0; i < b.N; i++ {
		dst := fmt.Sprintf("./targz/test_%d.tar.gz", i)
		_, _ = KFile.TarGz(src, dst)
	}
}

func BenchmarkUnTarGz(b *testing.B) {
	b.ResetTimer()
	var src, dst string
	for i := 0; i < b.N; i++ {
		src = fmt.Sprintf("./targz/test_%d.tar.gz", i)
		dst = fmt.Sprintf("./targz/test_%d", i)
		_, _ = KFile.UnTarGz(src, dst)
	}
}

func TestSafeFileName(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"", "."},
		{"abc", "abc"},
		{"123456789     '_-?ASDF@£$%£%^é.html", "123456789-asdf.html"},
		{"ReadMe.md", "readme.md"},
		{"file:///c:/test.go", "test.go"},
		{"../../../Hello World!.txt", "hello-world.txt"},
	}
	for _, test := range tests {
		actual := KFile.SafeFileName(test.param)
		if actual != test.expected {
			t.Errorf("Expected SafeFileName(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkSafeFileName(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.SafeFileName("../../../Hello World!.txt")
	}
}

func TestChmodBatch(t *testing.T) {
	dir := "/tmp/kgotest/test"
	err := os.MkdirAll(dir, 0766)
	if err == nil {
		file1 := "/tmp/kgotest/test/chmod/1.log"
		file2 := "/tmp/kgotest/test/chmod/2.log"
		file3 := "/tmp/kgotest/test/chmod/hehe/3.log"
		file4 := "/tmp/kgotest/test/chmod/hehe/4.log"

		KFile.Touch(file1, 0)
		KFile.Touch(file2, 0)
		KFile.Touch(file3, 0)
		KFile.Touch(file4, 0)

		res := KFile.ChmodBatch(dir, 0777, 0755)
		if !res {
			t.Error("ChmodBatch fail")
		}
	}

	res1 := KFile.ChmodBatch("/hello/world/123456", 0766, 0755)
	res2 := KFile.ChmodBatch("/root", 0766, 0755)
	if res1 || res2 {
		t.Error("ChmodBatch fail")
	}
}

func BenchmarkChmodBatch(b *testing.B) {
	b.ResetTimer()
	dir := "/tmp/kgotest/test"
	for i := 0; i < b.N; i++ {
		KFile.ChmodBatch(dir, 0777, 0755)
	}
}

func TestReadInArray(t *testing.T) {
	filepath := "./testdata/dante.txt"
	arr, err := KFile.ReadInArray(filepath)
	if err != nil || len(arr) != 19568 {
		t.Error("ReadInArray fail")
		return
	}

	_, _ = KFile.ReadInArray("./hello")
}

func BenchmarkReadInArray(b *testing.B) {
	b.ResetTimer()
	filepath := "./testdata/dante.txt"
	for i := 0; i < b.N; i++ {
		_, _ = KFile.ReadInArray(filepath)
	}
}

func TestCountLines(t *testing.T) {
	filepath := "./testdata/dante.txt"
	res, err := KFile.CountLines(filepath, 0)
	if err != nil || res == 0 {
		t.Error("CountLines fail")
		return
	}

	res, err = KFile.CountLines("./hello", 8)
	if err == nil || res != -1 {
		t.Error("CountLines fail")
		return
	}

	filepath = "./testdata/gopher10th-small.jpg"
	backup := "./testdata/gopher10th-small.jpg-bak"
	_, _ = KFile.FastCopy(filepath, backup)
	go func() {
		_ = KFile.Unlink(backup)
	}()
	go func() {
		res, err = KFile.CountLines(backup, 1)
	}()
}

func BenchmarkCountLines(b *testing.B) {
	b.ResetTimer()
	filepath := "./testdata/dante.txt"
	for i := 0; i < b.N; i++ {
		_, _ = KFile.CountLines(filepath, 0)
	}
}

func TestZipIszipUnzip(t *testing.T) {
	zfile := "./zip/test.zip"
	var res bool
	var err error

	_, err = KFile.Zip(zfile)
	if err == nil {
		t.Error("Zip fail")
		return
	}

	_, err = KFile.Zip(zfile, "hello-world")
	if err == nil {
		t.Error("Zip fail")
		return
	}

	_, err = KFile.Zip("", "./README.md", "/root")
	if err == nil {
		t.Error("Zip fail")
		return
	}

	res, err = KFile.Zip(zfile, "./README.md", "./testdata/dante.txt")
	if !res || err != nil {
		t.Error("Zip fail")
		return
	}

	svgFile := "./testdata/jetbrains.svg-bak"
	_, _ = KFile.FastCopy("./testdata/jetbrains.svg", svgFile)
	cmd := exec.Command("/bin/bash", "-c", "ln -sf ./testdata/jetbrains.svg-bak ./testdata/svg-lnk")
	_ = cmd.Run()
	_, err = KFile.Zip(zfile, "./testdata")
	if err == nil {
		t.Error("Zip fail")
		return
	}

	_, err = KFile.Zip(zfile, "./testdata/dante.txt", "./testdata/gopher10th-small.jpg", "./docs", "./docs")
	if err != nil {
		t.Error("Zip fail")
		return
	}

	//判断
	chk1 := KFile.IsZip("abc.txt")
	chk2 := KFile.IsZip("abc.zip")
	chk3 := KFile.IsZip(zfile)
	if chk1 || chk2 || !chk3 {
		t.Error("IsZip fail")
		return
	}

	//解压
	dstdir := "./zip/unzip/"
	_, err = KFile.UnZip("hello", dstdir)
	if err == nil {
		t.Error("UnZip fail")
		return
	}

	_, err = KFile.UnZip(zfile, "/root/hello")
	if err == nil {
		t.Error("UnZip fail")
		return
	}

	_, err = KFile.UnZip(zfile, dstdir)
	if err != nil {
		t.Error("UnZip fail")
		return
	}
}

func BenchmarkZip(b *testing.B) {
	b.ResetTimer()
	src := "./README.md"
	for i := 0; i < b.N; i++ {
		dst := fmt.Sprintf("./zip/test_%d.zip", i)
		_, _ = KFile.Zip(dst, src)
	}
}

func BenchmarkUnzip(b *testing.B) {
	b.ResetTimer()
	var src, dst string
	for i := 0; i < b.N; i++ {
		src = fmt.Sprintf("./zip/test_%d.zip", i)
		dst = fmt.Sprintf("./zip/unzip/test_%d", i)
		_, _ = KFile.UnZip(src, dst)
	}
}

func TestGetFileMode(t *testing.T) {
	filepath := "./testdata/dante.txt"
	perm, err := KFile.GetFileMode(filepath)
	if err != nil || perm <= 0 {
		t.Error("GetFileMode fail")
		return
	}

	perm, err = KFile.GetFileMode("/root/hello/world")
	if err == nil || perm != 0 {
		t.Error("GetFileMode fail")
		return
	}
}

func BenchmarkGetFileMode(b *testing.B) {
	b.ResetTimer()
	filepath := "./testdata/dante.txt"
	for i := 0; i < b.N; i++ {
		_, _ = KFile.GetFileMode(filepath)
	}
}

func TestAppendFile(t *testing.T) {
	cont := []byte("hello world.")
	err := KFile.AppendFile("", cont)
	if err == nil {
		t.Error("AppendFile fail")
		return
	}

	pth := "./testdata/append.txt"
	err = KFile.AppendFile(pth, cont)
	if err != nil {
		t.Error("AppendFile fail")
		return
	}

	err = KFile.AppendFile(pth, []byte("how are you?"))
	if err != nil {
		t.Error("AppendFile fail")
		return
	}

	_ = KFile.AppendFile("/root/hello/world", []byte("how are you?"))
	_ = KFile.AppendFile(pth, []byte(""))
}

func BenchmarkAppendFile(b *testing.B) {
	b.ResetTimer()
	pth := "./testdata/append.txt"
	cont := []byte("hello world.\r\n")
	for i := 0; i < b.N; i++ {
		_ = KFile.AppendFile(pth, cont)
	}
}

func TestReadFirstLine(t *testing.T) {
	var tests = []struct {
		file     string
		expected string
	}{
		{"", ""},
		{"./testdata/firstline.log", ""},
		{"./testdata/dante.txt", "LA DIVINA COMMEDIA"},
		{"docs/changelog.md", "# Changelog"},
	}
	for _, test := range tests {
		actual := KFile.ReadFirstLine(test.file)
		if actual != test.expected {
			t.Errorf("Expected FirstLine(%q) to be %v, got %v", test.file, test.expected, actual)
		}
	}
}

func BenchmarkReadFirstLine(b *testing.B) {
	b.ResetTimer()
	fpath := "./testdata/dante.txt"
	for i := 0; i < b.N; i++ {
		KFile.ReadFirstLine(fpath)
	}
}

func TestReadLastLine(t *testing.T) {
	tfile1 := "./testdata/lastline1.log"
	tfile2 := "./testdata/lastline2.log"
	str := "hello World"
	KFile.Touch(tfile1, 0)
	_ = KFile.WriteFile(tfile2, []byte(str))

	var tests = []struct {
		file     string
		expected string
	}{
		{"", ""},
		{"./testdata/firstline.log", ""},
		{tfile1, ""},
		{tfile2, str},
		{"./testdata/dante.txt", ""},
		{"docs/changelog.md", "*--end of file--*"},
	}
	for _, test := range tests {
		actual := KFile.ReadLastLine(test.file)
		if KStr.Trim(actual) != test.expected {
			t.Errorf("Expected FirstLine(%q) to be [%s], got [%s]", test.file, test.expected, actual)
		}
	}
}

func BenchmarkReadLastLine(b *testing.B) {
	b.ResetTimer()
	fpath := "./testdata/dante.txt"
	for i := 0; i < b.N; i++ {
		KFile.ReadLastLine(fpath)
	}
}
