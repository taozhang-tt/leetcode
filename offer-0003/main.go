package main

import (
	"sort"
)

// 剑指 Offer 03. 数组中重复的数字
// https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
// tip: 数组、哈希、原地交换

// 暴力法
func findRepeatNumber(nums []int) int {
	cached := make(map[int]bool)
	for _, num := range nums {
		if _, ok := cached[num]; ok {
			return num
		} else {
			cached[num] = true
		}
	}

	return 0
}

// 原地交换
func findRepeatNumber2(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i] == i {
			i++
			continue
		}
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
	}

	return -1
}

// 先排序再遍历
func findRepeatNumber3(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}

	return -1
}
