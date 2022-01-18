#### [47. 全排列 II](https://leetcode-cn.com/problems/permutations-ii/)

> 难度中等

给定一个可包含重复数字的序列 `nums` ，**按任意顺序** 返回所有不重复的全排列。

**示例 1：**

```
输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]
```

**示例 2：**

```
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
```

**提示：**

- `1 <= nums.length <= 8`
- `-10 <= nums[i] <= 10`

#### 解题思路

回溯算法去重，主要是同一树层上的节点去重

![47.全排列II1](https://img-blog.csdnimg.cn/20201124201331223.png)

#### 代码

```go
package leetcode
import "sort"
func permuteUnique(nums []int) [][]int {
	var res [][]int
	var path []int
	length := len(nums)
	used := make([]bool, length)
	sort.Ints(nums)
	var backTrack func(int)
	backTrack = func(index int) {
		//满足条件
		if len(path) == length {
			res = append(res, append([]int(nil), path...))
		}
		for i := 0; i < length; i++ {
			if used[i] || i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backTrack(index + 1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return res
}
```

