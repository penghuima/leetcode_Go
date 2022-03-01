#### [209. 长度最小的子数组](https://leetcode-cn.com/problems/minimum-size-subarray-sum/)

> 难度中等

给定一个含有 `n` 个正整数的数组和一个正整数 `target` **。**

找出该数组中满足其和 `≥ target` 的长度最小的 **连续子数组** `[numsl, numsl+1, ..., numsr-1, numsr]` ，并返回其长度**。**如果不存在符合条件的子数组，返回 `0` 。

**示例 1：**

```
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
```

**示例 2：**

```
输入：target = 4, nums = [1,4,4]
输出：1
```

**示例 3：**

```
输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0
```

**提示：**

- `1 <= target <= 109`
- `1 <= nums.length <= 105`
- `1 <= nums[i] <= 105`

#### 解题思路

滑动窗口

定义两个指针 start 和 end 分别表示子数组（滑动窗口窗口）的开始位置和结束位置，维护变量 sum 存储子数组中的元素和（即从nums[start] 到 nums[end] 的元素和）。

初始状态下，start 和 end 都指向下标 0，sum 的值为 0。

每一轮迭代，将 nums[end] 加到sum，如果 sum≥target，则更新子数组的最小长度（此时子数组的长度是end−start+1），然后将 nums[start] 从 sum 中减去并将 start 右移，直到 sum<target，在此过程中同样更新子数组的最小长度。在每一轮迭代的最后，将 end 右移。

#### 代码

```go
func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	sum := 0
	ans := math.MaxInt64
	for right < len(nums) {
		//先将元素添加进来窗口     right指针右边移动
		sum += nums[right]
		right++
		//满足缩减窗口的条件   和>=target    left指针左边移动
		for sum >= target {
			ans = min(ans, right-left)
			sum -= nums[left]
			left++
		}
	}
	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
```

