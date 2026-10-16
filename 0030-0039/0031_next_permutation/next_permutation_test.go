package nextpermutation

import (
	"slices"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	// Pre-build 100-element boundary arrays (N = 100 is max constraint)
	largeAscending := make([]int, 100)
	wantLargeAscending := make([]int, 100)
	for i := 0; i < 100; i++ {
		largeAscending[i] = i
		wantLargeAscending[i] = i
	}
	wantLargeAscending[98], wantLargeAscending[99] = 99, 98

	largeDescending := make([]int, 100)
	wantLargeDescending := make([]int, 100)
	for i := 0; i < 100; i++ {
		largeDescending[i] = 99 - i
		wantLargeDescending[i] = i
	}

	tests := []struct {
		name string
		nums []int
		want []int
	}{
		// --- 1. LeetCode Official Examples ---
		{
			name: "Example 1: [1, 2, 3] -> [1, 3, 2]",
			nums: []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			name: "Example 2: [3, 2, 1] -> [1, 2, 3] (wrap-around to lowest)",
			nums: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		{
			name: "Example 3: [1, 1, 5] -> [1, 5, 1]",
			nums: []int{1, 1, 5},
			want: []int{1, 5, 1},
		},

		// --- 2. Minimal Length Boundary Cases ---
		{
			name: "Single element: [1] -> [1]",
			nums: []int{1},
			want: []int{1},
		},
		{
			name: "Single zero element: [0] -> [0]",
			nums: []int{0},
			want: []int{0},
		},
		{
			name: "Two elements ascending: [1, 2] -> [2, 1]",
			nums: []int{1, 2},
			want: []int{2, 1},
		},
		{
			name: "Two elements descending: [2, 1] -> [1, 2]",
			nums: []int{2, 1},
			want: []int{1, 2},
		},
		{
			name: "Two identical elements: [1, 1] -> [1, 1]",
			nums: []int{1, 1},
			want: []int{1, 1},
		},

		// --- 3. All Elements Identical ---
		{
			name: "Three identical elements: [2, 2, 2] -> [2, 2, 2]",
			nums: []int{2, 2, 2},
			want: []int{2, 2, 2},
		},
		{
			name: "Five identical elements: [7, 7, 7, 7, 7] -> [7, 7, 7, 7, 7]",
			nums: []int{7, 7, 7, 7, 7},
			want: []int{7, 7, 7, 7, 7},
		},

		// --- 4. Complete Permutation Cycle (3 elements) ---
		{
			name: "Cycle step 1: [1, 2, 3] -> [1, 3, 2]",
			nums: []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			name: "Cycle step 2: [1, 3, 2] -> [2, 1, 3]",
			nums: []int{1, 3, 2},
			want: []int{2, 1, 3},
		},
		{
			name: "Cycle step 3: [2, 1, 3] -> [2, 3, 1]",
			nums: []int{2, 1, 3},
			want: []int{2, 3, 1},
		},
		{
			name: "Cycle step 4: [2, 3, 1] -> [3, 1, 2]",
			nums: []int{2, 3, 1},
			want: []int{3, 1, 2},
		},
		{
			name: "Cycle step 5: [3, 1, 2] -> [3, 2, 1]",
			nums: []int{3, 1, 2},
			want: []int{3, 2, 1},
		},
		{
			name: "Cycle step 6: [3, 2, 1] -> [1, 2, 3]",
			nums: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},

		// --- 5. Pivot at Different Positions ---
		{
			name: "Pivot at second-to-last index: [1, 2, 4, 3] -> [1, 3, 2, 4]",
			nums: []int{1, 2, 4, 3},
			want: []int{1, 3, 2, 4},
		},
		{
			name: "Pivot at index 0: [1, 4, 3, 2] -> [2, 1, 3, 4]",
			nums: []int{1, 4, 3, 2},
			want: []int{2, 1, 3, 4},
		},
		{
			name: "Pivot in middle: [1, 3, 5, 4, 2] -> [1, 4, 2, 3, 5]",
			nums: []int{1, 3, 5, 4, 2},
			want: []int{1, 4, 2, 3, 5},
		},
		{
			name: "Pivot before long descending tail: [2, 4, 6, 5, 3, 1] -> [2, 5, 1, 3, 4, 6]",
			nums: []int{2, 4, 6, 5, 3, 1},
			want: []int{2, 5, 1, 3, 4, 6},
		},

		// --- 6. Duplicate Elements & Plateaus ---
		{
			name: "Duplicates at suffix: [1, 5, 1] -> [5, 1, 1]",
			nums: []int{1, 5, 1},
			want: []int{5, 1, 1},
		},
		{
			name: "Decreasing with duplicates: [5, 1, 1] -> [1, 1, 5]",
			nums: []int{5, 1, 1},
			want: []int{1, 1, 5},
		},
		{
			name: "Duplicate candidate for swap: [1, 3, 2, 2] -> [2, 1, 2, 3]",
			nums: []int{1, 3, 2, 2},
			want: []int{2, 1, 2, 3},
		},
		{
			name: "Multiple duplicates in tail: [2, 3, 1, 3, 3] -> [2, 3, 3, 1, 3]",
			nums: []int{2, 3, 1, 3, 3},
			want: []int{2, 3, 3, 1, 3},
		},
		{
			name: "Plateau at start: [3, 3, 1, 2] -> [3, 3, 2, 1]",
			nums: []int{3, 3, 1, 2},
			want: []int{3, 3, 2, 1},
		},
		{
			name: "Large plateau after single lower element: [1, 5, 5, 5, 5] -> [5, 1, 5, 5, 5]",
			nums: []int{1, 5, 5, 5, 5},
			want: []int{5, 1, 5, 5, 5},
		},
		{
			name: "Complex duplicates: [2, 2, 7, 5, 4, 3, 2, 2, 1] -> [2, 3, 1, 2, 2, 2, 4, 5, 7]",
			nums: []int{2, 2, 7, 5, 4, 3, 2, 2, 1},
			want: []int{2, 3, 1, 2, 2, 2, 4, 5, 7},
		},

		// --- 7. Strictly Monotonic Sequences ---
		{
			name: "Strictly increasing: [1, 2, 3, 4, 5] -> [1, 2, 3, 5, 4]",
			nums: []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 5, 4},
		},
		{
			name: "Strictly decreasing: [5, 4, 3, 2, 1] -> [1, 2, 3, 4, 5]",
			nums: []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Decreasing with zero: [4, 3, 2, 1, 0] -> [0, 1, 2, 3, 4]",
			nums: []int{4, 3, 2, 1, 0},
			want: []int{0, 1, 2, 3, 4},
		},

		// --- 8. Boundary Values (0 and 100 per constraints) ---
		{
			name: "Constraint boundary values: [0, 100, 0] -> [100, 0, 0]",
			nums: []int{0, 100, 0},
			want: []int{100, 0, 0},
		},
		{
			name: "Constraint boundary values: [100, 0, 100] -> [100, 100, 0]",
			nums: []int{100, 0, 100},
			want: []int{100, 100, 0},
		},
		{
			name: "Constraint boundary values wrap: [100, 100, 0] -> [0, 100, 100]",
			nums: []int{100, 100, 0},
			want: []int{0, 100, 100},
		},

		// --- 9. Large Arrays (Constraint limit: len == 100) ---
		{
			name: "Large ascending (len 100)",
			nums: largeAscending,
			want: wantLargeAscending,
		},
		{
			name: "Large descending wraps to ascending (len 100)",
			nums: largeDescending,
			want: wantLargeDescending,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clone input slice since nextPermutation mutates in-place
			got := slices.Clone(tt.nums)
			nextPermutation(got)

			if !slices.Equal(got, tt.want) {
				t.Errorf("nextPermutation(%v) = %v; want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func BenchmarkNextPermutation(b *testing.B) {
	template := []int{2, 4, 6, 8, 10, 9, 7, 5, 3, 1}
	arr := make([]int, len(template))

	b.ResetTimer()
	for b.Loop() {
		copy(arr, template)
		nextPermutation(arr)
	}
}
