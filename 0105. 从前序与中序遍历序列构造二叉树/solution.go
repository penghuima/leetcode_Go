package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var build func(pre []int, in []int, preStart, preEnd, inStart, inEnd int) *TreeNode
	//划分 左子树，右子树
	build = func(pre []int, in []int, preStart, preEnd, inStart, inEnd int) *TreeNode {
		if preStart>preEnd{
			return nil
		}
		rootVal := pre[preStart]
		//在中序遍历列表中的根节点的位置
		index := 0
		for i, v := range inorder {
			if v == rootVal {
				index = i
				break
			}
		}
		root := TreeNode{rootVal, nil, nil}
		root.Left = build(pre, in, preStart+1, preStart+index-inStart, inStart, index-1)
		root.Right = build(pre, in, preStart+index-inStart+1, preEnd, index+1, inEnd)
		return &root
	}
	return build(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}
