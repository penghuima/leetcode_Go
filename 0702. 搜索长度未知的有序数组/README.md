#### [702. 搜索长度未知的有序数组](https://leetcode-cn.com/problems/search-in-a-sorted-array-of-unknown-size/)

> 难度中等

这是一个**交互问题**。

您有一个**升序**整数数组，其**长度未知**。您没有访问数组的权限，但是可以使用 `ArrayReader `接口访问它。你可以调用 `ArrayReader.get(i)`:

- 返回数组第`ith`个索引(**0-indexed**)处的值(即`secret[i]`)，或者
- 如果 `i` 超出了数组的边界，则返回 `231 - 1`

你也会得到一个整数 `target`。

如果存在`secret[k] == target`，请返回索引 `k` 的值；否则返回 `-1`

你必须写一个时间复杂度为 `O(log n)` 的算法。

**示例 1：**

```
输入: secret = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 存在在 nums 中，下标为 4
```

**示例 2：**

```
输入: secret = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不在数组中所以返回 -1
```

**提示：**

- `1 <= secret.length <= 104`
- `-104 <= secret[i], target <= 104`
- `secret` 严格递增

#### 解题思路

想用二分搜索的话，需要知道区间的长度，这个区间不能直接认为是数组的长度，而是包含target元素的数组区间的长度，而这个长度的最大值才是数组长度。

所以我们需要先找到包含 `target` 元素的数组区间的长度，怎么找呢？

正常来想，是不是要先确定一个右侧区间，通过判断该位置元素和 `target` 的值之间的关系来确定是否还需要继续往后探测区间。
那 `target` 和当前右侧区间位置的元素存在什么关系的时候我们需要探测呢？
当 `target > numbers.get(right)` 时，说明当前 `[left,right]` 的区间内，所有数都是小于 `target` 的，所以我们应当向后继续探测。
我们还是定义左右两个指针，令 `left = 0, right = 1`,利用给定的方法取 `right` 指针位置的元素，与 `target` 进行比较。

当 `reader.get(right) < target` 时，`right` 就向右扩展`，left` 取代 `right` 的位置，那 `right` 需要扩展多少呢？

为了保证对数级的复杂度，`right` 按照 2 倍进行扩展，即：`right = right * 2`。

#### 代码

```go
var nums []int
type ArrayReader struct{}
func (a *ArrayReader) get(index int) int {
	if index < len(nums) {
		return nums[index]
	}
	return math.MaxInt32
}
//上述一个结构体和一个方法是为了防止报错
//下面是直接提交的代码
func search(reader ArrayReader, target int) int {
	if reader.get(0) == target {
		return 0
	}
	left, right := 0, 1
	//确定搜索边界
	for reader.get(right) <= target {  //防止相等的时候  right没有扩容
		left = right
		right *= 2
	}
	for left < right {
		mid := left + (right-left)>>1
		if reader.get(mid) < target {
			left = left + 1
		} else if reader.get(mid) > target {
			right = mid
		} else {
			return mid
		}
	}
	return -1
}
```

