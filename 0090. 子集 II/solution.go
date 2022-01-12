package leetcode

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var path []int
	used := make([]bool, len(nums))
	var backTrack func(int)
	backTrack = func(index int) {
		//满足条件
		res = append(res, append([]int(nil), path...))
		for i := index; i < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backTrack(i + 1)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return res
}