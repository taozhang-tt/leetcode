package main

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func Test(t *testing.T) {
	Convey("Test", t, func() {
		nums := []int{0, 1, 2, 3}
		target := 1
		expect := 1
		So(binarySearch(nums, target), ShouldEqual, expect)
	})

	Convey("Test First", t, func() {
		nums := []int{5, 7, 7, 8, 8, 10}
		target := 8
		expect := 3
		So(binarySearchFirst(nums, target), ShouldEqual, expect)
	})

}
