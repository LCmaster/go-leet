# 14. Longest Common Prefix

**Difficulty:** Easy

## Description

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string `""`.

## Solution
We can use vertical scanning across all strings:

- Determine the length of the shortest string in `strs` to set the maximum possible prefix length.
- Iterate column by column (index `i` from `0` to `shortest - 1`):
  - Compare character `strs[0][i]` against `strs[j][i]` for every remaining string.
  - If a mismatch is found, return `strs[0][:i]`.
- If no mismatch occurs, return `strs[0][:shortest]`.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(S)$ | $O(1)$ |
| Where $S$ is the sum of characters across all strings (at most $n \cdot \min(\text{len})$) | Constant extra space |
