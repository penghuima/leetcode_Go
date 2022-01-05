#### [77. 组合](https://leetcode-cn.com/problems/combinations/)[labuladong 题解](https://labuladong.gitee.io/plugin-v2/?qno=77)[思路](https://leetcode-cn.com/problems/combinations/#)

> 难度中等

给定两个整数 `n` 和 `k`，返回范围 `[1, n]` 中所有可能的 `k` 个数的组合。

你可以按 **任何顺序** 返回答案。

**示例 1：**

```
输入：n = 4, k = 2
输出：
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]
```

**示例 2：**

```
输入：n = 1, k = 1
输出：[[1]]
```

**提示：**

- `1 <= n <= 20`
- `1 <= k <= n`

#### 解决思路

回溯算法的框架：

```go
路径 path []int
选择列表 nums []int
标记是否选择 visited map[int]bool

func backtrack(路径, 选择列表) (result [][]int){
    if 满足结束条件:
        result.add(路径)
        return
    
    for 选择 in 选择列表:
        #做选择
        将该选择从选择列表中移除
        路径.add(选择)
        backtrack(路径, 选择列表)
        #撤销选择
        路径.remove(选择)
        将该选择恢复到选择列表
}
```

**其核心就是 for 循环里面的递归，在递归调用之前「做选择」，在递归调用之后「撤销选择」**，特别简单。

使用一个start参数控制递归，生成如下这样一棵树，改造回溯算法模板

![image-20220105093206989](https://cdn.jsdelivr.net/gh/penghuima/ImageBed@master/img/blog_file/PicGo-Github-ImgBed20220105093214.png)

#### 代码

```go
func combine(n int, k int) [][]int {
	var res [][]int
	var path []int

	var backTrack func(int)
	backTrack = func(start int) {
		//符合条件
		if len(path) == k {
			temp := make([]int, k)
			copy(temp, path)
			res = append(res, temp)
		}
		for i := start; i <= n; i++ { //i从1开始
			//做出选择
			path = append(path, i)
			backTrack(i + 1)
			//回溯
			path = path[:len(path)-1]
		}
	}
	backTrack(1)
	return res
}
```

