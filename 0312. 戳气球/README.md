#### [312. 戳气球](https://leetcode-cn.com/problems/burst-balloons/)

> 难度困难

有 `n` 个气球，编号为`0` 到 `n - 1`，每个气球上都标有一个数字，这些数字存在数组 `nums` 中。

现在要求你戳破所有的气球。戳破第 `i` 个气球，你可以获得 `nums[i - 1] * nums[i] * nums[i + 1]` 枚硬币。 这里的 `i - 1` 和 `i + 1` 代表和 `i` 相邻的两个气球的序号。如果 `i - 1`或 `i + 1` 超出了数组的边界，那么就当它是一个数字为 `1` 的气球。

求所能获得硬币的最大数量。

**示例 1：**

```
输入：nums = [3,1,5,8]
输出：167
解释：
nums = [3,1,5,8] --> [3,5,8] --> [3,8] --> [8] --> []
coins =  3*1*5    +   3*5*8   +  1*3*8  + 1*8*1 = 167
```

**示例 2：**

```
输入：nums = [1,5]
输出：10
```

**提示：**

- `n == nums.length`
- `1 <= n <= 500`
- `0 <= nums[i] <= 100`

#### 解题思路

方法1：暴力回溯

我们需要穷举戳气球的顺序，就很类似于 “全排列” 问题。每次从选择列表里选择戳哪个气球（做选择），然后从选择列表中删除该选择，直到选择列表为空。回溯的时候，将选择再添加到选择列表。

方法2：动态规划

在使用动态规划方法的时候主要有以下难点

**难点1**

- 数组dp 如何定义，才能避免元素 nums[i] 和 nums[i-1]，nums[i+1] 关联，基本上这一步就卡死了

为了简化问题，我们将数组头尾两边各添加一个虚拟气球，数值为1。那么问题就转化为：**给你一组气球，请你戳破气球`0`和气球`n+1`之间的所有气球（不包括`0`和`n+1`），使得最终只剩下气球`0`和气球`n+1`两个气球，最多能够得到多少分？**

**解决难点1**

这激发我们定义数组 `dp[i][j]` 表示**戳破气球`i`和气球`j`之间（开区间，不包括`i`和`j`）的所有气球，获得的最高分数为`dp[i][j]`**。

**难点2**

- 动态递归方程式如何写

在上面我们已经定义了dp数组的意义。但如果我们正向思考，先戳破哪个气球的话，还是会造成元素 nums[i] 和 nums[i-1]，nums[i+1] 关联。但如果我们反向思考，假设前面如何戳破气球的顺序不管，我们只在意 `（i，j)` 区间内，最后被戳破的气球是 `k` ,那么戳破该气球可以获得的分数就是 `nums[i]*nums[k]*nums[j] `从而实现元素相邻之间的依赖问题。我们只需要遍历 （i，j）之间符合的k,然后取最大值就好了

**解决难点2**

```go
for k := i + 1; k < j; k++ {
    dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+nums[i]*nums[k]*nums[j])
}
```

**难点3**

- 如何初始化dp 数组

因为 `dp[i][j]` 是开区间，所以`dp[i][j] = 0`，其中`0 <= i <= n+1, j <= i+1`，因为这种情况下，开区间`(i, j)`中间根本没有气球可以戳。

**难点4**

- 如何遍历，即遍历顺序

通过我们观察动态递归方程式，我们可以确定遍历顺序是 从下到上，然后从左到右

#### 代码

> 回溯暴力破解  不能全部ac 超时

```go
func maxCoins(nums []int) int {
	res := 0
	sum := 0
	var backTrack func()
	backTrack = func() {
		//如果满足结束条件
		if len(nums) == 0 {
			res = max(res, sum) //记录最大值
			return
		}
		//回溯
		for i, v := range nums {
			//做选择
			temp := 0
			//计算乘积  nums[i-1] * v * nums[i+1] 考虑边界
			if i-1 < 0 {
				if i+1 > len(nums)-1 {
					temp = v
				} else {
					temp = v * nums[i+1]
				}
			} else if i+1 > len(nums)-1 {
				temp = nums[i-1] * v
			} else {
				temp = nums[i-1] * v * nums[i+1]
			}
			sum += temp
			//删除切片中的v
			nums = append(nums[:i], nums[i+1:]...)
			backTrack()
			//撤销选择
			sum -= temp
			//将v再加回到切片指定位置
			nums = append(nums[:i], append([]int{v}, nums[i:]...)...)
		}
	}
	backTrack()
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
```

> 动态规划

```go
//戳气球 动态规划
func maxCoins(nums []int) int {
	//将数组头尾添加元素1
	nums = append([]int{1}, nums[:]...)
	nums = append(nums, 1)
	dp := make([][]int, len(nums))
	for i, _ := range dp {
		dp[i] = make([]int, len(nums))
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max(dp[i][j], dp[i][k]+dp[k][j]+nums[i]*nums[k]*nums[j])
			}
		}
	}
	return dp[0][len(nums)-1]
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
```

