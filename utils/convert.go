package utils

import (
	"fmt"
	//"strconv"
)

// //字符串转长整型
// func Str2int64(s string) (int64, error) {
// 	return strconv.ParseInt(s, 10, 64)
// }

// //字符串转整形
// func Str2int(s string) (int, error) {
// 	return strconv.Atoi(s)
// }

// //整形转字符串
// func Int2str(i int) string {
// 	return strconv.Itoa(i)
// }
// func Int642str(i int64) string {
// 	return strconv.FormatInt(i, 10)
// }

//字符串接口 转 字符串
func Interface2str(v interface{}) string {
	switch v.(type) {
	case float64:
		return fmt.Sprintf("%g", v)
	case int, int32, int64:
		return fmt.Sprintf("%d", v)
	case nil:
		return "null"
	default:
		return fmt.Sprintf("%s", v)
	}

}
