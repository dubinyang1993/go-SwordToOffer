package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(majorityElement1([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
	fmt.Println(majorityElement2([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
	fmt.Println(majorityElement3([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
	fmt.Println(majorityElement4([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
}

// 排序取中间值
func majorityElement1(nums []int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}

	sort.Ints(nums)
	mid := length / 2

	// 如果题目不确定是否有超过数组一半长度的数字，则需要判断一下
	count := 0
	for _, v := range nums {
		if v == nums[mid] {
			count++
		}
	}
	if count > mid {
		return nums[mid]
	} else {
		return -1
	}
}

// 哈希表
func majorityElement2(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num] += 1
	}

	key := 0
	value := 0
	for k, v := range m {
		if v > value {
			value = v
			key = k
		}
	}

	// 判断个数
	mid := len(nums) / 2
	if value > mid {
		return key
	} else {
		return -1
	}
}

// 投票法，循环前处理count=0
func majorityElement3(nums []int) int {
	result := -1
	count := 0
	for i, v := range nums {
		if i == 0 {
			count = 1
			result = v
		} else {
			if count == 0 {
				result = v
				count = 1
			} else {
				if result == v {
					count++
				} else {
					count--
				}
			}
		}
	}

	// 判断个数
	count = 0
	for _, v := range nums {
		if v == result {
			count++
		}
	}
	mid := len(nums) / 2
	if count > mid {
		return result
	} else {
		return -1
	}
}

// 投票法，循环后处理count=0
func majorityElement4(nums []int) int {
	result := -1
	count := 0
	for i, v := range nums {
		if i == 0 {
			count = 1
			result = v
		} else {
			if result == v {
				count++
			} else {
				count--
			}
		}
		if count == 0 {
			result = v
			count = 1
		}
	}

	// 判断个数
	count = 0
	for _, v := range nums {
		if v == result {
			count++
		}
	}
	mid := len(nums) / 2
	if count > mid {
		return result
	} else {
		return -1
	}
}
