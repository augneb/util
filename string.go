package util

import (
	"regexp"
	"strings"
	"unicode"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")

// 驼峰转换为下划线
func ToSnakeCase(str string, toLower bool) string {
	snake := matchAllCap.ReplaceAllString(matchFirstCap.ReplaceAllString(str, "${1}_${2}"), "${1}_${2}")

	if toLower {
		return strings.ToLower(snake)
	}

	return snake
}

func Ucfirst(s string) string {
	if s == "" {
		return s
	}

	a := []rune(s)
	a[0] = unicode.ToUpper(a[0])

	return string(a)
}
