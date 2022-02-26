package leetcode

func moveZeroes(nums []int)  {
	var k int =0
	for i:=0;i<len(nums);{
		if nums[i]!=0{
			nums[k],nums[i] = nums[i],nums[k]
			i++
			k++
		}else {
			i++
		}
	}
}
