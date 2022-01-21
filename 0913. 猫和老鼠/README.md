#### [913. 猫和老鼠](https://leetcode-cn.com/problems/cat-and-mouse/)

> 难度困难

两位玩家分别扮演猫和老鼠，在一张 **无向** 图上进行游戏，两人轮流行动。

图的形式是：`graph[a]` 是一个列表，由满足 `ab` 是图中的一条边的所有节点 `b` 组成。

老鼠从节点 `1` 开始，第一个出发；猫从节点 `2` 开始，第二个出发。在节点 `0` 处有一个洞。

在每个玩家的行动中，他们 **必须** 沿着图中与所在当前位置连通的一条边移动。例如，如果老鼠在节点 `1` ，那么它必须移动到 `graph[1]` 中的任一节点。

此外，猫无法移动到洞中（节点 `0`）。

然后，游戏在出现以下三种情形之一时结束：

- 如果猫和老鼠出现在同一个节点，猫获胜。
- 如果老鼠到达洞中，老鼠获胜。
- 如果某一位置重复出现（即，玩家的位置和移动顺序都与上一次行动相同），游戏平局。

给你一张图 `graph` ，并假设两位玩家都都以最佳状态参与游戏：

- 如果老鼠获胜，则返回 `1`；
- 如果猫获胜，则返回 `2`；
- 如果平局，则返回 `0` 。

**示例 1：**

![img](https://assets.leetcode.com/uploads/2020/11/17/cat1.jpg)

```
输入：graph = [[2,5],[3],[0,4,5],[1,4,5],[2,3],[0,2,3]]
输出：0
```

**示例 2：**

![img](https://assets.leetcode.com/uploads/2020/11/17/cat2.jpg)

```
输入：graph = [[1,3],[0],[3],[0,2]]
输出：1 
```

**提示：**

- `3 <= graph.length <= 50`
- `1 <= graph[i].length < graph.length`
- `0 <= graph[i][j] < graph.length`
- `graph[i][j] != i`
- `graph[i]` 互不相同
- 猫和老鼠在游戏中总是移动

#### [解决思路](https://leetcode-cn.com/problems/cat-and-mouse/solution/mao-he-lao-shu-by-leetcode-solution-444x/)

**前言**
这道题是博弈问题，猫和老鼠都按照最优策略参与游戏。

在阐述具体解法之前，首先介绍博弈问题中的三个概念：必胜状态、必败状态与必和状态。

- 对于特定状态，如果游戏已经结束，则根据结束时的状态决定必胜状态、必败状态与必和状态。
  - 如果分出胜负，则该特定状态对于获胜方为必胜状态，对于落败方为必败状态。
  - 如果是平局，则该特定状态对于双方都为必和状态。

- 从特定状态开始，如果存在一种操作将状态变成必败状态，则当前玩家可以选择该操作，将必败状态留给对方玩家，因此该特定状态对于当前玩家为必胜状态。

- 从特定状态开始，如果所有操作都会将状态变成必胜状态，则无论当前玩家选择哪种操作，都会将必胜状态留给对方玩家，因此该特定状态对于当前玩家为必败状态。

- 从特定状态开始，如果任何操作都不能将状态变成必败状态，但是存在一种操作将状态变成必和状态，则当前玩家可以选择该操作，将必和状态留给对方玩家，因此该特定状态对于双方玩家都为必和状态。


对于每个玩家，最优策略如下：

争取将必胜状态留给自己，将必败状态留给对方玩家。

在自己无法到达必胜状态的情况下，争取将必和状态留给自己。

#### 代码

```go
const (
    draw     = 0
    mouseWin = 1
    catWin   = 2
)

func catMouseGame(graph [][]int) int {
    n := len(graph)
    dp := make([][][]int, n)
    for i := range dp {
        dp[i] = make([][]int, n)
        for j := range dp[i] {
            dp[i][j] = make([]int, n*2)
            for k := range dp[i][j] {
                dp[i][j][k] = -1
            }
        }
    }

    var getResult, getNextResult func(int, int, int) int
    getResult = func(mouse, cat, turns int) int {
        if turns == n*2 {
            return draw
        }
        res := dp[mouse][cat][turns]
        if res != -1 {
            return res
        }
        if mouse == 0 {
            res = mouseWin
        } else if cat == mouse {
            res = catWin
        } else {
            res = getNextResult(mouse, cat, turns)
        }
        dp[mouse][cat][turns] = res
        return res
    }
    getNextResult = func(mouse, cat, turns int) int {
        curMove := mouse
        if turns%2 == 1 {
            curMove = cat
        }
        defaultRes := mouseWin
        if curMove == mouse {
            defaultRes = catWin
        }
        res := defaultRes
        for _, next := range graph[curMove] {
            if curMove == cat && next == 0 {
                continue
            }
            nextMouse, nextCat := mouse, cat
            if curMove == mouse {
                nextMouse = next
            } else if curMove == cat {
                nextCat = next
            }
            nextRes := getResult(nextMouse, nextCat, turns+1)
            if nextRes != defaultRes {
                res = nextRes
                if res != draw {
                    break
                }
            }
        }
        return res
    }
    return getResult(1, 2, 0)
}
```

