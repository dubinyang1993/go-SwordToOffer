package main

import "fmt"

func main() {
	fmt.Println(printNumbers1(1))
}

// 循环打印
func printNumbers1(n int) []int {
	num := 1
	for n > 0 {
		num = num * 10
		n--
	}

	var res []int
	for i := 1; i < num; i++ {
		res = append(res, i)
	}
	return res
}

// 考虑大数问题
// todo
func printNumbers2(n int) []int {
	var res []int

	return res
}