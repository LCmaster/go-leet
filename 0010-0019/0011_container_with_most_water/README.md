# 11. Container With Most Water

**Difficulty:** Medium

## Description

You are given an integer array `height` of length `n`. There are `n` vertical lines drawn such that the two endpoints of the `ith` line are `(i, 0)` and `(i, height[i])`.

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

## Solution
We can use the two-pointer technique moving inward:

- Initialize `left = 0` and `right = len(height) - 1`.
- In each step, compute `area = (right - left) * min(height[left], height[right])` and update `total`.
- Move the pointer with the smaller height inward (if `height[left] < height[right]`, `left++`; else `right--`) since moving the taller boundary could never yield a larger area with a smaller width.
- Return `total`.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(n)$ | $O(1)$ |
| Single pass where pointers meet in $n - 1$ steps | Constant extra space |
