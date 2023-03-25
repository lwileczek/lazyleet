package main

import "fmt"

func main() {
	test0 := "III"
	test1 := "LVIII"
	test2 := "MCMXCIV"
	fmt.Printf("0. %s: %d\n", test0, romanToInt(test0))
	fmt.Printf("1. %s: %d\n", test1, romanToInt(test1))
	fmt.Printf("2. %s: %d\n", test2, romanToInt(test2))
}
func romanToInt(s string) int {
	total := 0
	last := 0
	current := 0
	for i := len(s) - 1; i > -1; i-- {
		current = bijection[string(s[i])]
		if current < last {
			total -= current
		} else {
			total += current
		}
		last = current
	}
	return total
}

var bijection = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}
