/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。



示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]


提示：

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
只会存在一个有效答案




来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
package main

import "fmt"

func main() {
	target := 9
	data := []int{2, 7, 11, 15}
	indexs := Normal(data, target)
	fmt.Println("传统方法:", indexs)
	indexV2s := NormalV2(data, target)
	fmt.Println("最优解法:", indexV2s)

	target = 6
	data = []int{3, 3}
	indexs = Normal(data, target)
	fmt.Println("传统方法:", indexs)
	indexV2s = NormalV2(data, target)
	fmt.Println("最优解法:", indexV2s)

	target = 6
	data = []int{3, 2, 4}
	indexs = Normal(data, target)
	fmt.Println("传统方法:", indexs)
	indexV2s = NormalV2(data, target)
	fmt.Println("最优解法:", indexV2s)

}

func Normal(data []int, target int) []int {
	l := len(data)
	for i := 0; i < l; i++ {
		v1 := data[i]
		for j := i + 1; j < l; j++ {
			v2 := data[j]
			if v1+v2 == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func NormalV2(data []int, target int) []int {
	tmpMap := make(map[int]int, len(data))
	for k, v := range data {
		if index, ok := tmpMap[target-v]; ok {
			return []int{index, k}
		} else {
			tmpMap[v] = k
		}
	}
	return nil
}
