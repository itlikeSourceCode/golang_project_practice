package main

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

import "fmt"

/*
执行用时 : 48 ms, 在所有 Go 提交中击败了30.34%的用户
内存消耗 : 2.9 MB, 在所有 Go 提交中击败了100.00%的用户
*/
func twoSum2(nums []int, target int) []int {
	// 时间复杂度: n^2
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

/*
执行用时 : 4 ms , 在所有 Go 提交中击败了 96.88%的用户
内存消耗 : 3.8 MB , 在所有 Go 提交中击败了 52.44%的用户
*/
func twoSum(nums []int, target int) []int {
	var needs = make(map[int]int)

	// 时间复杂度：O(n)
	for i := 0; i < len(nums); i++ {
		need := target - nums[i]

		if index, ok := needs[need]; ok {
			return []int{index, i}
		}

		needs[nums[i]] = i
	}

	return []int{}
}

// 执行入口
func main() {
	var result = twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(result)

	result = twoSum([]int{1, 2, 3, 4}, 4)
	fmt.Println(result)
}
