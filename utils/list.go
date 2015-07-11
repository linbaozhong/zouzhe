package utils

import (
	"sort"
)

//列表是否包含给定项
func ListContains(list []interface{}, key interface{}) (finded bool) {

	for _, v := range list {
		if v == key {
			finded = true
			break
		}
	}
	return
}

//字符串数组中是否包含给定项
func StringsContains(list []string, key string) (finded bool) {
	for _, v := range list {
		if v == key {
			finded = true
			break
		}
	}
	return
}

//移除slice中的元素
func RemoveStringSlice(s string, slice []string) []string {
	sort.Strings(slice)
	i := sort.SearchStrings(slice, s)

	if i < len(slice) && slice[i] == s {
		return append(slice[:i], slice[i+1:]...)
	}
	return slice
}

//int slice to string slice
func Int64s2Strings(i64 []int64) []string {
	s := make([]string, len(i64))
	for index, i := range i64 {
		s[index] = Int642str(i)
	}
	return s
}
