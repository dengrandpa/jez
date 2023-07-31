package fileutil

import (
	"fmt"
	"os"
)

func ExampleFilterMap() {

	dir := "./testdata/TestFilterMap/"

	dirPath := dir + "test-dir"
	_ = CreateDirs(dirPath)

	filePath := dir + "test-file.txt"
	_ = CreateFiles(filePath)

	resInt, _ := FilterMap(dir, func(entry os.DirEntry) (int, bool) {
		if !entry.IsDir() {
			return 1, true
		}

		return 0, false
	})

	fmt.Println(resInt)

	// Output:
	// [1]

	_ = DeleteDirs(dir)

}

func ExampleFilterMapWalk() {

	dir := "./testdata/TestFilterMapWalk/"

	path := dir + "test-dir/test.txt"

	_ = CreateFilesWithDirs(path)

	resInt, _ := FilterMapWalk(path, func(path string, entry os.DirEntry) (int, bool) {
		if entry.Name() == "test.txt" {
			return 1, true
		}

		return 0, false
	})

	fmt.Println(resInt)

	// Output:
	// [1]

	_ = DeleteDirs(dir)
}

func ExampleIsDir() {

	dir := "./testdata/TestIsDir/"

	_ = CreateDirs(dir)

	exist, _ := IsDir(dir)

	fmt.Println(exist)

	// Output:
	// true

	_ = DeleteDirs(dir)
}

func ExampleIsEmptyDir() {

	dir := "./testdata/TestIsEmptyDir/"

	_ = CreateDirs(dir)

	exist, _ := IsEmptyDir(dir)

	fmt.Println(exist)

	// Output:
	// true

	_ = DeleteDirs(dir)
}

func ExampleFileExists() {

	dir := "./testdata/TestFileExists/"

	path := dir + "test-file-exist.txt"

	_ = CreateFilesWithDirs(path)

	exist, _ := FileExists(path)

	fmt.Println(exist)

	// Output:
	// true

	_ = DeleteDirs(dir)
}

func ExampleDirExists() {

	dir := "./testdata/TestDirExists/"

	_ = CreateDirs(dir)

	exist, _ := DirExists(dir)

	fmt.Println(exist)

	// Output:
	// true

	_ = DeleteDirs(dir)
}

func ExampleCreateFiles() {

	dir := "./testdata/TestCreateFiles/"

	paths := []string{
		dir + "test-file-create1.txt",
		dir + "test-file-create2.txt",
	}

	_ = CreateDirs(dir)

	_ = CreateFiles(paths...)

	for _, path := range paths {
		fmt.Println(FileExists(path))
	}

	// Output:
	// true <nil>
	// true <nil>

	_ = DeleteDirs(dir)
}

func ExampleCreateFileWithData() {

	dir := "./testdata/TestCreateFileWithData/"

	path := dir + "test-file-data.txt"

	data := "test"

	_ = CreateDirs(dir)

	_ = CreateFileWithData(path, data)

	data2, _ := ReadAll(path)

	fmt.Println(data2)

	// Output:
	// test

	_ = DeleteDirs(dir)

}

func ExampleOverwriteFiles() {

	dir := "./testdata/TestOverwriteFiles/"

	paths := []string{
		dir + "test-file-create1.txt",
		dir + "test-file-create2.txt",
	}

	_ = CreateDirs(dir)

	_ = OverwriteFiles(paths...)

	for _, path := range paths {
		fmt.Println(FileExists(path))
	}

	// Output:
	// true <nil>
	// true <nil>

	_ = DeleteDirs(dir)
}

func ExampleCreateDirs() {

	dir := "./testdata/TestCreateDirs/"

	paths := []string{
		dir + "test-dir-create1",
		dir + "test-dir-create2",
	}

	_ = CreateDirs(paths...)

	for _, path := range paths {
		fmt.Println(DirExists(path))
	}

	// Output:
	// true <nil>
	// true <nil>

	_ = DeleteDirs(dir)
}

func ExampleCreateFilesWithDirs() {

	dir := "./testdata/TestCreateFilesWithDirs/"

	paths := []string{
		dir + "test-file-create1.txt",
		dir + "test-file-create2.txt",
	}

	_ = CreateFilesWithDirs(paths...)

	for _, path := range paths {
		fmt.Println(FileExists(path))
	}

	// Output:
	// true <nil>
	// true <nil>

	_ = DeleteDirs(dir)
}

func ExampleOverwriteFilesWithDirs() {

	dir := "./testdata/TestOverwriteFilesWithDirs/"

	paths := []string{
		dir + "test-file-create1.txt",
		dir + "test-file-create2.txt",
	}

	_ = OverwriteFilesWithDirs(paths...)

	for _, path := range paths {
		fmt.Println(FileExists(path))
	}

	// Output:
	// true <nil>
	// true <nil>

	_ = DeleteDirs(dir)
}

