# 26. Remove Duplicates from Sorted Array

**Difficulty:** Easy

## Description

Given an integer array `nums` sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same.

Return the number of unique elements `k`.

## Solution
We can use a two-pointer approach keeping track of the insertion position:

- If `nums` is empty, return `0`.
- Initialize `k = 1` for the write index of unique elements.
- Iterate `i` from `1` to `len(nums) - 1`:
  - When `nums[i] != nums[k - 1]`, assign `nums[k] = nums[i]` and increment `k`.
- Return `k`.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(n)$ | $O(1)$ |
| Single pass over the array of length $n$ | Modifies the array in-place with $O(1)$ extra memory |
