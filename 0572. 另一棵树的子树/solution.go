package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	var isSameTree func(p *TreeNode, q *TreeNode) bool
	isSameTree = func(p *TreeNode, q *TreeNode) bool {
		if p == nil || q == nil {
			return p == q
		}
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return isSameTree(root, subRoot) || isSameTree(root.Left, subRoot) || isSameTree(root.Right, subRoot)
}
