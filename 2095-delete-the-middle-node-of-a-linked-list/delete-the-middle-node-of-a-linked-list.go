package deletethemiddlenodeofalinkedlist

// 2095. 删除链表的中间节点

type ListNode struct {
	Val  int
	Next *ListNode
}

// 删除链表的中间节点，如果中间节点有两个，删除后一个
// 等价于查找链表的中间节点的前一个节点，如果中间节点有两个，以后一个为准
// 进一步等价于查找链表 prev->head 链表的中间节点，如果有两个，返回前一个
func deleteMiddle(head *ListNode) *ListNode {
	prev := &ListNode{Next: &ListNode{Next: head}}
	fast, slow := prev, prev
	for fast != nil {
		if fast.Next == nil {
			break
		}
		fast, slow = fast.Next.Next, slow.Next
	}
	slow.Next = slow.Next.Next
	return prev.Next.Next
}

func list2Slice(head *ListNode) []int {
	ret := make([]int, 0)
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}
	return ret
}
