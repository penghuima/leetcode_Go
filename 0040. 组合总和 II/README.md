#### [40. 组合总和 II](https://leetcode-cn.com/problems/combination-sum-ii/)

> 难度中等

给你一个由候选元素组成的集合 `candidates` 和一个目标数 `target` ，找出 `candidates` 中所有可以使数字和为 `target` 的组合。

`candidates` 中的每个元素在每个组合中只能使用 **一次** 。

**注意：**解集不能包含重复的组合。 

**示例 1:**

```
输入: candidates = [10,1,2,7,6,1,5], target = 8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
```

**示例 2:**

```
输入: candidates = [2,5,2,1,2], target = 5,
输出:
[
[1,2,2],
[5]
]
```

#### 解题思路

难点，主要是由于`candidates` 中有重复元素，会导致最终求得解中，会出现重复元素，如 [1,2,5] 和 [2,1,5]，而在回溯算法框架中并没有去重措施

看一下 代码随想录的[题解](https://programmercarl.com/0040.%E7%BB%84%E5%90%88%E6%80%BB%E5%92%8CII.html#%E5%9B%9E%E6%BA%AF%E4%B8%89%E9%83%A8%E6%9B%B2)

#### 代码

```go
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	sum := 0
	used := make(map[int]int, len(candidates))
	//排序
	sort.Ints(candidates)
	freq := make(map[int]int, len(candidates))
	for _, v := range candidates {
		freq[v]++
	}
	var backTrack func(int)
	backTrack = func(index int) {
		//满足条件
		if sum > target {
			return
		}
		if sum == target {
			//res=append(res,append([]int(nil),path...))//这三个点貌似就将一个切片变成各个int值了
			//貌似可以以一种更优的方式添加切片，而不是以复制的方式
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}
		for i := index; i < len(candidates); i++ {
			//我在思考这样一个问题如果第一个条件不成立，后续判断条件还会判断吗
			if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == 0 {
				continue
			}
			path = append(path, candidates[i])
			sum += candidates[i]
			used[i] = 1
			backTrack(i + 1)
			path = path[:len(path)-1]
			sum -= candidates[i]
			used[i] = 0
		}
	}
	backTrack(0)
	return res
}
```

