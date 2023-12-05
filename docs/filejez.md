# Filejez

------

提供了一些文件操作的工具函数，包括文件遍历、文件创建、文件删除等

------

## 用法

```go
import "github.com/dengrandpa/jez/filejez"
```

------

## 目录

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

------

### FilterMap
遍历当前目录，对每个文件调用 iteratee，如果返回 true，则将结果放入结果集中。

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
返回遍历所有目录、子目录，对每个文件调用 iteratee，如果返回 true，则将结果放入结果集中。

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
判断是否是目录。

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
判断目录是否为空。

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
判断文件是否存在。

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
判断目录是否存在。

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
等同于 os.Create,创建文件，如果文件已存在，则忽略，使用完毕后需要关闭。

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
创建文件，如果文件已存在，则忽略。

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
创建文件，如果文件已存在，则覆盖。

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
创建目录，包含子目录，如果目录已存在，则忽略。

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
创建文件，如果文件已存在，则忽略，同时创建目录，包含子目录。

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
创建文件，如果文件已存在，则覆盖，同时创建目录，包含子目录。

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
创建文件并写入字符串数据。

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
拷贝文件。

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
遍历目录、子目录，查找文件。

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
遍历目录、子目录，查找文件，对每个文件调用 iteratee 函数，如果返回 true，则表示找到了。

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
返回目录下的文件名切片

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
遍历目录下的文件，对每个文件调用 iteratee 函数，如果返回 true，则将文件名添加到切片中。

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
遍历目录下的文件，对每个文件调用 iteratee 函数，将返回的字符串添加到切片中。

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

### 。FilenamesWalk
返回目录下的文件名切片，包含子目录。

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
返回目录下的文件，包含子目录，对每个文件调用 iteratee 函数，如果返回 true，则将文件名添加到切片中。

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
返回目录下的文件，包含子目录，对每个文件调用 iteratee 函数，将返回的字符串添加到切片中。

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
删除文件。

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
删除目录。

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
返回删除空目录，包含子目录，例如：/a/b，当删除 b 时，如果 a 也是一个空目录，也会被删除。

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
递归删除指定目录下的文件和子目录

- 参数：
  - dirPath: 目录路径
  - iteratee: 遍历函数，用于处理每个文件（不包括目录）的逻辑，接收文件路径和 os.DirEntry 实例作为参数
  - withEmptyDir: 可选参数，指定是否删除空目录，默认为 false

- 返回值：
  - bool: 当前目录是否已被删除
  - error: 错误

- 注意事项：
  - 如果要删除空目录，请注意 bool 返回值，错误的 bool 返回值可能导致文件意外删除或空目录删除失败。
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
将目录或文件压缩为 zip 文件，如果zip已存在，则会被覆盖。

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
对每个文件或目录调用 iteratee 函数，如果返回 true，则将其压缩到 zip 文件中，如果zip文件已存在，则会被覆盖。

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
解压 zip 文件到指定目录，如果目录不存在，则会被创建。

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
将文件的所有内容读取为字符串。

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
读取文件的前 n 行，如果 n < 0，则读取所有行。

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
