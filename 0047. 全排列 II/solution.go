package leetcode

import "sort"

func permuteUnique(nums []int) [][]int {
	var res [][]int
	var path []int
	length := len(nums)
	used := make([]bool, length)
	sort.Ints(nums)
	var backTrack func(int)
	backTrack = func(index int) {
		//满足条件
		if len(path) == length {
			res = append(res, append([]int(nil), path...))
		}
		for i := 0; i < length; i++ {
			if used[i] || i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backTrack(index + 1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return res
}