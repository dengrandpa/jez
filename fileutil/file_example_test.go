package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ExampleFilterMap() {

	path := "./testdata"

	resInt, _ := FilterMap(path, func(entry os.DirEntry) (int, bool) {
		if !entry.IsDir() {
			return 1, true
		}

		return 0, false
	})

	fmt.Println(resInt)

	// Output:
	// [1 1]

}

func ExampleFilterMapWalk() {

	path := "./"

	resInt, _ := FilterMapWalk(path, func(path string, entry os.DirEntry) (int, bool) {
		if !entry.IsDir() {
			return 1, true
		}

		return 0, false
	})

	fmt.Println(resInt)

	// Output:
	// [1 1 1 1 1]
}

func ExampleIsDir() {
	path := "./testdata"

	fmt.Println(IsDir(path))

	// Output:
	// true <nil>
}

func ExampleIsEmptyDir() {
	path := "./testdata/test-dir-empty"

	fmt.Println(IsEmptyDir(path))

	// Output:
	// true <nil>
}

func ExampleFileExists() {
	path := "./testdata/test-file-exist.txt"

	fmt.Println(FileExists(path))

	// Output:
	// true <nil>
}

func ExampleDirExists() {
	path := "./testdata"

	fmt.Println(DirExists(path))

	// Output:
	// true <nil>
}

func ExampleCreateFiles() {
	paths := []string{
		"./testdata/test-file-create1.txt",
		"./testdata/test-file-create2.txt",
	}
	_ = CreateFiles(paths...)

	for _, path := range paths {
		fmt.Println(FileExists(path))
	}

	_ = DeleteFiles(paths...)

	// Output:
	// true <nil>
	// true <nil>

}

func ExampleCreateFileWithData() {

	path := "./testdata/test-file-data.txt"

	data := "test"

	_ = CreateFileWithData(path, data)

	data2, _ := ReadAll(path)

	fmt.Println(data2)

	// Output:
	// test

	_ = DeleteFiles(path)
}

func ExampleOverwriteFiles() {
	paths := []string{
		"./testdata/test-file-create1.txt",
		"./testdata/test-file-create2.txt",
	}

	_ = OverwriteFiles(paths...)

	for _, path := range paths {
		fmt.Println(FileExists(path))
	}

	_ = DeleteFiles(paths...)

	// Output:
	// true <nil>
	// true <nil>
}

func ExampleCreateDirs() {
	paths := []string{
		"./testdata/test-dir-create1",
		"./testdata/test-dir-create2",
	}

	_ = CreateDirs(paths...)

	for _, path := range paths {
		fmt.Println(DirExists(path))
	}

	_ = DeleteDirs(paths...)

	// Output:
	// true <nil>
	// true <nil>
}

func ExampleCreateFilesWithDirs() {
	paths := []string{
		"./testdata/test-dir-create1/test-file-create1.txt",
		"./testdata/test-dir-create2/test-file-create2.txt",
	}

	_ = CreateFilesWithDirs(paths...)

	for _, path := range paths {
		fmt.Println(FileExists(path))
	}

	for _, path := range paths {
		_ = DeleteDirs(filepath.Dir(path))
	}

	// Output:
	// true <nil>
	// true <nil>
}

func ExampleOverwriteFilesWithDirs() {
	paths := []string{
		"./testdata/test-dir-create1/test-file-create1.txt",
		"./testdata/test-dir-create2/test-file-create2.txt",
	}

	_ = OverwriteFilesWithDirs(paths...)

	for _, path := range paths {
		fmt.Println(FileExists(path))
	}

	for _, path := range paths {
		_ = DeleteDirs(filepath.Dir(path))
	}

	// Output:
	// true <nil>
	// true <nil>
}

func ExampleFindFileWalk() {
	path := "./testdata/"

	fmt.Println(FindFileWalk(path, "test-file-exist.txt"))

	// Output:
	// true <nil>
}

func ExampleFindFileWalkFilter() {
	path := "./testdata/"

	fmt.Println(FindFileWalkFilter(path, func(path string, entry os.DirEntry) bool {
		return entry.Name() == "test-file-exist.txt"
	}))

	// Output:
	// true <nil>
}

func ExampleFilenames() {
	fmt.Println(Filenames("./testdata"))

	// Output:
	// [test-file-exist.txt test-file-zip.zip] <nil>
}

func ExampleFilenamesFilter() {
	fmt.Println(FilenamesFilter("./testdata", func(path string, entry os.DirEntry) bool {
		return strings.Contains(entry.Name(), "test-file-exist")
	}))

	// Output:
	// [test-file-exist.txt test-file-zip.zip] <nil>
}

