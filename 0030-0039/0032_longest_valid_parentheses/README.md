# 32. Longest Valid Parentheses

**Difficulty:** Hard

## Description

Given a string containing just the characters `'('` and `')'`, return the length of the longest valid (well-formed) parentheses substring.

## Solution
We can use a stack storing indices to compute valid substring lengths:

- Preallocate and push `-1` onto the index stack as a base boundary for valid substrings.
- Iterate through index `i` of the string:
  - If `s[i] == '('`, push `i` onto the stack.
  - If `s[i] == ')'`, pop the top index:
    - If the stack is now empty, push current index `i` as the new base boundary.
    - Otherwise, compute `length = i - stack.top()` and update `maxValid = max(maxValid, length)`.
- Return `maxValid`.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(n)$ | $O(n)$ |
| Single pass over the string of length $n$ | Preallocated stack storing at most $n + 1$ indices |
