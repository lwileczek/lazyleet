# 9. Palindrome Number
Difficulty: Easy

My submission is in [submit.bak](./submit.bak)
Afterwords, the "best" solution was copied to [main](./main.go)

## Problem
Link: https://leetcode.com/problems/palindrome-number/  

Given an integer x, return true if x is palindrome integer.
An integer is a palindrome when it reads the same backward as forward.

> For example, 121 is a palindrome while 123 is not.

 
## Examples

Example 1:
Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

Example 2:
Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:
Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

 

## Constraints:

>   -2^31 <= x <= 2^31 - 1

 
## Bonus
Could you solve it without converting the integer to a string?

# Special Notes
There is some debat on the problem that there is not enough memory allocated to convert all possible integers to a string so you have to solve it while keeping the input an
integer. So, my solution may not work even though it passed the tests.

