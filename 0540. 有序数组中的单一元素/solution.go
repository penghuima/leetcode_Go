package leetcode


func singleNonDuplicate(nums []int) int {
	i, j := 0, len(nums)-1
	mid := 0
	//确定好搜索区间,只有一个元素的话，肯定是单的没必要搜索了啊
	for i < j {
		mid = i + (j-i)>>1 //防止加和溢出
		//偶数
		if mid%2 == 0 {
			if mid+1 < len(nums) && nums[mid] == nums[mid+1] {
				//左边没有问题
				i = mid + 2
			} else {
				//左边有问题
				j = mid
			}
		} else { //奇数
			if mid-1 >= 0 && nums[mid] == nums[mid-1] {
				//左边没有问题
				i = mid + 1
			} else {
				//右边没有问题
				j = mid - 2
			}
		}
	}
	//思考为什么是返回 nums[i] 因为判断区间是 i<j 肯定是 i=j的时候结束判断啊。
	//有的人可能会说不会j小于i的时候结束判断吗，因为判断区间始终是长度为 1 或者 3 ，5 这样长度的！
	return nums[i]
}