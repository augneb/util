package utils

import (
	"strings"
	"sort"
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

func InSlice(target string, array []string) bool {
	sort.Strings(array)
	i := sort.SearchStrings(array, target)
	if i < len(array) && array[i] == target {
		return true
	}

	return false
}

