package splitlinkedlistinparts

// 725. 分隔链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	// 统计节点总数
	root, cnt := head, 0
	for root != nil {
		cnt, root = cnt+1, root.Next
	}
	// 节点数分组，平均每组avg个，多出more个，将多出部分给前more组每组分配一个
	avg, more, ret := cnt/k, cnt%k, make([]*ListNode, k)
	for i := 0; i < k && head != nil; i++ {
		ret[i] = head
		step := avg
		if i < more {
			step++
		}
		for i := 0; i < step-1; i++ { // 将head移动到本组最后一个节点
			head = head.Next
		}
		head, head.Next = head.Next, nil
	}
	return ret
}

func print(list []*ListNode) [][]int {
	ret := make([][]int, len(list))
	for i, head := range list {
		ret[i] = make([]int, 0)
		for head != nil {
			ret[i], head = append(ret[i], head.Val), head.Next
		}
	}
	return ret
}
