package fileutil

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_osReadDir(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/osReadDir/"

	path := dir + "test-file.txt"
	_ = CreateFilesWithDirs(path)

	list, err := osReadDir(dir)

	ass.Nil(err)

	var ok bool

	for _, entry := range list {
		if entry.Name() == "test-file.txt" {
			ok = true
		}
	}

	_ = DeleteDirs(dir)

	ass.True(ok)
}

func TestFilterMap(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestFilterMap/"

	dirPath := dir + "test-dir"
	_ = CreateDirs(dirPath)

	filePath := dir + "test-file.txt"
	_ = CreateFiles(filePath)

	resInt, err := FilterMap(dir, func(entry os.DirEntry) (int, bool) {
		if !entry.IsDir() {
			return 1, true
		}

		return 0, false
	})

	ass.Nil(err)

	ass.Equal([]int{1}, resInt)

	_ = DeleteDirs(dir)

}

func TestFilterMapWalk(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestFilterMapWalk/"

	path := dir + "test-dir/test.txt"

	_ = CreateFilesWithDirs(path)

	resInt, err := FilterMapWalk(path, func(path string, entry os.DirEntry) (int, bool) {
		if entry.Name() == "test.txt" {
			return 1, true
		}

		return 0, false
	})

	ass.Nil(err)

	ass.Equal([]int{1}, resInt)

	_ = DeleteDirs(dir)
}

func TestIsDir(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestIsDir/"

	_ = CreateDirs(dir)

	exist, err := IsDir(dir)

	ass.Nil(err)

	ass.Equal(ass.DirExists(dir), exist)

	_ = DeleteDirs(dir)
}

func TestIsEmptyDir(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestIsEmptyDir/"

	_ = CreateDirs(dir)

	exist, err := IsEmptyDir(dir)

	ass.Nil(err)
	ass.True(exist)

	_ = DeleteDirs(dir)
}

func TestFileExists(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestFileExists/"

	path := dir + "test-file-exist.txt"

	_ = CreateFilesWithDirs(path)

	exist, err := FileExists(path)

	ass.Nil(err)
	ass.Equal(ass.FileExists(path), exist)

	_ = DeleteDirs(dir)
}

func TestDirExists(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestDirExists/"

	_ = CreateDirs(dir)

	exist, err := DirExists(dir)

	ass.Nil(err)
	ass.Equal(ass.DirExists(dir), exist)

	_ = DeleteDirs(dir)
}

func TestOsCreate(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestCreateFileWithOS/"

	path := dir + "test-file.txt"

	_ = CreateDirs(dir)

	file, err := OsCreate(path)

	ass.Nil(err)
	ass.FileExists(path)

	file.Close()

	_ = DeleteDirs(dir)

}

func TestCreateFiles(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestCreateFiles/"

	paths := []string{
		dir + "test-file-create1.txt",
		dir + "test-file-create2.txt",
	}

	_ = CreateDirs(dir)

	ass.Nil(CreateFiles(paths...))

	for _, path := range paths {
		ass.FileExists(path)
	}

	_ = DeleteDirs(dir)
}

func TestCreateFileWithData(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestCreateFileWithData/"

	path := dir + "test-file-data.txt"

	data := "test"

	_ = CreateDirs(dir)

	ass.Nil(CreateFileWithData(path, data))

	data2, _ := ReadAll(path)

	ass.Equal(data, data2)

	_ = DeleteDirs(dir)

}

func TestOverwriteFiles(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestOverwriteFiles/"

	paths := []string{
		dir + "test-file-create1.txt",
		dir + "test-file-create2.txt",
	}

	_ = CreateDirs(dir)

	ass.Nil(OverwriteFiles(paths...))

	for _, path := range paths {
		ass.FileExists(path)
	}

	_ = DeleteDirs(dir)
}

func TestCreateDirs(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestCreateDirs/"

	paths := []string{
		dir + "test-dir-create1",
		dir + "test-dir-create2",
	}

	ass.Nil(CreateDirs(paths...))

	for _, path := range paths {
		ass.DirExists(path)
	}

	_ = DeleteDirs(dir)
}

