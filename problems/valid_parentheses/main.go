package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	pairs := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	var needs rune
	var stack []rune
	for _, char := range s {
		if char != needs {
			if value, ok := pairs[char]; ok {
				stack = append(stack, needs)
				needs = value
			} else {
				return false
			}
		} else if len(stack) > 0 {
			needs, stack = stack[len(stack)-1], stack[:len(stack)-1]
		} else {
			needs = 0
		}
	}

	return len(stack) < 1
}
