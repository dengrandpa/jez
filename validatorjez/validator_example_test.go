package validatorjez

import (
	"fmt"
)

func ExampleIsBase64() {

	fmt.Println(IsBase64("SGVsbG8sIHdvcmxkIQ=="))
	fmt.Println(IsBase64(""))
	fmt.Println(IsBase64("123"))

	// Output:
	// true
	// false
	// false
}

func ExampleIsBase64URL() {
	fmt.Println(IsBase64URL("SGVsbG8sIHdvcmxkIQ=="))
	fmt.Println(IsBase64URL(""))
	fmt.Println(IsBase64URL("123"))

	// Output:
	// true
	// false
	// false
}

func ExampleIsChineseMainlandIDCard() {
	fmt.Println(IsChineseMainlandIDCard("210911192105130714"))
	fmt.Println(IsChineseMainlandIDCard("11010519491231002X"))
	fmt.Println(IsChineseMainlandIDCard("11010519491231002x"))
	fmt.Println(IsChineseMainlandIDCard("123456"))
	fmt.Println(IsChineseMainlandIDCard("990911192105130714"))
	fmt.Println(IsChineseMainlandIDCard("210911189905130714"))
	fmt.Println(IsChineseMainlandIDCard("210911222205130714"))

	// Output:
	// true
	// true
	// true
	// false
	// false
	// false
	// false
}

func ExampleIsChineseMainlandPhoneNumber() {
	fmt.Println(IsChineseMainlandPhoneNumber("13666666666"))
	fmt.Println(IsChineseMainlandPhoneNumber("+8613666666666", true))
	fmt.Println(IsChineseMainlandPhoneNumber("+8613666666666"))
	fmt.Println(IsChineseMainlandPhoneNumber("13666"))
	fmt.Println(IsChineseMainlandPhoneNumber("12666666666"))

	// Output:
	// true
	// true
	// false
	// false
	// false
}

func ExampleIsFloat() {
	fmt.Println(IsFloat("3.14"))
	fmt.Println(IsFloat("3"))
	fmt.Println(IsFloat(".3"))
	fmt.Println(IsFloat("3."))
	fmt.Println(IsFloat("3.1.1"))

	// Output:
	// true
	// true
	// true
	// true
	// false
}

func ExampleIsFloatType() {
	fmt.Println(IsFloatType(3.14))
	fmt.Println(IsFloatType(float64(3)))
	fmt.Println(IsFloatType(3.))
	fmt.Println(IsFloatType(.3))
	fmt.Println(IsFloatType("3.14"))

	// Output:
	// true
	// true
	// true
	// true
	// false
}

func ExampleIsIP() {
	fmt.Println(IsIP("1.1.1.1"))
	fmt.Println(IsIP("127.0.0.1"))
	fmt.Println(IsIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334"))
	fmt.Println(IsIP("127.0.0.1:8080"))
	fmt.Println(IsIP("localhost"))

	// Output:
	// true
	// true
	// true
	// false
	// false
}

func ExampleIsIPv4() {
	fmt.Println(IsIPv4("1.1.1.1"))
	fmt.Println(IsIPv4("127.0.0.1"))
	fmt.Println(IsIPv4("2001:0db8:85a3:0000:0000:8a2e:0370:7334"))
	fmt.Println(IsIPv4("127.0.0.1:8080"))
	fmt.Println(IsIPv4("localhost"))

	// Output:
	// true
	// true
	// false
	// false
	// false
}

func ExampleIsIPv6() {
	fmt.Println(IsIPv6("127.0.0.1"))
	fmt.Println(IsIPv6("2001:0db8:85a3:0000:0000:8a2e:0370:7334"))
	fmt.Println(IsIPv6("127.0.0.1:8080"))
	fmt.Println(IsIPv6("localhost"))

	// Output:
	// false
	// true
	// false
	// false
}

func ExampleIsIn() {
	fmt.Println(IsIn(1, 1, 2, 3))
	fmt.Println(IsIn(4, 1, 2, 3))

	// Output:
	// true
	// false
}

func ExampleIsInt() {
	fmt.Println(IsInt("1"))
	fmt.Println(IsInt("01"))
	fmt.Println(IsInt("1.1"))

	// Output:
	// true
	// true
	// false
}

func ExampleIsIntType() {
	fmt.Println(IsIntType(1))
	fmt.Println(IsIntType(1.1))
	fmt.Println(IsIntType("1"))

	// Output:
	// true
	// false
	// false
}

func ExampleIsJSON() {
	fmt.Println(IsJSON(`{"name":"tom"}`))
	fmt.Println(IsJSON(`abc`))

	// Output:
	// true
	// false
}

func ExampleIsLatitude() {
	fmt.Println(IsLatitude("-74.0060"))
	fmt.Println(IsLatitude("123.123"))

	// Output:
	// true
	// false
}

func ExampleIsLongitude() {
	fmt.Println(IsLongitude("-45.12345"))
	fmt.Println(IsLongitude("123.123.123"))

	// Output:
	// true
	// false
}

func ExampleIsNum() {
	fmt.Println(IsNum("123"))
	fmt.Println(IsNum("123.123"))
	fmt.Println(IsNum("123.123.123"))

	// Output:
	// true
	// true
	// false
}

func ExampleIsNumType() {
	fmt.Println(IsNumType(1))
	fmt.Println(IsNumType(1.1))
	fmt.Println(IsNumType("1"))

	// Output:
	// true
	// true
	// false
}

func ExampleIsPort() {
	fmt.Println(IsPort("65535"))
	fmt.Println(IsPort("65536"))
	fmt.Println(IsPort("0"))

	// Output:
	// true
	// false
	// false
}

func ExampleIsPrefixOrSuffix() {
	fmt.Println(IsPrefixOrSuffix("123", "12"))
	fmt.Println(IsPrefixOrSuffix("123", "23"))
	fmt.Println(IsPrefixOrSuffix("123", "2"))

	// Output:
	// true
	// true
	// false
}

func ExampleIsRange() {
	fmt.Println(IsRange(1, 1, 2))
	fmt.Println(IsRange(2, 1, 2))
	fmt.Println(IsRange(3, 1, 2))

	// Output:
	// true
	// true
	// false
}

func ExampleIsRegex() {
	fmt.Println(IsRegex(`^[0-9]+$`))
	fmt.Println(IsRegex(`abc(`))

	// Output:
	// true
	// false
}

func ExampleIsRegexMatch() {
	fmt.Println(IsRegexMatch("1", `^[0-9]+$`))
	fmt.Println(IsRegexMatch("2", `abc(`))
	fmt.Println(IsRegexMatch("a", `^[0-9]+$`))

	// Output:
	// true
	// false
	// false
}

func ExampleRuneLength() {
	fmt.Println(RuneLength("1", 1, 2))
	fmt.Println(RuneLength("123", 1, 2))

	// Output:
	// true
	// false
}

func ExampleStringLength() {
	fmt.Println(StringLength("1", 1, 2))
	fmt.Println(StringLength("2", 1, 2))
	fmt.Println(StringLength("3", 3, 4))

	// Output:
	// true
	// true
	// false
}
