package main

import "fmt"

func main() {
	fmt.Println(exchange1([]int{1, 2, 3, 4}))
	fmt.Println(exchange2([]int{1, 2, 3, 4}))
	fmt.Println(exchange3([]int{1, 2, 3, 4}))
	fmt.Println(exchange4([]int{1, 2, 3, 4}))

}

// 首尾指针，首尾一偶一奇则交换
func exchange1(nums []int) []int {
	low := 0
	high := len(nums) - 1
	for low < high {
		if nums[low]&1 == 1 {
			low++
			continue
		}
		if nums[high]&1 == 0 {
			high--
			continue
		}
		nums[low], nums[high] = nums[high], nums[low]
		low++
		high--
	}
	return nums
}

// 快慢指针，快的向前找奇数位置，慢的指向下个奇数应当存放的位置
func exchange2(nums []int) []int {
	var slow, fast int
	for fast < len(nums) {
		if nums[fast]&1 == 1 {
			if slow != fast {
				nums[slow], nums[fast] = nums[fast], nums[slow]
			}
			slow++
		}
		fast++
	}
	return nums
}

// 如果题目要求移动后，奇数偶数元素相对位置不变，可参考冒泡排序
// 不过这里提交会超时
func exchange3(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		flag := true
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j]&1 == 0 && nums[j+1]&1 == 1 {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
	return nums
}

// 暴力遍历，拼接奇数偶数
func exchange4(nums []int) []int {
	var slice []int
	for i := 0; i < len(nums); {
		if nums[i] & 1 == 0 {
			slice = append(slice, nums[i])
			nums = append(nums[:i], nums[i+1:]...)
		}  else {
			i++
		}
	}
	nums = append(nums, slice...)
	return nums
}
