# jez - 一些简单且常用的包...

![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/dengrandpa/jez)
[![codecov](https://codecov.io/gh/dengrandpa/jez/branch/main/graph/badge.svg?token=E3LRLIYTK2)](https://codecov.io/gh/dengrandpa/jez)
[![Go Report Card](https://goreportcard.com/badge/github.com/dengrandpa/jez)](https://goreportcard.com/report/github.com/dengrandpa/jez)
[![test](https://github.com/dengrandpa/jez/actions/workflows/workflows.yml/badge.svg)](https://github.com/dengrandpa/jez/actions/workflows/workflows.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/dengrandpa/jez.svg)](https://pkg.go.dev/github.com/dengrandpa/jez)
![GitHub](https://img.shields.io/github/license/dengrandpa/jez)

------


### 简体中文 | [English](./README_en.md) 

------

### 特别感谢ChatGPT提供技术支持！！

------

## 安装

```sh
go get github.com/dengrandpa/jez
```

## 用法

使用时导入相对应的包名。例如：

### slice工具包
```go
import "github.com/dengrandpa/jez/slicejez"
```

### map工具包
```go
import "github.com/dengrandpa/jez/mapjez"
```

…………

------

## 目录

-   [Filejez](#filejez)：提供了一些文件操作的工具函数，包括文件遍历、文件创建、文件删除等
-   [Mapjez](#mapjez)：提供了一些常用的 map 操作函数。
-   [Randomjez](#randomjez)：提供了一些随机数生成的函数。
-   [Slicejez](#slicejez)：提供了一些对切片的操作，包括遍历、映射、过滤、去重、求差集、求交集等。
-   [Timejez](#timejez)：提供了时间相关的一些操作，包括时间格式化、时间计算、时间区间计算、时间转换等
-   [Validatorjez](#validatorjez)：提供了一些常用的效验操作函数。

------

## [Filejez](./docs/filejez.md#Filejez)
提供了一些文件操作的工具函数，包括文件遍历、文件创建、文件删除等

```go
import "github.com/dengrandpa/jez/filejez"
```
### 函数

-   [FilterMap](./docs/filejez.md#filterMap)：遍历当前目录，对每个文件调用 iteratee，如果返回 true，则将结果放入结果集中
-   [FilterMapWalk](./docs/filejez.md#filterMapWalk)：返回遍历所有目录、子目录，对每个文件调用 iteratee，如果返回 true，则将结果放入结果集中
-   [IsDir](./docs/filejez.md#isDir)：判断是否是目录
-   [IsEmptyDir](./docs/filejez.md#isEmptyDir)：判断目录是否为空
-   [FileExists](./docs/filejez.md#fileExists)：判断文件是否存在
-   [DirExists](./docs/filejez.md#dirExists)：判断目录是否存在
-   [OsCreate](./docs/filejez.md#osCreate)：等同于 os.Create,创建文件，如果文件已存在，则忽略，使用完毕后需要关闭
-   [CreateFiles](./docs/filejez.md#createFiles)：创建文件，如果文件已存在，则忽略
-   [OverwriteFiles](./docs/filejez.md#overwriteFiles)：创建文件，如果文件已存在，则覆盖
-   [CreateDirs](./docs/filejez.md#createDirs)：创建目录，包含子目录，如果目录已存在，则忽略
-   [CreateFilesWithDirs](./docs/filejez.md#createFilesWithDirs)：创建文件，如果文件已存在，则忽略，同时创建目录，包含子目录
-   [OverwriteFilesWithDirs](./docs/filejez.md#overwriteFilesWithDirs)：创建文件，如果文件已存在，则覆盖，同时创建目录，包含子目录
-   [CreateFileWithData](./docs/filejez.md#createFileWithData)：创建文件并写入字符串数据
-   [CopyFile](./docs/filejez.md#copyFile)：拷贝文件
-   [FindFileWalk](./docs/filejez.md#findFileWalk)：遍历目录、子目录，查找文件
-   [FindFileWalkFilter](./docs/filejez.md#findFileWalkFilter)：遍历目录、子目录，查找文件，对每个文件调用 iteratee 函数，如果返回 true，则表示找到了
-   [Filenames](./docs/filejez.md#filenames)：返回目录下的文件名切片
-   [FilenamesFilter](./docs/filejez.md#filenamesFilter)：遍历目录下的文件，对每个文件调用 iteratee 函数，如果返回 true，则将文件名添加到切片中
-   [FilenamesBy](./docs/filejez.md#filenamesBy)：遍历目录下的文件，对每个文件调用 iteratee 函数，将返回的字符串添加到切片中
-   [FilenamesWalk](./docs/filejez.md#filenamesWalk)：返回目录下的文件名切片，包含子目录
-   [FilenamesWalkFilter](./docs/filejez.md#filenamesWalkFilter)：返回目录下的文件，包含子目录，对每个文件调用 iteratee 函数，如果返回 true，则将文件名添加到切片中
-   [FilenamesWalkBy](./docs/filejez.md#filenamesWalkBy)：返回目录下的文件，包含子目录，对每个文件调用 iteratee 函数，将返回的字符串添加到切片中
-   [DeleteFiles](./docs/filejez.md#deleteFiles)：删除文件
-   [DeleteDirs](./docs/filejez.md#deleteDirs)：删除目录
-   [DeleteEmptyDirWalk](./docs/filejez.md#deleteEmptyDirWalk)：返回删除空目录，包含子目录
-   [DeleteWalkBy](./docs/filejez.md#deleteWalkBy)：递归删除指定目录下的文件和子目录
-   [Zip](./docs/filejez.md#zip)：将目录或文件压缩为 zip 文件，如果zip已存在，则会被覆盖。
-   [ZipFilter](./docs/filejez.md#zipFilter)：对每个文件或目录调用 iteratee 函数，如果返回 true，则将其压缩到 zip 文件中，如果zip文件已存在，则会被覆盖。
-   [Unzip](./docs/filejez.md#unzip)：解压 zip 文件到指定目录，如果目录不存在，则会被创建。
-   [ReadAll](./docs/filejez.md#readAll)：将文件的所有内容读取为字符串。
-   [ReadLines](./docs/filejez.md#readLines)：读取文件的前 n 行，如果 n < 0，则读取所有行。

------

## [Mapjez](./docs/mapjez.md#Mapjez)
提供了一些常用的 map 操作函数。

```go
import "github.com/dengrandpa/jez/mapjez"
```
### 函数

-   [ForEach](./docs/mapjez.md#forEach)：遍历map，对每个元素调用 iteratee 函数。
-   [Filter](./docs/mapjez.md#filter)：遍历map，对每个元素调用 iteratee 函数，如果 iteratee 返回 true，则将该元素添加到结果map中。
-   [Keys](./docs/mapjez.md#keys)：遍历map，将每个key添加到结果slice中。
-   [KeysBy](./docs/mapjez.md#keysBy)：遍历map，对每个元素调用 iteratee 函数，并返回调用后结果。
-   [Values](./docs/mapjez.md#values)：返回map中所有的value。
-   [ValuesBy](./docs/mapjez.md#valuesBy)：遍历map，对每个元素调用 iteratee 函数，并返回调用后结果。
-   [ValuesUnique](./docs/mapjez.md#valuesUnique)：返回map中所有的value，结果去重。
-   [KeysAndValues](./docs/mapjez.md#keysAndValues)：返回map中所有的key和value。
-   [KeysAndValuesFilter](./docs/mapjez.md#keysAndValuesFilter)：遍历map，对每个元素调用 iteratee 函数，如果 iteratee 返回true，则将该元素添加到结果slice中。
-   [Deletes](./docs/mapjez.md#deletes)：通过key删除多个元素。
-   [DeleteByValues](./docs/mapjez.md#deleteByValues)：通过value删除多个元素。
-   [DeleteFilter](./docs/mapjez.md#deleteFilter)：遍历map，对每个元素调用 iteratee 函数，如果 iteratee 返回true，则删除该元素。
-   [ReplaceValue](./docs/mapjez.md#replaceValue)：替换所有value等于 old 的元素。
-   [MapToSliceBy](./docs/mapjez.md#mapToSliceBy)：map转切片，遍历map，对每个元素调用 iteratee 函数，并返回调用后结果切片。
-   [MapToSliceFilter](./docs/mapjez.md#mapToSliceFilter)：map转切片，遍历map，对每个元素调用 iteratee 函数，如果 iteratee 返回true，则将该元素添加到结果切片中。

------

## [Randomjez](./docs/randomjez.md#Randomjez)
提供了一些随机数生成的函数。

```go
import "github.com/dengrandpa/jez/randomjez"
```
### 函数

-   [Random](./docs/randomjez.md#random)：随机生成字符串
-   [RandomLower](./docs/randomjez.md#randomLower)：随机生成小写字母字符串
-   [RandomUpper](./docs/randomjez.md#randomUpper)：随机生成大写字母字符串
-   [RandomNumeral](./docs/randomjez.md#randomNumeral)：随机生成数字字符串
-   [RandomCaseLetters](./docs/randomjez.md#randomCaseLetters)：随机生成大小写字母字符串。
-   [RandomLowerNumeral](./docs/randomjez.md#randomLowerNumeral)：随机生成小写字母和数字字符串
-   [RandomUpperNumeral](./docs/randomjez.md#randomUpperNumeral)：随机生成大写字母和数字字符串
-   [RandomCharset](./docs/randomjez.md#randomCharset)：随机生成字符串，包含数字、大小写字母
-   [RandomInt](./docs/randomjez.md#randomInt)：随机生成整数，包含 min，不包含 max，即 [min,max)。
-   [RandomIntSlice](./docs/randomjez.md#randomIntSlice)：随机生成整数切片
-   [RandomIntSliceUnique](./docs/randomjez.md#randomIntSliceUnique)：随机生成不重复的整数切片
-   [RandomBytes](./docs/randomjez.md#randomBytes)：随机生成字节切片
-   [UUIDv4](./docs/randomjez.md#uUIDv4)：根据 RFC4122 生成 UUID v4版本
-   [Shuffle](./docs/randomjez.md#shuffle)：打乱切片中元素的顺序。
-   [Sample](./docs/randomjez.md#sample)：从切片中随机返回一个元素。
-   [Samples](./docs/randomjez.md#samples)：从切片中随机返回n个元素，结果不去重。

------

## [Slicejez](./docs/slicejez.md#Slicejez)
提供了一些对切片的操作，包括遍历、映射、过滤、去重、求差集、求交集等。

```go
import "github.com/dengrandpa/jez/slicejez"
```
### 函数

-   [ForEach](./docs/slicejez.md#forEach)：遍历切片并为每个元素调用 iteratee 函数。
-   [ForEachWithBreak](./docs/slicejez.md#forEachWithBreak)：遍历切片并为每个元素调用 iteratee 函数，如果返回 false，则停止遍历。
-   [Filter](./docs/slicejez.md#filter)：遍历切片并为每个元素调用 iteratee 函数，只返回调用结果为true的元素。
-   [Map](./docs/slicejez.md#map)：遍历切片并为每个元素调用 iteratee 函数，并返回调用后结果。
-   [Contain](./docs/slicejez.md#contain)：效验切片是否包含目标元素。
-   [ContainAll](./docs/slicejez.md#containAll)：效验切片是否包含所有的目标元素。
-   [FilterMap](./docs/slicejez.md#filterMap)：遍历切片并为每个元素调用 iteratee 函数，如果调用结果为true，则返回调用后元素。
-   [AppendIfNotDuplicate](./docs/slicejez.md#appendIfNotDuplicate)：添加元素到切片，如果元素已经存在，则不添加。
-   [AppendMultipleIfNotDuplicate](./docs/slicejez.md#appendMultipleIfNotDuplicate)：添加多个元素到切片，如果元素已经存在，则不添加。
-   [Remove](./docs/slicejez.md#remove)：从切片中删除元素。
-   [RemoveFilter](./docs/slicejez.md#removeFilter)：遍历切片并为每个元素调用 iteratee 函数，如果调用结果为true，则删除该元素。
-   [Unique](./docs/slicejez.md#unique)：去重。
-   [UniqueBy](./docs/slicejez.md#uniqueBy)：遍历切片并为每个元素调用 iteratee 函数，返回唯一的元素。
-   [UniqueNonzero](./docs/slicejez.md#uniqueNonzero)：删除重复元素及零值。
-   [UniqueNonzeroBy](./docs/slicejez.md#uniqueNonzeroBy)：遍历切片并为每个元素调用 iteratee 函数，返回唯一的、非零值的元素。
-   [Nonzero](./docs/slicejez.md#nonzero)：删除零值。
-   [Replace](./docs/slicejez.md#replace)：将切片中的元素 old 替换为 new ，最多替换 n 次，如果 n 为-1，则替换所有的 old 元素。
-   [ReplaceAll](./docs/slicejez.md#replaceAll)：将切片中的元素 old 替换为 new ，替换所有的 old 元素。
-   [Difference](./docs/slicejez.md#difference)：差集，结果不去重。
-   [DifferenceUnique](./docs/slicejez.md#differenceUnique)：差集，结果去重。
-   [Intersection](./docs/slicejez.md#intersection)：交集，结果元素唯一。
-   [MutualDifference](./docs/slicejez.md#mutualDifference)：差异，结果不去重。
-   [ToMapBy](./docs/slicejez.md#toMapBy)：遍历切片，将切片中的元素转换为map的key和value。
-   [Repeat](./docs/slicejez.md#repeat)：返回包含 n 个 item 的切片。
-   [Equal](./docs/slicejez.md#equal)：长度、顺序、值都相等时返回 true 。
-   [EqualElement](./docs/slicejez.md#equalElement)：长度、值相等时返回 true ，不考虑顺序。
-   [FindIndex](./docs/slicejez.md#findIndex)：返回第一个匹配的元素的索引，不存在则返回 -1 。
-   [FindIndexFilter](./docs/slicejez.md#findIndexFilter)：返回调用 iteratee 函数返回 true 的第一个元素的索引，不存在则返回 -1 。
-   [FindDuplicates](./docs/slicejez.md#findDuplicates)：返回切片中所有重复的元素，结果不去重。
-   [FindUniqueDuplicates](./docs/slicejez.md#findUniqueDuplicates)：返回切片中所有重复的元素，结果去重。
-   [Min](./docs/slicejez.md#min)：返回最小值
-   [Max](./docs/slicejez.md#max)：返回最大值
-   [Drop](./docs/slicejez.md#drop)：返回从开头删除n个元素的切片，如果 n 大于切片的长度，则返回空切片。
-   [DropLast](./docs/slicejez.md#dropLast)：返回从末尾删除n个元素的切片，如果 n 大于切片的长度，则返回空切片。
-   [Slice](./docs/slicejez.md#slice)：返回索引从 n 到 m 的切片，但不包括 m，等同于 slice[n:m]，即[min,max)，但不会在溢出时panic。
-   [IsSorted](./docs/slicejez.md#isSorted)：判断切片是否已排序。
-   [IsSortedBy](./docs/slicejez.md#isSortedBy)：遍历切片并为每个元素调用 iteratee 函数，以确定它是否已排序。
-   [Reverse](./docs/slicejez.md#reverse)：将切片中的元素顺序反转。
-   [Flatten](./docs/slicejez.md#flatten)：将二维切片转换为一维切片。
-   [InsertAt](./docs/slicejez.md#insertAt)：在切片的指定索引处插入值，如果索引大于切片的长度或小于 0，则将值附加到切片的末尾。
-   [NewSafeSlice](./docs/slicejez.md#newsafeslice)：创建一个并发安全的切片。
-   [SafeSlice_ForEach](./docs/slicejez.md#safeSliceforeach)：遍历切片并为每个元素调用 iteratee 函数。
-   [SafeSlice_ForEachWithBreak](./docs/slicejez.md#safesliceforeachwithbreak)：遍历切片并为每个元素调用 iteratee 函数，如果返回 false，则停止遍历。
-   [SafeSlice_Filter](./docs/slicejez.md#safeslicefilter)：遍历切片并为每个元素调用 iteratee 函数，只返回调用结果为true的元素。
-   [SafeSlice_Append](./docs/slicejez.md#safesliceappend)：添加元素到切片。
-   [SafeSlice_AppendIfNotDuplicate](./docs/slicejez.md#safesliceappendifnotduplicate)：添加元素到切片，如果元素已经存在，则不添加。
-   [SafeSlice_AppendMultipleIfNotDuplicate](./docs/slicejez.md#safesliceappendmultipleifnotduplicate)：添加多个元素到切片，如果元素已经存在，则不添加。
-   [SafeSlice_Load](./docs/slicejez.md#safesliceload)：返回切片的副本。
-   [SafeSlice_LoadByIndex](./docs/slicejez.md#safesliceloadbyindex)：返回指定索引位置的元素，-1 则返回最后一个元素，如果索引超出范围，panic。
-   [SafeSlice_Index](./docs/slicejez.md#safesliceindex)：返回指定元素在切片中的索引位置。
-   [SafeSlice_Insert](./docs/slicejez.md#safesliceinsert)：在指定索引位置插入元素。
-   [SafeSlice_Len](./docs/slicejez.md#safeslicelen)：返回切片的长度。
-   [SafeSlice_Remove](./docs/slicejez.md#safesliceremove)：从切片中移除元素。
-   [SafeSlice_RemoveByIndex](./docs/slicejez.md#safesliceremovebyindex)：从切片中移除指定索引位置的元素。
-   [SafeSlice_Replace](./docs/slicejez.md#safeslicereplace)：将切片中的元素 old 替换为 new ，最多替换 n 次，如果 n 为-1，则替换所有的 old 元素。
-   [SafeSlice_ReplaceByIndex](./docs/slicejez.md#safeslicereplacebyindex)：将指定索引位置的元素替换为 new 。
-   [SafeSlice_Slice](./docs/slicejez.md#safesliceslice)：返回索引从 n 到 m 的切片，但不包括 m，等同于 slice[n:m]，即[min,max)，但不会在溢出时panic。

------

## [Timejez](./docs/timejez.md#Timejez)
提供了时间相关的一些操作，包括时间格式化、时间计算、时间区间计算、时间转换等

```go
import "github.com/dengrandpa/jez/timejez"
```
### 函数

-   [ParseTime](./docs/timejez.md#parseTime)：将字符串转换为时间，默认格式为 YYYYMMDDHHMMSS
-   [ParseTimestamp](./docs/timejez.md#parseTimestamp)：将时间戳转换为时间
-   [StartOfMinute](./docs/timejez.md#startOfMinute)：返回时间 t 所在分钟的开始时间 yyyy-mm-dd hh:mm:00
-   [EndOfMinute](./docs/timejez.md#endOfMinute)：返回时间 t 所在分钟的结束时间 yyyy-mm-dd hh:mm:59
-   [StartOfHour](./docs/timejez.md#startOfHour)：返回时间 t 所在小时的开始时间 yyyy-mm-dd hh:00:00
-   [EndOfHour](./docs/timejez.md#endOfHour)：返回时间 t 所在小时的结束时间 yyyy-mm-dd hh:59:59
-   [StartOfDay](./docs/timejez.md#startOfDay)：返回时间 t 所在天的开始时间 yyyy-mm-dd 00:00:00
-   [EndOfDay](./docs/timejez.md#endOfDay)：返回时间 t 所在天的结束时间 yyyy-mm-dd 23:59:59
-   [StartOfWeekMonday](./docs/timejez.md#startOfWeekMonday)：返回时间 t 所在周的开始时间，周一为第一天 yyyy-mm-dd 00:00:00
-   [EndOfWeekSunday](./docs/timejez.md#endOfWeekSunday)：返回时间 t 所在周的结束时间，周日为最后一天 yyyy-mm-dd 23:59:59
-   [StartOfWeekSunday](./docs/timejez.md#startOfWeekSunday)：返回时间 t 所在周的开始时间，周日为第一天 yyyy-mm-dd 00:00:00
-   [EndOfWeekMonday](./docs/timejez.md#endOfWeekMonday)：返回时间 t 所在周的结束时间，周一为最后一天 yyyy-mm-dd 23:59:59
-   [StartOfMonth](./docs/timejez.md#startOfMonth)：返回时间 t 所在月的开始时间 yyyy-mm-01 00:00:00
-   [EndOfMonth](./docs/timejez.md#endOfMonth)：返回时间 t 所在月的结束时间 yyyy-mm-dd 23:59:59
-   [StartOfYear](./docs/timejez.md#startOfYear)：返回时间 t 所在年的开始时间 yyyy-01-01 00:00:00
-   [EndOfYear](./docs/timejez.md#endOfYear)：返回时间 t 所在年的结束时间 yyyy-12-31 23:59:59
-   [AddSecond](./docs/timejez.md#addSecond)：添加或删除秒数
-   [AddMinute](./docs/timejez.md#addMinute)：添加或删除分钟数
-   [AddHour](./docs/timejez.md#addHour)：添加或删除小时数
-   [AddDay](./docs/timejez.md#addDay)：添加或删除天数
-   [AddWeek](./docs/timejez.md#addWeek)：添加或删除周数
-   [AddMonth](./docs/timejez.md#addMonth)：添加或删除月数
-   [AddYear](./docs/timejez.md#addYear)：添加或删除年数
-   [FormatTime](./docs/timejez.md#formatTime)：将时间格式化为字符串，默认格式为 yyyy-mm-dd hh:mm:ss
-   [FormatTimestamp](./docs/timejez.md#formatTimestamp)：将时间戳格式化为字符串，默认格式为 yyyy-mm-dd hh:mm:ss
-   [FormatNow](./docs/timejez.md#formatNow)：返回当前时间的字符串格式，默认格式为 yyyy-mm-dd hh:mm:ss
-   [IsLeapYear](./docs/timejez.md#isLeapYear)：判断年份 year 是否为闰年
-   [RangeHours](./docs/timejez.md#rangeHours)：返回两个时间之间的所有小时的切片，包含 start 和 end，即[start,end]，如果start和end结果一样，则只返回1个
-   [RangeDays](./docs/timejez.md#rangeDays)：返回两个时间之间的所有天的切片，包含 start 和 end，即[start,end]，如果start和end结果一样，则只返回1个
-   [RangeMonths](./docs/timejez.md#rangeMonths)：返回两个时间之间的所有月的切片，包含 start 和 end，即[start,end]，如果start和end结果一样，则只返回1个
-   [RangeYears](./docs/timejez.md#rangeYears)：返回两个时间之间的所有年的切片，包含 start 和 end，即[start,end]，如果start和end结果一样，则只返回1个
-   [NewTime](./docs/timejez.md#newTime)：相当于 time.Date，如果不传参数则相当于 time.Now。
-   [SubTime](./docs/timejez.md#subTime)：返回 t1 和 t2 之间的差值的日、小时、分钟和秒
-   [SubTimestamp](./docs/timejez.md#subTimestamp)：返回 t1 和 t2 之间的差值的日、小时、分钟和秒
-   [CST](./docs/timejez.md#cST)：返回中国时区
-   [ToCST](./docs/timejez.md#toCST)：将时间转换为中国时区

------

## [Validatorjez](./docs/validatorjez.md#validatorjez)
提供了一些常用的效验操作函数

```go
import "github.com/dengrandpa/jez/validatorjez"
```
### 函数

-   [IsBase64](./docs/validatorjez.md#isBase64)：是否为base64
-   [IsBase64URL](./docs/validatorjez.md#isBase64URL)：是否为base64url
-   [IsChineseMainlandIDCard](./docs/validatorjez.md#isChineseMainlandIDCard)：是否为中国大陆身份证号码
-   [IsChineseMainlandPhoneNumber](./docs/validatorjez.md#isChineseMainlandPhoneNumber)：是否为中国大陆手机号码, withCode为是否可包含国家代码 86 / +86
-   [IsFloat](./docs/validatorjez.md#isFloat)：是否为浮点数
-   [IsFloatType](./docs/validatorjez.md#isFloatType)：是否为浮点数类型
-   [IsIP](./docs/validatorjez.md#isIP)：是否为IP地址
-   [IsIPv4](./docs/validatorjez.md#isIPv4)：是否为IPv4地址
-   [IsIPv6](./docs/validatorjez.md#isIPv6)：是否为IPv6地址
-   [IsIn](./docs/validatorjez.md#isIn)：是否在指定列表中
-   [IsInt](./docs/validatorjez.md#isInt)：是否为整数
-   [IsIntType](./docs/validatorjez.md#isIntType)：是否为整数类型
-   [IsJSON](./docs/validatorjez.md#isJSON)：是否为json
-   [IsLongitude](./docs/validatorjez.md#isLongitude)：是否为经度
-   [IsLatitude](./docs/validatorjez.md#isLatitude)：是否为纬度
-   [IsNum](./docs/validatorjez.md#isNum)：是否为数字(包含浮点数)
-   [IsNumType](./docs/validatorjez.md#isNumType)：是否为数字类型(包含浮点数)
-   [IsPort](./docs/validatorjez.md#isPort)：是否为端口号
-   [IsPrefixOrSuffix](./docs/validatorjez.md#isPrefixOrSuffix)：是否以指定字符串开头或结尾
-   [IsRange](./docs/validatorjez.md#isRange)：数值范围
-   [IsRegex](./docs/validatorjez.md#isRegex)：是否为正则表达式
-   [IsRegexMatch](./docs/validatorjez.md#isRegexMatch)：是否匹配正则表达式
-   [RuneLength](./docs/validatorjez.md#runeLength)：字符长度
-   [StringLength](./docs/validatorjez.md#stringLength)：字符串长度

