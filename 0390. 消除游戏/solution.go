package main

func LastRemaining(n int) int {
	head := 1
	left := true
	step := 1
	for n > 1 {
		if left || n%2 == 1 {
			head += step
		}
		step <<= 1 // 步长*2
		left = !left
		n >>= 1 // 总数/2
	}
	return head
}
