package removenthnodefromendoflist

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func Test(t *testing.T) {
	Convey("Test removeNthFromEnd", t, func() {
		expect := []int{1}
		l := &ListNode{1, &ListNode{2, nil}}
		s := list2Slice(removeNthFromEnd(l, 1))
		So(s, ShouldResemble, expect)

		expect = []int{2}
		l = &ListNode{1, &ListNode{2, nil}}
		s = list2Slice(removeNthFromEnd(l, 2))
		So(s, ShouldResemble, expect)

		expect = []int{}
		l = &ListNode{1, nil}
		s = list2Slice(removeNthFromEnd(l, 1))
		So(s, ShouldResemble, expect)
	})
}
