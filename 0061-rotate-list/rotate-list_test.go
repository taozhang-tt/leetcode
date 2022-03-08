package rotatelist

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func Test(t *testing.T) {
	Convey("Test rotateRight", t, func() {
		head := &ListNode{0, &ListNode{1, &ListNode{2, nil}}}
		root := rotateRight(head, 1)
		ret := []int{2, 0, 1}
		So(List2Slice(root), ShouldResemble, ret)

		head = &ListNode{0, &ListNode{1, &ListNode{2, nil}}}
		root = rotateRight(head, 2)
		ret = []int{1, 2, 0}
		So(List2Slice(root), ShouldResemble, ret)

		head = &ListNode{0, &ListNode{1, &ListNode{2, nil}}}
		root = rotateRight(head, 3)
		ret = []int{0, 1, 2}
		So(List2Slice(root), ShouldResemble, ret)
	})
}