func TestCreateFilesWithDirs(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestCreateFilesWithDirs/"

	paths := []string{
		dir + "test-file-create1.txt",
		dir + "test-file-create2.txt",
	}

	ass.Nil(CreateFilesWithDirs(paths...))

	for _, path := range paths {
		ass.FileExists(path)
	}

	_ = DeleteDirs(dir)
}

func TestOverwriteFilesWithDirs(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestOverwriteFilesWithDirs/"

	paths := []string{
		dir + "test-file-create1.txt",
		dir + "test-file-create2.txt",
	}

	ass.Nil(OverwriteFilesWithDirs(paths...))

	for _, path := range paths {
		ass.FileExists(path)
	}

	_ = DeleteDirs(dir)
}

func TestCopyFile(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestCopyFile/"

	src := dir + "test-file-exist.txt"
	dst := dir + "test-file-exist2.txt"

	data := "test"

	_ = CreateDirs(dir)
	_ = CreateFileWithData(src, data)

	_ = CopyFile(src, dst)

	ass.FileExists(dst)
	res, _ := ReadAll(dst)

	ass.Equal(data, res)

	_ = DeleteDirs(dir)
}

func TestFindFileWalk(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestFindFileWalk/"

	path := dir + "/test-dir/test-file.txt"

	_ = CreateFilesWithDirs(path)

	exist, _ := FindFileWalk(dir, "test-file.txt")

	ass.True(exist)

	_ = DeleteDirs(dir)

}

func TestFindFileWalkFilter(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestFindFileWalkFilter/"

	path := dir + "/test-dir/test-file.txt"
	path1 := dir + "/test-dir/test-file1.txt"

	_ = CreateFilesWithDirs(path, path1)

	exist, _ := FindFileWalkFilter(dir, func(path string, entry os.DirEntry) bool {
		return entry.Name() == "test-file.txt"
	})

	ass.True(exist)

	_ = DeleteDirs(dir)
}

func TestFilenames(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestFilenames/"

	path := dir + "/test-dir/test-file.txt"
	path1 := dir + "test-file1.txt"

	_ = CreateFilesWithDirs(path, path1)

	filenames, _ := Filenames(dir)

	ass.Contains(filenames, "test-file1.txt")
	ass.NotContains(filenames, "test-file.txt")

	_ = DeleteDirs(dir)
}

func TestFilenamesFilter(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestFilenamesFilter/"

	path := dir + "/test-dir/test-file.txt"
	path1 := dir + "test-file1.txt"

	_ = CreateFilesWithDirs(path, path1)

	filenames, _ := FilenamesFilter(dir, func(path string, entry os.DirEntry) bool {
		return true
	})

	ass.Contains(filenames, "test-file1.txt")
	ass.NotContains(filenames, "test-file.txt")

	_ = DeleteDirs(dir)
}

func TestFilenamesBy(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestFilenamesBy/"

	path := dir + "/test-dir/test-file.txt"
	path1 := dir + "test-file1.txt"

	_ = CreateFilesWithDirs(path, path1)

	filenames, _ := FilenamesBy(dir, func(path string, entry os.DirEntry) string {
		return "by-" + entry.Name()
	})

	ass.Contains(filenames, "by-test-file1.txt")

	_ = DeleteDirs(dir)

}

func TestFilenamesWalk(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestFilenamesWalk/"

	path := dir + "/test-dir/test-file.txt"
	path1 := dir + "test-file1.txt"

	_ = CreateFilesWithDirs(path, path1)

	filenames, _ := FilenamesWalk(dir)

	ass.Equal(2, len(filenames))

	_ = DeleteDirs(dir)
}

func TestFilenamesWalkFilter(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestFilenamesWalkFilter/"

	path := dir + "/test-dir/test-file.txt"
	path1 := dir + "test-file1.txt"

	_ = CreateFilesWithDirs(path, path1)

	filenames, _ := FilenamesWalkFilter(dir, func(path string, entry os.DirEntry) bool {
		return true
	})

	ass.Equal(2, len(filenames))

	_ = DeleteDirs(dir)
}

