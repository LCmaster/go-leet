package longestvalidparentheses

// 32. Longest Valid Parentheses
//
// Given a string containing just the characters `'('` and `')'`, return *the length of the longest valid (well-formed) parentheses **substring*.
//
// Example 1:
//
// ```text
// Input: s = "(()"
// Output: 2
// Explanation: The longest valid parentheses substring is "()".
// ```
//
// Example 2:
//
// ```text
// Input: s = ")()())"
// Output: 4
// Explanation: The longest valid parentheses substring is "()()".
// ```
//
// Example 3:
//
// ```text
// Input: s = ""
// Output: 0
// ```
//
// **Constraints:**
// - `0 <= s.length <= 3 * 10^4`
// - `s[i]` is `'('`, or `')'`.
func longestValidParentheses(s string) int {
	maxValid := 0
	buff := []int{-1}
	for i := range len(s) {
		if s[i] == '(' {
			buff = append(buff, i)
		} else {
			buff = buff[:len(buff)-1]
			if len(buff) == 0 {
				buff = append(buff, i)
			} else {
				length := i - buff[len(buff)-1]
				if length > maxValid {
					maxValid = length
				}
			}
		}
	}
	return maxValid
}
