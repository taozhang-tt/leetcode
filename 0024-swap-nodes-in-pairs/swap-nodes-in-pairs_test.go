package swapnodesinpairs

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func TestSwapPairs(t *testing.T) {
	head1 := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					nil,
				},
			},
		},
	}

	head2 := &ListNode{
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
	Convey("TestSwapPairs1", t, func() {
		curr := swapPairs(head1)
		curr = swapPairs(curr)
		So(curr, ShouldResemble, head1)

		curr = swapPairs(head2)
		curr = swapPairs(curr)
		So(curr, ShouldResemble, head2)
	})

	Convey("TestSwapPairs2", t, func() {
		curr := swapPairs2(head1)
		curr = swapPairs2(curr)
		So(curr, ShouldResemble, head1)

		curr = swapPairs2(head2)
		curr = swapPairs2(curr)
		So(curr, ShouldResemble, head2)
	})
}
