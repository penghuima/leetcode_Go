package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	Q := []*TreeNode{root}
	for {
		if len(Q) == 0 {
			break
		}
		size := len(Q) //记录上一层 节点的个数
		var path []int
		for j := 0; j < size; j++ {
			path = append(path, Q[j].Val) //就是不知道的改变会不会影响path的值
			if Q[j].Left != nil {
				//入队
				Q = append(Q, Q[j].Left)
			}
			if Q[j].Right != nil {
				//入队
				Q = append(Q, Q[j].Right)
			}
		}
		res = append(res, path)
		//出队
		Q = Q[size:]
	}
	return res
}
