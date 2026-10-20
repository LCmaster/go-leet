# 13. Roman to Integer

**Difficulty:** Easy

## Description

Roman numerals are represented by seven different symbols: `I` (1), `V` (5), `X` (10), `L` (50), `C` (100), `D` (500), and `M` (1000).

Given a roman numeral, convert it to an integer.

## Solution
We can iterate through the Roman numeral string and apply subtractive rules:

- Map characters to their integer values using a lookup table.
- Iterate from left to right:
  - If the current numeral value is less than the next numeral value (`curr < next`), subtract `curr` from `result` (e.g., `IV` -> $-1 + 5 = 4$).
  - Otherwise, add `curr` to `result`.
- Return `result`.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(n)$ | $O(1)$ |
| Single pass over the Roman numeral string of length $n$ ($\le 15$) | Fixed-size lookup table |
