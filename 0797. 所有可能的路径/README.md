#### [797. 所有可能的路径](https://leetcode-cn.com/problems/all-paths-from-source-to-target/)

> 难度中等

给你一个有 `n` 个节点的 **有向无环图（DAG）**，请你找出所有从节点 `0` 到节点 `n-1` 的路径并输出（**不要求按特定顺序**）

二维数组的第 `i` 个数组中的单元都表示有向图中 `i` 号节点所能到达的下一些节点，空就是没有下一个结点了。

译者注：有向图是有方向的，即规定了 a→b 你就不能从 b→a 。

 

**示例 1：**

![img](https://assets.leetcode.com/uploads/2020/09/28/all_1.jpg)

```
输入：graph = [[1,2],[3],[3],[]]
输出：[[0,1,3],[0,2,3]]
解释：有两条路径 0 -> 1 -> 3 和 0 -> 2 -> 3
```

**示例 2：**

![img](https://assets.leetcode.com/uploads/2020/09/28/all_2.jpg)

```
输入：graph = [[4,3,1],[3,2,4],[3],[4],[]]
输出：[[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
```

**示例 3：**

```
输入：graph = [[1],[]]
输出：[[0,1]]
```

**示例 4：**

```
输入：graph = [[1,2,3],[2],[3],[]]
输出：[[0,1,2,3],[0,2,3],[0,3]]
```

**示例 5：**

```
输入：graph = [[1,3],[2],[3],[]]
输出：[[0,1,2,3],[0,3]]
```

**提示：**

- `n == graph.length`
- `2 <= n <= 15`
- `0 <= graph[i][j] < n`
- `graph[i][j] != i`（即，不存在自环）
- `graph[i]` 中的所有元素 **互不相同**
- 保证输入为 **有向无环图（DAG）**

#### 解题思路

解法很简单，以 `0` 为起点遍历图，同时记录遍历过的路径，当遍历到终点时将路径记录下来即可。

在  `backTrack func(index int)`   参数 index 表示节点序号

**既然输入的图是无环的，我们就不需要 `visited` 数组辅助了**

#### 代码

```go
func allPathsSourceTarget(graph [][]int) (res [][]int) {
	length := len(graph)
    path := []int{0}
	var backTrack func(index int) //此时这个index的意思是节点序号的意思
	backTrack = func(index int) {
		//满足条件
		if index == length-1 {
			res = append(res, append([]int(nil), path...))
			return
		}
		//单层逻辑搜素  遍历所有 index 节点可以到达的节点
		for i := 0; i < len(graph[index]); i++ {
			path = append(path, graph[index][i])
			backTrack(graph[index][i])
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return
}
```

