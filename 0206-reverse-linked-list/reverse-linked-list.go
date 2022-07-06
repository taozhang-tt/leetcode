package reverselinkedlist

// 206. 翻转链表

type ListNode struct {
	Val  int
	Next *ListNode
}

// 直接翻转
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	final := reverseList2(head.Next)
	head.Next.Next, head.Next = head, nil
	return final
}
