package leetcode

import "strings"

func simplifyPath(path string) string {
	stack := []string{}
	for _, name := range strings.Split(path, "/") {
		if name == ".." {
			if len(stack) > 0 {
				//弹栈
				stack = stack[:len(stack)-1]
			}
		} else if name != "" && name != "." {
			stack = append(stack, name)
		}
		//像空格和.就不做处理了
	}

	return "/" + strings.Join(stack, "/")
}
