package validatorjez

import "regexp"

var (
	rxChinesePhoneNumber    = regexp.MustCompile(`^1[3-9]\d{9}$`)
	rxChineseMainlandIDCard = regexp.MustCompile(`^(\d{17})([0-9]|X|x)$`)
	rxLatitude              = regexp.MustCompile(`^[-+]?([1-8]?\d(\.\d+)?|90(\.0+)?)$`)
	rxLongitude             = regexp.MustCompile(`^[-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?)$`)
)

var (
	// 身份证公式
	factor = [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	// 身份证效验位
	verifyStr = [11]string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
	// 身份证起始年份
	birthStartYear = 1900
	// 省份代码
	provinceKv = map[string]struct{}{
		"11": {},
		"12": {},
		"13": {},
		"14": {},
		"15": {},
		"21": {},
		"22": {},
		"23": {},
		"31": {},
		"32": {},
		"33": {},
		"34": {},
		"35": {},
		"36": {},
		"37": {},
		"41": {},
		"42": {},
		"43": {},
		"44": {},
		"45": {},
		"46": {},
		"50": {},
		"51": {},
		"52": {},
		"53": {},
		"54": {},
		"61": {},
		"62": {},
		"63": {},
		"64": {},
		"65": {},
		//"71": {},
		//"81": {},
		//"82": {},
	}
)
