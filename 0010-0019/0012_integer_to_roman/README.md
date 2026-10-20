# 12. Integer to Roman

**Difficulty:** Medium

## Description

Seven different symbols represent Roman numerals: `I` (1), `V` (5), `X` (10), `L` (50), `C` (100), `D` (500), and `M` (1000).

Given an integer, convert it to a Roman numeral.

## Solution
We can use a greedy approach by matching against descending values:

- Define parallel lookup lists of standard values and subtractive pairs (`1000: M`, `900: CM`, `500: D`, `400: CD`, `100: C`, `90: XC`, `50: L`, `40: XL`, `10: X`, `9: IX`, `5: V`, `4: IV`, `1: I`).
- Iterate through values from largest to smallest.
- While `num >= value[i]`, append `symbol[i]` to the result builder and subtract `value[i]` from `num`.
- Return the string representation.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(1)$ | $O(1)$ |
| Bounded by input range ($num \le 3999$), generating at most 15 characters | Fixed-size lookup table and minimal buffer |
