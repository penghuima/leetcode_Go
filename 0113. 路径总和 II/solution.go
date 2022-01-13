package leetcode


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func pathSum(root *TreeNode, targetSum int) [][]int {
	sum := 0
	var res [][]int
	var path []int
	var backTrack func(node *TreeNode)
	backTrack = func(node *TreeNode) {
		//满足条件
		if node == nil {
			return
		}
		path = append(path, node.Val)
		sum += node.Val
		if sum == targetSum && node.Left == nil && node.Right == nil {
			res = append(res, append([]int(nil), path...))
		}
		//单层遍历
		if node.Left != nil {
			backTrack(node.Left)
			path = path[:len(path)-1]
			sum -= node.Left.Val
		}
		if node.Right != nil {
			backTrack(node.Right)
			path = path[:len(path)-1]
			sum -= node.Right.Val
		}

	}
	backTrack(root)
	return res
}