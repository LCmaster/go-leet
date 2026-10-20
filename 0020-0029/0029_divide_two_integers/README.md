# 29. Divide Two Integers

**Difficulty:** Medium

## Description

Given two integers `dividend` and `divisor`, divide two integers without using multiplication, division, and mod operator.

The integer division should truncate toward zero, which means losing its fractional part.

Return the quotient after dividing `dividend` by `divisor`. If the quotient exceeds the 32-bit signed integer range `[-2^31, 2^31 - 1]`, clamp it to remain within bounds.

## Solution
We can use bit manipulation to simulate long division:

- Record the sign and convert `dividend` and `divisor` to `uint32` using absolute values.
- If `dividend < divisor`, return `0`.
- Calculate maximum bit-shift offset via bit length (`bits.Len32(dvd) - bits.Len32(dvs)`).
- Loop downwards from `shift` to `0`:
  - If `(dvd >> i) >= dvs`, subtract `dvs << i` from `dvd` and add `1 << i` to `quotient`.
- Apply the sign and clamp results overflowing `math.MaxInt32`.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(\log(\text{dividend}))$ | $O(1)$ |
| At most 32 bit-shift iterations for 32-bit integers | Constant extra space using integer bit shifts |
