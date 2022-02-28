#### [11. 盛最多水的容器](https://leetcode-cn.com/problems/container-with-most-water/)

> 难度中等

给定一个长度为 `n` 的整数数组 `height` 。有 `n` 条垂线，第 `i` 条线的两个端点是 `(i, 0)` 和 `(i, height[i])` 。

找出其中的两条线，使得它们与 `x` 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

**说明：**你不能倾斜容器。

**示例 1：**

![img](https://aliyun-lc-upload.oss-cn-hangzhou.aliyuncs.com/aliyun-lc-upload/uploads/2018/07/25/question_11.jpg)

```
输入：[1,8,6,2,5,4,8,3,7]
输出：49 
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
```

**示例 2：**

```
输入：height = [1,1]
输出：1
```

**提示：**

- `n == height.length`
- `2 <= n <= 105`
- `0 <= height[i] <= 104`

#### 解题思路

我们在由线段长度构成的数组中使用两个指针，一个放在开始，一个置于末尾。 此外，我们会使用变量 maxarea 来持续存储到目前为止所获得的最大面积。 在每一步中，我们会找出指针所指向的两条线段形成的区域，更新 maxarea，并将指向较短线段的指针向较长线段那端移动一步。

每次移动指针的时候是移动短板侧的指针

#### 代码

```go
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	areaMax := 0
	for i < j {
		width := j - i
		if height[i] > height[j] {
			areaMax = max(areaMax, height[j]*width)
			j--
		} else {
			areaMax = max(areaMax, height[i]*width)
			i++
		}
	}
	return areaMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

