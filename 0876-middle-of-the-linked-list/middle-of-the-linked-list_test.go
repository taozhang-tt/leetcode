package middleofthelinkedlist

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func Test(t *testing.T) {
	Convey("Test middleNode", t, func() {
		head := &ListNode{1, nil}
		middle := middleNode(head)
		So(middle.Val, ShouldEqual, 1)

		head = &ListNode{1, &ListNode{2, nil}}
		middle = middleNode(head)
		So(middle.Val, ShouldEqual, 2)

		head = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
		middle = middleNode(head)
		So(middle.Val, ShouldEqual, 2)

		head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
		middle = middleNode(head)
		So(middle.Val, ShouldEqual, 3)
	})

	Convey("Test middleNode2", t, func() {
		head := &ListNode{1, nil}
		middle := middleNode2(head)
		So(middle.Val, ShouldEqual, 1)

		head = &ListNode{1, &ListNode{2, nil}}
		middle = middleNode2(head)
		So(middle.Val, ShouldEqual, 1)

		head = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
		middle = middleNode2(head)
		So(middle.Val, ShouldEqual, 2)

		head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
		middle = middleNode2(head)
		So(middle.Val, ShouldEqual, 2)
	})

	Convey("Test middleNode3", t, func() {

		head := &ListNode{1, &ListNode{2, nil}}
		middle := middleNode3(head)
		So(middle.Val, ShouldEqual, 1)

		head = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
		middle = middleNode3(head)
		So(middle.Val, ShouldEqual, 1)

		head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
		middle = middleNode3(head)
		So(middle.Val, ShouldEqual, 2)
	})
}
