package removeduplicatesfromsortedlistii

// 82. 删除排序链表中的重复元素 II

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	virtual, sign := &ListNode{Next: head}, false // 添加一个虚拟头结点
	prev := virtual
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val { // head == head.Next
			head.Next = head.Next.Next // 删除head.Next
			sign = true                // 标记head应该被删除
		} else { // head != head.Next
			if sign { // 当前head应该被删除
				prev.Next, head, sign = head.Next, head.Next, false // 删除head
			} else { // 当前head不用删除
				prev, head = prev.Next, head.Next
			}
		}
	}
	if sign { // 删除最后一个节点
		prev.Next = head.Next
	}
	return virtual.Next
}

func list2Slice(head *ListNode) []int {
	ret := make([]int, 0)
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}
	return ret
}

func delete(head *ListNode) *ListNode {
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
