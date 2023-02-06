package main

import (
	"fmt"
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func Test(t *testing.T) {
	Convey("Test Reverse", t, func() {
		head := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		}
		expect := reverse(reverse(head))
		So(head, ShouldEqual, expect)

		expect = reverseInRecursion(reverseInRecursion(head))
		So(head, ShouldEqual, expect)

	})
}

func print(head *ListNode) {
	for head != nil {
		fmt.Printf("%v, ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
