package reverselinkedlist

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func TestReverseList(t *testing.T) {
	Convey("TestReverseList", t, func() {
		var head = &ListNode{
			1,
			&ListNode{
				2,
				&ListNode{
					3,
					nil,
				},
			},
		}

		Convey("TestReverseList1", func() {
			curr := reverseList(head)
			curr = reverseList(curr)
			So(head, ShouldResemble, curr)
		})

		Convey("TestReverseList2", func() {
			curr := reverseList2(head)
			curr = reverseList2(curr)
			So(head, ShouldResemble, curr)
		})
	})
}
