# 7. Reverse Integer

**Difficulty:** Medium

## Description

Given a signed 32-bit integer `x`, return `x` with its digits reversed. If reversing `x` causes the value to go outside the signed 32-bit integer range `[-2^31, 2^31 - 1]`, then return `0`.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

## Solution
We can pop and push digits using modulo and integer division:

- Repeatedly extract the last digit with `x % 10` and append it to `result` via `result * 10 + (x % 10)`.
- Divide `x` by 10 in each iteration until `x == 0`.
- Check if `result` overflows 32-bit signed integer limits (`math.MinInt32` to `math.MaxInt32`); if so, return `0`.
- Return `result`.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(\log_{10}(\|x\|))$ | $O(1)$ |
| Number of digits in $x$ (at most 10 iterations for a 32-bit integer) | Constant extra space |
