#### [744. 寻找比目标字母大的最小字母](https://leetcode-cn.com/problems/find-smallest-letter-greater-than-target/)

> 难度简单

给你一个排序后的字符列表 `letters` ，列表中只包含小写英文字母。另给出一个目标字母 `target`，请你寻找在这一有序列表里比目标字母大的最小字母。

在比较时，字母是依序循环出现的。举个例子：

- 如果目标字母 `target = 'z'` 并且字符列表为 `letters = ['a', 'b']`，则答案返回 `'a'`

**示例 1：**

```
输入: letters = ["c", "f", "j"]，target = "a"
输出: "c"
```

**示例 2:**

```
输入: letters = ["c","f","j"], target = "c"
输出: "f"
```

**示例 3:**

```
输入: letters = ["c","f","j"], target = "d"
输出: "f"
```

**提示：**

- `2 <= letters.length <= 104`
- `letters[i]` 是一个小写字母
- `letters` 按非递减顺序排序
- `letters` 最少包含两个不同的字母
- `target` 是一个小写字母

#### 解题思路

我们直接设计一个算法 ，时间复杂度是 O(logn) 考虑直接使用二分搜索

我们先搜索出 target 字符出现的最右边的位置    然后返回该位置的后一个字符就好了

#### 代码

```go
//这个题不就是找到 target 字符出现的最右边的位置    然后返回该位置的后一个字符就好了
func nextGreatestLetter(letters []byte, target byte) byte {
	//题目已经给了长度大于2
	index := searchTargetRight(letters, target)
	if index == len(letters) {
		return letters[0]
	}
	return letters[index+1]

}
func searchTargetRight(letters []byte, target byte) int {
	left, right := 0, len(letters)
	for left < right {
		mid := left + (right-left)>>1
		if letters[mid] > target {
			right = mid
		} else {
			left = mid + 1 //即使找到了相等 也将left +1  最后返回的时候返回 left-1即可
		}
	}
	//一定是存在的这个字符
	return left - 1
}

```

