package permutations

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	var backtrace func(path, opts []int)
	backtrace = func(path, opts []int) {
		if len(opts) == 0 {
			ans = append(ans, path)
			return
		}
		for i := 0; i < len(opts); i++ {
			currPath := append(path, opts[i])

			currOpts := append(opts[:i:i], opts[i+1:]...)
			backtrace(currPath, currOpts)
		}
	}
	backtrace([]int{}, nums)
	return ans
}
