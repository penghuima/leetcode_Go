package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func findMode(root *TreeNode) []int {
	var res []int
	hashMap := make(map[int]int, 1)
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		hashMap[root.Val]++
		inorder(root.Right)
	}
	inorder(root) //遍历结束
	//maxNum
	//如何将频率最高的元素全部打印出来呢，一次遍历只能找到最大值，但之前重复的呢
	//两次遍历
	maxNum := -1
	index := -1
	for k, v := range hashMap {
		if v > maxNum {
			maxNum = v
			index = k
		}
	}
	for _, v := range hashMap {
		if v == hashMap[index] {
			res = append(res, v)
		}
	}
	return res
}
