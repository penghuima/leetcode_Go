package leetcode

func pushDominoes(dominoes string) string {
	//将左右两边添加元素 L,R
	newDominoes := []byte(dominoes)
	newDominoes = append(newDominoes, 'R')
	newDominoes = append([]byte{'L'}, newDominoes...)
	left := byte('L')
	for i := 1; i <= len(newDominoes)-2; {
		j := i
		for j < len(newDominoes) && newDominoes[j] == '.' {
			j++
		}
		right := newDominoes[j]
		if left == right {
			for i < j {
				newDominoes[i] = right
				i++
			}
		} else if left == 'R' && right == 'L' {
			k := j - 1
			for i < k {
				newDominoes[i] = left
				newDominoes[k] = right
				i++
				k--
			}
		}
		left = right
		i = j + 1
	}
	//返回
	newDominoes = newDominoes[1 : len(newDominoes)-1]
	return string(newDominoes)
}
