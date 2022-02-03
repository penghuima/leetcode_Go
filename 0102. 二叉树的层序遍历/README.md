#### [102. 二叉树的层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)

> 难度中等

给你二叉树的根节点 `root` ，返回其节点值的 **层序遍历** 。 （即逐层地，从左到右访问所有节点）。

**示例 1：**

![img](https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg)

```
输入：root = [3,9,20,null,null,15,7]
输出：[[3],[9,20],[15,7]]
```

**示例 2：**

```
输入：root = [1]
输出：[[1]]
```

**示例 3：**

```
输入：root = []
输出：[]
```

**提示：**

- 树中节点数目在范围 `[0, 2000]` 内
- `-1000 <= Node.val <= 1000`

#### 解题思路

我们可以用广度优先搜索解决这个问题。

我们可以用一种巧妙的方法修改广度优先搜索：

- 首先根元素入队
- 当队列不为空的时候
  - 求当前队列的长度 `si`
  - 依次从队列中取 `si` 个元素进行拓展，然后进入下一次迭代

**它和普通广度优先搜索的区别在于，普通广度优先搜索每次只取一个元素拓展，而这里每次取`si`个元素。**

第 i 次迭代前，队列中的所有元素就是第 i 层的所有元素，并且按照从左向右的顺序排列

#### 代码

> 第一版

```go
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	Q := []*TreeNode{root}
	for ; len(Q) != 0; {
		p := make([]*TreeNode, len(Q))
		copy(p, Q)
		Q = Q[:0]
		var path []int
		for j := 0; j < len(p); j++ {
			path = append(path, p[j].Val)
			if p[j].Left != nil {
				Q = append(Q, p[j].Left)
			}
			if p[j].Right != nil {
				Q = append(Q, p[j].Right)
			}
		}
		res = append(res, path)
	}
	return res
}
```

> 第二版

```go
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
```

