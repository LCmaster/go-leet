package reverseinteger

import "math"

// 7. Reverse Integer
//
// Given a signed 32-bit integer `x`, return `x`* with its digits reversed*. If reversing `x` causes the value to go outside the signed 32-bit integer range `[-2^31, 2^31 - 1]`, then return `0`.
//
// **Assume the environment does not allow you to store 64-bit integers (signed or unsigned).**
//
// Example 1:
//
// ```text
// Input: x = 123
// Output: 321
// ```
//
// Example 2:
//
// ```text
// Input: x = -123
// Output: -321
// ```
//
// Example 3:
//
// ```text
// Input: x = 120
// Output: 21
// ```
//
// **Constraints:**
// - `-2^31 <= x <= 2^31 - 1`

func reverse(x int) int {
	result := 0
	for x != 0 {
		result = result*10 + (x % 10)
		x /= 10
	}
	if result < math.MinInt32 || result > math.MaxInt32 {
		return 0
	}
	return result
}
