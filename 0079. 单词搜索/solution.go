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