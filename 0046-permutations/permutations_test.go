package permutations

import (
	"fmt"
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func Test(t *testing.T) {
	Convey("Test permute", t, func() {
		nums := []int{}
		ans := permute(nums)
		ret := [][]int{{}}
		So(ans, ShouldResemble, ret)

		nums = []int{1}
		ans = permute(nums)
		ret = [][]int{{1}}
		So(ans, ShouldResemble, ret)

		nums = []int{1, 2}
		ans = permute(nums)
		fmt.Println(ans)
		ret = [][]int{{1, 2}, {2, 1}}
		So(ans, ShouldResemble, ret)

	})
}
