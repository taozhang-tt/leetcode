package rotatearray

// 189. 轮转数组

// 直接翻转
func rotate(nums []int, k int) {
	k %= len(nums)
	if k == 0 {
		return
	}
	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
}

// 全部翻转
// 翻转[0, k % n)
// 翻转[k % n, n)
func reverse(nums []int) {
	for i, n := 0, len(nums); i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}

func rotate2(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}
