# Fileutil

Package fileutil provides some tool functions for file operations, including file traversal, file creation, file deletion, etc.

## Usage

```go
import "github.com/dengrandpa/jez/fileutil"
```

## Index

-   [FilterMap](#filterMap)
-   [FilterMapWalk](#filterMapWalk)
-   [IsDir](#isDir)
-   [IsEmptyDir](#isEmptyDir)
-   [FileExists](#fileExists)
-   [DirExists](#dirExists)
-   [OsCreate](#osCreate)
-   [CreateFiles](#createFiles)
-   [OverwriteFiles](#overwriteFiles)
-   [CreateDirs](#createDirs)
-   [CreateFilesWithDirs](#createFilesWithDirs)
-   [OverwriteFilesWithDirs](#overwriteFilesWithDirs)
-   [CreateFileWithData](#createFileWithData)
-   [CopyFile](#copyFile)
-   [FindFileWalk](#findFileWalk)
-   [FindFileWalkFilter](#findFileWalkFilter)
-   [Filenames](#filenames)
-   [FilenamesFilter](#filenamesFilter)
-   [FilenamesBy](#filenamesBy)
-   [FilenamesWalk](#filenamesWalk)
-   [FilenamesWalkFilter](#filenamesWalkFilter)
-   [FilenamesWalkBy](#filenamesWalkBy)
-   [DeleteFiles](#deleteFiles)
-   [DeleteDirs](#deleteDirs)
-   [DeleteEmptyDirWalk](#deleteEmptyDirWalk)
-   [DeleteWalkBy](#deleteWalkBy)
-   [Zip](#zip)
-   [ZipFilter](#zipFilter)
-   [Unzip](#unzip)
-   [ReadAll](#readAll)
-   [ReadLines](#readLines)



### FilterMap
Traverse the current directory, calling iteratee on each file, and if it returns true, put the result in the result set.

```go
list, _ := fileutil.FilterMap(dir, func(entry os.DirEntry) (string, bool) {
  return entry.Name(), true
})

fmt.Println(list)

// Output:
// [test-file.txt test-file1.txt]
`````

### FilterMapWalk
Returns to traverse all directories and subdirectories, call iteratee for each file, and if it returns true, put the result into the result set.

```go
list, _ := fileutil.FilterMapWalk(path, func(path string, entry os.DirEntry) (int, bool) {
  return entry.Name(), true
})

fmt.Println(list)

// Output:
// [test-file.txt test-file1.txt]
```

### IsDir
Determine whether it is a directory.

```go
exist, _ := fileutil.IsDir(dir)
fmt.Println(exist)

// Output:
// true
```

### IsEmptyDir
Determine whether the directory is empty.

```go
exist, _ := fileutil.IsEmptyDir(dir)
fmt.Println(exist)

// Output:
// true
```

### FileExists
Determine whether the file exists.

```go
exist, _ := fileutil.FileExists(path)
fmt.Println(exist)

// Output:
// true
```

### DirExists
Determine whether the directory exists.

```go
exist, _ := fileutil.DirExists(dir)
fmt.Println(exist)

// Output:
// true
```

### OsCreate
Equivalent to os.Create, create a file. If the file already exists, it will be ignored. After use, it needs to be closed.

```go
file, _ := fileutil.OsCreate(path)
_ = file.Close()
fmt.Println(FileExists(path))

// Output:
// true <nil>
```

### CreateFiles
Create a file, if the file already exists, ignore it.

```go
_ = fileutil.CreateFiles(path)
fmt.Println(FileExists(path))

// Output:
// true <nil>
```

### OverwriteFiles
Create a file, and if the file already exists, overwrite it.

```go
_ = fileutil.OverwriteFiles(paths)
fmt.Println(FileExists(path))

// Output:
// true
```

### CreateDirs
Create a directory, including subdirectories. If the directory already exists, ignore it.

```go
_ := fileutil.CreateDirs(paths)
fmt.Println(DirExists(path))

// Output:
// true
```

### CreateFilesWithDirs
Create a file, if the file already exists, ignore it, and create a directory at the same time, including subdirectories.

```go
_ = fileutil.CreateFilesWithDirs(paths...)
fmt.Println(FileExists(path))

// Output:
// true
```

### OverwriteFilesWithDirs
Create a file, if the file already exists, overwrite it, and create a directory at the same time, including subdirectories.

```go
_ = fileutil.OverwriteFilesWithDirs(paths...)
fmt.Println(FileExists(path))

// Output:
// true
```

### CreateFileWithData
Create a file and write string data.

```go
_ = fileutil.CreateFileWithData(path, data)
data2, _ := ReadAll(path)
fmt.Println(data2)

// Output:
// test
```

### CopyFile
Copy the file.

```go
_ = fileutil.CopyFile(src, dst)
res, _ := ReadAll(dst)
fmt.Println(res)

// Output:
// test
```

### FindFileWalk
Traverse directories and subdirectories to find files.

```go
exist, _ := fileutil.FindFileWalk(dir, "test-file.txt")
fmt.Println(exist)

// Output:
// true
```

### FindFileWalkFilter
Traverse directories and subdirectories, find files, and call the iteratee function for each file. If it returns true, it means that it has been found.

```go
exist, _ := fileutil.FindFileWalkFilter(dir, func(path string, entry os.DirEntry) bool {
  return entry.Name() == "test-file.txt"
})
fmt.Println(exist)

// Output:
// true
```

### Filenames
Return to the file name slice under the directory.

```go
filenames, _ := fileutil.Filenames(dir)
fmt.Println(filenames)

// Output:
// [test-file.txt]
```

### FilenamesFilter
Traverse the files in the directory, call the iteratee function for each file, and add the file name to the slice if it returns true.

```go
filenames, _ := fileutil.FilenamesFilter(dir, func(path string, entry os.DirEntry) bool {
  return true
})
fmt.Println(filenames)

// Output:
// [test-file.txt]
```

### FilenamesBy
Traverse the files in the directory, call the iteratee function for each file, and add the returned string to the slice.

```go
filenames, _ := fileutil.FilenamesBy(dir, func(path string, entry os.DirEntry) string {
  return "by-" + entry.Name()
})
fmt.Println(filenames)

// Output:
// [by-test-file.txt]
```

### FilenamesWalk
Return the file name slice under the directory, including the subdirectory.

```go
filenames, _ := fileutil.FilenamesWalk(dir)
fmt.Println(filenames)

// Output:
// [test-file.txt test-file1.txt]
```

### FilenamesWalkFilter
Return the files in the directory, including subdirectories, call the iteratee function for each file, and add the file name to the slice if it returns true.

```go
filenames, _ := fileutil.FilenamesWalkFilter(dir, func(path string, entry os.DirEntry) bool {
		return true
	})
fmt.Println(filenames)

// Output:
// [test-file.txt test-file1.txt]
```

### FilenamesWalkBy
Return the files in the directory, including subdirectories, call the iteratee function for each file, and add the returned string to the slice.

```go
filenames, _ := fileutil.FilenamesWalkBy(dir, func(path string, entry os.DirEntry) string {
  return "by-" + entry.Name()
})
fmt.Println(filenames)

// Output:
// [by-test-file.txt by-test-file1.txt]
```

### DeleteFiles
Delete the file.

```go
_ = fileutil.DeleteFiles(paths)
fmt.Println(FileExists(paths))

// Output:
// false <nil>
```

### DeleteDirs
Delete the directory.

```go
_ = fileutil.DeleteDirs(paths)
fmt.Println(DirExists(paths))

// Output:
// false <nil>
```

### DeleteEmptyDirWalk
Return to delete the empty directory, including subdirectories, for example: /a/b. When deleting b, if a is also an empty directory, it will also be deleted.

```go
_ = fileutil.DeleteEmptyDirWalk(dir)
fmt.Println(DirExists(dir))

// Output:
// false <nil>
```

### DeleteWalkBy
Recursively delete files and subdirectories in the specified directory, itratee: used to process the logic of each file (excluding the directory), receive the file path and the os.DirEntry instance as parameters.

```go
isDelete, _ := fileutil.DeleteWalkBy(dir, func(path string, entry os.DirEntry) (bool, error) {
  return true, nil
}, true)
fmt.Println(isDelete)
fmt.Println(DirExists(dir))

// Output:
// true <nil>
// true <nil>
```

### Zip
Compress the directory or file into a zip file, and if zip already exists, it will be overwritten.

```go
_ = fileutil.Zip("./test-dir-walk", target)
fmt.Println(FileExists(target))

// Output:
// true <nil>

```

### ZipFilter
Call the iteratee function for each file or directory. If it returns true, it will be compressed into a zip file. If the zip file already exists, it will be overwritten.

```go
_ = fileutil.ZipFilter(".", target, func(path string, entry os.DirEntry) bool {
  return path == "./test-dir-walk"
})
fmt.Println(FileExists(target))

// Output:
// true <nil>
```

### Unzip
Unzip the zip file to the specified directory, and if the directory does not exist, it will be created.

```go
_ = fileutil.Unzip(src, dst)
fmt.Println(DirExists(dst))

// Output:
// true <nil>
```

### ReadAll
Read all the contents of the file as a string.

```go
data2, _ := fileutil.ReadAll(path)
fmt.Println(data2)

// Output:
// test
```

### ReadLines
Read the first n lines of the file, and if n < 0, read all lines.

```go
data2, _ := fileutil.ReadLines(path, 1)
fmt.Println(data2)

// Output:
// [test]
```
