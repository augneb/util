package utils

import (
	"strings"
	"reflect"
)

// 去除空行
func ReduceEmptyElements(items []string) []string {
	result := []string{}
	for _, text := range items {
		if strings.Trim(text, " ") != "" {
			result = append(result, text)
		}
	}

	return result
}

func InSlice(elt, slice interface{}) bool {
	v := reflect.Indirect(reflect.ValueOf(slice))

	for i := 0; i < v.Len(); i++ {
		if reflect.DeepEqual(v.Index(i).Interface(), elt) {
			return true
		}
	}

	return false
}
