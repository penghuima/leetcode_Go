package leetcode


func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	sum := 0
	var backTrack func(int)
	backTrack = func(index int) {
		//满足条件
		if sum > target {
			return
		}
		if sum == target {
			//作为声明temp和copy操作是为了避免后续改变切片的值对结果的影响
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}
		//
		for i := index; i < len(candidates); i++ {
			path = append(path, candidates[i])
			sum += candidates[i]
			backTrack(i)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}
	backTrack(0)
	return res
}