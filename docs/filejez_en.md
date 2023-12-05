# Filejez

------

Provides some tool functions for file operations, including file traversal, file creation, file deletion, etc.

------

## Usage

```go
import "github.com/dengrandpa/jez/filejez"
```

------

## Index

-   [FilterMap](#filtermap)
-   [FilterMapWalk](#filtermapwalk)
-   [IsDir](#isdir)
-   [IsEmptyDir](#isemptydir)
-   [FileExists](#fileexists)
-   [DirExists](#direxists)
-   [OsCreate](#oscreate)
-   [CreateFiles](#createfiles)
-   [OverwriteFiles](#overwritefiles)
-   [CreateDirs](#createdirs)
-   [CreateFilesWithDirs](#createfileswithdirs)
-   [OverwriteFilesWithDirs](#overwritefileswithdirs)
-   [CreateFileWithData](#createfilewithdata)
-   [CopyFile](#copyfile)
-   [FindFileWalk](#findfilewalk)
-   [FindFileWalkFilter](#findfilewalkfilter)
-   [Filenames](#filenames)
-   [FilenamesFilter](#filenamesfilter)
-   [FilenamesBy](#filenamesby)
-   [FilenamesWalk](#filenameswalk)
-   [FilenamesWalkFilter](#filenameswalkfilter)
-   [FilenamesWalkBy](#filenameswalkby)
-   [DeleteFiles](#deletefiles)
-   [DeleteDirs](#deletedirs)
-   [DeleteEmptyDirWalk](#deleteemptydirwalk)
-   [DeleteWalkBy](#deletewalkby)
-   [Zip](#zip)
-   [ZipFilter](#zipfilter)
-   [Unzip](#unzip)
-   [ReadAll](#readall)
-   [ReadLines](#readlines)

------

### FilterMap
Traverse the current directory, calling iteratee on each file, and if it returns true, put the result in the result set.

```go
package main

import (
	"fmt"
	"os"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  list, _ := filejez.FilterMap("dir", func(entry os.DirEntry) (string, bool) {
    return entry.Name(), true
  })
  
  fmt.Println(list)

  // Output:
  // [test-file.txt test-file1.txt]
}

```

### FilterMapWalk
Returns to traverse all directories and subdirectories, call iteratee for each file, and if it returns true, put the result into the result set.

```go
package main

import (
	"fmt"
	"os"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  list, _ := filejez.FilterMapWalk("path", func(path string, entry os.DirEntry) (string, bool) {
    return entry.Name(), true
  })
  
  fmt.Println(list)

  // Output:
  // [test-file.txt test-file1.txt]
}

```

### IsDir
Determine whether it is a directory.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  exist, _ := filejez.IsDir("dir")
  fmt.Println(exist)
  
  // Output:
  // true
}

```

### IsEmptyDir
Determine whether the directory is empty.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  exist, _ := filejez.IsEmptyDir("dir")
  fmt.Println(exist)
  
  // Output:
  // true
}

```

### FileExists
Determine whether the file exists.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  exist, _ := filejez.FileExists("path")
  fmt.Println(exist)
  
  // Output:
  // true
}

```

### DirExists
Determine whether the directory exists.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  exist, _ := filejez.DirExists("dir")
  fmt.Println(exist)
  
  // Output:
  // true
}

```

### OsCreate
Equivalent to os.Create, create a file. If the file already exists, it will be ignored. After use, it needs to be closed.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  file, _ := filejez.OsCreate("path")
  _ = file.Close()
  fmt.Println(filejez.FileExists("path"))
  
  // Output:
  // true <nil>
}

```

### CreateFiles
Create a file, if the file already exists, ignore it.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  _ = filejez.CreateFiles("path")
  fmt.Println(filejez.FileExists("path"))
  
  // Output:
  // true <nil>
}

```

### OverwriteFiles
Create a file, and if the file already exists, overwrite it.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  _ = filejez.OverwriteFiles("path")
  fmt.Println(filejez.FileExists("path"))
  
  // Output:
  // true
}

```

### CreateDirs
Create a directory, including subdirectories. If the directory already exists, ignore it.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  _ = filejez.CreateDirs("path")
  fmt.Println(filejez.DirExists("path"))
  
  // Output:
  // true
}

```

### CreateFilesWithDirs
Create a file, if the file already exists, ignore it, and create a directory at the same time, including subdirectories.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  _ = filejez.CreateFilesWithDirs("path")
  fmt.Println(filejez.FileExists("path"))
  
  // Output:
  // true
}

```

### OverwriteFilesWithDirs
Create a file, if the file already exists, overwrite it, and create a directory at the same time, including subdirectories.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  _ = filejez.OverwriteFilesWithDirs("path")
  fmt.Println(filejez.FileExists("path"))
  
  // Output:
  // true
}

```

### CreateFileWithData
Create a file and write string data.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  _ = filejez.CreateFileWithData("path", "test")
  data2, _ := filejez.ReadAll("path")
  fmt.Println(data2)
  
  // Output:
  // test
}

```

### CopyFile
Copy the file.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  _ = filejez.CopyFile("src", "dst")
  res, _ := filejez.ReadAll("dst")
  fmt.Println(res)
  
  // Output:
  // test
}

```

### FindFileWalk
Traverse directories and subdirectories to find files.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  exist, _ := filejez.FindFileWalk("dir", "test-file.txt")
  fmt.Println(exist)
  
  // Output:
  // true
}

```

### FindFileWalkFilter
Traverse directories and subdirectories, find files, and call the iteratee function for each file. If it returns true, it means that it has been found.

```go
package main

import (
	"fmt"
	"os"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  exist, _ := filejez.FindFileWalkFilter("dir", func(path string, entry os.DirEntry) bool {
    return entry.Name() == "test-file.txt"
  })
  fmt.Println(exist)
  
  // Output:
  // true
}

```

### Filenames
Return to the file name slice under the directory.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  filenames, _ := filejez.Filenames("dir")
  fmt.Println(filenames)
  
  // Output:
  // [test-file.txt]
}

```

### FilenamesFilter
Traverse the files in the directory, call the iteratee function for each file, and add the file name to the slice if it returns true.

```go
package main

import (
	"fmt"
	"os"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  filenames, _ := filejez.FilenamesFilter("dir", func(path string, entry os.DirEntry) bool {
    return true
  })
  fmt.Println(filenames)
  
  // Output:
  // [test-file.txt]
}

```

### FilenamesBy
Traverse the files in the directory, call the iteratee function for each file, and add the returned string to the slice.

```go
package main

import (
	"fmt"
	"os"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  filenames, _ := filejez.FilenamesBy("dir", func(path string, entry os.DirEntry) string {
    return "by-" + entry.Name()
  })
  fmt.Println(filenames)
  
  // Output:
  // [by-test-file.txt]
}

```

### FilenamesWalk
Return the file name slice under the directory, including the subdirectory.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  filenames, _ := filejez.FilenamesWalk("dir")
  fmt.Println(filenames)
  
  // Output:
  // [test-file.txt test-file1.txt]
}

```

### FilenamesWalkFilter
Return the files in the directory, including subdirectories, call the iteratee function for each file, and add the file name to the slice if it returns true.

```go
package main

import (
	"fmt"
	"os"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  filenames, _ := filejez.FilenamesWalkFilter("dir", func(path string, entry os.DirEntry) bool {
          return true
      })
  fmt.Println(filenames)
  
  // Output:
  // [test-file.txt test-file1.txt]
}

```

### FilenamesWalkBy
Return the files in the directory, including subdirectories, call the iteratee function for each file, and add the returned string to the slice.

```go
package main

import (
	"fmt"
	"os"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  filenames, _ := filejez.FilenamesWalkBy("dir", func(path string, entry os.DirEntry) string {
    return "by-" + entry.Name()
  })
  fmt.Println(filenames)
  
  // Output:
  // [by-test-file.txt by-test-file1.txt]
}

```

### DeleteFiles
Delete the file.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  _ = filejez.DeleteFiles("path")
  fmt.Println(filejez.FileExists("path"))
  
  // Output:
  // false <nil>
}

```

### DeleteDirs
Delete the directory.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  _ = filejez.DeleteDirs("path")
  fmt.Println(filejez.DirExists("path"))
  
  // Output:
  // false <nil>
}

```

### DeleteEmptyDirWalk
Return to delete the empty directory, including subdirectories, for example: /a/b. When deleting b, if a is also an empty directory, it will also be deleted.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  _ = filejez.DeleteEmptyDirWalk("dir")
  fmt.Println(filejez.DirExists("dir"))
  
  // Output:
  // false <nil>
}

```

### DeleteWalkBy
Recursively delete files and subdirectories in the specified directory, itratee: used to process the logic of each file (excluding the directory), receive the file path and the os.DirEntry instance as parameters.

- Parameters:
    - DirPath: directory path
    - Iterate: A traversal function used to process the logic of each file (excluding directories), receive file paths, and os DirEntry instance as a parameter
    - With EmptyDir: optional parameter, specifying whether to delete empty directories, default to false
- Return value:
    - Bool: Has the current directory been deleted
    - Error: Error
- Notes:
    - If you want to delete an empty directory, please pay attention to the bool return value. Incorrect bool return values may cause unexpected file deletion or empty directory deletion failure.

```go
package main

import (
	"fmt"
	"os"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  isDelete, _ := filejez.DeleteWalkBy("dir", func(path string, entry os.DirEntry) (bool, error) {
    return true, nil
  }, true)
  fmt.Println(isDelete)
  fmt.Println(filejez.DirExists("dir"))
  
  // Output:
  // true <nil>
  // true <nil>
}

```

### Zip
Compress the directory or file into a zip file, and if zip already exists, it will be overwritten.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  _ = filejez.Zip("./test-dir-walk", "target")
  fmt.Println(filejez.FileExists("target"))
  
  // Output:
  // true <nil>
}

```

### ZipFilter
Call the iteratee function for each file or directory. If it returns true, it will be compressed into a zip file. If the zip file already exists, it will be overwritten.

```go
package main

import (
	"fmt"
	"os"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  _ = filejez.ZipFilter(".", "target", func(path string, entry os.DirEntry) bool {
    return path == "./test-dir-walk"
  })
  fmt.Println(filejez.FileExists("target"))
  
  // Output:
  // true <nil>
}

```

### Unzip
Unzip the zip file to the specified directory, and if the directory does not exist, it will be created.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  _ = filejez.Unzip("src", "dst")
  fmt.Println(filejez.DirExists("dst"))
  
  // Output:
  // true <nil>
}

```

### ReadAll
Read all the contents of the file as a string.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  data2, _ := filejez.ReadAll("path")
  fmt.Println(data2)
  
  // Output:
  // test
}

```

### ReadLines
Read the first n lines of the file, and if n < 0, read all lines.

```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/filejez"
)

func main(){
  data2, _ := filejez.ReadLines("path", 1)
  fmt.Println(data2)
  
  // Output:
  // [test]
}

```
