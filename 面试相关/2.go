// 判断字符串中字符是否全都不同问题描述请实现一个算法，确定一个字符串的所有字符【是否全都不同】.
// 这里我们要求【不允许使用额外的存储结构】。
// 给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。 保证字符串中的字符为【ASCII字符】。字符串的长度小于等于【3000】。

package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "afdeq"
	fmt.Print(StringUnique(s))

}

func StringUnique(s string) bool {

	if strings.Count(s, "") > 3000 {
		return false
	}

	for _, c := range s {
		if strings.Count(s, string(c)) > 1 {
			return false
		}
	}

	return true

}
