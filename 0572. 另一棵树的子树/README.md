#### [572. 另一棵树的子树](https://leetcode-cn.com/problems/subtree-of-another-tree/)[思路](https://leetcode-cn.com/problems/subtree-of-another-tree/#)

> 难度简单

给你两棵二叉树 `root` 和 `subRoot` 。检验 `root` 中是否包含和 `subRoot` 具有相同结构和节点值的子树。如果存在，返回 `true` ；否则，返回 `false` 。

二叉树 `tree` 的一棵子树包括 `tree` 的某个节点和这个节点的所有后代节点。`tree` 也可以看做它自身的一棵子树。

**示例 1：**

![img](https://assets.leetcode.com/uploads/2021/04/28/subtree1-tree.jpg)

```
输入：root = [3,4,5,1,2], subRoot = [4,1,2]
输出：true
```

**示例 2：**

![img](https://assets.leetcode.com/uploads/2021/04/28/subtree2-tree.jpg)

```
输入：root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
输出：false
```

**提示：**

- `root` 树上的节点数量范围是 `[1, 2000]`
- `subRoot` 树上的节点数量范围是 `[1, 1000]`
- `-104 <= root.val <= 104`
- `-104 <= subRoot.val <= 104`

#### 解题思路

深度优先搜索暴力匹配

这是一种最朴素的方法——深度优先搜索枚举 s中的每一个节点，判断这个点的子树是否和 t 相等。如何判断一个节点的子树是否和 t 相等呢，我们又需要做一次深度优先搜索来检查，即让两个指针一开始先指向该节点和 t 的根，然后「同步移动」两根指针来「同步遍历」这两棵树，判断对应位置是否相等。

#### 代码

```go
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
    // 判断以 root 为根的二叉树是否和 subRoot 相同
    if (isSameTree(root, subRoot)) {
        return true;
    }
    // 去左右子树中判断是否有和 subRoot 相同的子树
    return isSubtree(root.left, subRoot) || isSubtree(root.right, subRoot);
}
```

