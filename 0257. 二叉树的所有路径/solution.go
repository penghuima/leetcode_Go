package leetcode

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	var path []int
	var backTrack func(root *TreeNode)
	backTrack = func(root *TreeNode) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil {
			//将路径转化为题目要求格式的样子 1—>2
			answer := ""
			for i, v := range path {
				if i != len(path)-1 { //不是最后一个元素
					answer += strconv.Itoa(v) + "->"
				} else { //最后一个元素
					answer += strconv.Itoa(v)
				}
			}
			res = append(res, answer)
		}
		backTrack(root.Left)
		backTrack(root.Right)
		path = path[:len(path)-1]
	}
	backTrack(root)
	return res
}
