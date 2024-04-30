package main


import "strings"

// simplifyPath 71.简化路径
func simplifyPath(path string) string {
	var stack []string
	for _, portion := range strings.Split(path, "/") {
		if portion == "" || portion == "." {
			continue
		} else if portion == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, portion)
		}
	}

	return "/" + strings.Join(stack, "/")
}
