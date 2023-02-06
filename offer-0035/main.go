package main

// [剑指 Offer 35. 复杂链表的复制](https://leetcode.cn/problems/fu-za-lian-biao-de-fu-zhi-lcof/)
// tag: 哈希表、链表
// tip: 回溯、哈希表

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 回溯+哈希
func copyRandomListByRecursion(head *Node) *Node {
	cachedNode := make(map[*Node]*Node)
	var deepCopy func(*Node) *Node

	deepCopy = func(n *Node) *Node {
		if n == nil {
			return nil
		}
		if node, ok := cachedNode[n]; ok {
			return node
		}
		node := &Node{Val: n.Val}
		cachedNode[n] = node
		node.Next = deepCopy(n.Next)
		node.Random = deepCopy(n.Random)
		return node
	}

	return deepCopy(head)
}

// 哈希、直接遍历
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	m := make(map[*Node]*Node)

	p := head
	for p != nil {
		m[p], p = &Node{Val: p.Val}, p.Next
	}

	p = head
	for p != nil {
		m[p].Next, m[p].Random, p = m[p.Next], m[p.Random], p.Next
	}

	return m[head]
}
