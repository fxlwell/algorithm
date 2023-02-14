package main

import "fmt"

func main() {
	nums := []int{2, 1, 9, 7, 11, 15}
	target := 9

	fmt.Println(FirstIdeas(nums, target))
	fmt.Println(MapIdeas(nums, target))

}

func FirstIdeas(nums []int, target int) []int {
	length := len(nums)
	if length == 0 {
		return nil
	}
	if length < 2 && nums[0] != target {
		return nil
	}

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func MapIdeas(nums []int, target int) []int {
	length := len(nums)
	if length == 0 {
		return nil
	}
	if length < 2 && nums[0] != target {
		return nil
	}

	m := make(map[int]int)

	for k, v := range nums {
		if i, ok := m[target-v]; ok {
			return []int{k, i}
		} else {
			m[v] = k
		}
	}

	return nil
}

func KnumSum(nums []int, k int, target int) []int {
}
