#### [131. 分割回文串](https://leetcode-cn.com/problems/palindrome-partitioning/)

> 难度中等

给你一个字符串 `s`，请你将 `s` 分割成一些子串，使每个子串都是 **回文串** 。返回 `s` 所有可能的分割方案。

**回文串** 是正着读和反着读都一样的字符串。

**示例 1：**

```
输入：s = "aab"
输出：[["a","a","b"],["aa","b"]]
```

**示例 2：**

```
输入：s = "a"
输出：[["a"]]
```

**提示：**

- `1 <= s.length <= 16`
- `s` 仅由小写英文字母组成

#### [解题思路](https://programmercarl.com/0131.%E5%88%86%E5%89%B2%E5%9B%9E%E6%96%87%E4%B8%B2.html)

> 难点：知道用回溯算法做，也知道回溯算法的模板，但就是不知道切割线(index)是什么，以及如何表示切割字串[index,i]

本题这涉及到两个关键问题：

1. 切割问题，有不同的切割方式
2. 判断回文

我们来分析一下切割，**其实切割问题类似组合问题**。

例如对于字符串abcdef：

- 组合问题：选取一个a之后，在bcdef中再去选取第二个，选取b之后在cdef中在选组第三个.....。
- 切割问题：切割一个a之后，在bcdef中再去切割第二段，切割b之后在cdef中在切割第三段.....。

所以切割问题，也可以抽象为一颗树形结构，如图：

![131.分割回文串](https://code-thinking.cdn.bcebos.com/pics/131.%E5%88%86%E5%89%B2%E5%9B%9E%E6%96%87%E4%B8%B2.jpg)

==递归用来纵向遍历，for循环用来横向遍历，==切割线（就是图中的红线）切割到字符串的结尾位置，说明找到了一个切割方法。

此时可以发现，切割问题的回溯搜索的过程和组合问题的回溯搜索的过程是差不多的。

**回溯三部曲**

- 递归函数参数

全局变量数组path存放切割后回文的子串，二维数组result存放结果集。 （这两个参数可以放到函数参数里）

代码如下：

```go
var res [][]string
var path []string //切割字符串集合
var backTrack func(int)
backTrack = func(index int) {
}
```

- 递归函数终止条件

![131.分割回文串](https://code-thinking.cdn.bcebos.com/pics/131.%E5%88%86%E5%89%B2%E5%9B%9E%E6%96%87%E4%B8%B2.jpg)

从树形结构的图中可以看出：切割线切到了字符串最后面，说明找到了一种切割方法，此时就是本层递归的终止终止条件。

**那么在代码里什么是切割线呢？**

在处理组合问题的时候，递归参数需要传入index，表示下一轮递归遍历的起始位置，这个 index就是切割线。

所以终止条件代码如下：

```go
if index == len(s) {
    res = append(res, append([]string(nil), path...))
}
```

- 单层搜索的逻辑

**来看看在递归循环，中如何截取子串呢？**

在`for int i = index; i < len(s); i++)`循环中，我们 定义了起始位置 index，那么 [index, i] 就是要截取的子串。

首先判断这个子串是不是回文，如果是回文，就加入在`path`中，path用来记录切割过的回文子串。

代码如下：

```go
for i := index; i < len(s); i++ {
			if isPartition(s, index, i) { // 是回文子串
				path = append(path, s[index:i+1]) // 获取[startIndex,i]在s中的子串
			}else {
				continue                 // 如果不是则直接跳过
			}
			backTrack(i + 1)            // 寻找i+1为起始位置的子串
			path = path[:len(path)-1]   // 回溯过程，弹出本次已经填在的子串
		}
```

**注意切割过的位置，不能重复切割，所以，backTrack( i + 1); 传入下一层的起始位置为i + 1**。

**判断回文子串**

可以使用双指针法，一个指针从前向后，一个指针从后先前，如果前后指针所指向的元素是相等的，就是回文字符串了。

那么判断回文的代码如下：

```go
func isPartition(s string, start, end int) bool {
	for ; start < end; {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
```

#### 代码

```go
func isPartition(s string, start, end int) bool {
	for ; start < end; {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func partition(s string) [][]string {
	var res [][]string
	var path []string //切割字符串集合
	var backTrack func(int)
	backTrack = func(index int) {
		//满足结束条件
		if index == len(s) {
			res = append(res, append([]string(nil), path...))
		}
		for i := index; i < len(s); i++ {
			if isPartition(s, index, i) {
				path = append(path, s[index:i+1])
			}else {
				continue
			}
			backTrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return res
}
```

