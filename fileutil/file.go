// Package fileutil provides some file utility functions.
package fileutil

import (
	"errors"
	"os"
	"path/filepath"
)

var walkFind = errors.New("find")

// removes sorting of os.ReadDir.
func osReadDir(dirPath string) ([]os.DirEntry, error) {
	f, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return f.ReadDir(-1)
}

// IsDir determines whether the path is a directory.
func IsDir(dirPath string) (bool, error) {
	info, err := os.Stat(dirPath)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

// IsEmptyDir determines whether the directory is empty.
func IsEmptyDir(dirPath string) (bool, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return false, err
	}

	return len(entries) == 0, nil
}

// FileExists determines whether the file exists.
func FileExists(filePath string) (bool, error) {
	isDir, err := IsDir(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return !isDir, nil
}

// DirExists determines whether the directory exists.
func DirExists(dirPath string) (bool, error) {
	isDir, err := IsDir(dirPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return isDir, nil
}

// CreateFiles create files, ignore them if they exist.
func CreateFiles(filePaths ...string) error {
	for _, filePath := range filePaths {
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		_ = file.Close()
	}
	return nil
}

// OverwriteFiles overwrite files, create them if they do not exist.
func OverwriteFiles(filePaths ...string) error {
	for _, filePath := range filePaths {
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
		if err != nil {
			return err
		}
		_ = file.Close()
	}

	return nil
}

// CreateDirs create directory with Unix permission bits, 0o777.
func CreateDirs(dirPaths ...string) error {
	for _, dirPath := range dirPaths {
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// CreateFilesWithDirs create files, ignore them if they exist.
func CreateFilesWithDirs(filePaths ...string) error {
	for _, filePath := range filePaths {
		if err := CreateDirs(filepath.Dir(filePath)); err != nil {
			return err
		}
	}
	return CreateFiles(filePaths...)
}

// OverwriteFilesWithDirs overwrite files, create them if they do not exist.
func OverwriteFilesWithDirs(filePaths ...string) error {
	for _, filePath := range filePaths {
		if err := CreateDirs(filepath.Dir(filePath)); err != nil {
			return err
		}
	}
	return OverwriteFiles(filePaths...)
}

// FindFileRecursively recursively find files in the directory.
func FindFileRecursively(dirPath, filename string) (bool, error) {

	err := filepath.WalkDir(dirPath, func(path string, entry os.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if !entry.IsDir() && entry.Name() == filename {
			return walkFind
		}

		return nil
	})

	if err != nil && errors.Is(err, walkFind) {
		return true, nil
	}

	return false, err
}

// FindFileRecursivelyFilter recursively find the iteratee function in the directory to return true filename.
func FindFileRecursivelyFilter(dirPath string, iteratee func(path string, entry os.DirEntry) bool) (bool, error) {

	err := filepath.WalkDir(dirPath, func(path string, entry os.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if !entry.IsDir() && iteratee(path, entry) {
			return walkFind
		}

		return nil
	})

	if err != nil && errors.Is(err, walkFind) {
		return true, nil
	}

	return false, err
}

// Filenames get all file names in the directory.
func Filenames(dirPath string) ([]string, error) {
	entries, err := osReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			names = append(names, entry.Name())
		}
	}

	return names, nil
}

// FilenamesFilter get the iteratee function in the directory to return true filename.
func FilenamesFilter(dirPath string, iteratee func(path string, entry os.DirEntry) bool) ([]string, error) {
	entries, err := osReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() && iteratee(filepath.Join(dirPath, entry.Name()), entry) {
			names = append(names, entry.Name())
		}
	}

	return names, nil
}

// FilenamesBy get the iteratee function in the directory to return filename.
func FilenamesBy(dirPath string, iteratee func(path string, entry os.DirEntry) string) ([]string, error) {
	entries, err := osReadDir(dirPath)
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			names = append(names, iteratee(filepath.Join(dirPath, entry.Name()), entry))
		}
	}

	return names, nil
}

