# 17. Letter Combinations of a Phone Number

**Difficulty:** Medium

## Description

Given a string containing digits from `2-9` inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (just like on telephone buttons) is given where 1 does not map to any letters.

## Solution
We can generate all combinations using backtracking:

- Map each digit `2-9` to its corresponding letters using a static dictionary.
- If `digits` is empty, return an empty slice.
- Calculate total combinations upfront and allocate the result slice capacity.
- Use a recursive `backtrack(index)` function with a buffer `buff` of length `len(digits)`:
  - If `index == len(digits)`, append `string(buff)` to `result`.
  - Otherwise, loop over all characters mapped to `digits[index]`, place each into `buff[index]`, and recurse to `index + 1`.
- Return `result`.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(4^n \cdot n)$ | $O(n)$ |
| Generates up to $4^n$ combinations of length $n$ | Recursion stack and buffer of size $n$ (excluding output slice) |
