#### [22. 括号生成](https://leetcode-cn.com/problems/generate-parentheses/)[labuladong 题解](https://labuladong.gitee.io/plugin-v2/?qno=22)[思路](https://leetcode-cn.com/problems/generate-parentheses/#)

> 难度中等

数字 `n` 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 **有效的** 括号组合。

**示例 1：**

```
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
```

**示例 2：**

```
输入：n = 1
输出：["()"]
```

**提示：**

- `1 <= n <= 8`

#### 解决思路

本题可以改写为：

现在有 `2n` 个位置，每个位置可以放置字符 `(` 或者 `)`，组成的所有括号组合中，有多少个是合法的？

这就是典型的回溯算法提醒，暴力穷举就行了。但是以拿n=3为例，如果暴力穷举，需要判别64种情况，这就可以**剪枝**了。

不过为了减少不必要的穷举，我们要知道合法括号串有以下性质：

**1、一个「合法」括号组合的左括号数量一定等于右括号数量，这个很好理解**。

**2、对于一个「合法」的括号字符串组合 `p`，必然对于任何 `0 <= i < len(p)` 都有：子串 `p[0..i]` 中左括号的数量都大于或等于右括号的数量**。

因为从左往右算的话，肯定是左括号多嘛，到最后左右括号数量相等，说明这个括号组合是合法的。

用 `left` 记录还可以使用多少个左括号，用 `right` 记录还可以使用多少个右括号，就可以直接套用回溯算法套路框架了。

#### 代码

**方法一**

首先如果不考虑剪枝的话，打印所有括号组合，然后在写一个函数来判断这个括号组合是否合理，就可以通过了，显得有点呆

```go
func vaild(letters []byte) bool {
	balance := 0
	for _, v := range letters {
		if v == '(' {
			balance++
		} else {
			balance--
		}
		if balance < 0 {
			return false
		}
	}
	return balance == 0
}
func generateParenthesis(n int) []string {
	var res []string
	var path []byte
	letters := []byte{'(', ')'}
	var backTrack func(int)
	backTrack = func(index int) {
		//如果满足条件
		if index == 2*n {
			if vaild(path) {
				res = append(res, string(path))
			}
			return
		}
		for i := 0; i < 2; i++ {
			path = append(path, letters[i])
			backTrack(index + 1)
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return res
}
```

**方法二**

对回溯算法进行剪枝操作

> 很疑惑，在leetcode上，两种方法执行用时和内存消耗一样啊

```go
func generateParenthesis(n int) []string {
	var res []string
	var path []byte
	//记录还可以使用的左括号和右括号的次数
	left, right := n, n
	var backTrack func(int, int)
	backTrack = func(left, right int) {
		//如果满足条件
		if right < left {
			return
		}
		if left < 0 || right < 0 {
			return
		}
		if left == 0 && right == 0 {
			res = append(res, string(path))
		}
		path = append(path, '(')
		backTrack(left-1, right)
		path = path[:len(path)-1]

		path = append(path, ')')
		backTrack(left, right-1)
		path = path[:len(path)-1]
	}
	backTrack(left, right)
	return res
}
```