func ExampleFilenamesBy() {
	fmt.Println(FilenamesBy("./testdata", func(path string, entry os.DirEntry) string {
		return "by-" + entry.Name()
	}))

	// Output:
	// [by-test-file-exist.txt] <nil>
}

func ExampleFilenamesWalk() {
	fmt.Println(FilenamesWalk("./testdata"))

	// Output:
	// [test-file-exist.txt test-file-zip.zip] <nil>
}

func ExampleFilenamesWalkFilter() {
	fmt.Println(FilenamesWalkFilter("./testdata/", func(path string, entry os.DirEntry) bool {
		return strings.Contains(entry.Name(), "test-file-exist")
	}))

	// Output:
	// [test-file-exist.txt] <nil>
}

func ExampleFilenamesWalkBy() {
	fmt.Println(FilenamesWalkBy("./testdata/", func(path string, entry os.DirEntry) string {
		return "by-" + entry.Name()
	}))

	// Output:
	// [by-test-file-exist.txt by-test-file-zip.zip] <nil>
}

func ExampleDeleteFiles() {
	paths := []string{
		"./testdata/test-file-delete1.txt",
		"./testdata/test-file-delete2.txt",
	}

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
}

func ExampleDeleteDirs() {
	paths := []string{
		"./testdata/test-dir-delete1",
		"./testdata/test-dir-delete2",
	}

	_ = CreateDirs(paths...)

	for _, path := range paths {
		fmt.Println(DirExists(path))
	}

	_ = DeleteDirs(paths...)

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
	paths := []string{
		"./testdata/test-dir-delete/test-dirs-delete1",
		"./testdata/test-dir-delete/test-dirs-delete2",
	}

	_ = CreateDirs(paths...)

	for _, path := range paths {
		fmt.Println(DirExists(path))
	}

	_ = DeleteEmptyDirWalk("./testdata/test-dir-delete")

	for _, path := range paths {
		fmt.Println(DirExists(path))
	}

	// Output:
	// true <nil>
	// true <nil>
	// false <nil>
	// false <nil>
}

func ExampleDeleteWalkBy() {
	filePaths := []string{
		"./testdata/test-dir-delete/test-dir-delete/test-dirs-delete.txt",
		"./testdata/test-dir-delete/test-dir-delete1/test-dirs-delete.txt",
	}

	dirPath := "./testdata/test-dir-delete/test-dir-delete2"

	_ = CreateFilesWithDirs(filePaths...)
	for _, path := range filePaths {
		fmt.Println(FileExists(path))
	}

	_ = CreateDirs(dirPath)
	fmt.Println(DirExists(dirPath))

	isDelete, _ := DeleteWalkBy("./testdata/test-dir-delete", func(path string, entry os.DirEntry) (bool, error) {
		return strings.Contains(path, "test-dir-delete1"), nil
	}, true)

	fmt.Println(isDelete)

	fmt.Println(DirExists("./testdata/test-dir-delete"))
	fmt.Println(DirExists("./testdata/test-dir-delete1"))
	fmt.Println(DirExists(dirPath))

	_ = DeleteDirs("./testdata/test-dir-delete")

	// Output:
	// true <nil>
	// true <nil>
	// true <nil>
	// false
	// true <nil>
	// false <nil>
	// false <nil>
}

func ExampleZip() {
	target := "./testdata/test-file-zip.zip"

	_ = Zip("./testdata", target)

	fmt.Println(FileExists(target))

	// Output:
	// true <nil>
}

func ExampleZipFilter() {

	target := "./testdata/test-file-zip.zip"

	_ = ZipFilter("./testdata", target, func(path string, entry os.DirEntry) bool {
		return entry.Name() != "test-file-zip.zip"
	})

	fmt.Println(FileExists(target))

	// Output:
	// true <nil>
}

func ExampleReadAll() {

	path := "./testdata/test-file-data.txt"

	data := "test"

	_ = CreateFileWithData(path, data)

	data2, _ := ReadAll(path)

	fmt.Println(data2)

	// Output:
	// test

	_ = DeleteFiles(path)
}

func ExampleReadLine() {

	path := "./testdata/test-file-data.txt"

	data := "test\ntest"

	_ = CreateFileWithData(path, data)

	data2, _ := ReadLine(path, 1)

	fmt.Println(data2)

	// Output:
	// [test]

	_ = DeleteFiles(path)
}
