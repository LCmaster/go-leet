# 22. Generate Parentheses

**Difficulty:** Medium

## Description

Given `n` pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

## Solution
We can generate all valid combinations using backtracking:

- Start with an empty string and track the count of open brackets `i` and closed brackets `j`.
- If `i < n`, we can place an opening bracket `'('` and recurse with `i + 1`.
- If `j < i`, we can place a closing bracket `')'` and recurse with `j + 1`.
- When `i == n && j == n`, append the constructed string to `result`.
- Return `result`.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(\frac{4^n}{\sqrt{n}})$ | $O(n)$ |
| Bounded by the $n^{\text{th}}$ Catalan number $\frac{1}{n+1}\binom{2n}{n}$ | Recursion stack depth is $2n$ (excluding output combinations) |
