package leetcode

import (
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	res := math.MaxInt64
	for i, v := range nums {
		//控制窗口左边界，防止下标溢出
		if i > len(nums)-k {
			break
		}
		temp := nums[i+k-1] - v
		if res > temp {
			res = temp
		}
	}
	return res
}
