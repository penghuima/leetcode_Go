#### [540. 有序数组中的单一元素](https://leetcode-cn.com/problems/single-element-in-a-sorted-array/)

> 难度中等

给你一个仅由整数组成的**有序数组**，其中每个元素都会出现两次，唯有一个数只会出现一次。

请你找出并返回只出现一次的那个数。

你设计的解决方案必须满足 `O(log n)` 时间复杂度和 `O(1)` 空间复杂度。

**示例 1:**

```
输入: nums = [1,1,2,3,3,4,4,8,8]
输出: 2
```

**示例 2:**

```
输入: nums =  [3,3,7,7,10,11,11]
输出: 10
```

**提示:**

- `1 <= nums.length <= 105`
- `0 <= nums[i] <= 105`

#### 解题思路

这道题因为加了十几件复杂度限制条件，因此注定使用二分搜索解答。

对于一个有序数组，如果整个数组每个元素都是成对出现的，**成对元素中的第一个所对应的下标必然是偶数，成对元素中的第二个所对应的下标必然是奇数** 。如果有某个元素插入进去的话，势必该元素左边的数据符合上述规矩，而该元素右边的数据必定违反上述规矩。

存在这样的二段性，指导我们根据当前二分点 mid 的奇偶性进行分情况讨论：

- mid 为偶数下标：根据上述结论，正常情况下偶数下标的值会与下一值相同，因此如果满足该条件，可以确保 mid 之前并没有插入单一元素。即需要将更新逻辑修改为 l = mid + 2 和 r =mid ；

- mid 为奇数下标：同理，根据上述结论，正常情况下奇数下标的值会与上一值相同，因此如果满足该条件，可以确保 mid 之前并没有插入单一元素，相应的更新 l=mid+1 和 r=mid。

其实有一种统一的做法，不用判断mid是否是奇偶，直接使用与1异或解决。

利用按位异或的性质，可以得到 mid 和相邻的数之间的如下关系

```go
func singleNonDuplicate(nums []int) int {
    low, high := 0, len(nums)-1
    for low < high {
        mid := low + (high-low)/2
        if nums[mid] == nums[mid^1] {  //异或 大大简化了代码简洁性
            low = mid + 1
        } else {
            high = mid
        }
    }
    return nums[low]
}
```

#### 代码

> 代码1

```go
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
```

> ~~代码~~2 都不需要加边界溢出条件，因为 i<j 而区间又一定是 1，3，5 这样的

```go
func singleNonDuplicate(nums []int) int {
	i, j := 0, len(nums)-1
	mid := 0
	//确定好搜索区间,只有一个元素的话，肯定是单的没必要搜索了啊
	for i < j {
		mid = i + (j-i)>>1 //防止加和溢出
		//偶数
		if mid%2 == 0 {
			if nums[mid] == nums[mid+1] {
				//左边没有问题
				i = mid + 2
			} else {
				//左边有问题
				j = mid
			}
		} else { //奇数
			if  nums[mid] == nums[mid-1] {
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
```

> 其实有一种统一的做法，不用判断mid是否是奇偶，直接使用与1异或解决。
>
> 利用按位异或的性质，可以得到 mid 和相邻的数之间的如下关系  `nums[mid] == nums[mid^1]`

```go
func singleNonDuplicate(nums []int) int {
    low, high := 0, len(nums)-1
    for low < high {
        mid := low + (high-low)/2
        if nums[mid] == nums[mid^1] {  //异或 大大简化了代码简洁性
            low = mid + 1
        } else {
            high = mid
        }
    }
    return nums[low]
}
```

