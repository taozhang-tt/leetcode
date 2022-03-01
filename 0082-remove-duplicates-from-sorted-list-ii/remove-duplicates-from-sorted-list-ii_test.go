package removeduplicatesfromsortedlistii

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func Test(t *testing.T) {
	Convey("Test list2Slice", t, func() {
		head := &ListNode{1, &ListNode{2, nil}}
		So(list2Slice(head), ShouldResemble, []int{1, 2})
	})
	Convey("Test deleteDuplicates", t, func() {
		s := []int{2}
		l := &ListNode{1, &ListNode{1, &ListNode{2, nil}}}
		l = deleteDuplicates(l)
		So(list2Slice(l), ShouldResemble, s)

		s = []int{2}
		l = &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}
		l = deleteDuplicates(l)
		So(list2Slice(l), ShouldResemble, s)

		s = []int{1}
		l = &ListNode{1, nil}
		l = deleteDuplicates(l)
		So(list2Slice(l), ShouldResemble, s)

		s = []int{}
		l = nil
		l = deleteDuplicates(l)
		So(list2Slice(l), ShouldResemble, s)
	})
}
