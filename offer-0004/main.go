package main

// [剑指 Offer 04. 二维数组中的查找](https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/?envType=study-plan&id=lcof&plan=lcof&plan_progress=bkioxdt)
// tip: 树，搜索，二叉搜索树

// 暴力
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

// 二叉搜索树
// 将二位矩阵逆时针旋转45度，像一棵二叉搜索树，根结点是二位矩阵的右上角节点
// i 代表行坐标，j 代表列坐标
// 从根结点开始查找，当前节点 matrix[i][j] < target 则向右节点查找 i+1; matrix[i][j] > target 向左节点查找 j-1
func findNumberIn2DArray2(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			i += 1
		} else {
			j -= 1
		}
	}
	return false
}
