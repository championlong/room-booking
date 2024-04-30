package main

// isValid 20.有效的括号
func isValid(s string) bool {
	if len(s) % 2 != 0 {
		return false
	}
	ruleMap := map[string]string{"(": ")", "[": "]", "{": "}"}

	var stack []string
	for _, str := range s {
		if right, ok := ruleMap[string(str)]; ok {
			stack = append(stack, right)
		} else if len(stack) == 0 || string(str) != stack[len(stack)-1] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

