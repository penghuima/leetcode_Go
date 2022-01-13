#### [113. 路径总和 II](https://leetcode-cn.com/problems/path-sum-ii/)[思路](https://leetcode-cn.com/problems/path-sum-ii/#)

> 难度中等
>

给你二叉树的根节点 `root` 和一个整数目标和 `targetSum` ，找出所有 ==**从根节点到叶子节点**==路径总和等于给定目标和的路径。

**叶子节点** 是指没有子节点的节点。

**示例 1：**

![img](https://assets.leetcode.com/uploads/2021/01/18/pathsumii1.jpg)

```
输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：[[5,4,11,2],[5,8,4,5]]
```

**示例 2：**

![img](https://assets.leetcode.com/uploads/2021/01/18/pathsum2.jpg)

```
输入：root = [1,2,3], targetSum = 5
输出：[]
```

**示例 3：**

```
输入：root = [1,2], targetSum = 0
输出：[]
```

**提示：**

- 树中节点总数在范围 `[0, 5000]` 内
- `-1000 <= Node.val <= 1000`
- `-1000 <= targetSum <= 1000`

#### 解题思路

还是回溯，注意一下题目的要求是**到叶子节点的路径和**，而且回溯的时候注意和的加减即代码26和31行，是减去左右子节点的值，而不是当前节点的值

#### 代码

```go
/*
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
*/
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
```

