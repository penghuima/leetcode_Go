package leetcode

//单调栈
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

//暴力法
func dailyTemperatures1(t []int) []int {
	size := len(t)
	res := make([]int, size)
	for i := 0; i < len(t)-1; i++ {
		for j := i + 1; j < len(t); j++ {
			// 如果之后出现更高，说明找到了
			if t[j] > t[i] {
				res[i] = j - i
				break
			}
		}
	}
	return res
}