func ExampleOsCreate() {

	dir := "./testdata/TestCreateFileWithOS/"

	path := dir + "test-file.txt"

	_ = CreateDirs(dir)

	file, _ := OsCreate(path)
	file.Close()

	fmt.Println(FileExists(path))

	// Output:
	// true <nil>

	_ = DeleteDirs(dir)

}

func ExampleCopyFile() {

	dir := "./testdata/TestCopyFile/"

	src := dir + "test-file-exist.txt"
	dst := dir + "test-file-exist2.txt"

	data := "test"

	_ = CreateDirs(dir)
	_ = CreateFileWithData(src, data)

	_ = CopyFile(src, dst)

	res, _ := ReadAll(dst)

	fmt.Println(res)

	// Output:
	// test

	_ = DeleteDirs(dir)
}

func ExampleFindFileWalk() {

	dir := "./testdata/TestFindFileWalk/"

	path := dir + "/test-dir/test-file.txt"

	_ = CreateFilesWithDirs(path)

	exist, _ := FindFileWalk(dir, "test-file.txt")

	fmt.Println(exist)

	// Output:
	// true

	_ = DeleteDirs(dir)

}

func ExampleFindFileWalkFilter() {

	dir := "./testdata/TestFindFileWalkFilter/"

	path := dir + "/test-dir/test-file.txt"
	path1 := dir + "/test-dir/test-file1.txt"

	_ = CreateFilesWithDirs(path, path1)

	exist, _ := FindFileWalkFilter(dir, func(path string, entry os.DirEntry) bool {
		return entry.Name() == "test-file.txt"
	})

	fmt.Println(exist)

	// Output:
	// true

	_ = DeleteDirs(dir)
}

func ExampleFilenames() {

	dir := "./testdata/TestFilenames/"

	path := dir + "/test-dir/test-file.txt"
	path1 := dir + "test-file1.txt"

	_ = CreateFilesWithDirs(path, path1)

	filenames, _ := Filenames(dir)

	fmt.Println(filenames)

	// Output:
	// [test-file1.txt]

	_ = DeleteDirs(dir)
}

func ExampleFilenamesFilter() {

	dir := "./testdata/TestFilenamesFilter/"

	path := dir + "/test-dir/test-file.txt"
	path1 := dir + "test-file1.txt"

	_ = CreateFilesWithDirs(path, path1)

	filenames, _ := FilenamesFilter(dir, func(path string, entry os.DirEntry) bool {
		return true
	})

	fmt.Println(filenames)

	// Output:
	// [test-file1.txt]

	_ = DeleteDirs(dir)
}

func ExampleFilenamesBy() {

	dir := "./testdata/TestFilenamesBy/"

	path := dir + "/test-dir/test-file.txt"
	path1 := dir + "test-file1.txt"

	_ = CreateFilesWithDirs(path, path1)

	filenames, _ := FilenamesBy(dir, func(path string, entry os.DirEntry) string {
		return "by-" + entry.Name()
	})

	fmt.Println(filenames)

	// Output:
	// [by-test-file1.txt]

	_ = DeleteDirs(dir)

}

func ExampleFilenamesWalk() {

	dir := "./testdata/TestFilenamesWalk/"

	path := dir + "/test-dir/test-file.txt"
	path1 := dir + "test-file1.txt"

	_ = CreateFilesWithDirs(path, path1)

	filenames, _ := FilenamesWalk(dir)

	fmt.Println(filenames)

	// Output:
	// [test-file.txt test-file1.txt]

	_ = DeleteDirs(dir)
}

func ExampleFilenamesWalkFilter() {

	dir := "./testdata/TestFilenamesWalkFilter/"

	path := dir + "/test-dir/test-file.txt"
	path1 := dir + "test-file1.txt"

	_ = CreateFilesWithDirs(path, path1)

	filenames, _ := FilenamesWalkFilter(dir, func(path string, entry os.DirEntry) bool {
		return true
	})

	fmt.Println(filenames)

	// Output:
	// [test-file.txt test-file1.txt]

	_ = DeleteDirs(dir)
}

func ExampleFilenamesWalkBy() {

	dir := "./testdata/TestFilenamesWalkBy/"

	path := dir + "/test-dir/test-file.txt"
	path1 := dir + "test-file1.txt"

	_ = CreateFilesWithDirs(path, path1)

	filenames, _ := FilenamesWalkBy(dir, func(path string, entry os.DirEntry) string {
		return "by-" + entry.Name()
	})

	fmt.Println(filenames)

	// Output:
	// [by-test-file.txt by-test-file1.txt]

	_ = DeleteDirs(dir)
}

