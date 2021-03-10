package main

import "fmt"

func main() {
	fmt.Println(minArray1([]int{3, 1, 3}))
	fmt.Println(minArray2([]int{3, 1, 3}))
}

// 循环比较
// 一般情况下原递增数组的尾部和头部连接，且尾部大于头部
// 所有元素相等或者直接输入的就是原递增数组，就需要比较到倒数第二位
func minArray1(numbers []int) int {
	length := len(numbers)

	for i := 0; i < length-1; i++ {
		if numbers[i] > numbers[i+1] {
			return numbers[i+1]
		} else {
			if i == length-2 {
				return numbers[0]
			}
		}
	}
	return numbers[0]
}

// 二分查找
// 原始递增数组是排序数组，要想到用二分
// 假设旋转数组最后一个元素为x，则最小值右侧元素小于等于x，最小值左侧元素大于等于x
// 比较mid和high
// mid<high则mid在最小值右侧，忽略mid右侧
// mid>high则mid在最小值左侧，忽略mid左侧
// mid=high无法判断左右，但因为重复，可以忽略右端点
func minArray2(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		mid := low + (high-low)/2
		if numbers[mid] < numbers[high] {
			high = mid
		} else if numbers[mid] > numbers[high] {
			low = mid + 1
		} else {
			high--
		}
	}
	return numbers[low]
}
