package removeduplicatesfromsorted

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func TestDeleteDuplicates(t *testing.T) {
	Convey("Test List2Slice", t, func() {
		head := &ListNode{1, &ListNode{2, nil}}
		So(List2Slice(head), ShouldResemble, []int{1, 2})
	})

	Convey("Test deleteDuplicates", t, func() {
		s := []int{1, 2}
		l := &ListNode{1, &ListNode{1, &ListNode{2, nil}}}
		l = deleteDuplicates(l)
		So(List2Slice(l), ShouldResemble, s)

		s = []int{1, 2, 3}
		l = &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}
		l = deleteDuplicates(l)
		So(List2Slice(l), ShouldResemble, s)

		s = []int{1}
		l = &ListNode{1, nil}
		l = deleteDuplicates(l)
		So(List2Slice(l), ShouldResemble, s)

		s = []int{}
		l = nil
		l = deleteDuplicates(l)
		So(List2Slice(l), ShouldResemble, s)
	})

}
