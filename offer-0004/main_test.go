package main

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func Test(t *testing.T) {
	Convey("Test", t, func() {
		matrix := [][]int{
			{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		}
		So(findNumberIn2DArray2(matrix, 5), ShouldBeTrue)
		So(findNumberIn2DArray(matrix, 5), ShouldBeTrue)

		matrix = [][]int{
			{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		}
		So(findNumberIn2DArray2(matrix, 20), ShouldBeFalse)
		So(findNumberIn2DArray(matrix, 20), ShouldBeFalse)

		matrix = [][]int{
			{-5},
		}
		So(findNumberIn2DArray2(matrix, -5), ShouldBeTrue)
		So(findNumberIn2DArray(matrix, -5), ShouldBeTrue)

		matrix = [][]int{
			{-1, 3},
		}
		So(findNumberIn2DArray2(matrix, 3), ShouldBeTrue)
		So(findNumberIn2DArray(matrix, 3), ShouldBeTrue)

		matrix = [][]int{
			{1, 4},
			{2, 5},
		}
		So(findNumberIn2DArray2(matrix, 2), ShouldBeTrue)
		So(findNumberIn2DArray(matrix, 2), ShouldBeTrue)
	})

}
