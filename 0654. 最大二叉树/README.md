#### [654. 最大二叉树](https://leetcode-cn.com/problems/maximum-binary-tree/)

> 难度中等

给定一个不含重复元素的整数数组 `nums` 。一个以此数组直接递归构建的 **最大二叉树** 定义如下：

1. 二叉树的根是数组 `nums` 中的最大元素。
2. 左子树是通过数组中 **最大值左边部分** 递归构造出的最大二叉树。
3. 右子树是通过数组中 **最大值右边部分** 递归构造出的最大二叉树。

返回有给定数组 `nums` 构建的 **最大二叉树** 。

**示例 1：**

![img](https://assets.leetcode.com/uploads/2020/12/24/tree1.jpg)

```
输入：nums = [3,2,1,6,0,5]
输出：[6,3,5,null,2,0,null,null,1]
解释：递归调用如下所示：
- [3,2,1,6,0,5] 中的最大值是 6 ，左边部分是 [3,2,1] ，右边部分是 [0,5] 。
    - [3,2,1] 中的最大值是 3 ，左边部分是 [] ，右边部分是 [2,1] 。
        - 空数组，无子节点。
        - [2,1] 中的最大值是 2 ，左边部分是 [] ，右边部分是 [1] 。
            - 空数组，无子节点。
            - 只有一个元素，所以子节点是一个值为 1 的节点。
    - [0,5] 中的最大值是 5 ，左边部分是 [0] ，右边部分是 [] 。
        - 只有一个元素，所以子节点是一个值为 0 的节点。
        - 空数组，无子节点。
```

**示例 2：**

![img](https://assets.leetcode.com/uploads/2020/12/24/tree2.jpg)

```
输入：nums = [3,2,1]
输出：[3,null,2,null,1]
```

**提示：**

- `1 <= nums.length <= 1000`
- `0 <= nums[i] <= 1000`
- `nums` 中的所有整数 **互不相同**

#### 解题思路

首先我们每次找到切片中的最大值和所在下标，值用其创建根节点，下标用其分割切片，来创建左子树和右子树

边界条件是如果子树的切片为空，则返回 `nil` 作为递归终止条件

#### 代码

```go
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
```

