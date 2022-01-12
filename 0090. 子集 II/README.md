#### [90. 子集 II](https://leetcode-cn.com/problems/subsets-ii/)

> 难度中等

给你一个整数数组 `nums` ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。

解集 **不能** 包含重复的子集。返回的解集中，子集可以按 **任意顺序** 排列。

**示例 1：**

```
输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
```

**示例 2：**

```
输入：nums = [0]
输出：[[],[0]]
```

**提示：**

- `1 <= nums.length <= 10`
- `-10 <= nums[i] <= 10`

#### 解题思路

回溯算法去重

![90.子集II](https://img-blog.csdnimg.cn/20201124195411977.png)

#### 代码

```go
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var path []int
	used := make([]bool, len(nums))
	var backTrack func(int)
	backTrack = func(index int) {
		//满足条件
		res = append(res, append([]int(nil), path...))
		for i := index; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backTrack(i + 1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return res
}
```

> 第16行是 i+1, 而不是 index+1