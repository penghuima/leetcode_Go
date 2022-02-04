#### [501. 二叉搜索树中的众数](https://leetcode-cn.com/problems/find-mode-in-binary-search-tree/)

> 难度简单

给定一个有相同值的二叉搜索树（BST），找出 BST 中的所有众数（出现频率最高的元素）。

假定 BST 有如下定义：

- 结点左子树中所含结点的值小于等于当前结点的值
- 结点右子树中所含结点的值大于等于当前结点的值
- 左子树和右子树都是二叉搜索树

例如：
给定 BST `[1,null,2,2]`,

```
   1
    \
     2
    /
   2
```

`返回[2]`.

**提示**：如果众数超过1个，不需考虑输出顺序

**进阶：**你可以不使用额外的空间吗？（假设由递归产生的隐式调用栈的开销不被计算在内）

#### 解题思路

就是中序遍历处理嘛，我们可以得到中序遍历数组。

难点：如何求最大频率的元素集合（注意是集合，不是一个元素，可以有多个众数）

- 方法1：遍历数组2次

先遍历一遍数组，找出最大频率（maxCount），然后再重新遍历一遍数组把出现频率为maxCount的元素放进集合。（因为众数有多个）

- 方法2：其实只需要遍历一次就可以找到所有的众数

如果频率 count 等于 maxCount（最大频率），当然要把这个元素加入到结果中。是不是感觉这里有问题，result怎么能轻易就把元素放进去了呢，万一，这个maxCount此时还不是真正最大频率呢。

所以下面要做如下操作

频率 count 大于 maxCount 的时候，不仅要更新 maxCount，**而且要清空结果集（以下代码为result数组）**，因为结果集之前的元素都失效了。

#### 代码

> 对于一般的二叉树，不考虑二叉搜索树的特点

```go
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
	//如何将频率最高的元素全部打印出来呢，一次遍历只能找到最大值，但重复的呢
	//直接暴力，两次遍历将重复的众数打印出来
	maxNum := -1
	index := -1
	for k, v := range hashMap {
		if v > maxNum {
			maxNum = v
			index = k
		}
	}
	for k, v := range hashMap {
		if v == hashMap[index] {
			res = append(res, k)
		}
	}
	return res
}
```

> 利用搜索二叉树的特性

```go
func findMode(root *TreeNode) []int {
	var res []int
	count := 1
	max := -1
	var pre *TreeNode
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		//中序遍历处理
		if pre != nil && pre.Val == root.Val {
			count++
		} else {
			//重新开始计数
			count = 1
		}
		if count >= max {
			if count > max {
				res = res[:0] //重新将结果切片清 0
				res = append(res, root.Val)
                max=count
			} else {
				res = append(res, root.Val)
			}
		}
		pre = root
		inorder(root.Right)
	}
	inorder(root) //遍历结束
	return res
}
```

