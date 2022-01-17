package leetcode

func maxDepth(s string) int {
	res, size := 0, 0
	for _, v := range s {
		if v == '(' {
			size++
			if size > res {
				res = size
			}
		} else if v == ')' {
			size--
		}
	}
	return res
}
