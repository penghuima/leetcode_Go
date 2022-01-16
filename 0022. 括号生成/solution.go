package leetcode

//先写一个函数将所有可能打印出来
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
func generateParenthesis1(n int) []string {
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
//对回溯算法进行剪枝操作
func generateParenthesis(n int) []string {
	var res []string
	var path []byte
	//记录还可以使用的左括号和右括号的次数
	left,right:=n,n
	var backTrack func(int,int)
	backTrack = func(left,right int) {
		//如果满足条件
		if right<left {
			return
		}
		if left<0 ||right<0{
			return
		}
		if left==0&&right==0{
			res=append(res,string(path))
		}
		path=append(path,'(')
		backTrack(left-1,right)
		path=path[:len(path)-1]

		path=append(path,')')
		backTrack(left,right-1)
		path=path[:len(path)-1]
	}
	backTrack(left,right)
	return res
}