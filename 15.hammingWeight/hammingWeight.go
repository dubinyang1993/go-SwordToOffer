package main

import "fmt"

func main() {
	fmt.Println(hammingWeight1(11))
	fmt.Println(hammingWeight2(00000000000000000000000000001011))

}

// num&1=1最右位必定为1，num&1=0最右位必定为0
// 处理完最右位，右移一位
func hammingWeight1(num uint32) int {
	var res int
	for num > 0 {
		if num&1 == 1 {
			res++
		}
		num >>= 1
	}
	return res
}

//       n = 11000
//     n-1 = 10111
// n&(n-1) = 10000
func hammingWeight2(num uint32) int{
	var res int
	for num > 0 {
		res++
		num &= (num-1)
	}
	return res
}
