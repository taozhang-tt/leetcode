package main

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func TestReverse(t *testing.T) {
	Convey("TestReverse", t, func() {
		var testCases = []*ListNode{
			{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						nil,
					},
				},
			},
			nil,
		}

		for _, head := range testCases {
			curr := reverse(head)
			curr = reverse(curr)
			So(head, ShouldResemble, curr)
		}
	})
}

func TestReverseK(t *testing.T) {
	Convey("TestReverseK", t, func() {
		var testCases = []*ListNode{
			{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						nil,
					},
				},
			},
		}

		for _, head := range testCases {
			curr := reverseK(head, 2)
			curr = reverseK(curr, 2)
			So(head, ShouldResemble, curr)

			curr = reverseK(head, 3)
			curr = reverseK(curr, 3)
			So(head, ShouldResemble, curr)
		}
	})
}

func TestReverseBetween(t *testing.T) {
	head := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					&ListNode{
						5,
						nil,
					},
				},
			},
		},
	}

	Convey("reverseBetween", t, func() {
		r1 := reverseBetween(head, 2, 4)
		r2 := reverseBetween(r1, 2, 4)
		So(r2, ShouldResemble, head)

		r1 = reverseBetween(head, 1, 1)
		r2 = reverseBetween(r1, 1, 1)
		So(r2, ShouldResemble, head)

		r1 = reverseBetween(head, 1, 5)
		r2 = reverseBetween(r1, 1, 5)
		So(r2, ShouldResemble, head)
	})
}
