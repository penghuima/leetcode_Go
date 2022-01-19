#### [39. 组合总和](https://leetcode-cn.com/problems/combination-sum/)[思路](https://leetcode-cn.com/problems/combination-sum/#)

> 难度中等

给你一个 **无重复元素** 的整数数组 `candidates` 和一个目标整数 `target` ，找出 `candidates` 中可以使数字和为目标数 `target` 的 **所有不同组合** ，并以列表形式返回。你可以按 **任意顺序** 返回这些组合。

`candidates` 中的 **同一个** 数字可以 **无限制重复被选取** 。如果至少一个数字的被选数量不同，则两种组合是不同的。 

对于给定的输入，保证和为 `target` 的不同组合数少于 `150` 个。



**示例 1：**

```
输入：candidates = [2,3,6,7], target = 7
输出：[[2,2,3],[7]]
解释：
2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
7 也是一个候选， 7 = 7 。
仅有这两种组合。
```

**示例 2：**

```
输入: candidates = [2,3,5], target = 8
输出: [[2,2,2,2],[2,3,3],[3,5]]
```

**示例 3：**

```
输入: candidates = [2], target = 1
输出: []
```

**示例 4：**

```
输入: candidates = [1], target = 1
输出: [[1]]
```

**示例 5：**

```
输入: candidates = [1], target = 2
输出: [[1,1]]
```

**提示：**

- `1 <= candidates.length <= 30`
- `1 <= candidates[i] <= 200`
- `candidate` 中的每个元素都 **互不相同**
- `1 <= target <= 500`

#### 解题思路

回溯算法，但在回溯的时候，需要注意避免重复解，如 [2,2,3] 和 [2,3,2] 就是重复解，此时需要加一个 index 参数来定位回溯到 “选择列表”哪一位了，如果不想该选择列表下标之前的元素出现在穷举范围内，则设置`backTrack(i+1)`，如果包含该下标元素则设置`backTrack(i)`好好体会 index的作用，以及代码18行和21行的写法。

#### 代码

```go
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	sum := 0
	var backTrack func(int)
	backTrack = func(index int) {
		//满足条件
		if sum > target {
			return
		}
		if sum == target {
			//作为声明temp和copy操作是为了避免后续改变切片的值对结果的影响
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}
		//
		for i := index; i < len(candidates); i++ {
			path = append(path, candidates[i])
			sum += candidates[i]
			backTrack(i)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}
	backTrack(0)
	return res
}
```

