package deletethemiddlenodeofalinkedlist

import (
	"testing"

	. "github.com/taozhang-tt/spadger/pkg/testing/xconvey"
)

func Test(t *testing.T) {
	Convey("Test deleteMiddle", t, func() {
		head := &ListNode{1, nil}
		head = deleteMiddle(head)
		s := list2Slice(head)
		ret := []int{}
		So(s, ShouldResemble, ret)

		head = &ListNode{1, &ListNode{2, nil}}
		head = deleteMiddle(head)
		s = list2Slice(head)
		ret = []int{1}
		So(s, ShouldResemble, ret)
	})

}
