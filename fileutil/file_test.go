package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterMap(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	path := "./testdata"

	resInt, _ := FilterMap(path, func(entry os.DirEntry) (int, bool) {
		if !entry.IsDir() {
			return 1, true
		}

		return 0, false
	})

	ass.Equal([]int{1, 1}, resInt)

}

func TestFilterMapWalk(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	path := "./"

	resInt, _ := FilterMapWalk(path, func(path string, entry os.DirEntry) (int, bool) {
		if !entry.IsDir() {
			return 1, true
		}

		return 0, false
	})

	ass.Equal([]int{1, 1, 1, 1, 1}, resInt)
}

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

func TestCreateFileWithData(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	path := "./testdata/test-file-data.txt"

	data := "test"

	_ = CreateFileWithData(path, data)

	data2, _ := ReadAll(path)

	ass.Equal(data, data2)

	_ = DeleteFiles(path)
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

func TestCreateFileWithOS(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	path := "./testdata/test-file-createwithos.txt"

	file, _ := OsCreate(path)

	ass.FileExists(path)

	file.Close()

	_ = DeleteFiles(path)

}

func TestCopyFile(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	src := "./testdata/test-file-exist.txt"
	dst := "./testdata/test-file-exist2.txt"

	_ = CopyFile(src, dst)

	ass.FileExists(dst)

	_ = DeleteFiles(dst)

}

func TestFindFileWalk(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	path := "./testdata/"

	exist, err := FindFileWalk(path, "test-file-exist.txt")

	ass.Nil(err)
	ass.True(exist)

}

func TestFindFileWalkFilter(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	path := "./testdata/"

	exist, err := FindFileWalkFilter(path, func(path string, entry os.DirEntry) bool {
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

func TestFilenamesWalk(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	filenames, err := FilenamesWalk("./testdata/")

	ass.Nil(err)

	ass.Contains(filenames, "test-file-exist.txt")
}

func TestFilenamesWalkFilter(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	filenames, err := FilenamesWalkFilter("./testdata/", func(path string, entry os.DirEntry) bool {
		return entry.Name() == "test-file-exist.txt"
	})
	ass.Nil(err)

	ass.Equal(filenames[0], "test-file-exist.txt")
}

func TestFilenamesWalkBy(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	filenames, err := FilenamesWalkBy("./testdata/", func(path string, entry os.DirEntry) string {
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

func TestDeleteEmptyDirWalk(t *testing.T) {
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

	ass.Nil(DeleteEmptyDirWalk("./testdata/test-dir-delete"))

	for _, path := range paths {
		ass.NoDirExists(path)
	}
}

func TestDeleteWalkBy(t *testing.T) {
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

	isDelete, err := DeleteWalkBy("./testdata/test-dir-delete", func(path string, entry os.DirEntry) (bool, error) {
		return strings.Contains(path, "test-dir-delete1"), nil
	}, true)

	ass.Nil(err)
	ass.False(isDelete)

	ass.DirExists("./testdata/test-dir-delete")
	ass.NoDirExists("./testdata/test-dir-delete1")
	ass.NoDirExists(dirPath)

	_ = DeleteDirs("./testdata/test-dir-delete")
}

func TestZip(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	target := "./testdata/test-file-zip.zip"

	_ = Zip("./testdata", target)

	ass.FileExists(target)
}

func TestZipFilter(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	target := "./testdata/test-file-zip.zip"

	_ = ZipFilter("./testdata", target, func(path string, entry os.DirEntry) bool {
		return entry.Name() != "test-file-zip.zip"
	})

	ass.FileExists(target)
}

func TestUnzip(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	src := "./testdata/test-file-zip.zip"

	dst := "./testdata/unzip"

	_ = Unzip(src, dst)

	ass.DirExists(dst)

	_ = DeleteDirs(dst)
}

func TestReadAll(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	path := "./testdata/test-file-data.txt"

	data := "test"

	_ = CreateFileWithData(path, data)

	data2, _ := ReadAll(path)

	ass.Equal(data, data2)

	_ = DeleteFiles(path)
}

func TestReadLine(t *testing.T) {
	t.Parallel()
	ass := assert.New(t)

	path := "./testdata/test-file-data.txt"

	data := "test\ntest"

	_ = CreateFileWithData(path, data)

	data2, _ := ReadLine(path, 1)

	ass.Equal([]string{"test"}, data2)

	_ = DeleteFiles(path)
}