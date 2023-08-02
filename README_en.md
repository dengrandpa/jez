# jez - Some simple and commonly used packages...

![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/dengrandpa/jez)
[![codecov](https://codecov.io/gh/dengrandpa/jez/branch/main/graph/badge.svg?token=E3LRLIYTK2)](https://codecov.io/gh/dengrandpa/jez)
[![Go Report Card](https://goreportcard.com/badge/github.com/dengrandpa/jez)](https://goreportcard.com/report/github.com/dengrandpa/jez)
[![test](https://github.com/dengrandpa/jez/actions/workflows/workflows.yml/badge.svg)](https://github.com/dengrandpa/jez/actions/workflows/workflows.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/dengrandpa/jez.svg)](https://pkg.go.dev/github.com/dengrandpa/jez)
![GitHub](https://img.shields.io/github/license/dengrandpa/jez)


English | [简体中文](./README.md)

### Special thanks to ChatGPT for providing technical support! !

## Install

```sh
go get github.com/dengrandpa/jez
```

## Usage

Import the corresponding package name when using it. For example:

### slice package
```go
import "github.com/dengrandpa/sliceutil"
```

### map package
```go
import "github.com/dengrandpa/maputil"
```

### …………

## Document
### Index
-   [Fileutil](./docs/fileutil.md)
    -   [FilterMap](./docs/fileutil.md#filterMap)
    -   [FilterMapWalk](./docs/fileutil.md#filterMapWalk)
    -   [IsDir](./docs/fileutil.md#isDir)
    -   [IsEmptyDir](./docs/fileutil.md#isEmptyDir)
    -   [FileExists](./docs/fileutil.md#fileExists)
    -   [DirExists](./docs/fileutil.md#dirExists)
    -   [OsCreate](./docs/fileutil.md#osCreate)
    -   [CreateFiles](./docs/fileutil.md#createFiles)
    -   [OverwriteFiles](./docs/fileutil.md#overwriteFiles)
    -   [CreateDirs](./docs/fileutil.md#createDirs)
    -   [CreateFilesWithDirs](./docs/fileutil.md#createFilesWithDirs)
    -   [OverwriteFilesWithDirs](./docs/fileutil.md#overwriteFilesWithDirs)
    -   [CreateFileWithData](./docs/fileutil.md#createFileWithData)
    -   [CopyFile](./docs/fileutil.md#copyFile)
    -   [FindFileWalk](./docs/fileutil.md#findFileWalk)
    -   [FindFileWalkFilter](./docs/fileutil.md#findFileWalkFilter)
    -   [Filenames](./docs/fileutil.md#filenames)
    -   [FilenamesFilter](./docs/fileutil.md#filenamesFilter)
    -   [FilenamesBy](./docs/fileutil.md#filenamesBy)
    -   [FilenamesWalk](./docs/fileutil.md#filenamesWalk)
    -   [FilenamesWalkFilter](./docs/fileutil.md#filenamesWalkFilter)
    -   [FilenamesWalkBy](./docs/fileutil.md#filenamesWalkBy)
    -   [DeleteFiles](./docs/fileutil.md#deleteFiles)
    -   [DeleteDirs](./docs/fileutil.md#deleteDirs)
    -   [DeleteEmptyDirWalk](./docs/fileutil.md#deleteEmptyDirWalk)
    -   [DeleteWalkBy](./docs/fileutil.md#deleteWalkBy)
    -   [Zip](./docs/fileutil.md#zip)
    -   [ZipFilter](./docs/fileutil.md#zipFilter)
    -   [Unzip](./docs/fileutil.md#unzip)
    -   [ReadAll](./docs/fileutil.md#readAll)
    -   [ReadLines](./docs/fileutil.md#readLines)
