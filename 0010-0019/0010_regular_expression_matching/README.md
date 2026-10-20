# 10. Regular Expression Matching

**Difficulty:** Hard

## Description

Given an input string `s` and a pattern `p`, implement regular expression matching with support for `'.'` and `'*'` where:
- `'.'` Matches any single character.
- `'*'` Matches zero or more of the preceding element.

Return a boolean indicating whether the matching covers the entire input string (not partial).

## Solution
We can solve this using dynamic programming with two rows (`prev` and `curr`):

- Initialize `prev[0] = true` and handle patterns with `*` matching empty strings via `prev[i] = prev[i - 2]`.
- For each character in `s` at index `i - 1`:
  - If `s[i - 1] == p[j - 1]` or `p[j - 1] == '.'`, inherit the match from `prev[j - 1]`.
  - If `p[j - 1] == '*'`, check zero occurrences (`curr[j - 2]`) or one/more occurrences if the preceding pattern matches `s[i - 1]` (`prev[j]`).
- Swap rows after each iteration and return `prev[n]`.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(m \cdot n)$ | $O(n)$ |
| Fills an $m \times n$ state grid where $m = \text{len}(s)$ and $n = \text{len}(p)$ | Space-optimized to use two rows of size $n + 1$ |
