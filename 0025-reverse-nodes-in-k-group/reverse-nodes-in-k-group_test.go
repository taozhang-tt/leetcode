package reversenodesinkgroup

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func Test(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	Convey("reverseK", t, func() {
		r1 := reverseK(head, 2)
		r2 := reverseK(r1, 2)
		So(r2, ShouldResemble, head)

		r1 = reverseK(head, 3)
		r2 = reverseK(r1, 3)
		So(r2, ShouldResemble, head)
	})
	Convey("cutK", t, func() {
		head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
		ret := cutK(head, 2)
		So(len(ret), ShouldEqual, 3)

		head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
		ret = cutK(head, 3)
		So(len(ret), ShouldEqual, 2)

		head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
		ret = cutK(head, 1)
		So(len(ret), ShouldEqual, 5)
	})

	Convey("reverseKGroup", t, func() {
		r1 := reverseKGroup(head, 1)
		r2 := reverseKGroup(r1, 1)
		So(r2, ShouldResemble, head)

		r1 = reverseKGroup(head, 2)
		r2 = reverseKGroup(r1, 2)
		So(r2, ShouldResemble, head)

		r1 = reverseKGroup(head, 3)
		r2 = reverseKGroup(r1, 3)
		So(r2, ShouldResemble, head)

		r1 = reverseKGroup(head, 4)
		r2 = reverseKGroup(r1, 4)
		So(r2, ShouldResemble, head)

		r1 = reverseKGroup(head, 5)
		r2 = reverseKGroup(r1, 5)
		So(r2, ShouldResemble, head)
	})
}
