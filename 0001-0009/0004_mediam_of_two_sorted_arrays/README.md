# 4. Median of Two Sorted Arrays

**Difficulty:** Hard

## Description

Given two sorted arrays `nums1` and `nums2` of size `m` and `n` respectively, return the median of the two sorted arrays.

The overall run time complexity should be `O(log (m+n))`.

## Solution
We can use binary search on the partition of the smaller array:

- Ensure `nums1` is the smaller array (`len(nums1) <= len(nums2)`) to minimize the search space.
- Binary search for a partition in `nums1` (`partitionX`) and derive `partitionY = (m + n + 1) / 2 - partitionX`.
- Check if the partition is valid (`maxLeftX <= minRightY` and `maxLeftY <= minRightX`).
- If valid, return $\max(\text{maxLeftX}, \text{maxLeftY})$ when $m + n$ is odd, or the average of $\max(\text{maxLeftX}, \text{maxLeftY})$ and $\min(\text{minRightX}, \text{minRightY})$ when even.
- Otherwise, adjust the binary search range towards the correct partition.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(\log(\min(m, n)))$ | $O(1)$ |
| Binary search on the smaller array of size $\min(m, n)$ | Constant extra space |
