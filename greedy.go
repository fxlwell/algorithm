package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 2, 3, 4, 3, 4, 5, 6, 7, 6, 7, 8, 9, 10}
	//data := []int{3, 4, 5, 6, 7}

	fmt.Println("最长连续递增字符串儿")
	fmt.Println("普通算法:", normal(data))
	fmt.Println("贪心算法:", greedy(data))
	fmt.Println("柠檬水找钱:", lemon([]int{5, 10, 5, 10, 10}))
}

func lemon(data []int) bool {
	five := 0
	ten := 0
	for _, m := range data {
		if m == 5 {
			five++
		} else if m == 10 {
			if five > 0 {
				ten++
				five--
			} else {
				return false
			}
		} else {
			if five > 0 && ten > 0 {
				five--
				ten--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func greedy(data []int) int {

	start := 0
	max := 0

	for i := 1; i < len(data); i++ {
		if data[i] <= data[i-1] {
			start = i
		}
		max = Max(max, i-start+1)
	}

	return max
}

func normal(data []int) [][]int {
	length := len(data)
	if length <= 1 {
		return [][]int{data}
	}

	markStart := false

	rod := make(map[int]int)

	tmp := data[0]
	indexStart := 0
	indexLen := 0

	for i := 1; i < length; i++ {
		if data[i]-tmp == 1 {
			if markStart == false {
				indexStart = i - 1
				markStart = true
			}
			indexLen++
			fmt.Println(indexStart, indexLen)
		} else {
			if markStart == true {
				rod[indexStart] = indexLen
				indexStart = 0
				indexLen = 0
			}
			markStart = false
		}
		tmp = data[i]
	}

	if markStart == true {
		rod[indexStart] = indexLen
		indexStart = 0
		indexLen = 0
	}

	maxLength := 0
	for _, l := range rod {
		if l > maxLength {
			maxLength = l
		}
	}

	var result [][]int

	for s, l := range rod {
		if l == maxLength {
			result = append(result, data[s:s+l+1])
		}
	}

	return result
}
