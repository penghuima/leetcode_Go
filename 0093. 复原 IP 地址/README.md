#### [93. 复原 IP 地址](https://leetcode-cn.com/problems/restore-ip-addresses/)

> 难度中等

**有效 IP 地址** 正好由四个整数（每个整数位于 `0` 到 `255` 之间组成，且不能含有前导 `0`），整数之间用 `'.'` 分隔。

- 例如："0.1.2.201" 和 "192.168.1.1" 是 **有效** IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 **无效** IP 地址。

给定一个只包含数字的字符串 `s` ，用以表示一个 IP 地址，返回所有可能的**有效 IP 地址**，这些地址可以通过在 `s` 中插入 `'.'` 来形成。你不能重新排序或删除 `s` 中的任何数字。你可以按 **任何** 顺序返回答案。

**示例 1：**

```
输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]
```

**示例 2：**

```
输入：s = "0000"
输出：["0.0.0.0"]
```

**示例 3：**

```
输入：s = "1111"
输出：["1.1.1.1"]
```

**示例 4：**

```
输入：s = "010010"
输出：["0.10.0.10","0.100.1.0"]
```

**示例 5：**

```
输入：s = "101023"
输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
```

**提示：**

- `0 <= s.length <= 20`
- `s` 仅由数字组成

#### 解题思路

这道题和分割回文串题目类似

[回溯切割问题](https://programmercarl.com/0093.%E5%A4%8D%E5%8E%9FIP%E5%9C%B0%E5%9D%80.html#%E5%9B%9E%E6%BA%AF%E4%B8%89%E9%83%A8%E6%9B%B2)

![93.复原IP地址](https://img-blog.csdnimg.cn/20201123203735933.png)

**难点：**

- 切割问题可以抽象为组合问题
- 如何模拟切割
- 在递归循环中切割字串如何表示
- 判断切割字符串是否符合要求
- 递归如何终止

**解答**

1. 关于模拟切割，其实 backTrack(index,int) 的 index 就是切割线，并在下一次递归时加一，避免重复切割 backTrack(index+1)
2. 切割字符串的表示就是 [index,i] 在代码中表示就是  `s[index:i+1]`
3. 写一个函数 `isNormalIp(s string, start, end int) bool` 判断切割字符串是否符合要求
4. 如果我们切割线移动到末尾并且我们正好切割4段，则是一次正确的切割，将这次结果追加到 result 切片中

**回溯三部曲**

- 递归函数参数

全局变量 path 存放切割后的字符串，index一定是需要的，用来记录下次切割的起始位置

```go
var res []string
var path []string
var backTrack func(index int)
```

- 递归终止条件

如果切割线移动到末尾，并且正好切割四段，则表示正确切割

```go
if index == len(s) && len(path) == 4 {
    temp := path[0] + "." + path[1] + "." + path[2] + "." + path[3] //拼成1个ip字符串
    res = append(res, temp)
}
```

- 单层搜索的逻辑

在`for (int i = index; i < s.size(); i++)`循环中 [index, i] 这个区间就是截取的子串，需要判断这个子串是否合法。

如果合法就表示已经分割，将切割的字符串追加到 path ,并进行递归。如果这个字串不合法,或者已经切割四段了，而index还没移动到末尾就直接 return 回溯

```go
for i := index; i < len(s); i++ {
			//i-index+1 <= 3这个条件不满足就不要执行 isNormalIp函数了
			if i-index+1 <= 3 && len(path) < 4 && isNormalIp(s, index, i) {
				path = append(path, s[index:i+1])
				backTrack(i + 1)
			} else {
				return
			}
			path = path[:len(path)-1]
		}
```

#### 代码

> strconv.Atoi(s[start : end+1]) 将字符串转变为int

```go
package leetcode
import "strconv"
//判断字串是否合理 符合 [0,255] 且数字不能含有前导0
func isNormalIp(s string, start, end int) bool {
	if end-start+1 > 1 && s[start] == '0' {
		return false
	}
	//将字符串转变为 int 如 ”255“-> 255
	checkInt, _ := strconv.Atoi(s[start : end+1])
	if checkInt > 255 {
		return false
	}
	return true
}
//需要优化，多递归了
func restoreIpAddresses(s string) []string {
	var res []string
	var path []string
	var backTrack func(int)
	backTrack = func(index int) {
		if index == len(s) && len(path) == 4 {
			temp := path[0] + "." + path[1] + "." + path[2] + "." + path[3] //拼成1个ip字符串
			res = append(res, temp)
		}
		for i := index; i < len(s); i++ {
			//i-index+1 <= 3这个条件不满足就不要执行 isNormalIp函数了
			if i-index+1 <= 3 && len(path) < 4 && isNormalIp(s, index, i) {
				path = append(path, s[index:i+1])
				backTrack(i + 1)
			} else {
				return
			}
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return res
}
```

