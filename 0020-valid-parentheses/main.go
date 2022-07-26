package main

// 20. 有效的括号
// https://leetcode.cn/problems/valid-parentheses/
/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
*/

func isValid(s string) bool {
	set := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if right, ok := set[s[i]]; ok {
			if len(stack) == 0 || right != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func isValid2(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) == 1 {
		return false
	}
	m := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack, idx := make([]byte, len(s)), -1
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			idx++
			stack[idx] = s[i]
			continue
		}
		if idx == -1 {
			return false
		}
		if m[stack[idx]] == s[i] {
			idx--
			continue
		}
		return false
	}
	if idx == -1 {
		return true
	}
	return false
}
