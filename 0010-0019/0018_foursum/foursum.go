package foursum

import "slices"

// 18. 4Sum
//
// Given an array `nums` of `n` integers, return *an array of all the **unique** quadruplets* `[nums[a], nums[b], nums[c], nums[d]]` such that:
// - `0 <= a, b, c, d< n`
// - `a`, `b`, `c`, and `d` are **distinct**.
// - `nums[a] + nums[b] + nums[c] + nums[d] == target`
//
// You may return the answer in **any order**.
//
// Example 1:
//
// ```text
// Input: nums = [1,0,-1,0,-2,2], target = 0
// Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
// ```
//
// Example 2:
//
// ```text
// Input: nums = [2,2,2,2,2], target = 8
// Output: [[2,2,2,2]]
// ```
//
// **Constraints:**
// - `1 <= nums.length <= 200`
// - `-10^9 <= nums[i] <= 10^9`
// - `-10^9 <= target <= 10^9`

func fourSum(nums []int, target int) [][]int {
	m := len(nums)
	result := [][]int{}
	if m < 4 {
		return result
	}
	slices.Sort(nums)
	if nums[0]+nums[1]+nums[2]+nums[3] > target {
		return result
	}
	if nums[m-4]+nums[m-3]+nums[m-2]+nums[m-1] < target {
		return result
	}
	for i := 0; i < m-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}
		if nums[i]+nums[m-3]+nums[m-2]+nums[m-1] < target {
			continue
		}
		for j := i + 1; j < m-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			k, l := j+1, m-1
			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				diff := target - sum
				if diff == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
					l--
					for k < l && nums[k] == nums[k-1] {
						k++
					}
					for k < l && nums[l] == nums[l+1] {
						l--
					}
				} else if diff > 0 {
					k++
				} else {
					l--
				}
			}
		}
	}
	return result
}