func TestFilenamesWalkBy(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestFilenamesWalkBy/"

	path := dir + "/test-dir/test-file.txt"
	path1 := dir + "test-file1.txt"

	_ = CreateFilesWithDirs(path, path1)

	filenames, _ := FilenamesWalkBy(dir, func(path string, entry os.DirEntry) string {
		return "by-" + entry.Name()
	})

	ass.Equal(2, len(filenames))

	_ = DeleteDirs(dir)
}

func TestDeleteFiles(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestDeleteFiles/"

	paths := []string{
		dir + "test-file-delete1.txt",
		dir + "test-file-delete2.txt",
	}

	_ = CreateDirs(dir)

	_ = CreateFiles(paths...)

	for _, path := range paths {
		ass.FileExists(path)
	}

	_ = DeleteFiles(paths...)

	for _, path := range paths {
		ass.NoFileExists(path)
	}

	_ = DeleteDirs(dir)

}

func TestDeleteDirs(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestDeleteDirs/"

	paths := []string{
		dir + "test-dir-delete1",
		dir + "test-dir-delete2",
	}

	_ = CreateDirs(paths...)

	for _, path := range paths {
		ass.DirExists(path)
	}

	_ = DeleteDirs(dir)

	for _, path := range paths {
		ass.NoDirExists(path)
	}
}

func TestDeleteEmptyDirWalk(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestDeleteEmptyDirWalk/"

	paths := []string{
		dir + "test-dirs-delete1",
		dir + "test-dirs-delete2",
	}

	_ = CreateDirs(paths...)

	for _, path := range paths {
		ass.DirExists(path)
	}

	_ = DeleteEmptyDirWalk(dir)

	for _, path := range paths {
		ass.NoDirExists(path)
	}
}

func TestDeleteWalkBy(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestDeleteWalkBy/"

	filePaths := []string{
		dir + "test-dirs-delete.txt",
		dir + "test-dirs-delete.txt",
	}

	dir2 := dir + "test-dirs"

	// 创建文件
	_ = CreateFilesWithDirs(filePaths...)

	for _, path := range filePaths {
		ass.FileExists(path)
	}

	// 创建空目录
	_ = CreateDirs(dir2)

	ass.DirExists(dir2)

	// 删除包含空目录
	isDelete, err := DeleteWalkBy(dir, func(path string, entry os.DirEntry) (bool, error) {
		return true, nil
	}, true)

	ass.Nil(err)
	ass.True(isDelete)

	ass.NoDirExists(dir2)
	ass.NoDirExists(dir)
	// _ = DeleteDirs("./test-dir-delete")
}

func TestZip(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestZip/"

	_ = CreateDirs(dir)

	// target := "./test-file-zip.zip"

	target := dir + "test-file-zip.zip"

	// 将test-dir-walk 压缩到 ./testdata/TestZip/test-file-zip.zip
	_ = Zip("./test-dir-walk", target)

	ass.FileExists(target)

	_ = DeleteDirs(dir)
}

func TestZipFilter(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestZipFilter/"

	_ = CreateDirs(dir)

	// target := "./test-file-zip.zip"

	target := dir + "test-file-zip.zip"

	// 只把 将test-dir-walk 压缩到 ./testdata/TestZip/test-file-zip.zip
	_ = ZipFilter(".", target, func(path string, entry os.DirEntry) bool {
		return path == "./test-dir-walk"
	})

	ass.FileExists(target)

	_ = DeleteDirs(dir)
}

func TestUnzip(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	src := "./test-file-zip.zip"

	dst := "./Unzip/"

	_ = CreateDirs(dst)

	_ = Unzip(src, dst)

	ass.DirExists(dst)

	_ = DeleteDirs(dst)
}

func TestReadAll(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestReadAll/"

	path := dir + "test-file-data.txt"

	data := "test"

	_ = CreateDirs(dir)
	_ = CreateFileWithData(path, data)

	data2, _ := ReadAll(path)

	ass.Equal(data, data2)

	_ = DeleteDirs(dir)
}

func TestReadLines(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	dir := "./testdata/TestReadLine/"

	path := dir + "test-file-data.txt"

	data := "test\ntest"

	_ = CreateDirs(dir)
	_ = CreateFileWithData(path, data)

	data2, _ := ReadLines(path, 1)

	ass.Equal([]string{"test"}, data2)

	_ = DeleteDirs(dir)
}
