package maps

import (
	"fmt"
	"sort"
)

// 基于map的排序
// map是无序的
func mapSort() {
	var a = map[string]int{
		"balana": 1,
		"apple":  3,
		"orange": 2,
		"plume":  8,
		"melon":  6,
		"cherry": 4,
	}

	for k, v := range a {
		fmt.Println(k, v)
	}

	var keys []string
	for k := range a {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, a[k])
	}
}
