# 1. Two Sum
**Difficulty:** Easy
## Description
Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.
You may assume that each input would have **exactly one solution**, and you may not use the same element twice.
You can return the answer in any order.
## Solution
We can solve this in a single pass using a hash map:

Maintain a map `seen` mapping number -> index.
For each number `num` at index i, compute `complement = target - num`.
If complement is already in `seen`, we found our pair: return `[seen[complement], i]`.
Otherwise, record the current element: `seen[num] = i`.

**Complexity:**
- Time: $O(n)$ — single pass over the array with $O(1)$ map lookups.
- Space: $O(n)$ — to store up to $n$ elements in the map.