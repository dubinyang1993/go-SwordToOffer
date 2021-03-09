package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findRepeatNumber1([]int{2, 3, 1, 0, 2, 5, 3}))
	fmt.Println(findRepeatNumber2([]int{2, 3, 1, 0, 2, 5, 3}))
	fmt.Println(findRepeatNumber3([]int{2, 3, 1, 0, 2, 5, 3}))
}

// 哈希表
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

// 排序，如果有重复，元素必定相邻
func findRepeatNumber2(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

// 原地置换，如果不重复，顺序情况下元素和小标一致，必定是0到n-1
func findRepeatNumber3(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for i != nums[i] {
			if nums[i]== nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}
