package threesumclosest

import "slices"

// 16. 3Sum Closest
//
// You are given an integer array `nums` of length `n` and an integer `target`.
//
// Find three integers at **distinct indices** in `nums` such that the sum is **closest** to `target`.
//
// Return the sum of the three integers.
//
// You may assume that each input would have **exactly** one solution.
//
// Example 1:
//
// ```text
// Input: nums = [-1,2,1,-4], target = 1
// Output: 2
// Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
// ```
//
// Example 2:
//
// ```text
// Input: nums = [0,0,0], target = 1
// Output: 0
// Explanation: The sum that is closest to the target is 0. (0 + 0 + 0 = 0).
// ```
//
// **Constraints:**
// - `3 <= nums.length <= 500`
// - `-1000 <= nums[i] <= 1000`
// - `-10^4 <= target <= 10^4`

func threeSumClosest(nums []int, target int) int {
	m := len(nums)
	slices.Sort(nums)
	result := nums[0] + nums[1] + nums[m-1]

	for i := 0; i < m-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		j, k := i+1, m-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			diff := target - sum
			if diff == 0 {
				return sum
			} else if diff > 0 {
				j++
			} else {
				k--
			}
			currDiff := diff
			prevDiff := target - result
			if currDiff < 0 {
				currDiff = -currDiff
			}
			if prevDiff < 0 {
				prevDiff = -prevDiff
			}

			if currDiff < prevDiff {
				result = sum
			}
		}
	}
	return result
}
