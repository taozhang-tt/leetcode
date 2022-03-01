package removeduplicatesfromsorted

// 83. 删除排序链表中的重复元素

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev, curr := head, head.Next
	for curr != nil {
		if curr.Val == prev.Val {
			prev.Next, curr = curr.Next, curr.Next
		} else {
			prev, curr = curr, curr.Next
		}
	}
	return head
}

func List2Slice(head *ListNode) []int {
	ret := make([]int, 0)
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}
	return ret
}
