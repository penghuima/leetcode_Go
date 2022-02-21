#### [838. 推多米诺](https://leetcode-cn.com/problems/push-dominoes/)

> 难度中等

`n` 张多米诺骨牌排成一行，将每张多米诺骨牌垂直竖立。在开始时，同时把一些多米诺骨牌向左或向右推。

每过一秒，倒向左边的多米诺骨牌会推动其左侧相邻的多米诺骨牌。同样地，倒向右边的多米诺骨牌也会推动竖立在其右侧的相邻多米诺骨牌。

如果一张垂直竖立的多米诺骨牌的两侧同时有多米诺骨牌倒下时，由于受力平衡， 该骨牌仍然保持不变。

就这个问题而言，我们会认为一张正在倒下的多米诺骨牌不会对其它正在倒下或已经倒下的多米诺骨牌施加额外的力。

给你一个字符串 `dominoes` 表示这一行多米诺骨牌的初始状态，其中：

- `dominoes[i] = 'L'`，表示第 `i` 张多米诺骨牌被推向左侧，
- `dominoes[i] = 'R'`，表示第 `i` 张多米诺骨牌被推向右侧，
- `dominoes[i] = '.'`，表示没有推动第 `i` 张多米诺骨牌。

返回表示最终状态的字符串。

**示例 1：**

```
输入：dominoes = "RR.L"
输出："RR.L"
解释：第一张多米诺骨牌没有给第二张施加额外的力。
```

**示例 2：**

![img](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/05/18/domino.png)

```
输入：dominoes = ".L.R...LR..L.."
输出："LL.RR.LLRRLL.."
```

**提示：**

- `n == dominoes.length`
- `1 <= n <= 105`
- `dominoes[i]` 为 `'L'`、`'R'` 或 `'.'`

#### 解题思路

> 双指针法

**一个推倒了的牌只能对另一个站着的牌起作用**


含义是：

- 两个相邻的被推倒的牌互不影响。
- 一张站立的牌（"."）的最终状态与离其两侧最近的 "L" 或 "R" 有关。

所以我们应该找出每个（"."）左右两边最近的两个被推倒了的牌，然后判断这两个牌是什么样子的即可，不用考虑这个区间以外的牌。因为这两张牌被推倒了，而这个区间外的其他牌不会对推倒了的牌起作用。

双指针
可以使用「双指针」的方式寻找 "."左右两边距离最近的被推倒的牌，形成"X....Y"型的区间。

在这两个被推倒了牌形成的区间里，根据左右两端的牌不同，有四种可能性：

```go
'R......R' => 'RRRRRRRR'
'R......L' => 'RRRRLLLL' or 'RRRR.LLLL'
'L......R' => 'L......R'
'L......L' => 'LLLLLLLL'
```


使用双指针算法：

l指向区间的开始（指向 "L" 或者 "R"）；
r跳过所有的 "."，指向区间的结束（也指向 "L" 或者 "R"）。
此时区间的形状为 "X....Y"，判断这个区间左右端点的 "X"、 "Y"是什么，确定中间的 "."的状态。


由于可能存在输入的dominoes的最左边和最右边都是 "."，那么形成不了"X....Y" 这样的区间。所以，下面的代码中，给dominoes最左边添加了一个 "L"，最右边添加了一个 "R"，添加的这两个虚拟的牌不会影响dominoes内部所有的牌的倒向，但是有助于我们形成区间，而且这两个添加的牌，也不会放到最终结果里。在每个 for 循环中，向 res 添加结果只添加区间的 `[l, r)` 部分，即左闭右开。而且注意当 l = 0 的位置，是我们虚拟的牌，不要向 res 中添加。

#### 代码

```go
func pushDominoes(dominoes string) string {
	//将左右两边添加元素 L,R
	newDominoes := []byte(dominoes)
	newDominoes = append(newDominoes, 'R')
	newDominoes = append([]byte{'L'}, newDominoes...)
	left := byte('L')
	for i := 1; i <= len(newDominoes)-2; {
		j := i
		for j < len(newDominoes) && newDominoes[j] == '.' {
			j++
		}
		right := newDominoes[j]
		if left == right {
			for i < j {
				newDominoes[i] = right
				i++
			}
		} else if left == 'R' && right == 'L' {
			k := j - 1
			for i < k {
				newDominoes[i] = left
				newDominoes[k] = right
				i++
				k--
			}
		}
		left = right
		i = j + 1
	}
	//返回
	newDominoes = newDominoes[1 : len(newDominoes)-1]
	return string(newDominoes)
}
```

