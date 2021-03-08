package main

import "fmt"

func main() {
	fmt.Println(findRepeatNumber1([]int{2, 3, 1, 0, 2, 5, 3}))
}

func findRepeatNumber1(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return num
		}
		m[num] = 1
	}
	return -1
}

