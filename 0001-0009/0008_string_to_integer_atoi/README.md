# 8. String to Integer (atoi)

**Difficulty:** Medium

## Description

Implement the `myAtoi(string s)` function, which converts a string to a 32-bit signed integer.

The algorithm for `myAtoi(string s)` is as follows:
- **Whitespace**: Ignore any leading whitespace (`" "`).
- **Signedness**: Determine the sign by checking if the next character is `'-'` or `'+'`, assuming positivity if neither present.
- **Conversion**: Read the integer by skipping leading zeros until a non-digit character is encountered or the end of the string is reached. If no digits were read, then the result is 0.
- **Rounding**: If the integer is out of the 32-bit signed integer range `[-2^31, 2^31 - 1]`, then round the integer to remain in the range. Specifically, integers less than `-2^31` should be rounded to `-2^31`, and integers greater than `2^31 - 1` should be rounded to `2^31 - 1`.

Return the integer as the final result.

## Solution
We can parse the string in a single linear scan:

- Skip all leading whitespace characters.
- Check the next character for an optional sign (`'+'` or `'-'`) and track `sign = 1` or `sign = -1`.
- Process consecutive digits `0-9`:
  - Before adding digit `d`, check for 32-bit signed integer overflow: if `result > math.MaxInt32 / 10` or (`result == math.MaxInt32 / 10 && d > 7`), return `math.MaxInt32` or `math.MinInt32` depending on `sign`.
  - Update `result = result * 10 + d`.
  - Stop at the first non-digit character.
- Return `result * sign`.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(n)$ | $O(1)$ |
| Single pass over the string of length $n$ | Constant extra space |
