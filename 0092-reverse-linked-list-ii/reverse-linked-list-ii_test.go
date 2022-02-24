package reverselinkedlistii

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func TestReverseK(t *testing.T) {
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
	Convey("reverseK", t, func() {
		So(reverseK(reverseK(head, 2), 2), ShouldResemble, head)
	})

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
