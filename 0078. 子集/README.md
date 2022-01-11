#### [78. 子集](https://leetcode-cn.com/problems/subsets/)

难度中等1439

给你一个整数数组 `nums` ，数组中的元素 **互不相同** 。返回该数组所有可能的子集（幂集）。

解集 **不能** 包含重复的子集。你可以按 **任意顺序** 返回解集。

**示例 1：**

```
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
```

**示例 2：**

```
输入：nums = [0]
输出：[[],[0]]
```

**提示：**

- `1 <= nums.length <= 10`
- `-10 <= nums[i] <= 10`
- `nums` 中的所有元素 **互不相同**

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

#### 代码

最终版代码：

```go
package leetcode

func subsets(nums []int) [][]int {
	var res [][]int
	var path []int
	length := len(nums)
	var backTrack func(int)
	backTrack = func(start int) {
		//满足结束条件
		temp := make([]int, len(path)) 
		copy(temp, path)
		res = append(res, temp)
		for i := start; i < length; i++ {
			//做选择
			path = append(path, nums[i])
			backTrack(i + 1) 
			//撤销选择
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return res
}
```

第一版摸爬滚打写的代码

```go
package main

import (
	"fmt"
)
var res [][]int
var path []int
func backTrack (nums []int ,temp int) {
	//满足结束条件
	temp1:=make([]int,len(path)) //为啥使用make就正确，如果使用变量 var temp1 []int,就输出错误呢
	copy(temp1,path)
	res=append(res,temp1)
	for i:=temp;i<len(nums);i++{
		//做选择
		path=append(path ,nums[i])
		backTrack(nums,i+1)
		//撤销选择
		//为什么感觉调试的时候是正确的，输出就不对了呢
		path=path[:len(path)-1]
/*
犯得一个经典错误，在使用切片过程中，忽略了后来切片值的更改对前面切片的影响。因为切片是指针形式的，依托于一个底座数组，因此数组值的更改
会影响到前面的切片。如果切片扩容的话，会重新创建一个底座数组
*/
	}

}

func subsets(nums []int) [][]int {
	backTrack(nums,0)
	return res
}
func main(){
	nums:=[]int{1,2,3}
	res:=subsets(nums)
	for k,v := range res{
		fmt.Println(k,v)
	}
}

```

