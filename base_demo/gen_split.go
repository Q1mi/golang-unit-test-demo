package base_demo

import "strings"

//// Split 把字符串s按照给定的分隔符sep进行分割返回字符串切片
//func Split(s, sep string) (result []string) {
//	i := strings.Index(s, sep)
//
//	for i > -1 {
//		result = append(result, s[:i])
//		s = s[i+1:]
//		i = strings.Index(s, sep)
//	}
//	result = append(result, s)
//	return
//}

func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)
	result = make([]string, 0, strings.Count(s, sep)+1)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):] // 这里使用len(sep)获取sep的长度
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
