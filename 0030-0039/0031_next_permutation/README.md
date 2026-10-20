# 31. Next Permutation

**Difficulty:** Medium

## Description

Given an array of integers `nums`, find the next permutation of `nums`.

The replacement must be in place and use only constant extra memory. If no greater arrangement is possible, rearrange it as the lowest possible order (i.e., sorted in ascending order).

## Solution
We can generate the next lexicographical permutation in-place using Narayana Pandita's algorithm:

- Scan from right to left to find the first decreasing element: the largest index `i` such that `nums[i] < nums[i+1]`.
- If no such index exists (array is in descending order), reverse the entire array and return.
- Scan from right to left to find the first element `nums[j]` larger than `nums[i]`.
- Swap `nums[i]` and `nums[j]`.
- Reverse the suffix starting at `i + 1` to the end of the array to minimize the newly formed prefix.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(n)$ | $O(1)$ |
| Two linear scans and a partial array reversal take $O(n)$ total time | Reorders elements in-place with constant extra memory |
