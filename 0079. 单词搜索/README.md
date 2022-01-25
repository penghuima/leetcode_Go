#### [79. 单词搜索](https://leetcode-cn.com/problems/word-search/)

> 难度中等

给定一个 `m x n` 二维字符网格 `board` 和一个字符串单词 `word` 。如果 `word` 存在于网格中，返回 `true` ；否则，返回 `false` 。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

**示例 1：**

![img](https://assets.leetcode.com/uploads/2020/11/04/word2.jpg)

```
输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
```

**示例 2：**

![img](https://assets.leetcode.com/uploads/2020/11/04/word-1.jpg)

```
输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
输出：true
```

**示例 3：**

![img](https://assets.leetcode.com/uploads/2020/10/15/word3.jpg)

```
输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
输出：false
```

**提示：**

- `m == board.length`
- `n = board[i].length`
- `1 <= m, n <= 6`
- `1 <= word.length <= 15`
- `board` 和 `word` 仅由**大小写英文字母**组成

#### 解决思路

> https://leetcode-cn.com/problems/word-search/solution/shou-hua-tu-jie-79-dan-ci-sou-suo-dfs-si-lu-de-cha/

**思路**

- 以"SEE"为例，首先要选起点：遍历矩阵，找到起点S。

- 起点可能不止一个，基于其中一个S，看看能否找出剩下的"EE"路径。

- 下一个字符E有四个可选点：当前点的上、下、左、右。

- 逐个尝试每一种选择。基于当前选择，为下一个字符选点，又有四种选择。

- 每到一个点做的事情是一样的。DFS 往下选点，构建路径。

- 当发现某个选择不对，不用继续选下去了，结束当前递归，考察别的选择。


**细节**

关注当前考察的点，处理它，其他丢给递归子调用去做。

- 判断当前选择的点，本身是不是一个错的点。

- 剩下的字符能否找到路径，交给递归子调用去做。


如果当前点是错的，不用往下递归了，返回false。否则继续递归四个方向，为剩下的字符选点。
那么，哪些情况说明这是一个错的点：

1. 当前的点，越出矩阵边界。
2. 当前的点，之前访问过，不满足「同一个单元格内的字母不允许被重复使用」。
3. 当前的点，不是目标点，比如你想找 E，却来到了 D

**记录访问过的点**

用一个二维矩阵 used，记录已经访问过的点，下次再选择访问这个点，就直接返回 false。

**为什么要回溯?**
有的选点是错的，选它就构建不出目标路径，不能继续选。要撤销这个选择，去尝试别的选择。

```go
//i,j表示基于当前选择的网格中的点[i,j]来判断后续能找到剩余字符的路径
flag := backTrack(i-1, j, k+1) || backTrack(i+1, j, k+1) 
		|| backTrack(i, j-1, k+1) || backTrack(i, j+1, k+1)
```


如果第一个递归调用返回 false，就会执行||后的下一个递归调用

- 这里暗含回溯：当前处在 [i,j]，选择 [i-1, j] 继续递归，返回 false 的话，会撤销 [i-1, j] 这个选择，回到 [i,j]，继续选择 [i+1, j] 递归。
  只要其中有一个递归调用返回 true，|| 后的递归就不会执行，即找到解就终止搜索。

  > 利用||的短路效应，把枝剪了。

如果求出 flag 为 false，说明基于当前点不能找到剩下的路径，所以当前递归要返回 false，还要在 used 矩阵中把当前点恢复为未访问，让它后续能正常被访问。

- 因为，基于当前路径，选当前点是不对的，但基于别的路径，走到这选它，有可能是对的。

#### 代码

> 递归终止的条件要按照一定顺序摆放，比如越界就要排放在判断当前找的字符是否正确前面，否则会报错，下标溢出

```go
package leetcode

//按照上下左右的顺序遍历
//其实这就是深度优先搜索了  DFS
func exist(board [][]byte, word string) bool {
	m,n:=len(board),len(board[0]) //网格 长和宽
	used:=make([][]bool,m)  //记录网格里的元素是否使用过
	for i:=range used{
		used[i]=make([]bool,n)
	}
	var backTrack func(int,int,int) bool
	//从网格下标i,j开始回溯,k记录找到了几个符合的字符
	backTrack= func(i,j,k int) bool {
		//满足条件
		if k==len(word){
			return true
		}
		//越界
		if i<0 ||j<0 ||i>=m ||j>=n{
			return false
		}
		//如果当前点是错的，或者该节点是正确的但已经使用过
		if board[i][j]!=word[k]||used[i][j] {
			return false
		}
		used[i][j]=true
		//利用||的短路效应使回溯搜索到解后就立马结束    上下左右搜索
		flag := backTrack(i-1, j, k+1) || backTrack(i+1, j, k+1) || backTrack(i, j-1, k+1) || backTrack(i, j+1, k+1)
		//当搜索到正确结果以后，就一层层函数栈往上返回 true ，结束递归
		if flag{
			return true
		}else{
			//如果flag为false，则说明基于当前路径，选当前节点是不对的，但基于别的路径，走到这选该节点，有可能是对的
			used[i][j]=false
			return false
		}
	}
	//遍历
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if backTrack(i,j,0){
				return true
			}
		}
	}
	return false
}
```

