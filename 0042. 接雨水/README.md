#### [42. 接雨水](https://leetcode-cn.com/problems/trapping-rain-water/)

> 难度困难

给定 `n` 个非负整数表示每个宽度为 `1` 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

**示例 1：**

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/rainwatertrap.png)

```
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 
```

**示例 2：**

```
输入：height = [4,2,0,3,2,5]
输出：9
```

**提示：**

- `n == height.length`
- `1 <= n <= 2 * 104`
- `0 <= height[i] <= 105`

#### 解题思路

按列来计算，每一列雨水的高度，取决于该列左侧最高的柱子和右侧最高的柱子中最矮的那个柱子的高度。

> 暴力破解

首先我们想到的是暴力破解，对于每根柱子（下标从1到n-2）,我们只要知道这根柱子其左边最高柱子的高度 left 和右边最高柱子的高度 right 就可以知道这根柱子能接多少雨水了  `min(left, right) - height[i]` 时间复杂度主要是重复求每根柱子的左右最大柱子的高度。

> 备忘录优化，动态规划

暴力破解就是会重复计算柱子左右高度，如果我们使用空间换时间的思想，可以提前计算好每根柱子的左右最高柱子高度

> 双指针

方法二中我们使用了空间换时间的方法，使用了两个数组来保存每根柱子的左右最高柱子高度，我们可以使用双指针边走边算，来节省空间复杂度。我们使用变量 leftMax 和 rightMax来分别记录 height[0...left] 和 height[right..n-1]的最高柱子的高度，然后移动左右指针 left 和 right 来遍历整个数组

> 单调栈  栈内存放数组下标  正序

首先要想明白为什么可以用单调栈解决，这个题乍一看不是要找每根柱子的左右两侧最高的柱子高度，并不是求每根柱子左右邻近的柱子，其实仔细想想在我们遍历数组的时候只要出现凹槽就可以接住雨水，凹槽不就是每根柱子左右两个邻近的高柱子吗。

我们维护一个单调栈，使其栈顶元素到栈底元素递增，这样才能访问数组元素i时，当i比栈顶元素大时，我们知道栈顶元素的右边邻近高柱是i。在遍历过程中，如果数组元素 > 栈顶元素 `top` ,那么我们需要 `top` 下面的一个元素 `left` , 根据单调栈的性质我们知道 `height[left]` 一定是大于 `height[top]` 的，这不就是凹槽出现了吗，计算这一局部凹槽的面积。宽度是 `i-left-1` ，高度是 `min(height[left],height[i])-height[top]` 面积就是宽*高。为了计算 `left` 需要将 top 弹栈，在计算了 `top` 下标的柱子能接的雨水后，`left` 就变成了新的 `top`

当然如果不存在 left 即单调栈为空嘛，即左边没有高柱，构不成凹槽。

**思考：这次是单调栈正序解决问题，可以倒序吗？**如果倒叙的话，当前元素的右邻近高柱是栈顶元素，但不知道当前元素的左邻近高柱啊！

#### 代码

> 暴力破解

```go
func trap(height []int) int {
	length := len(height)
	res := 0
	for i := 1; i < length-1; i++ {
		left, right := 0, 0
		//不是找右边第一个大的元素，是找右边最大的元素
		for j := i; j < length; j++ {
			right = max(right, height[j])
		}
		//找左边最大的元素
		for j := i; j >= 0; j-- {
			left = max(left, height[j])
		}
		res += min(left, right) - height[i]
	}
	return res
}
func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
```

> 备忘录

```go
func trap(height []int) int {
	length := len(height)
	res := 0
	left := make([]int, length)
	right := make([]int, length)
	//初始化
	left[0] = height[0]
	right[length-1] = height[length-1]
	for i, v := range height[1:] {
		left[i+1] = max(left[i], v)
	}
	for i := length - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}
	for i, v := range height {
		res += min(left[i], right[i]) - v
	}
	return res
}
func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
```

> 双指针

```go
func trap2(height []int) int {
	res := 0
	length := len(height)
	left, right := 0, length-1
	leftMax := 0
	rightMax := 0
	for left <= right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if leftMax < rightMax {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}
	return res
}
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
```

> 单调栈

```go
func trap(height []int) int {
	singleStack := []int{}
	res := 0
	for i, v := range height {
		//如果当前元素大于栈顶元素
		for len(singleStack) > 0 && v > height[singleStack[len(singleStack)-1]] {
			top := singleStack[len(singleStack)-1]         //栈顶元素
			singleStack = singleStack[:len(singleStack)-1] //弹栈
			//如果栈为空，说明栈顶元素的左边没有邻近高柱，构不成凹槽
			if len(singleStack) == 0 {
				break
			}
			//构成凹槽
			left := singleStack[len(singleStack)-1]
			w := i - left - 1
			h := min(height[left], v) - height[top]
			res += w * h
		}
		//入栈
		singleStack = append(singleStack, i)
	}
	return res
}
func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
```

