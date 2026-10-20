# 18. 4Sum

**Difficulty:** Medium

## Description

Given an array `nums` of `n` integers, return an array of all unique quadruplets `[nums[a], nums[b], nums[c], nums[d]]` such that:
- `0 <= a, b, c, d < n`
- `a`, `b`, `c`, and `d` are distinct.
- `nums[a] + nums[b] + nums[c] + nums[d] == target`

You may return the answer in any order.

## Solution
We can sort the array and use two nested loops combined with two pointers and pruning:

- Sort `nums` in ascending order.
- Prune immediately if the smallest 4 elements sum to greater than `target` or the largest 4 elements sum to less than `target`.
- Loop `i` from `0` to `m - 4` and `j` from `i + 1` to `m - 3`, skipping duplicates and pruning out-of-range partial sums.
- Use two pointers `k = j + 1` and `l = m - 1`:
  - If `sum == target`, append the quadruplet and skip duplicate values for both `k` and `l`.
  - If `sum < target`, increment `k`; if `sum > target`, decrement `l`.
- Return the resulting quadruplets.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(n^3)$ | $O(1)$ |
| Sorting takes $O(n \log n)$ and the two nested loops with two-pointer search take $O(n^3)$ | Auxiliary space (excluding output quadruplets) |
