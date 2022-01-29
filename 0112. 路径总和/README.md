#### [112. 路径总和](https://leetcode-cn.com/problems/path-sum/)

> 难度简单

给你二叉树的根节点 `root` 和一个表示目标和的整数 `targetSum` 。判断该树中是否存在 **根节点到叶子节点** 的路径，这条路径上所有节点值相加等于目标和 `targetSum` 。如果存在，返回 `true` ；否则，返回 `false` 。

**叶子节点** 是指没有子节点的节点。

**示例 1：**

![img](https://assets.leetcode.com/uploads/2021/01/18/pathsum1.jpg)

```
输入：root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
输出：true
解释：等于目标和的根节点到叶节点路径如上图所示。
```

**示例 2：**

![img](https://assets.leetcode.com/uploads/2021/01/18/pathsum2.jpg)

```
输入：root = [1,2,3], targetSum = 5
输出：false
解释：树中存在两条根节点到叶子节点的路径：
(1 --> 2): 和为 3
(1 --> 3): 和为 4
不存在 sum = 5 的根节点到叶子节点的路径。
```

**示例 3：**

```
输入：root = [], targetSum = 0
输出：false
解释：由于树是空的，所以不存在根节点到叶子节点的路径。
```

**提示：**

- 树中节点的数目在范围 `[0, 5000]` 内
- `-1000 <= Node.val <= 1000`
- `-1000 <= targetSum <= 1000`

#### 解题思路

观察要求我们完成的函数，我们可以归纳出它的功能：询问是否存在从当前节点 root 到叶子节点的路径，满足其路径和为 sum。

假定从根节点到当前节点的值之和为 `val`，我们可以将这个大问题转化为一个小问题：是否存在从当前节点的子节点到叶子的路径，满足其路径和为 `sum - val`。

不难发现这满足递归的性质，若当前节点就是叶子节点，那么我们直接判断 sum 是否等于 val 即可（因为路径和已经确定，就是当前节点的值，我们只需要判断该路径和是否满足条件）。若当前节点不是叶子节点，我们只需要递归地询问它的子节点是否能满足条件即可。

#### 代码

> 回溯隐藏在参数 targetSum中

```go
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		if targetSum == root.Val {
			return true
		}
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
```

> 为了把回溯的过程体现出来，可以改为如下代码：

```go
func hasPathSum(root *TreeNode, targetSum int) bool {
	var backTrack func(root *TreeNode, targetSum int) bool
	backTrack = func(root *TreeNode, targetSum int) bool {
		if root == nil {
			return false
		}
		if root.Left == nil && root.Right == nil {
			if targetSum == root.Val {
				return true
			}
		}
		if root.Left != nil {
			targetSum -= root.Val
			if backTrack(root.Left, targetSum) {
				return true
			}
			targetSum += root.Val
		}
		if root.Right != nil {
			targetSum -= root.Val
			if backTrack(root.Right, targetSum) {
				return true
			}
			targetSum += root.Val
		}
		return false
	}
	return backTrack(root, targetSum)
}
```

