package containerwithmostwater

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		// --- LeetCode Official Examples ---
		{
			name:   "Example 1 - standard case with varying heights",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "Example 2 - minimum length array with equal heights",
			height: []int{1, 1},
			want:   1,
		},

		// --- Minimum Length Edge Cases (n = 2) ---
		{
			name:   "Two elements identical height",
			height: []int{5, 5},
			want:   5,
		},
		{
			name:   "Two elements, left much taller",
			height: []int{100, 2},
			want:   2,
		},
		{
			name:   "Two elements, right much taller",
			height: []int{2, 100},
			want:   2,
		},

		// --- Three Elements Cases ---
		{
			name:   "Three elements with peak in middle",
			height: []int{1, 10, 1},
			want:   2,
		},
		{
			name:   "Three elements with trough in middle (endpoints win)",
			height: []int{10, 1, 10},
			want:   20,
		},

		// --- Monotonic Arrays ---
		{
			name:   "Strictly increasing heights",
			height: []int{1, 2, 3, 4, 5},
			want:   6,
		},
		{
			name:   "Strictly decreasing heights",
			height: []int{5, 4, 3, 2, 1},
			want:   6,
		},
		{
			name:   "Step staircase pattern",
			height: []int{2, 3, 4, 5, 18, 17, 6},
			want:   17,
		},

		// --- Uniform / Flat Heights ---
		{
			name:   "All elements identical",
			height: []int{4, 4, 4, 4, 4},
			want:   16,
		},
		{
			name:   "Large uniform array",
			height: []int{10, 10, 10, 10, 10, 10, 10},
			want:   60,
		},

		// --- Zeros in Array ---
		{
			name:   "All zeros",
			height: []int{0, 0, 0, 0},
			want:   0,
		},
		{
			name:   "Zeroes at endpoints with tall elements inside",
			height: []int{0, 5, 5, 0},
			want:   5,
		},
		{
			name:   "Zero in the middle",
			height: []int{5, 0, 5},
			want:   10,
		},
		{
			name:   "Alternating zeros and heights",
			height: []int{0, 2, 0, 3, 0, 2, 0},
			want:   8,
		},

		// --- Symmetry & Shape Variations ---
		{
			name:   "Pyramid shape (symmetric peak)",
			height: []int{1, 2, 3, 4, 5, 4, 3, 2, 1},
			want:   12,
		},
		{
			name:   "Tall walls at outermost endpoints",
			height: []int{100, 1, 1, 1, 100},
			want:   400,
		},
		{
			name:   "Very tall walls adjacent in middle",
			height: []int{1, 100, 100, 1},
			want:   100,
		},
		{
			name:   "Wide distance with low height beats narrow with high height",
			height: []int{2, 1, 1, 1, 1, 1, 1, 2},
			want:   14,
		},

		// --- Maximum Constraint Heights (up to 10^4) ---
		{
			name:   "Max height values at minimum length",
			height: []int{10000, 10000},
			want:   10000,
		},
		{
			name:   "Max height values separated by zeros",
			height: []int{10000, 0, 0, 10000},
			want:   30000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.height)
			if got != tt.want {
				t.Errorf("maxArea(%v) = %d, want %d", tt.height, got, tt.want)
			}
		})
	}
}
