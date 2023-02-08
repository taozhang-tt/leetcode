package main

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func Test(t *testing.T) {
	Convey("Test Reverse", t, func() {
		nums := []int{2, 3, 1, 0, 2, 5, 3}
		num := findRepeatNumber2(nums)
		So(num, ShouldEqual, 2)

		nums = []int{0, 1, 2, 3, 4, 11, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
		num = findRepeatNumber2(nums)
		So(num, ShouldEqual, 11)

		nums = []int{0, 1, 2, 3, 4, 11, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
		num = findRepeatNumber3(nums)
		So(num, ShouldEqual, 11)
	})
}
