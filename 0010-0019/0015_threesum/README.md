# 15. 3Sum

**Difficulty:** Medium

## Description

Given an integer array `nums`, return all the triplets `[nums[i], nums[j], nums[k]]` such that `i != j`, `i != k`, and `j != k`, and `nums[i] + nums[j] + nums[k] == 0`.

Notice that the solution set must not contain duplicate triplets.

## Solution
We can sort the array and use the two-pointer approach with pruning:

- Sort `nums` in ascending order.
- Iterate `i` from `0` to `n - 3`:
  - Skip duplicate values for `nums[i]`.
  - Prune early if `nums[i] + nums[i+1] + nums[i+2] > 0` (no further sums can equal 0) or `nums[i] + nums[n-2] + nums[n-1] < 0` (current number is too small).
  - Use two pointers `j = i + 1` and `k = n - 1` to find pairs where `nums[i] + nums[j] + nums[k] == 0`.
  - When a triplet is found, append it and advance `j` and `k` past duplicate elements.
  - If the sum is less than 0, increment `j`; if greater than 0, decrement `k`.
- Return the list of unique triplets.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(n^2)$ | $O(1)$ |
| Sorting takes $O(n \log n)$ and the two-pointer search takes $O(n^2)$ | Auxiliary space (excluding the output triplets list) |
