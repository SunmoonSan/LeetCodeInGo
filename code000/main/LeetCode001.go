package main

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的 两个 整数。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

func twoSum2(nums []int, target int) []int {
	kvs := map[int]int{} // 初始化一个map
	for i, v := range nums {
		_, ok := kvs[target-v]
		if ok {
			return []int{kvs[target-v], i}
		} else {
			kvs[v] = i
		}
	}
	return []int{-1, -1}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum2(nums, target))
}
