package main

import "fmt"

func main() {
	fmt.Println(fib1(4))
	fmt.Println(fib2(4))
	fmt.Println(fib3(4))
}

// 递归，不推荐，提交会超时
func fib1(n int) int {
	if n <= 1 {
		return n
	}
	return (fib1(n-1) + fib1(n-2)) % 1000000007
}

// 动态规划
func fib2(n int) int {
	if n <= 1 {
		return n
	}

	slice := make([]int, n+1)
	slice[0] = 0
	slice[1] = 1
	for i := 2; i <= n; i++ {
		slice[i] = (slice[i-1] + slice[i-2]) % 1000000007
	}
	return slice[n]
}

// 动态规划
func fib3(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, (a+b)%1000000007
	}
	return b
}
