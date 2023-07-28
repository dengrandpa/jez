package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDir(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	path := "./testdata"
	exist, err := IsDir(path)

	ass.Nil(err)
	ass.Equal(ass.DirExists(path), exist)
}

func TestIsEmptyDir(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	path := "./testdata/test-dir-empty"

	exist, err := IsEmptyDir(path)
	ass.Nil(err)
	ass.True(exist)
}

func TestFileExists(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	path := "./testdata/test-file-exist.txt"

	exist, err := FileExists(path)
	ass.Nil(err)
	ass.Equal(ass.FileExists(path), exist)
}

func TestDirExists(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	path := "./testdata"

	exist, err := DirExists(path)
	ass.Nil(err)

	ass.Equal(ass.DirExists(path), exist)
}

func TestCreateFiles(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	paths := []string{
		"./testdata/test-file-create1.txt",
		"./testdata/test-file-create2.txt",
	}

	ass.Nil(CreateFiles(paths...))

	for _, path := range paths {
		ass.FileExists(path)
	}

	_ = DeleteFiles(paths...)

}

func TestOverwriteFiles(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	paths := []string{
		"./testdata/test-file-create1.txt",
		"./testdata/test-file-create2.txt",
	}

	ass.Nil(OverwriteFiles(paths...))

	for _, path := range paths {
		ass.FileExists(path)
	}

	_ = DeleteFiles(paths...)
}

func TestCreateDirs(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	paths := []string{
		"./testdata/test-dir-create1",
		"./testdata/test-dir-create2",
	}

	ass.Nil(CreateDirs(paths...))

	for _, path := range paths {
		ass.DirExists(path)
	}

	_ = DeleteDirs(paths...)
}

func TestCreateFilesWithDirs(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	paths := []string{
		"./testdata/test-dir-create1/test-file-create1.txt",
		"./testdata/test-dir-create2/test-file-create2.txt",
	}

	ass.Nil(CreateFilesWithDirs(paths...))

	for _, path := range paths {
		ass.FileExists(path)
	}

	fmt.Println()

	for _, path := range paths {
		_ = DeleteDirs(filepath.Dir(path))
	}
}

func TestOverwriteFilesWithDirs(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	paths := []string{
		"./testdata/test-dir-create1/test-file-create1.txt",
		"./testdata/test-dir-create2/test-file-create2.txt",
	}

	ass.Nil(OverwriteFilesWithDirs(paths...))

	for _, path := range paths {
		ass.FileExists(path)
	}

	for _, path := range paths {
		_ = DeleteDirs(filepath.Dir(path))
	}
}

func TestFindFileRecursively(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	path := "./testdata/"

	exist, err := FindFileRecursively(path, "test-file-exist.txt")

	ass.Nil(err)
	ass.True(exist)

}

func TestFindFileRecursivelyFilter(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	path := "./testdata/"

	exist, err := FindFileRecursivelyFilter(path, func(path string, entry os.DirEntry) bool {
		return entry.Name() == "test-file-exist.txt"
	})

	ass.Nil(err)
	ass.True(exist)
}

func TestFilenames(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	filenames, err := Filenames("./testdata")

	ass.Nil(err)

	ass.Contains(filenames, "test-file-exist.txt")
}

func TestFilenamesFilter(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	filenames, err := FilenamesFilter("./testdata", func(path string, entry os.DirEntry) bool {
		return entry.Name() == "test-file-exist.txt"
	})
	ass.Nil(err)
	ass.Equal(filenames[0], "test-file-exist.txt")
}

func TestFilenamesBy(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	filenames, err := FilenamesBy("./testdata", func(path string, entry os.DirEntry) string {
		return "by-" + entry.Name()
	})
	ass.Nil(err)

	ass.Contains(filenames, "by-test-file-exist.txt")

}

func TestFilenamesRecursively(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	filenames, err := FilenamesRecursively("./testdata/")

	ass.Nil(err)

	ass.Contains(filenames, "test-file-exist.txt")
}

func TestFilenamesRecursivelyFilter(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	filenames, err := FilenamesRecursivelyFilter("./testdata/", func(path string, entry os.DirEntry) bool {
		return entry.Name() == "test-file-exist.txt"
	})
	ass.Nil(err)

	ass.Equal(filenames[0], "test-file-exist.txt")
}

func TestFilenamesRecursivelyBy(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	filenames, err := FilenamesRecursivelyBy("./testdata/", func(path string, entry os.DirEntry) string {
		return "by-" + entry.Name()
	})
	ass.Nil(err)

	ass.Contains(filenames, "by-test-file-exist.txt")
}

func TestDeleteFiles(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	paths := []string{
		"./testdata/test-file-delete1.txt",
		"./testdata/test-file-delete2.txt",
	}

	ass.Nil(CreateFiles(paths...))

	for _, path := range paths {
		ass.FileExists(path)
	}

	ass.Nil(DeleteFiles(paths...))

	for _, path := range paths {
		ass.NoFileExists(path)
	}

}

func TestDeleteDirs(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	paths := []string{
		"./testdata/test-dir-delete1",
		"./testdata/test-dir-delete2",
	}

	ass.Nil(CreateDirs(paths...))

	for _, path := range paths {
		ass.DirExists(path)
	}

	ass.Nil(DeleteDirs(paths...))

	for _, path := range paths {
		ass.NoDirExists(path)
	}
}

func TestDeleteEmptyDirRecursively(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	paths := []string{
		"./testdata/test-dir-delete/test-dirs-delete1",
		"./testdata/test-dir-delete/test-dirs-delete2",
	}

	ass.Nil(CreateDirs(paths...))

	for _, path := range paths {
		ass.DirExists(path)
	}

	ass.Nil(DeleteEmptyDirRecursively("./testdata/test-dir-delete"))

	for _, path := range paths {
		ass.NoDirExists(path)
	}
}

func TestDeleteRecursivelyBy(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	filePaths := []string{
		"./testdata/test-dir-delete/test-dir-delete/test-dirs-delete.txt",
		"./testdata/test-dir-delete/test-dir-delete1/test-dirs-delete.txt",
	}

	dirPath := "./testdata/test-dir-delete/test-dir-delete2"

	ass.Nil(CreateFilesWithDirs(filePaths...))
	for _, path := range filePaths {
		ass.FileExists(path)
	}

	ass.Nil(CreateDirs(dirPath))
	ass.DirExists(dirPath)

	isDelete, err := DeleteRecursivelyBy("./testdata/test-dir-delete", func(path string, entry os.DirEntry) (bool, error) {
		return strings.Contains(path, "test-dir-delete1"), nil
	}, true)

	ass.Nil(err)
	ass.False(isDelete)

	ass.DirExists("./testdata/test-dir-delete")
	ass.NoDirExists("./testdata/test-dir-delete1")
	ass.NoDirExists(dirPath)

	_ = DeleteDirs("./testdata/test-dir-delete")
}
