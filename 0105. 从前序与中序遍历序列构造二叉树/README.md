#### [105. 从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

> 难度中等

给定一棵树的前序遍历 `preorder` 与中序遍历 `inorder`。请构造二叉树并返回其根节点。

**注意:**
你可以假设树中没有重复的元素。

**示例 1:**

![img](https://assets.leetcode.com/uploads/2021/02/19/tree.jpg)

```
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
```

**示例 2:**

```
Input: preorder = [-1], inorder = [-1]
Output: [-1]
```

**提示:**

- `1 <= preorder.length <= 3000`
- `inorder.length == preorder.length`
- `-3000 <= preorder[i], inorder[i] <= 3000`
- `preorder` 和 `inorder` 均无重复元素
- `inorder` 均出现在 `preorder`
- `preorder` 保证为二叉树的前序遍历序列
- `inorder` 保证为二叉树的中序遍历序列

#### 解题思路

对于任意一颗树而言，

前序遍历的形式总是

`[ 根节点, [左子树的前序遍历结果], [右子树的前序遍历结果] ]`

即根节点总是前序遍历中的第一个节点。

而中序遍历的形式总是

`[ [左子树的中序遍历结果], 根节点, [右子树的中序遍历结果] ]`

只要我们在中序遍历中定位到根节点，那么我们就可以分别知道左子树和右子树中的节点数目。由于同一颗子树的前序遍历和中序遍历的长度显然是相同的，因此我们就可以对应到前序遍历的结果中，对上述形式中的所有左右括号进行定位。

这样以来，我们就知道了左子树的前序遍历和中序遍历结果，以及右子树的前序遍历和中序遍历结果，我们就可以递归地对构造出左子树和右子树，再将这两颗子树接到根节点的左右位置。

**细节**

在中序遍历中对根节点进行定位时，一种简单的方法是直接扫描整个中序遍历的结果并找出根节点，但这样做的时间复杂度较高。我们可以考虑使用哈希表来帮助我们快速地定位根节点。对于哈希映射中的每个键值对，键表示一个元素（节点的值），值表示其在中序遍历中的出现位置。在构造二叉树的过程之前，我们可以对中序遍历的列表进行一遍扫描，就可以构造出这个哈希映射。在此后构造二叉树的过程中，我们就只需要 O(1) 的时间对根节点进行定位了。

#### 代码

```go
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
```

