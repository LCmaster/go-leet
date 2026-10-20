# 16. 3Sum Closest

**Difficulty:** Medium

## Description

You are given an integer array `nums` of length `n` and an integer `target`.

Find three integers at **distinct indices** in `nums` such that the sum is **closest** to `target`.

Return the sum of the three integers.

You may assume that each input would have **exactly** one solution.

## Solution
We can sort the array and use the two-pointer technique to track the closest sum:

- Sort `nums` in ascending order and initialize `result = nums[0] + nums[1] + nums[m - 1]`.
- Iterate `i` from `0` to `m - 3`, skipping duplicates.
- Use two pointers `j = i + 1` and `k = m - 1`:
  - Compute `sum = nums[i] + nums[j] + nums[k]`.
  - If `sum == target`, return `sum` immediately.
  - If `|target - sum| < |target - result|`, update `result = sum`.
  - If `sum < target`, increment `j`; if `sum > target`, decrement `k`.
- Return `result`.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(n^2)$ | $O(1)$ |
| Sorting takes $O(n \log n)$ and the two-pointer traversal takes $O(n^2)$ | Constant extra space |
