package util

type ListNode struct {
	Val  int
	Next *ListNode
}

func List2Slice(head *ListNode) []int {
	ret := make([]int, 0)
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}
	return ret
}
