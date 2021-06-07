package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}
	fmt.Println(findNumberIn2DArray(matrix, 8))
	fmt.Println(findNumberIn2DArray2(matrix, 10))
}

// 从右下角开始比较
func findNumberIn2DArray(matrix [][]int, target int) bool {
	line := len(matrix) - 1 // 行
	j := 0                  // 列
	for line >= 0 && j < len(matrix[line]) {
		if matrix[line][j] == target {
			return true
		} else if matrix[line][j] > target {
			line--
		} else if matrix[line][j] < target {
			j++
		}
	}
	return false
}

// 二分查找
func findNumberIn2DArray2(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		j := len(matrix[i]) - 1
		if j >= 0 && matrix[i][0] <= target && matrix[i][j] >= target {
			if binarySearch(matrix[i], target) {
				return true
			}
		}
	}
	return false
}

func binarySearch(nums []int, target int) bool {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
