#### [46. 全排列](https://leetcode-cn.com/problems/permutations/)

> 难度中等

给定一个不含重复数字的数组 `nums` ，返回其 **所有可能的全排列** 。你可以 **按任意顺序** 返回答案。

**示例 1：**

```
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
```

**示例 2：**

```
输入：nums = [0,1]
输出：[[0,1],[1,0]]
```

**示例 3：**

```
输入：nums = [1]
输出：[[1]]
```

**提示：**

- `1 <= nums.length <= 6`
- `-10 <= nums[i] <= 10`
- `nums` 中的所有整数 **互不相同**

#### 解决思路

废话不多说，直接上回溯算法框架。**解决一个回溯问题，实际上就是一个决策树的遍历过程**。你只需要思考 3 个问题：

1、路径：也就是已经做出的选择。(记录已经做出的选择)

2、选择列表：也就是你当前可以做的选择。

3、结束条件：也就是到达决策树底层，无法再做选择的条件。

如果你不理解这三个词语的解释，没关系，我们后面会用「全排列」和「N 皇后问题」这两个经典的回溯算法问题来帮你理解这些词语是什么意思，现在你先留着印象。

代码方面，回溯算法的框架：

```go
路径 path []int
选择列表 nums []int
标记是否选择 visited map[int]bool

func backtrack(选择列表) (result [][]int){
    if 满足结束条件:
        result.add(路径)
        return
    
    for 选择 in 选择列表:
        #做选择
        将该选择从选择列表中移除
        路径.add(选择)
        backtrack(选择列表)
        #撤销选择
        路径.remove(选择)
        将该选择恢复到选择列表
}
```

**其核心就是 for 循环里面的递归，在递归调用之前「做选择」，在递归调用之后「撤销选择」**，特别简单。

什么叫做选择和撤销选择呢，这个框架的底层原理是什么呢？下面我们就通过「全排列」这个问题来解开之前的疑惑，详细探究一下其中的奥妙！

![image-20220104105746108](https://cdn.jsdelivr.net/gh/penghuima/ImageBed@master/img/blog_file/PicGo-Github-ImgBed20220104105746.png)

#### 代码

```go
func permute(nums []int) [][]int {
	var res [][]int                  //最终要返回的二维数组切片
	var path []int                   //记录路径
	var visited = make(map[int]bool) //对已经做过的选择进行标记
	length := len(nums)
	var backTrack func()            //定义函数变量类型
	backTrack = func() {
		//如果满足结束条件，则将路径加入到二维数组切片中
		if len(path) == length {
			temp := make([]int, length)
			copy(temp, path)        //path的值存到中
			res = append(res, temp)
			return
		}
		//回溯
		for i := 0; i < length; i++ {
			//做选择
			if visited[nums[i]] == true {
				continue
			}
			path = append(path, nums[i])
			visited[nums[i]] = true
			backTrack()
			//撤销选择
			path = path[:len(path)-1]
			visited[nums[i]] = false
		}
	}
	backTrack()                    //开始执行函数
	return res
}
```

