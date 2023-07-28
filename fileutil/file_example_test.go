package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

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

func ExampleFindFileRecursively() {
	path := "./testdata/"

	fmt.Println(FindFileRecursively(path, "test-file-exist.txt"))

	// Output:
	// true <nil>
}

func ExampleFindFileRecursivelyFilter() {
	path := "./testdata/"

	fmt.Println(FindFileRecursivelyFilter(path, func(path string, entry os.DirEntry) bool {
		return entry.Name() == "test-file-exist.txt"
	}))

	// Output:
	// true <nil>
}

func ExampleFilenames() {
	fmt.Println(Filenames("./testdata"))

	// Output:
	// [test-file-exist.txt] <nil>
}

func ExampleFilenamesFilter() {
	fmt.Println(FilenamesFilter("./testdata", func(path string, entry os.DirEntry) bool {
		return strings.Contains(entry.Name(), "test-file-exist")
	}))

	// Output:
	// [test-file-exist.txt] <nil>
}

func ExampleFilenamesBy() {
	fmt.Println(FilenamesBy("./testdata", func(path string, entry os.DirEntry) string {
		return "by-" + entry.Name()
	}))

	// Output:
	// [by-test-file-exist.txt] <nil>
}

func ExampleFilenamesRecursively() {
	fmt.Println(FilenamesRecursively("./testdata"))

	// Output:
	// [test-file-exist.txt] <nil>
}

func ExampleFilenamesRecursivelyFilter() {
	fmt.Println(FilenamesRecursivelyFilter("./testdata/", func(path string, entry os.DirEntry) bool {
		return strings.Contains(entry.Name(), "test-file-exist")
	}))

	// Output:
	// [test-file-exist.txt] <nil>
}

func ExampleFilenamesRecursivelyBy() {
	fmt.Println(FilenamesRecursivelyBy("./testdata/", func(path string, entry os.DirEntry) string {
		return "by-" + entry.Name()
	}))

	// Output:
	// [by-test-file-exist.txt] <nil>
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

func ExampleDeleteEmptyDirRecursively() {
	paths := []string{
		"./testdata/test-dir-delete/test-dirs-delete1",
		"./testdata/test-dir-delete/test-dirs-delete2",
	}

	_ = CreateDirs(paths...)

	for _, path := range paths {
		fmt.Println(DirExists(path))
	}

	_ = DeleteEmptyDirRecursively("./testdata/test-dir-delete")

	for _, path := range paths {
		fmt.Println(DirExists(path))
	}

	// Output:
	// true <nil>
	// true <nil>
	// false <nil>
	// false <nil>
}

func ExampleDeleteRecursivelyBy() {
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

	isDelete, _ := DeleteRecursivelyBy("./testdata/test-dir-delete", func(path string, entry os.DirEntry) (bool, error) {
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
