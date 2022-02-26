#### [283. 移动零](https://leetcode-cn.com/problems/move-zeroes/)

> 难度简单

给定一个数组 `nums`，编写一个函数将所有 `0` 移动到数组的末尾，同时保持非零元素的相对顺序。

**请注意** ，必须在不复制数组的情况下原地对数组进行操作。

**示例 1:**

```
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
```

**示例 2:**

```
输入: nums = [0]
输出: [0]
```

**提示**:

- `1 <= nums.length <= 104`
- `-231 <= nums[i] <= 231 - 1`

#### 解题思路

> 双指针
>

![img](https://pic.leetcode-cn.com/1631236885-mZOtaR-file_1631236885904)

如样例所示，数组nums = [0,1,0,3,12]，移动完成后变成nums = [1,3,12,0,0] ，下面来讲解双指针的做法。

我们定义两个指针，i指针和k指针，i指针用来遍历整个nums数组，k指针用来放置nums数组元素。然后将非0元素按照原有的相对顺序都放置到nums数组前面，剩下的位置都置为0。这样我们就完成了0元素的移动，同时也保持了非0元素的相对顺序。

具体过程如下：

1、定义两个指针i和k，初始化i = 0，k = 0。
2、i指针向后移动，遍整个nums数组，如果 nums[i] != 0，也就是说遇到了非0元素，此时我们就将nums[i]元素放置到nums[k]位置，同时k++后一位。
3、最后**将k位置之后的元素都赋值为0**。

#### 代码

```go
//写法1
func moveZeroes1(nums []int) {
	var k int = 0
	for i := 0; i < len(nums); {
		if nums[i] != 0 {
			nums[k] = nums[i]
			i++
			k++
		} else {
			i++
		}
	}
	for j := k; j < len(nums); j++ {
		nums[j] = 0
	}
}
//写法2
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
```

