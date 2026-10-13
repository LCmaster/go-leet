package palindromenumber

// 9. Palindrome Number
//
// Given an integer `x`, return `true` if `x` is a **palindrome**, and `false` otherwise.
//
// Example 1:
//
// ```text
// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.
// ```
//
// Example 2:
//
// ```text
// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
// ```
//
// Example 3:
//
// ```text
// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
// ```
//
// **Constraints:**
// - `-231<= x <= 231- 1`
//
// **Follow up:** Could you solve it without converting the integer to a string?

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reverted := 0
	for reverted < x {
		reverted = reverted*10 + (x % 10)
		x /= 10
	}

	return reverted == x || reverted/10 == x
}
