package leetcode

func findTargetSumWays(nums []int, target int) int {
	res,sum:=0,0
	lenth:=len(nums)
	var backTrack func(index int)
	backTrack= func(index int) {
		if index==lenth && sum==target{
			res++
		}
		for i:=index;i<lenth;i++{
			//先对每个元素执行 + 操作
			sum+=nums[i]
			backTrack(i+1)
			sum-=nums[i]
			//再对每个元素执行 — 操作
			sum-=nums[i]
			backTrack(i+1)
			sum+=nums[i]
		}
	}
	backTrack(0)
	return res
}
