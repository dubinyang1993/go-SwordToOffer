package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(replaceSpace("We are happy."))
}

// 利用strings.Builder长度可变来处理
// 当然也可以拼接字符串，或用[]rune处理
func replaceSpace(s string) string {
	var b strings.Builder
	for _, v := range s {
		if v == ' ' {
			b.WriteString("%20")
		} else {
			b.WriteRune(v)
		}
	}
	return b.String()
}
