package validatorjez

import (
	"testing"

	"golang.org/x/exp/constraints"
)

func TestIsBase64(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "正确", args: args{s: "SGVsbG8sIHdvcmxkIQ=="}, want: true},
		{name: "空字符串", args: args{s: ""}, want: false},
		{name: "无效", args: args{s: "123"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBase64(tt.args.s); got != tt.want {
				t.Errorf("IsBase64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBase64URL(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "正确", args: args{s: "SGVsbG8sIHdvcmxkIQ=="}, want: true},
		{name: "空字符串", args: args{s: ""}, want: false},
		{name: "无效", args: args{s: "123"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBase64URL(tt.args.s); got != tt.want {
				t.Errorf("IsBase64URL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsChineseMainlandIDCard(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "无X或x", args: args{s: "210911192105130714"}, want: true},
		{name: "X", args: args{s: "11010519491231002X"}, want: true},
		{name: "x", args: args{s: "11010519491231002x"}, want: true},
		{name: "位数不对", args: args{s: "123456"}, want: false},
		{name: "省份不对", args: args{s: "990911192105130714"}, want: false},
		{name: "生日年份小于1900", args: args{s: "210911189905130714"}, want: false},
		{name: "生日年份大于今日", args: args{s: "210911222205130714"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsChineseMainlandIDCard(tt.args.s); got != tt.want {
				t.Errorf("IsChineseMainlandIDCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsChineseMainlandPhoneNumber(t *testing.T) {
	type args struct {
		s        string
		withCode []bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "正确", args: args{s: "13666666666"}, want: true},
		{name: "+86 ok", args: args{s: "+8613666666666", withCode: []bool{true}}, want: true},
		{name: "+86 not ok", args: args{s: "+8613666666666"}, want: false},
		{name: "位数不对", args: args{s: "13666"}, want: false},
		{name: "无区号", args: args{s: "12666666666"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsChineseMainlandPhoneNumber(tt.args.s, tt.args.withCode...); got != tt.want {
				t.Errorf("IsChineseMainlandPhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFloat(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "3.14", args: args{s: "3.14"}, want: true},
		{name: "3", args: args{s: "3"}, want: true},
		{name: ".3", args: args{s: ".3"}, want: true},
		{name: "3.", args: args{s: "3."}, want: true},
		{name: "3.1.1", args: args{s: "3.1.1"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFloat(tt.args.s); got != tt.want {
				t.Errorf("IsFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFloatType(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "3.14", args: args{v: 3.14}, want: true},
		{name: "3", args: args{v: float64(3)}, want: true},
		{name: "3.", args: args{v: 3.}, want: true},
		{name: ".3", args: args{v: .3}, want: true},
		{name: "字符串", args: args{v: "3.14"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFloatType(tt.args.v); got != tt.want {
				t.Errorf("IsFloatType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIP(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1.1.1.1", args: args{s: "1.1.1.1"}, want: true},
		{name: "127.0.0.1", args: args{s: "127.0.0.1"}, want: true},
		{name: "ip6", args: args{s: "2001:0db8:85a3:0000:0000:8a2e:0370:7334"}, want: true},
		{name: "127.0.0.1:8080", args: args{s: "127.0.0.1:8080"}, want: false},
		{name: "localhost", args: args{s: "localhost"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIP(tt.args.s); got != tt.want {
				t.Errorf("IsIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIPv4(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1.1.1.1", args: args{s: "1.1.1.1"}, want: true},
		{name: "127.0.0.1", args: args{s: "127.0.0.1"}, want: true},
		{name: "ip6", args: args{s: "2001:0db8:85a3:0000:0000:8a2e:0370:7334"}, want: false},
		{name: "127.0.0.1:8080", args: args{s: "127.0.0.1:8080"}, want: false},
		{name: "localhost", args: args{s: "localhost"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIPv4(tt.args.s); got != tt.want {
				t.Errorf("IsIPv4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIPv6(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "127.0.0.1", args: args{s: "127.0.0.1"}, want: false},
		{name: "ip6", args: args{s: "2001:0db8:85a3:0000:0000:8a2e:0370:7334"}, want: true},
		{name: "127.0.0.1:8080", args: args{s: "127.0.0.1:8080"}, want: false},
		{name: "localhost", args: args{s: "localhost"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIPv6(tt.args.s); got != tt.want {
				t.Errorf("IsIPv6() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIn(t *testing.T) {
	type args[T comparable] struct {
		s    T
		list []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{name: "1 in []int{1,2,3}", args: args[int]{s: 1, list: []int{1, 2, 3}}, want: true},
		{name: "4 in []int{1,2,3}", args: args[int]{s: 4, list: []int{1, 2, 3}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIn(tt.args.s, tt.args.list...); got != tt.want {
				t.Errorf("IsIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{s: "1"}, want: true},
		{name: "01", args: args{s: "01"}, want: true},
		{name: "1.1", args: args{s: "1.1"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInt(tt.args.s); got != tt.want {
				t.Errorf("IsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIntType(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{v: 1}, want: true},
		{name: "1.1", args: args{v: 1.1}, want: false},
		{name: "字符串", args: args{v: "1"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIntType(tt.args.v); got != tt.want {
				t.Errorf("IsIntType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsJSON(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "json", args: args{s: `{"name":"tom"}`}, want: true},
		{name: "无效", args: args{s: `abc`}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsJSON(tt.args.s); got != tt.want {
				t.Errorf("IsJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLatitude(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "123.123", args: args{str: "-74.0060"}, want: true},
		{name: "123.123.123", args: args{str: "123.123.123"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLatitude(tt.args.str); got != tt.want {
				t.Errorf("IsLatitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLongitude(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "123.123", args: args{str: "-45.12345"}, want: true},
		{name: "123.123.123", args: args{str: "123.123.123"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLongitude(tt.args.str); got != tt.want {
				t.Errorf("IsLongitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNum(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "123", args: args{s: "123"}, want: true},
		{name: "123.123", args: args{s: "123.123"}, want: true},
		{name: "123.123.123", args: args{s: "123.123.123"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNum(tt.args.s); got != tt.want {
				t.Errorf("IsNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNumType(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{v: 1}, want: true},
		{name: "1.1", args: args{v: 1.1}, want: true},
		{name: "字符串", args: args{v: "1"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumType(tt.args.v); got != tt.want {
				t.Errorf("IsNumType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "65535", args: args{s: "65535"}, want: true},
		{name: "65536", args: args{s: "65536"}, want: false},
		{name: "0", args: args{s: "0"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPort(tt.args.s); got != tt.want {
				t.Errorf("IsPort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrefixOrSuffix(t *testing.T) {
	type args struct {
		s   string
		pos string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "123", args: args{s: "123", pos: "12"}, want: true},
		{name: "123", args: args{s: "123", pos: "23"}, want: true},
		{name: "123", args: args{s: "123", pos: "2"}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrefixOrSuffix(tt.args.s, tt.args.pos); got != tt.want {
				t.Errorf("IsPrefixOrSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRange(t *testing.T) {
	type args[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		s   T
		min T
		max T
	}
	type testCase[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{name: "1", args: args[int]{s: 1, min: 1, max: 2}, want: true},
		{name: "2", args: args[int]{s: 2, min: 1, max: 2}, want: true},
		{name: "3", args: args[int]{s: 3, min: 1, max: 2}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRange(tt.args.s, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("IsRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRegex(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "regex", args: args{s: `^[0-9]+$`}, want: true},
		{name: "无效", args: args{s: "abc("}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRegex(tt.args.s); got != tt.want {
				t.Errorf("IsRegex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRegexMatch(t *testing.T) {
	type args struct {
		s     string
		regex string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{s: "1", regex: `^[0-9]+$`}, want: true},
		{name: "错误表达式", args: args{s: "2", regex: `abc(`}, want: false},
		{name: "a", args: args{s: "a", regex: `^[0-9]+$`}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRegexMatch(tt.args.s, tt.args.regex); got != tt.want {
				t.Errorf("IsRegexMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneLength(t *testing.T) {
	type args struct {
		s   string
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{s: "1", min: 1, max: 2}, want: true},
		{name: "123", args: args{s: "123", min: 1, max: 2}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RuneLength(tt.args.s, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("RuneLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringLength(t *testing.T) {
	type args struct {
		s   string
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{s: "1", min: 1, max: 2}, want: true},
		{name: "2", args: args{s: "2", min: 1, max: 2}, want: true},
		{name: "3", args: args{s: "3", min: 3, max: 4}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringLength(tt.args.s, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("IsStringLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
