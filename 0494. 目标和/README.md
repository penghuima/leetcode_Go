#### [494. 目标和](https://leetcode-cn.com/problems/target-sum/)

> 难度中等

给你一个整数数组 `nums` 和一个整数 `target` 。

向数组中的每个整数前添加 `'+'` 或 `'-'` ，然后串联起所有整数，可以构造一个 **表达式** ：

- 例如，`nums = [2, 1]` ，可以在 `2` 之前添加 `'+'` ，在 `1` 之前添加 `'-'` ，然后串联起来得到表达式 `"+2-1"` 。

返回可以通过上述方法构造的、运算结果等于 `target` 的不同 **表达式** 的数目。

**示例 1：**

```
输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
```

**示例 2：**

```
输入：nums = [1], target = 1
输出：1
```

**提示：**

- `1 <= nums.length <= 20`
- `0 <= nums[i] <= 1000`
- `0 <= sum(nums[i]) <= 1000`
- `-1000 <= target <= 1000`

#### 解题思路

两种思路，一种是回溯法，一种是动态递归

> 回溯

关键就是搞清楚什么是「选择」，而对于这道题，「选择」不是明摆着的吗？**对于每个数字 `nums[i]`，我们可以选择给一个正号 `+` 或者一个负号 `-`**，然后利用回溯模板穷举出来所有可能的结果，数一数到底有几种组合能够凑出 `target` 不就行了嘛？==（其实观察选择就是看每个父节点的子节点都有什么选项）==

#### 代码

> 回溯

```go
func findTargetSumWays(nums []int, target int) int {
	res, sum := 0, 0
	lenth := len(nums)
	var backTrack func(index int)
	backTrack = func(index int) {
		if index == lenth {
			if sum == target {
				res++
			}
			return
		}
		//先对每个元素执行 + 操作
		sum += nums[index]
		backTrack(index + 1)
		sum -= nums[index]
		//再对每个元素执行 — 操作
		sum -= nums[index]
		backTrack(index + 1)
		sum += nums[index]
	}
	backTrack(0)
	return res
}
```

