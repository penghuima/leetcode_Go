package leetcode

import "sort"

//由于切片中元素可以重复，造成解里会出现重复,
//如果要去除重复，要对进行排序
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	sum := 0
	used := make(map[int]int, len(candidates))
	//排序
	sort.Ints(candidates)
	freq := make(map[int]int, len(candidates))
	for _, v := range candidates {
		freq[v]++
	}
	var backTrack func(int)
	backTrack = func(index int) {
		//满足条件
		if sum > target {
			return
		}
		if sum == target {
			//res=append(res,append([]int(nil),path...))//这三个点貌似就将一个切片变成各个int值了
			//貌似可以以一种更优的方式添加切片，而不是以复制的方式
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}
		for i := index; i < len(candidates); i++ {
			//我在思考这样一个问题如果第一个条件不成立，后续判断条件还会判断吗
			if i > 0 && candidates[i] == candidates[i-1] && used[i-1] == 0 {
				continue
			}
			path = append(path, candidates[i])
			sum += candidates[i]
			used[i] = 1
			backTrack(i + 1)
			path = path[:len(path)-1]
			sum -= candidates[i]
			used[i] = 0
		}
	}
	backTrack(0)
	return res
}
