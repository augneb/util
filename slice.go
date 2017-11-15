package util

import (
	"strings"
	"reflect"
)

// 去除空行
func SliceStringFilter(items []string) []string {
	result := []string{}
	for _, text := range items {
		if strings.Trim(text, " ") != "" {
			result = append(result, text)
		}
	}

	return result
}

func SliceIn(elt, slice interface{}) bool {
	v := reflect.Indirect(reflect.ValueOf(slice))

	for i := 0; i < v.Len(); i++ {
		if reflect.DeepEqual(v.Index(i).Interface(), elt) {
			return true
		}
	}

	return false
}

func SliceStringUnique(s *[]string) {
	f := make(map[string]bool, len(*s))
	t := 0
	for i, v := range *s {
		if _, ok := f[v]; !ok {
			f[v] = true
			(*s)[t] = (*s)[i]
			t++
		}
	}

	*s = (*s)[:t]
}
