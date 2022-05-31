package main

import (
	"fmt"
)

//二分查找,实现sqrt()函数，返回整数部分
func main() {

	fmt.Println(bs(9))
	fmt.Println(bs(10))
	fmt.Println(bs(11))
	fmt.Println(bs(12))
	fmt.Println(bs(13))
	fmt.Println(bs(14))
	fmt.Println(bs(15))
	fmt.Println(bs(16))
}

func bs(num int) int {
	l := 0
	r := num
	mid := 0

	for {
		mid = (l + r) / 2
		if mid*mid > num {
			r = mid - 1
		} else if mid*mid == num {
			break
		} else {
			l = mid + 1
		}

		if l > r {
			break
		}
	}

	return mid
}
