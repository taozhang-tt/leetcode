package mergetwosortedlists

import (
	"testing"

	. "github.com/taozhang-tt/leetcode/util"
	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func Test(t *testing.T) {
	Convey("Test mergeTwoLists", t, func() {
		list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
		list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
		list3 := mergeTwoLists(list1, list2)
		expect := []int{1, 1, 2, 3, 4, 4}
		actual := List2Slice(list3)
		So(actual, ShouldResemble, expect)
	})

	Convey("Test mergeTwoLists2", t, func() {
		list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
		list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
		list3 := mergeTwoLists2(list1, list2)
		expect := []int{1, 1, 2, 3, 4, 4}
		actual := List2Slice(list3)
		So(actual, ShouldResemble, expect)
	})
}
