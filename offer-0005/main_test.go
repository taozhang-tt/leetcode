package main

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func TestDeleteDuplicates(t *testing.T) {
	Convey("Test ReplaceSpace", t, func() {
		var testCases = []struct {
			s      string
			expect string
		}{
			{"We are happy.", "We%20are%20happy."},
			{"我 很 开 心。", "我%20很%20开%20心。"},
		}
		for _, test := range testCases {
			So(replaceSpace(test.s), ShouldEqual, test.expect)
		}

		for _, test := range testCases {
			So(replaceSpace2(test.s), ShouldEqual, test.expect)
		}
	})
}
