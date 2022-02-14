package leetcode

//暴力方法
func trap(height []int) int {
	length := len(height)
	res := 0
	for i := 1; i < length-1; i++ {
		left, right := 0, 0
		//不是找右边第一个大的元素，是找右边最大的元素
		for j := i; j < length; j++ {
			right = max(right, height[j])
		}
		//找左边第一个大的元素
		for j := i; j >= 0; j-- {
			left = max(left, height[j])
		}
		res += min(left, right) - height[i]
	}
	return res
}

//备忘录
func trap1(height []int) int {
	length := len(height)
	res := 0
	left := make([]int, length)
	right := make([]int, length)
	//初始化
	left[0] = height[0]
	right[length-1] = height[length-1]
	for i, v := range height[1:] {
		left[i+1] = max(left[i], v)
	}
	for i := length - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}
	for i, v := range height {
		res += min(left[i], right[i]) - v
	}
	return res
}

//双指针
func trap2(height []int) int {
	res := 0
	length := len(height)
	left, right := 0, length-1
	leftMax := 0
	rightMax := 0
	for left <= right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if leftMax < rightMax {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}
	return res
}
//单调栈 正序
func trap3(height []int) int {
	singleStack := []int{}
	res := 0
	for i, v := range height {
		//如果当前元素大于栈顶元素
		for len(singleStack) > 0 && v > height[singleStack[len(singleStack)-1]] {
			top := singleStack[len(singleStack)-1]         //栈顶元素
			singleStack = singleStack[:len(singleStack)-1] //弹栈
			//如果栈为空，说明栈顶元素的左边没有邻近高柱，构不成凹槽
			if len(singleStack) == 0 {
				break
			}
			//构成凹槽
			left := singleStack[len(singleStack)-1]
			w := i - left - 1
			h := min(height[left], v) - height[top]
			res += w * h
		}
		//入栈
		singleStack = append(singleStack, i)
	}
	return res
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
