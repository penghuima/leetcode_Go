#### [739. 每日温度](https://leetcode-cn.com/problems/daily-temperatures/)

> 难度中等

请根据每日 `气温` 列表 `temperatures` ，请计算在每一天需要等几天才会有更高的温度。如果气温在这之后都不会升高，请在该位置用 `0` 来代替。

**示例 1:**

```
输入: temperatures = [73,74,75,71,69,72,76,73]
输出: [1,1,4,2,1,1,0,0]
```

**示例 2:**

```
输入: temperatures = [30,40,50,60]
输出: [1,1,1,0]
```

**示例 3:**

```
输入: temperatures = [30,60,90]
输出: [1,1,0]
```

**提示：**

- 1 <= temperatures.length <= 10^5^
- `30 <= temperatures[i] <= 100`

#### 解题思路

首先想到的当然是暴力解法，两层for循环。时间复杂度是$O(n^2)$

那么接下来在来看看使用单调栈的解法。

**单调栈通常用于是一维数组，要寻找任一个元素的右边或者左边第一个比自己大或者小的元素的位置。**

更多关于单调栈的介绍 

> [讲解1](https://programmercarl.com/0739.%E6%AF%8F%E6%97%A5%E6%B8%A9%E5%BA%A6.html#%E6%80%9D%E8%B7%AF)
>
> [讲解2](https://labuladong.gitee.io/algo/2/20/50/)

#### 代码

> 暴力方法，双重for循环

```go
func dailyTemperatures(t []int) []int {
	size:=len(t)
	res:=make([]int,size)
	for i := 0; i < len(t)-1; i++ {
		for j := i + 1; j < len(t); j++ {
			// 如果之后出现更高，说明找到了
			if t[j] > t[i] {
				res[i]=j-i
				break
			}
		}
	}
	return res
}
```

> 单调栈

```go
func dailyTemperatures(temperatures []int) []int {
	size := len(temperatures)
	res := make([]int, size)
	var singleStack []int
	for i, v := range temperatures {
		for len(singleStack) != 0 && v > temperatures[singleStack[len(singleStack)-1]]{
			//弹栈
			top := singleStack[len(singleStack)-1]
			singleStack = singleStack[:len(singleStack)-1]
			//更新
			res[top] = i - top
		}
		//入栈
		singleStack = append(singleStack,i)
	}
	return res
}
```

