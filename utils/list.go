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
