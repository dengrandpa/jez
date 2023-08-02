// Package fileutil 文件相关函数
package fileutil

import (
	"archive/zip"
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
)

var walkFind = errors.New("find")

// 等同于 os.ReadDir， 删除了排序
func osReadDir(dirPath string) ([]os.DirEntry, error) {
	f, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return f.ReadDir(-1)
}

// FilterMap 遍历当前目录，对每个文件调用 iteratee，如果返回 true，则将结果放入结果集中
func FilterMap[T any](dirPath string, iteratee func(entry os.DirEntry) (T, bool)) ([]T, error) {
	entries, err := osReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	result := make([]T, 0, len(entries))

	for _, entry := range entries {
		if r, ok := iteratee(entry); ok {
			result = append(result, r)
		}
	}

	return result, nil
}

// FilterMapWalk 返回遍历所有目录、子目录，对每个文件调用 iteratee，如果返回 true，则将结果放入结果集中
func FilterMapWalk[T any](dirPath string, iteratee func(path string, d os.DirEntry) (T, bool)) ([]T, error) {

	var result []T

	err := filepath.WalkDir(dirPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if r, ok := iteratee(path, d); ok {
			result = append(result, r)
		}
		return nil
	})

	return result, err
}

// IsDir 判断是否是目录
func IsDir(dirPath string) (bool, error) {
	info, err := os.Stat(dirPath)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

// IsEmptyDir 判断目录是否为空
func IsEmptyDir(dirPath string) (bool, error) {
	entries, err := osReadDir(dirPath)
	if err != nil {
		return false, err
	}

	return len(entries) == 0, nil
}

// FileExists 判断文件是否存在
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

// DirExists 判断目录是否存在
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

// OsCreate 等同于 os.Create,创建文件，如果文件已存在，则忽略，使用完毕后需要关闭
func OsCreate(filePath string) (*os.File, error) {
	return os.Create(filePath)
}

// CreateFiles 创建文件，如果文件已存在，则忽略
func CreateFiles(filePaths ...string) error {
	for _, filePath := range filePaths {
		file, err := OsCreate(filePath)
		if err != nil {
			return err
		}
		_ = file.Close()
	}
	return nil
}

// OverwriteFiles 创建文件，如果文件已存在，则覆盖
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

// CreateDirs 创建目录，包含子目录，如果目录已存在，则忽略
func CreateDirs(dirPaths ...string) error {
	for _, dirPath := range dirPaths {
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// CreateFilesWithDirs 创建文件，如果文件已存在，则忽略，同时创建目录，包含子目录
func CreateFilesWithDirs(filePaths ...string) error {
	for _, filePath := range filePaths {
		if err := CreateDirs(filepath.Dir(filePath)); err != nil {
			return err
		}
	}
	return CreateFiles(filePaths...)
}

// OverwriteFilesWithDirs 创建文件，如果文件已存在，则覆盖，同时创建目录，包含子目录
func OverwriteFilesWithDirs(filePaths ...string) error {
	for _, filePath := range filePaths {
		if err := CreateDirs(filepath.Dir(filePath)); err != nil {
			return err
		}
	}
	return OverwriteFiles(filePaths...)
}

// CreateFileWithData 创建文件并写入字符串数据
func CreateFileWithData(filePath, data string) error {
	file, err := OsCreate(filePath)
	if err != nil {
		return err
	}

	_, err = file.WriteString(data)
	_ = file.Close()

	return err
}

// CopyFile 拷贝文件
func CopyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := OsCreate(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

// FindFileWalk 遍历目录、子目录，查找文件
func FindFileWalk(dirPath, filename string) (bool, error) {

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

// FindFileWalkFilter 遍历目录、子目录，查找文件，对每个文件调用 iteratee 函数，如果返回 true，则表示找到了
func FindFileWalkFilter(dirPath string, iteratee func(path string, entry os.DirEntry) bool) (bool, error) {

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

// Filenames 返回目录下的文件名切片
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

// FilenamesFilter 遍历目录下的文件，对每个文件调用 iteratee 函数，如果返回 true，则将文件名添加到切片中
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

// FilenamesBy 遍历目录下的文件，对每个文件调用 iteratee 函数，将返回的字符串添加到切片中
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

// FilenamesWalk 返回目录下的文件名切片，包含子目录
func FilenamesWalk(dirPath string) ([]string, error) {
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

// FilenamesWalkFilter 返回目录下的文件，包含子目录，对每个文件调用 iteratee 函数，如果返回 true，则将文件名添加到切片中
func FilenamesWalkFilter(dirPath string, iteratee func(path string, entry os.DirEntry) bool) ([]string, error) {
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

// FilenamesWalkBy 返回目录下的文件，包含子目录，对每个文件调用 iteratee 函数，将返回的字符串添加到切片中
func FilenamesWalkBy(dirPath string, iteratee func(path string, entry os.DirEntry) string) ([]string, error) {
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

// DeleteFiles 删除文件
func DeleteFiles(filePaths ...string) error {
	for _, filePath := range filePaths {
		if err := os.Remove(filePath); err != nil {
			return err
		}
	}
	return nil
}

// DeleteDirs 删除目录
func DeleteDirs(dirPaths ...string) error {
	for _, dirPath := range dirPaths {
		if err := os.RemoveAll(dirPath); err != nil {
			return err
		}
	}
	return nil
}

// DeleteEmptyDirWalk 返回删除空目录，包含子目录
//
// 例如：
//   - /a/b，当删除 b 时，如果 a 也是一个空目录，也会被删除
func DeleteEmptyDirWalk(dirPath string) error {
	_, err := DeleteWalkBy(dirPath, nil, true)
	return err
}

// DeleteWalkBy 递归删除指定目录下的文件和子目录
//
// 参数：
//   - dirPath: 目录路径
//   - iteratee: 遍历函数，用于处理每个文件（不包括目录）的逻辑，接收文件路径和 os.DirEntry 实例作为参数
//   - withEmptyDir: 可选参数，指定是否删除空目录，默认为 false
//
// 返回值：
//   - bool: 当前目录是否已被删除
//   - error: 错误
//
// 注意事项：
//   - 如果要删除空目录，请注意 bool 返回值，错误的 bool 返回值可能导致文件意外删除或空目录删除失败。
func DeleteWalkBy(
	dirPath string, iteratee func(path string, entry os.DirEntry) (bool, error), withEmptyDir ...bool) (bool, error) {

	var withEmpty bool
	if len(withEmptyDir) > 0 {
		withEmpty = withEmptyDir[0]
	}

	entries, err := osReadDir(dirPath)
	if err != nil {
		return false, err
	}

	// 当前目录中剩余文件的数量
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

			// 递归子目录
			if isDelete, err = DeleteWalkBy(path, iteratee, withEmptyDir...); err != nil {
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

// Zip 将目录或文件压缩为 zip 文件，如果zip已存在，则会被覆盖。
func Zip(src, dst string) error {
	zipFile, err := OsCreate(dst)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	return filepath.WalkDir(src, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// 在 zip 文件中创建文件或目录
		zipPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		// 如果是目录，则创建目录
		if info.IsDir() {
			_, err = zipWriter.Create(zipPath + "/")
			if err != nil {
				return err
			}
		} else {
			// 如果是文件，则创建文件并将文件内容写入 zip 文件
			var file *os.File
			if file, err = os.Open(path); err != nil {
				return err
			}
			defer file.Close()

			var writer io.Writer
			if writer, err = zipWriter.Create(zipPath); err != nil {
				return err
			}

			if _, err = io.Copy(writer, file); err != nil {
				return err
			}
		}

		return nil
	})
}

// ZipFilter 对每个文件或目录调用 iteratee 函数，如果返回 true，则将其压缩到 zip 文件中，如果zip文件已存在，则会被覆盖。
func ZipFilter(src, dst string, iteratee func(path string, entry os.DirEntry) bool) error {
	zipFile, err := OsCreate(dst)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	return filepath.WalkDir(src, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !iteratee(path, info) {
			return nil
		}

		// 在 zip 文件中创建文件或目录
		zipPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		// 如果是文件，则创建文件并将文件内容写入 zip 文件
		if info.IsDir() {
			_, err = zipWriter.Create(zipPath + "/")
			if err != nil {
				return err
			}
		} else {
			// 如果是文件，则创建文件并将文件内容写入 zip 文件
			var file *os.File
			if file, err = os.Open(path); err != nil {
				return err
			}
			defer file.Close()

			var writer io.Writer
			if writer, err = zipWriter.Create(zipPath); err != nil {
				return err
			}

			if _, err = io.Copy(writer, file); err != nil {
				return err
			}
		}

		return nil
	})
}

// Unzip 解压 zip 文件到指定目录，如果目录不存在，则会被创建。
func Unzip(src, dst string) error {

	reader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer reader.Close()

	// 遍历 zip 文件中的每个文件/目录
	for _, file := range reader.File {
		// 构建解压文件路径
		extractedFilePath := filepath.Join(dst, file.Name)

		if file.FileInfo().IsDir() {
			if err = CreateDirs(extractedFilePath); err != nil {
				return err
			}
		} else {
			if err = CreateDirs(filepath.Dir(extractedFilePath)); err != nil {
				return err
			}

			var zFile io.ReadCloser
			if zFile, err = file.Open(); err != nil {
				return err
			}
			defer zFile.Close()

			var extractedFile *os.File
			extractedFile, err = os.OpenFile(extractedFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				return err
			}
			defer extractedFile.Close()

			_, err = io.Copy(extractedFile, zFile)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

// ReadAll 将文件的所有内容读取为字符串。
func ReadAll(filePath string) (string, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// ReadLines 读取文件的前 n 行，如果 n < 0，则读取所有行。
func ReadLines(filePath string, n int) ([]string, error) {
	if n == 0 {
		return []string{}, nil
	}

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		n--

		if n == 0 {
			break
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
