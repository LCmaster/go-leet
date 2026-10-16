package dividetwointegers

import (
	"math"
	"math/bits"
)

// 29. Divide Two Integers
//
// Given two integers `dividend` and `divisor`, divide two integers **without** using multiplication, division, and mod operator.
//
// The integer division should truncate toward zero, which means losing its fractional part. For example, `8.345` would be truncated to `8`, and `-2.7335` would be truncated to `-2`.
//
// Return *the **quotient** after dividing *`dividend`* by *`divisor`.
//
// **Note: **Assume we are dealing with an environment that could only store integers within the **32-bit** signed integer range: `[−2^31, 2^31 − 1]`. For this problem, if the quotient is **strictly greater than** `2^31 - 1`, then return `2^31 - 1`, and if the quotient is **strictly less than** `-2^31`, then return `-2^31`.
//
// Example 1:
//
// ```text
// Input: dividend = 10, divisor = 3
// Output: 3
// Explanation: 10/3 = 3.33333.. which is truncated to 3.
// ```
//
// Example 2:
//
// ```text
// Input: dividend = 7, divisor = -3
// Output: -2
// Explanation: 7/-3 = -2.33333.. which is truncated to -2.
// ```
//
// **Constraints:**
// - `-2^31 <= dividend, divisor <= 2^31 - 1`
// - `divisor != 0`
func absU32(n int) uint32 {
	if n < 0 {
		return uint32(-n)
	}
	return uint32(n)
}

func divide(dividend int, divisor int) int {
	isNegative := (dividend < 0) != (divisor < 0)
	dvd, dvs := absU32(dividend), absU32(divisor)
	if dvd < dvs {
		return 0
	}

	var quotient uint32
	shift := bits.Len32(dvd) - bits.Len32(dvs)
	for i := shift; i >= 0; i-- {
		if (dvd >> i) >= dvs {
			dvd -= dvs << i
			quotient += 1 << i
		}
	}

	if isNegative {
		return -int(quotient)
	}
	if quotient > math.MaxInt32 {
		return math.MaxInt32
	}
	return int(quotient)
}
