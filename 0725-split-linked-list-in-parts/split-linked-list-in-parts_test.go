package splitlinkedlistinparts

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func Test(t *testing.T) {
	Convey("Test print", t, func() {
		list := []*ListNode{
			{1, &ListNode{2, nil}},
			{1, nil},
		}
		ret := print(list)
		So(ret, ShouldResemble, [][]int{
			{1, 2},
			{1},
		})
	})

	Convey("Test splitListToParts", t, func() {
		head := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
		list := splitListToParts(head, 5)
		ret := print(list)
		So(ret, ShouldResemble, [][]int{
			{1},
			{2},
			{3},
			{},
			{},
		})

		head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, &ListNode{7, &ListNode{8, &ListNode{9, &ListNode{10, nil}}}}}}}}}}
		list = splitListToParts(head, 3)
		ret = print(list)
		So(ret, ShouldResemble, [][]int{
			{1, 2, 3, 4},
			{5, 6, 7},
			{8, 9, 10},
		})
	})

}