func ExampleDeleteFiles() {

	dir := "./testdata/TestDeleteFiles/"

	paths := []string{
		dir + "test-file-delete1.txt",
		dir + "test-file-delete2.txt",
	}

	_ = CreateDirs(dir)

	_ = CreateFiles(paths...)

	for _, path := range paths {
		fmt.Println(FileExists(path))
	}

	_ = DeleteFiles(paths...)

	for _, path := range paths {
		fmt.Println(FileExists(path))
	}

	// Output:
	// true <nil>
	// true <nil>
	// false <nil>
	// false <nil>

	_ = DeleteDirs(dir)

}

func ExampleDeleteDirs() {

	dir := "./testdata/TestDeleteDirs/"

	paths := []string{
		dir + "test-dir-delete1",
		dir + "test-dir-delete2",
	}

	_ = CreateDirs(paths...)

	for _, path := range paths {
		fmt.Println(DirExists(path))
	}

	_ = DeleteDirs(dir)

	for _, path := range paths {
		fmt.Println(DirExists(path))
	}

	// Output:
	// true <nil>
	// true <nil>
	// false <nil>
	// false <nil>
}

func ExampleDeleteEmptyDirWalk() {

	dir := "./testdata/TestDeleteEmptyDirWalk/"

	paths := []string{
		dir + "test-dirs-delete1",
		dir + "test-dirs-delete2",
	}

	_ = CreateDirs(paths...)

	for _, path := range paths {
		fmt.Println(DirExists(path))
	}

	_ = DeleteEmptyDirWalk(dir)

	for _, path := range paths {
		fmt.Println(DirExists(path))
	}

	fmt.Println(DirExists(dir))

	// Output:
	// true <nil>
	// true <nil>
	// false <nil>
	// false <nil>
	// false <nil>
}

func ExampleDeleteWalkBy() {

	dir := "./testdata/TestDeleteWalkBy/"

	filePaths := []string{
		dir + "test-dirs-delete.txt",
		dir + "test-dirs-delete.txt",
	}

	dir2 := dir + "test-dirs"

	// 创建文件
	_ = CreateFilesWithDirs(filePaths...)

	for _, path := range filePaths {
		fmt.Println(FileExists(path))
	}

	// 创建空目录
	_ = CreateDirs(dir2)

	fmt.Println(DirExists(dir2))

	// 删除包含空目录
	isDelete, _ := DeleteWalkBy(dir, func(path string, entry os.DirEntry) (bool, error) {
		return true, nil
	}, true)

	fmt.Println(isDelete)
	fmt.Println(DirExists(dir2))
	fmt.Println(DirExists(dir))

	// Output:
	// true <nil>
	// true <nil>
	// true <nil>
	// true
	// false <nil>
	// false <nil>
}

func ExampleZip() {

	dir := "./testdata/TestZip/"

	_ = CreateDirs(dir)

	// target := "./test-file-zip.zip"

	target := dir + "test-file-zip.zip"

	// 将test-dir-walk 压缩到 ./testdata/TestZip/test-file-zip.zip
	_ = Zip("./test-dir-walk", target)

	fmt.Println(FileExists(target))

	// Output:
	// true <nil>

	_ = DeleteDirs(dir)
}

func ExampleZipFilter() {

	dir := "./testdata/TestZipFilter/"

	_ = CreateDirs(dir)

	// target := "./test-file-zip.zip"

	target := dir + "test-file-zip.zip"

	// 只把 将test-dir-walk 压缩到 ./testdata/TestZip/test-file-zip.zip
	_ = ZipFilter(".", target, func(path string, entry os.DirEntry) bool {
		return path == "./test-dir-walk"
	})

	fmt.Println(FileExists(target))

	// Output:
	// true <nil>

	_ = DeleteDirs(dir)
}

func ExampleUnzip() {

	src := "./test-file-zip.zip"

	dst := "./Unzip/"

	_ = CreateDirs(dst)

	_ = Unzip(src, dst)

	fmt.Println(DirExists(dst))

	// Output:
	// true <nil>

	_ = DeleteDirs(dst)
}

func ExampleReadAll() {

	dir := "./testdata/TestReadAll/"

	path := dir + "test-file-data.txt"

	data := "test"

	_ = CreateDirs(dir)
	_ = CreateFileWithData(path, data)

	data2, _ := ReadAll(path)

	fmt.Println(data2)

	// Output:
	// test

	_ = DeleteDirs(dir)
}

func ExampleReadLine() {

	dir := "./testdata/TestReadLine/"

	path := dir + "test-file-data.txt"

	data := "test\ntest"

	_ = CreateDirs(dir)
	_ = CreateFileWithData(path, data)

	data2, _ := ReadLine(path, 1)

	fmt.Println(data2)

	// Output:
	// [test]

	_ = DeleteDirs(dir)
}
