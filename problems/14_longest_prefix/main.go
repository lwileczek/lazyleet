package main

import (
	"fmt"
)

func main() {
	fmt.Println("begin tests")
	test := []string{
		"abab", "aba", "abc",
	}

	fmt.Println("Test4 - ab ==", longestCommonPrefix(test))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	if strs[0] == "" {
		return ""
	}

	var bytefix []byte
search:
	for k := range strs[0] {
		for _, str := range strs[1:] {
			if len(str) <= k {
				break search
			}
			if strs[0][k] != str[k] {
				break search
			}
		}
		bytefix = append(bytefix, strs[0][k])
	}

	return string(bytefix[:])
}