// FilenamesRecursively recursively get all file names in the directory.
func FilenamesRecursively(dirPath string) ([]string, error) {
	var names []string

	err := filepath.WalkDir(dirPath, func(path string, entry os.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if !entry.IsDir() {
			names = append(names, entry.Name())
		}

		return nil
	})

	return names, err
}

// FilenamesRecursivelyFilter recursively get the iteratee function in the directory to return true filename.
func FilenamesRecursivelyFilter(dirPath string, iteratee func(path string, entry os.DirEntry) bool) ([]string, error) {
	var names []string

	err := filepath.WalkDir(dirPath, func(path string, entry os.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if !entry.IsDir() && iteratee(path, entry) {
			names = append(names, entry.Name())
		}

		return nil
	})

	return names, err
}

// FilenamesRecursivelyBy recursively get the iteratee function in the directory to return filename.
func FilenamesRecursivelyBy(dirPath string, iteratee func(path string, entry os.DirEntry) string) ([]string, error) {
	var names []string

	err := filepath.WalkDir(dirPath, func(path string, entry os.DirEntry, err error) error {

		if err != nil {
			return err
		}

		if !entry.IsDir() {
			names = append(names, iteratee(path, entry))
		}

		return nil
	})

	return names, err
}

// DeleteFiles delete the file.
func DeleteFiles(filePaths ...string) error {
	for _, filePath := range filePaths {
		if err := os.Remove(filePath); err != nil {
			return err
		}
	}
	return nil
}

// DeleteDirs delete the directory.
func DeleteDirs(dirPaths ...string) error {
	for _, dirPath := range dirPaths {
		if err := os.RemoveAll(dirPath); err != nil {
			return err
		}
	}
	return nil
}

// DeleteEmptyDirRecursively recursively delete all empty directories.
// Note: For example "/a/b" when "b" is deleted, if "a" is also an empty directory, it will also be deleted.
func DeleteEmptyDirRecursively(dirPath string) error {
	_, err := DeleteRecursivelyBy(dirPath, nil, true)
	return err
}

// DeleteRecursivelyBy is a recursive-deleted implementation.
// It deletes files and directories recursively starting from the given directory path.
// The fileDo function is called for each file encountered during the deletion process.
// The fileDo function takes the file path and the os.DirEntry instance as parameters,
// and returns whether the current directory has been deleted.
// If an error occurs during deletion, return
// If you want to delete an empty directory, be careful with the bool return value.
// Incorrect bool return value may result in accidental deletion of a file or unsuccessful deletion of an empty directory.
//
// Parameters:
// - dirPath: The path of the directory to delete recursively.
// - fileDo: The function to be called for each file encountered during the deletion process.
// It takes the file path and the os.DirEntry instance as parameters,
// and returns a boolean value indicating whether the current directory has been deleted,
// and an error, if any.
// - withEmptyDir: Optional parameter to specify whether to delete empty directories.
// If set to true, empty directories will be deleted.
// If not provided or set to false, empty directories will not be deleted.
//
// Returns:
// - bool: Whether the current directory has been deleted.
// - error: An error if any occurred during the deletion process.
func DeleteRecursivelyBy(dirPath string, iteratee func(path string, entry os.DirEntry) (bool, error), withEmptyDir ...bool) (bool, error) {

	var withEmpty bool
	if len(withEmptyDir) > 0 {
		withEmpty = withEmptyDir[0]
	}

	entries, err := osReadDir(dirPath)
	if err != nil {
		return false, err
	}

	// the number of remaining files in the current directory
	num := len(entries)

	for _, entry := range entries {

		path := filepath.Join(dirPath, entry.Name())

		var isDelete bool

		if !entry.IsDir() {

			if iteratee == nil {
				continue
			}

			if isDelete, err = iteratee(path, entry); err != nil {
				return false, err
			}

		} else {

			// recursive subdirectories
			if isDelete, err = DeleteRecursivelyBy(path, iteratee, withEmptyDir...); err != nil {
				return false, err
			}
		}

		if isDelete {
			num--
		}
	}

	if !withEmpty || num != 0 {
		return false, nil
	}

	return true, DeleteDirs(dirPath)
}
