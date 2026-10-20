# 27. Remove Element

**Difficulty:** Easy

## Description

Given an integer array `nums` and an integer `val`, remove all occurrences of `val` in `nums` in-place. The order of the elements may be changed. Then return the number of elements in `nums` which are not equal to `val`.

## Solution
We can use a two-pointer approach swapping target elements with the end of the array:

- Initialize `i = 0` and `j = len(nums)`.
- While `i < j`:
  - If `nums[i] == val`, swap `nums[i]` with the last unexamined element `nums[j - 1]` and decrement `j`.
  - Otherwise, increment `i`.
- Return `j`, which represents the count of valid elements.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(n)$ | $O(1)$ |
| Pointers `i` and `j` traverse at most $n$ elements combined | In-place swaps with no extra memory allocated |
