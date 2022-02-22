package linkedlistcycle

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func TestHasCycle(t *testing.T) {
	var (
		node1 = &ListNode{Val: 1}
		node2 = &ListNode{Val: 2}
		node3 = &ListNode{Val: 3}
		node4 = &ListNode{Val: 4}
	)
	Convey("HasCycle1", t, func() {
		withoutCycle := node1
		node1.Next = node2
		node2.Next = node3
		node3.Next = node4
		So(hasCycle(withoutCycle), ShouldBeFalse)
		So(hasCycle2(withoutCycle), ShouldBeFalse)
	})
	Convey("HasCycle2", t, func() {
		withCycle := node1
		node1.Next = node2
		node2.Next = node3
		node3.Next = node4
		node4.Next = node2
		So(hasCycle(withCycle), ShouldBeTrue)
		So(hasCycle2(withCycle), ShouldBeTrue)
	})
}
