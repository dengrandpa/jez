package validatorjez

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/exp/constraints"
)

// IsBase64 是否为base64
func IsBase64(s string) bool {
	if len(s) == 0 {
		return false
	}
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}

// IsBase64URL 是否为base64url
func IsBase64URL(s string) bool {
	if len(s) == 0 {
		return false
	}
	_, err := base64.URLEncoding.DecodeString(s)
	return err == nil
}

// IsChineseMainlandIDCard 是否为中国大陆身份证号码
func IsChineseMainlandIDCard(s string) bool {
	// 全部字符应为数字，最后一位可为x或X
	if !rxChineseMainlandIDCard.MatchString(s) {
		return false
	}

	// 验证省份代码, 根据GB/T2260 补全全部省份代码
	_, ok := provinceKv[s[0:2]]
	if !ok {
		return false
	}

	// 验证生日，必须大于 minYear 小于当前年份
	birthStr := fmt.Sprintf("%s-%s-%s", s[6:10], s[10:12], s[12:14])
	birthday, err := time.Parse("2006-01-02", birthStr)
	if err != nil || birthday.After(time.Now()) || birthday.Year() < birthStartYear {
		return false
	}

	// 验证校验码
	sum := 0
	for i, c := range s[:17] {
		v, _ := strconv.Atoi(string(c))
		sum += v * factor[i]
	}

	return verifyStr[sum%11] == strings.ToUpper(s[17:18])
}

// IsChineseMainlandPhoneNumber 是否为中国大陆手机号码, withCode为是否可包含国家代码 86 / +86
func IsChineseMainlandPhoneNumber(s string, withCode ...bool) bool {
	if len(withCode) > 0 && withCode[0] {
		s = strings.TrimPrefix(strings.TrimPrefix(s, "+"), "86")
	}

	return rxChinesePhoneNumber.MatchString(s)
}

// IsFloat 是否为浮点数
func IsFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// IsFloatType 是否为浮点数类型
func IsFloatType(v any) bool {
	switch v.(type) {
	case float32, float64:
		return true
	}
	return false
}

// IsIP 是否为IP地址
func IsIP(s string) bool {
	ip := net.ParseIP(s)
	return ip != nil
}

// IsIPv4 是否为IPv4地址
func IsIPv4(s string) bool {
	if !IsIP(s) {
		return false
	}
	return strings.Contains(s, ".")
}

// IsIPv6 是否为IPv6地址
func IsIPv6(s string) bool {
	if !IsIP(s) {
		return false
	}
	return strings.Contains(s, ":")
}

// IsIn 是否在指定列表中
func IsIn[T comparable](s T, list ...T) bool {
	for _, v := range list {
		if s == v {
			return true
		}
	}
	return false
}

// IsInt 是否为整数
func IsInt(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

// IsIntType 是否为整数类型
func IsIntType(v any) bool {
	switch v.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		return true
	}
	return false
}

// IsJSON 是否为json
func IsJSON(s string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(s), &js) == nil
}

// IsLongitude 是否为经度
func IsLongitude(str string) bool {
	return rxLongitude.MatchString(str)
}

// IsLatitude 是否为纬度
func IsLatitude(str string) bool {
	return rxLatitude.MatchString(str)
}

// IsNum 是否为数字(包含浮点数)
func IsNum(s string) bool {
	return IsInt(s) || IsFloat(s)
}

// IsNumType 是否为数字类型(包含浮点数)
func IsNumType(v any) bool {
	return IsIntType(v) || IsFloatType(v)
}

// IsPort 是否为端口号
func IsPort(s string) bool {
	i, err := strconv.ParseInt(s, 10, 64)
	return err == nil && i > 0 && i <= 65535
}

// IsPrefixOrSuffix 是否以指定字符串开头或结尾
func IsPrefixOrSuffix(s, pos string) bool {
	return strings.HasPrefix(s, pos) || strings.HasSuffix(s, pos)
}

// IsRange 数值范围
func IsRange[T constraints.Integer | constraints.Float](s, min, max T) bool {
	return s >= min && s <= max
}

// IsRegex 是否为正则表达式
func IsRegex(s string) bool {
	_, err := regexp.Compile(s)
	return err == nil
}

// IsRegexMatch 是否匹配正则表达式
func IsRegexMatch(s, regex string) bool {
	r, err := regexp.Compile(regex)
	if err != nil {
		return false
	}

	return r.MatchString(s)
}

// RuneLength 字符长度
func RuneLength(s string, min, max int) bool {
	return utf8.RuneCountInString(s) >= min && utf8.RuneCountInString(s) <= max
}

// StringLength 字符串长度
func StringLength(s string, min, max int) bool {
	return len(s) >= min && len(s) <= max
}
