package main

// [1233. 删除子文件夹](https://leetcode.cn/problems/remove-sub-folders-from-the-filesystem/)
// 字典树、双指针

import (
	"sort"
	"strings"
)

// 排序+双指针
func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	ret := []string{folder[0]}
	if len(folder) == 1 {
		return ret
	}
	for l, r := 0, 1; r < len(folder); {
		if strings.HasPrefix(folder[r], folder[l]) && folder[r][len(folder[l])] == '/' {
			r++
			continue
		}
		ret = append(ret, folder[r])
		l = r
		r++
	}

	return ret
}
