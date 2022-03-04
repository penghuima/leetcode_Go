package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func closestValue(root *TreeNode, target float64) int {
	res := root.Val
	//需要一个函数比较左节点的差值和右节点的差值啊
	var traverse func(root *TreeNode, target float64)
	traverse = func(root *TreeNode, target float64) {
		if root == nil {
			return
		}
		// 一边搜索一边更新离 target 最近的值
		if math.Abs(float64(root.Val)-target) < math.Abs(float64(res)-target) {
			res = root.Val
		}
		if float64(root.Val) < target {
			traverse(root.Right, target)
		} else {
			traverse(root.Left, target)
		}
	}
	traverse(root, target)
	return res
}