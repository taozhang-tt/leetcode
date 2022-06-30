package mergetwosortedlists

// 21. 合并两个有序链表

import . "github.com/taozhang-tt/leetcode/util"

// 添加虚拟头结点，直接遍历
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	virtual := &ListNode{}
	prev := virtual
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}
	if list1 != nil {
		prev.Next = list1
	} else {
		prev.Next = list2
	}
	return virtual.Next
}

// 递归
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}
