# Validatorjez

------

提供了一些常用的 效验 操作函数。

------

## 用法

```go
import "github.com/dengrandpa/jez/validatorjez"
```

------

## 目录

-   [IsBase64](#isbase64)
-   [IsBase64URL](#isbase64url)
-   [IsChineseMainlandIDCard](#ischinesemainlandidcard)
-   [IsChineseMainlandPhoneNumber](#ischinesemainlandphonenumber)
-   [IsFloat](#isfloat)
-   [IsFloatType](#isfloattype)
-   [IsIP](#isip)
-   [IsIPv4](#isipv4)
-   [IsIPv6](#isipv6)
-   [IsIn](#isin)
-   [IsInt](#isint)
-   [IsIntType](#isinttype)
-   [IsJSON](#isjson)
-   [IsLongitude](#islongitude)
-   [IsLatitude](#islatitude)
-   [IsNum](#isnum)
-   [IsNumType](#isnumtype)
-   [IsPort](#isport)
-   [IsPrefixOrSuffix](#isprefixorsuffix)
-   [IsRange](#isrange)
-   [IsRegex](#isregex)
-   [IsRegexMatch](#isregexmatch)
-   [RuneLength](#runelength)
-   [StringLength](#stringlength)

------


### IsBase64
是否为base64
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {

	fmt.Println(validatorjez.IsBase64("SGVsbG8sIHdvcmxkIQ=="))
	fmt.Println(validatorjez.IsBase64(""))
	fmt.Println(validatorjez.IsBase64("123"))

	// Output:
	// true
	// false
	// false
}


```

### IsBase64URL
是否为base64url
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.IsBase64URL("SGVsbG8sIHdvcmxkIQ=="))
	fmt.Println(validatorjez.IsBase64URL(""))
	fmt.Println(validatorjez.IsBase64URL("123"))

	// Output:
	// true
	// false
	// false
}


```

### IsChineseMainlandIDCard
是否为中国大陆身份证号码
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.IsChineseMainlandIDCard("210911192105130714"))
	fmt.Println(validatorjez.IsChineseMainlandIDCard("11010519491231002X"))
	fmt.Println(validatorjez.IsChineseMainlandIDCard("11010519491231002x"))
	fmt.Println(validatorjez.IsChineseMainlandIDCard("123456"))
	fmt.Println(validatorjez.IsChineseMainlandIDCard("990911192105130714"))
	fmt.Println(validatorjez.IsChineseMainlandIDCard("210911189905130714"))
	fmt.Println(validatorjez.IsChineseMainlandIDCard("210911222205130714"))

	// Output:
	// true
	// true
	// true
	// false
	// false
	// false
	// false
}


```

### IsChineseMainlandPhoneNumber
是否为中国大陆手机号码, withCode为是否可包含国家代码 86 / +86
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.IsChineseMainlandPhoneNumber("13666666666"))
	fmt.Println(validatorjez.IsChineseMainlandPhoneNumber("+8613666666666", true))
	fmt.Println(validatorjez.IsChineseMainlandPhoneNumber("+8613666666666"))
	fmt.Println(validatorjez.IsChineseMainlandPhoneNumber("13666"))
	fmt.Println(validatorjez.IsChineseMainlandPhoneNumber("12666666666"))

	// Output:
	// true
	// true
	// false
	// false
	// false
}


```

### IsFloat
是否为浮点数
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.IsFloat("3.14"))
	fmt.Println(validatorjez.IsFloat("3"))
	fmt.Println(validatorjez.IsFloat(".3"))
	fmt.Println(validatorjez.IsFloat("3."))
	fmt.Println(validatorjez.IsFloat("3.1.1"))

	// Output:
	// true
	// true
	// true
	// true
	// false
}


```

### IsFloatType
是否为浮点数类型
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.IsFloatType(3.14))
	fmt.Println(validatorjez.IsFloatType(float64(3)))
	fmt.Println(validatorjez.IsFloatType(3.))
	fmt.Println(validatorjez.IsFloatType(.3))
	fmt.Println(validatorjez.IsFloatType("3.14"))

	// Output:
	// true
	// true
	// true
	// true
	// false
}


```

### IsIP
是否为IP地址
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.IsIP("1.1.1.1"))
	fmt.Println(validatorjez.IsIP("127.0.0.1"))
	fmt.Println(validatorjez.IsIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334"))
	fmt.Println(validatorjez.IsIP("127.0.0.1:8080"))
	fmt.Println(validatorjez.IsIP("localhost"))

	// Output:
	// true
	// true
	// true
	// false
	// false
}


```

### IsIPv4
是否为IPv4地址
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.IsIPv4("1.1.1.1"))
	fmt.Println(validatorjez.IsIPv4("127.0.0.1"))
	fmt.Println(validatorjez.IsIPv4("2001:0db8:85a3:0000:0000:8a2e:0370:7334"))
	fmt.Println(validatorjez.IsIPv4("127.0.0.1:8080"))
	fmt.Println(validatorjez.IsIPv4("localhost"))

	// Output:
	// true
	// true
	// false
	// false
	// false
}


```

### IsIPv6
是否为IPv6地址
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.IsIPv6("127.0.0.1"))
	fmt.Println(validatorjez.IsIPv6("2001:0db8:85a3:0000:0000:8a2e:0370:7334"))
	fmt.Println(validatorjez.IsIPv6("127.0.0.1:8080"))
	fmt.Println(validatorjez.IsIPv6("localhost"))

	// Output:
	// false
	// true
	// false
	// false
}


```

### IsIn
是否在指定列表中
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.IsIn(1, 1, 2, 3))
	fmt.Println(validatorjez.IsIn(4, 1, 2, 3))

	// Output:
	// true
	// false
}


```

### IsInt
是否为整数
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.IsInt("1"))
	fmt.Println(validatorjez.IsInt("01"))
	fmt.Println(validatorjez.IsInt("1.1"))

	// Output:
	// true
	// true
	// false
}


```

### IsIntType
是否为整数类型
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.IsIntType(1))
	fmt.Println(validatorjez.IsIntType(1.1))
	fmt.Println(validatorjez.IsIntType("1"))

	// Output:
	// true
	// false
	// false
}


```

### IsJSON
是否为json
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.IsJSON(`{"name":"tom"}`))
	fmt.Println(validatorjez.IsJSON(`abc`))

	// Output:
	// true
	// false
}


```

### IsLongitude
是否为经度
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.IsLatitude("-74.0060"))
	fmt.Println(validatorjez.IsLatitude("123.123"))

	// Output:
	// true
	// false
}


```

### IsLatitude
是否为纬度
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.IsLongitude("-45.12345"))
	fmt.Println(validatorjez.IsLongitude("123.123.123"))

	// Output:
	// true
	// false
}


```

### IsNum
是否为数字(包含浮点数)
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.IsNum("123"))
	fmt.Println(validatorjez.IsNum("123.123"))
	fmt.Println(validatorjez.IsNum("123.123.123"))

	// Output:
	// true
	// true
	// false
}


```

### IsNumType
是否为数字类型(包含浮点数)
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.IsNumType(1))
	fmt.Println(validatorjez.IsNumType(1.1))
	fmt.Println(validatorjez.IsNumType("1"))

	// Output:
	// true
	// true
	// false
}


```

### IsPort
是否为端口号
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.IsPort("65535"))
	fmt.Println(validatorjez.IsPort("65536"))
	fmt.Println(validatorjez.IsPort("0"))

	// Output:
	// true
	// false
	// false
}


```

### IsPrefixOrSuffix
是否以指定字符串开头或结尾
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.IsPrefixOrSuffix("123", "12"))
	fmt.Println(validatorjez.IsPrefixOrSuffix("123", "23"))
	fmt.Println(validatorjez.IsPrefixOrSuffix("123", "2"))

	// Output:
	// true
	// true
	// false
}


```

### IsRange
数值范围
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.IsRange(1, 1, 2))
	fmt.Println(validatorjez.IsRange(2, 1, 2))
	fmt.Println(validatorjez.IsRange(3, 1, 2))

	// Output:
	// true
	// true
	// false
}


```

### IsRegex
是否为正则表达式
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.IsRegex(`^[0-9]+$`))
	fmt.Println(validatorjez.IsRegex(`abc(`))

	// Output:
	// true
	// false
}


```

### IsRegexMatch
是否匹配正则表达式
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.IsRegexMatch("1", `^[0-9]+$`))
	fmt.Println(validatorjez.IsRegexMatch("2", `abc(`))
	fmt.Println(validatorjez.IsRegexMatch("a", `^[0-9]+$`))

	// Output:
	// true
	// false
	// false
}


```

### RuneLength
字符长度
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.RuneLength("1", 1, 2))
	fmt.Println(validatorjez.RuneLength("123", 1, 2))

	// Output:
	// true
	// false
}


```

### StringLength
字符串长度
```go
package main

import (
	"fmt"
	
	"github.com/dengrandpa/jez/validatorjez"
)

func main() {
	fmt.Println(validatorjez.StringLength("1", 1, 2))
	fmt.Println(validatorjez.StringLength("2", 1, 2))
	fmt.Println(validatorjez.StringLength("3", 3, 4))

	// Output:
	// true
	// true
	// false
}


```


