package main

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func TestSearch(t *testing.T) {
	type testCase struct {
		nums   []int
		target int
		expect int
	}
	Convey("TestSearch", t, func() {
		testcases := []testCase{
			{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
			{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		}

		for _, testcase := range testcases {
			So(search(testcase.nums, testcase.target), ShouldEqual, testcase.expect)
			So(search2(testcase.nums, testcase.target), ShouldEqual, testcase.expect)
		}
	})
}
