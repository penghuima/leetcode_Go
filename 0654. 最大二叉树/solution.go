package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var maxBuild func(nums []int) *TreeNode
	maxBuild = func(nums []int) *TreeNode {
		//边界条件
		if len(nums) == 0 {
			return nil
		}
		index := -1
		maxNum := -1
		//寻找最大值
		for i, v := range nums {
			if v > maxNum {
				maxNum = v
				index = i
			}
		}
		root := &TreeNode{
			maxNum,
			maxBuild(nums[:index]),
			maxBuild(nums[index+1:]),
		}
		return root
	}
	return maxBuild(nums)
}
