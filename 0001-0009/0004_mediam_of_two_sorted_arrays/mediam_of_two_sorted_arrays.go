package mediamoftwosortedarrays

import "math"

// 4. Median of Two Sorted Arrays
//
// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
//
// The overall run time complexity should be O(log (m+n)).
//
// Examples:
// - Input: nums1 = [1,3], nums2 = [2]   | Output: 2.00000 | Explanation: merged array = [1,2,3] and median is 2.
// - Input: nums1 = [1,2], nums2 = [3,4] | Output: 2.50000 | Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
//
// Constraints:
// - nums1.length == m
// - nums2.length == n
// - 0 <= m <= 1000
// - 0 <= n <= 1000
// - 1 <= m + n <= 2000
// - -10^6 <= nums1[i], nums2[i] <= 10^6

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	low, high := 0, m

	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (m+n+1)/2 - partitionX

		maxLeftX := math.MinInt64
		if partitionX > 0 {
			maxLeftX = nums1[partitionX-1]
		}

		minRightX := math.MaxInt64
		if partitionX < m {
			minRightX = nums1[partitionX]
		}

		maxLeftY := math.MinInt64
		if partitionY > 0 {
			maxLeftY = nums2[partitionY-1]
		}

		minRightY := math.MaxInt64
		if partitionY < n {
			minRightY = nums2[partitionY]
		}

		// Check if we found the correct partition
		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (m+n)%2 == 0 {
				leftMax := max(maxLeftX, maxLeftY)
				rightMin := min(minRightX, minRightY)
				return float64(leftMax+rightMin) / 2.0
			} else {
				return float64(max(maxLeftX, maxLeftY))
			}
		} else if maxLeftX > minRightY {
			// We are too far on the right side for partitionX, move left
			high = partitionX - 1
		} else {
			// We are too far on the left side for partitionX, move right
			low = partitionX + 1
		}
	}

	return 0.0
}
