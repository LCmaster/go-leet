package threesum

// 15. 3Sum
//
// Given an integer array nums, return all the triplets `[nums[i], nums[j], nums[k]]` such that `i != j`, `i != k`, and `j != k`, and `nums[i] + nums[j] + nums[k] == 0`.
//
// Notice that the solution set must not contain duplicate triplets.
//
// Example 1:
//
// ```text
// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Explanation:
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
// The distinct triplets are [-1,0,1] and [-1,-1,2].
// Notice that the order of the output and the order of the triplets does not matter.
// ```
//
// Example 2:
//
// ```text
// Input: nums = [0,1,1]
// Output: []
// Explanation: The only possible triplet does not sum up to 0.
// ```
//
// Example 3:
//
// ```text
// Input: nums = [0,0,0]
// Output: [[0,0,0]]
// Explanation: The only possible triplet sums up to 0.
// ```
//
// **Constraints:**
// - `3 <= nums.length <= 3000`
// - `-105 <= nums[i] <= 105`

import "slices"

func threeSum(nums []int) [][]int {
	n := len(nums)
	slices.Sort(nums)
	result := [][]int{}
	if n < 3 {
		return result
	}
	// Check edge cases
	if nums[0] > 0 || nums[n-1] < 0 {
		return result
	}

	for i := 0; i < n-2; i++ {
		// Skip duplicates in nums[i]
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// If smallest numbers > 0 then there's no point in continuing
		if (nums[i] + nums[i+1] + nums[i+2]) > 0 {
			break
		}
		// If smallest and greatest are < 0 then smallest are contributing too much to the sum
		if nums[i]+nums[n-2]+nums[n-1] < 0 {
			continue
		}
		j, k := i+1, n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return result
}
