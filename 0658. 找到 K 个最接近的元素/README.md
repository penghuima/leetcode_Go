#### [658. 找到 K 个最接近的元素](https://leetcode-cn.com/problems/find-k-closest-elements/)

> 难度中等

给定一个 **排序好** 的数组 `arr` ，两个整数 `k` 和 `x` ，从数组中找到最靠近 `x`（两数之差最小）的 `k` 个数。返回的结果必须要是按升序排好的。

整数 `a` 比整数 `b` 更接近 `x` 需要满足：

- `|a - x| < |b - x|` 或者
- `|a - x| == |b - x|` 且 `a < b`

**示例 1：**

```
输入：arr = [1,2,3,4,5], k = 4, x = 3
输出：[1,2,3,4]
```

**示例 2：**

```
输入：arr = [1,2,3,4,5], k = 4, x = -1
输出：[1,2,3,4]
```

**提示：**

- `1 <= k <= arr.length`
- `1 <= arr.length <= 104`
- `arr` 按 **升序** 排列
- `-104 <= arr[i], x <= 104`

#### 解题思路

首先返回的结果一定是 k 个连续的，那么我们只需要使用二分法找到这段连续数组的左侧边界就好了！

#### 代码

```go
//返回的结果肯定是 k 个连续的 利用二分法找这段数组的左边界
func findClosestElements(arr []int, k int, x int) []int {
	res := []int{}
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)>>1
		//防止越界 mid+k
		if mid+k < len(arr) && x-arr[mid] > arr[mid+k]-x { //右边界比左边界更靠近x 向右寻找左边界
			left = mid + 1
		} else {
			right = mid
		}
	}
	//结果存在
	res = append(res, arr[left:left+k]...)
	return res
}
```

