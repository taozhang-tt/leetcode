package rotatearray

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func Test(t *testing.T) {
	Convey("Test rotate", t, func() {
		nums := []int{1, 2, 3, 4, 5}
		expect := []int{4, 5, 1, 2, 3}
		rotate(nums, 2)
		So(nums, ShouldResemble, expect)
	})

	Convey("Test rotate2", t, func() {
		nums := []int{1, 2, 3, 4, 5}
		expect := []int{4, 5, 1, 2, 3}
		rotate2(nums, 2)
		So(nums, ShouldResemble, expect)
	})
}
