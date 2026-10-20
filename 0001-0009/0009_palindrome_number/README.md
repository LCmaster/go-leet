# 9. Palindrome Number

**Difficulty:** Easy

## Description

Given an integer `x`, return `true` if `x` is a palindrome, and `false` otherwise.

## Solution
We can revert only the second half of the number to avoid overflow:

- Negative numbers cannot be palindromes due to the leading minus sign.
- Numbers ending with `0` (except `0` itself) cannot be palindromes.
- Repeatedly pop the last digit from `x` and append it to `reverted` until `reverted >= x`.
- Check if `x == reverted` (for even number of digits) or `x == reverted / 10` (for odd number of digits, discarding the middle digit).

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(\log_{10}(x))$ | $O(1)$ |
| Only iterates through half of the digits of $x$ | Constant extra space |
