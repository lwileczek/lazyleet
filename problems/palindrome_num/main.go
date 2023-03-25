package main

// Submittion: https://leetcode.com/submissions/detail/832069327/

import "fmt"

func main() {
	test0 := 121
	test1 := -121
	test2 := 10
	test3 := -2
	test4 := 52125
	test5 := 0
	test6 := 1
	test7 := 678
	test8 := 103
	test9 := 298371871
	test10 := 8521331258

	fmt.Println("0. true:  ", isPalindrome(test0))
	fmt.Println("1. false: ", isPalindrome(test1))
	fmt.Println("2. false: ", isPalindrome(test2))
	fmt.Println("3. false: ", isPalindrome(test3))
	fmt.Println("4. 52125: true - ", isPalindrome(test4))
	fmt.Println("5. true:  ", isPalindrome(test5))
	fmt.Println("6. true:  ", isPalindrome(test6))
	fmt.Println("7. false: ", isPalindrome(test7))
	fmt.Println("8. false: ", isPalindrome(test8))
	fmt.Println("9. false: ", isPalindrome(test9))
	fmt.Println("10. true: ", isPalindrome(test10))
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNum := 0
	for x > revertedNum {
		revertedNum = revertedNum*10 + (x % 10)
		x /= 10
	}

	return x == revertedNum || x == (revertedNum/10)
}
