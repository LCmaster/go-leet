package reverseinteger

import (
	"math"
)

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
	if x == 0 {
		return x
	}

	lowerLimit := math.MinInt32
	upperLimit := math.MaxInt32

	n := x
	isNegative := false
	if n < 0 {
		isNegative = true
		n *= -1
	}

	i, result := 0, 0
	digit := []int{}
	for n/PowInt(10, i) >= 1 {
		d := (n / PowInt(10, i)) % 10
		digit = append(digit, d)
		i++
	}
	for j := 0; i > 0; j++ {
		result += PowInt(10, i-1) * digit[j]
		i--
	}

	if isNegative {
		result *= -1
	}

	if result < lowerLimit || result > upperLimit {
		result = 0
	}

	return result
}

func PowInt(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp&1 == 1 {
			result *= base
		}
		base *= base
		exp >>= 1
	}
	return result
}
