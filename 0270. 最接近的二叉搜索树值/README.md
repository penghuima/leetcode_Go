#### [270. 最接近的二叉搜索树值](https://leetcode-cn.com/problems/closest-binary-search-tree-value/)

> 难度简单

给定一个不为空的二叉搜索树和一个目标值 target，请在该二叉搜索树中找到最接近目标值 target 的数值。

**注意：**

- 给定的目标值 target 是一个浮点数
- 题目保证在该二叉搜索树中只会存在一个最接近目标值的数

**示例：**

```
输入: root = [4,2,5,1,3]，目标值 target = 3.714286

    4
   / \
  2   5
 / \
1   3

输出: 4
```

#### 解题思路

二分搜索，边搜索边记录最接近的值

#### 代码

```go
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
```

