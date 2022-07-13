package main

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func TestDetectCycle(t *testing.T) {
	head := &ListNode{Val: 3}
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 0}
	node3 := &ListNode{Val: -4}
	head.Next = node1
	node1.Next = node2
	node2.Next = node3
	node3.Next = node1

	Convey("TestHasCycle", t, func() {
		node := detectCycle(head)
		So(node, ShouldEqual, node1)
	})

	Convey("TestHasCycle2", t, func() {
		node := detectCycle2(head)
		So(node, ShouldEqual, node1)
	})
}
