// 翻转字符串
// 问题描述

// 请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。

// 给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。

package main

import (
	"fmt"
)

func main() {

	s := "asdfsdf"
	fmt.Print(StringReverse(s))

}

func StringReverse(s string) string {

	C := len(s)

	if C <= 1 {
		return s
	}

	ss := []rune(s)

	k := 0
	v := C - 1

	for k < v {
		ss[k], ss[v] = ss[v], ss[k]
		k++
		v--
	}

	return string(ss)

}
