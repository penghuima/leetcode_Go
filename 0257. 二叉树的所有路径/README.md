#### [257. 二叉树的所有路径](https://leetcode-cn.com/problems/binary-tree-paths/)

> 难度简单

给你一个二叉树的根节点 `root` ，按 **任意顺序** ，返回所有从根节点到叶子节点的路径。

**叶子节点** 是指没有子节点的节点。

**示例 1：**

![img](https://assets.leetcode.com/uploads/2021/03/12/paths-tree.jpg)

```
输入：root = [1,2,3,null,5]
输出：["1->2->5","1->3"]
```

**示例 2：**

```
输入：root = [1]
输出：["1"]
```

**提示：**

- 树中节点的数目在范围 `[1, 100]` 内
- `-100 <= Node.val <= 100`

#### 解题思路

最直观的方法是使用深度优先搜索。在深度优先搜索遍历二叉树时，我们需要考虑当前的节点以及它的孩子节点。

- 如果当前节点不是叶子节点，则在当前的路径末尾添加该节点，并继续递归遍历该节点的每一个孩子节点。
- 如果当前节点是叶子节点，则在当前路径末尾添加该节点后我们就得到了一条从根节点到叶子节点的路径，将该路径加入到答案即可。

如此，当遍历完整棵二叉树以后我们就得到了所有从根节点到叶子节点的路径。当然，深度优先搜索也可以使用非递归的方式实现，这里不再赘述。

#### 代码

> 运行时间太长，耗时击败5%，内存消耗击败 7%

```go
func binaryTreePaths(root *TreeNode) []string {
	var res []string
	var path []int
	var backTrack func(root *TreeNode)
	backTrack = func(root *TreeNode) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		if root.Left == nil && root.Right == nil {
			//将路径转化为题目要求格式的样子 1—>2
			temp:=append([]int(nil),path...)  //担心后来对path切片值的修改会影响结果
			answer := ""
			for i, v := range temp{
				if i != len(temp)-1 { //不是最后一个元素
					answer += strconv.Itoa(v) + "->"
				} else { //最后一个元素
					answer += strconv.Itoa(v)
				}
			}
			res = append(res, answer)
		}
		backTrack(root.Left)
		backTrack(root.Right)
		path = path[:len(path)-1]
	}
	backTrack(root)
	return res
}
```

> 100%，95%
>
> **貌似没有看到回溯的逻辑，其实不然，回溯就隐藏在参数 s 上了**

```go
func binaryTreePaths(root *TreeNode) []string {
    res := make([]string, 0)
    var travel func(node *TreeNode, s string)
    travel = func(node *TreeNode, s string) {
        if node.Left == nil && node.Right == nil {
            v := s + strconv.Itoa(node.Val)
            res = append(res, v)
            return
        }
        s = s + strconv.Itoa(node.Val) + "->"
        if node.Left != nil {
            travel(node.Left, s)
        }
        if node.Right != nil {
            travel(node.Right, s)
        }
    }
    travel(root, "")
    return res
}
```

