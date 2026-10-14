package threesum

import (
	"cmp"
	"reflect"
	"slices"
	"testing"
)

// normalizeTriplets sorts each individual triplet, and then sorts the list of triplets
// lexicographically so test comparisons are independent of output order.
func normalizeTriplets(triplets [][]int) [][]int {
	if len(triplets) == 0 {
		return [][]int{}
	}
	result := make([][]int, len(triplets))
	for i, t := range triplets {
		cp := make([]int, len(t))
		copy(cp, t)
		slices.Sort(cp)
		result[i] = cp
	}
	slices.SortFunc(result, func(a, b []int) int {
		for i := 0; i < len(a) && i < len(b); i++ {
			if a[i] != b[i] {
				return cmp.Compare(a[i], b[i])
			}
		}
		return cmp.Compare(len(a), len(b))
	})
	return result
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		// --- LeetCode Official Examples ---
		{
			name: "Example 1 - standard mixed numbers with multiple triplets",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "Example 2 - no triplet sums to 0",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "Example 3 - all zeros",
			nums: []int{0, 0, 0},
			want: [][]int{
				{0, 0, 0},
			},
		},

		// --- Minimum Length Constraint (n = 3) ---
		{
			name: "Length 3 - sums to zero",
			nums: []int{-1, 0, 1},
			want: [][]int{
				{-1, 0, 1},
			},
		},
		{
			name: "Length 3 - does not sum to zero",
			nums: []int{1, 2, 3},
			want: [][]int{},
		},
		{
			name: "Length 3 - negative sum",
			nums: []int{-5, -2, 7},
			want: [][]int{
				{-5, -2, 7},
			},
		},

		// --- Zero Handling & Deduplication ---
		{
			name: "Four zeros - must return single [0, 0, 0]",
			nums: []int{0, 0, 0, 0},
			want: [][]int{
				{0, 0, 0},
			},
		},
		{
			name: "Many zeros - must return single [0, 0, 0]",
			nums: []int{0, 0, 0, 0, 0, 0},
			want: [][]int{
				{0, 0, 0},
			},
		},
		{
			name: "Single zero with opposites",
			nums: []int{-2, 0, 2},
			want: [][]int{
				{-2, 0, 2},
			},
		},

		// --- Strict Duplicate Suppression ---
		{
			name: "Duplicate pairs producing duplicate triplets",
			nums: []int{-1, -1, -1, 2, 2},
			want: [][]int{
				{-1, -1, 2},
			},
		},
		{
			name: "Duplicates around zero",
			nums: []int{-1, -1, 0, 1, 1},
			want: [][]int{
				{-1, 0, 1},
			},
		},
		{
			name: "Multiple duplicates of same values",
			nums: []int{-2, 0, 0, 2, 2},
			want: [][]int{
				{-2, 0, 2},
			},
		},
		{
			name: "Highly duplicated array",
			nums: []int{-1, -1, -1, -1, 2, 2, 2, 2},
			want: [][]int{
				{-1, -1, 2},
			},
		},

		// --- All Positive / All Negative (No Solution) ---
		{
			name: "All positive integers",
			nums: []int{1, 2, 3, 4, 5, 6},
			want: [][]int{},
		},
		{
			name: "All negative integers",
			nums: []int{-6, -5, -4, -3, -2, -1},
			want: [][]int{},
		},

		// --- Multiple Triplets ---
		{
			name: "Symmetric range from -3 to 3",
			nums: []int{-3, -2, -1, 0, 1, 2, 3},
			want: [][]int{
				{-3, 0, 3},
				{-3, 1, 2},
				{-2, -1, 3},
				{-2, 0, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "One negative element, multiple positive complements",
			nums: []int{-5, 1, 2, 3, 4},
			want: [][]int{
				{-5, 1, 4},
				{-5, 2, 3},
			},
		},
		{
			name: "One positive element, multiple negative complements",
			nums: []int{5, -1, -2, -3, -4},
			want: [][]int{
				{-4, -1, 5},
				{-3, -2, 5},
			},
		},
		{
			name: "Consecutive triplets with overlapping numbers",
			nums: []int{-4, -1, -1, 0, 1, 2},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},

		// --- Large / Boundary Numbers (-10^5 to 10^5) ---
		{
			name: "Large numbers cancellation",
			nums: []int{-100000, 50000, 50000},
			want: [][]int{
				{-100000, 50000, 50000},
			},
		},
		{
			name: "Max constraint with zero",
			nums: []int{100000, -100000, 0},
			want: [][]int{
				{-100000, 0, 100000},
			},
		},
		{
			name: "Large numbers multiple combinations",
			nums: []int{-100000, -50000, 0, 50000, 100000},
			want: [][]int{
				{-100000, 0, 100000},
				{-50000, 0, 50000},
			},
		},

		// --- Array with No Triplets Despite Pairs Summing to 0 ---
		{
			name: "Opposite pairs without a third number summing to 0",
			nums: []int{-3, 3, -2, 2, 10},
			want: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			normGot := normalizeTriplets(got)
			normWant := normalizeTriplets(tt.want)

			if !reflect.DeepEqual(normGot, normWant) {
				t.Errorf("threeSum(%v)\ngot:  %v\nwant: %v", tt.nums, normGot, normWant)
			}
		})
	}
}
