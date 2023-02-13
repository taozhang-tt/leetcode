package main

// [剑指 Offer 53 - I. 在排序数组中查找数字 I](https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/solution/mian-shi-ti-53-i-zai-pai-xu-shu-zu-zhong-cha-zha-5/)
// tag: 二分查找

// 暴力法
func search(nums []int, target int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			cnt++
		}
	}
	return cnt
}

func search2(nums []int, target int) int {
	return binarySearchLast(nums, target) - binarySearchFirst(nums, target) + 1
}

// 二分查找
func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

// 二分查找，存在则返回索引，不存在则返回应该插入的位置
func binarySearchIndex(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}

// 二分查找，第一次出现的位置
func binarySearchFirst(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if target <= nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}

// 二分查找，最后一次出现的位置
func binarySearchLast(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if target >= nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return r
}
