package threesumclosest

import (
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		// --- 1. Official LeetCode Examples ---
		{
			name:   "Example 1: mixed positive and negative",
			nums:   []int{-1, 2, 1, -4},
			target: 1,
			want:   2, // (-1 + 2 + 1 = 2, distance |1 - 2| = 1)
		},
		{
			name:   "Example 2: all zeros",
			nums:   []int{0, 0, 0},
			target: 1,
			want:   0, // (0 + 0 + 0 = 0, distance |1 - 0| = 1)
		},

		// --- 2. Minimum Array Length (n = 3) ---
		{
			name:   "Minimum length 3 with positive integers",
			nums:   []int{1, 2, 3},
			target: 100,
			want:   6, // only one possible sum: 1 + 2 + 3 = 6
		},
		{
			name:   "Minimum length 3 with negative integers",
			nums:   []int{-5, -3, -2},
			target: -1,
			want:   -10, // -5 + -3 + -2 = -10
		},
		{
			name:   "Minimum length 3 with zeros and negative",
			nums:   []int{-1, 0, 0},
			target: 5,
			want:   -1,
		},

		// --- 3. Exact Sum Matches (diff == 0) ---
		{
			name:   "Exact match at zero",
			nums:   []int{-1, 0, 1, 2},
			target: 0,
			want:   0, // -1 + 0 + 1 = 0
		},
		{
			name:   "Exact match positive target",
			nums:   []int{1, 2, 3, 4},
			target: 6,
			want:   6, // 1 + 2 + 3 = 6
		},
		{
			name:   "Exact match negative target",
			nums:   []int{-4, -2, 1, 5},
			target: -5,
			want:   -5, // -4 + -2 + 1 = -5
		},
		{
			name:   "Exact match with multiple duplicate values",
			nums:   []int{4, 0, 5, -5, 3, 3, 0, -4, -5},
			target: -2,
			want:   -2, // -5 + 0 + 3 = -2
		},

		// --- 4. Target Outside Array Possible Sum Range ---
		{
			name:   "Target far greater than maximum possible sum",
			nums:   []int{1, 2, 3, 4, 5},
			target: 1000,
			want:   12, // max triplet is 3 + 4 + 5 = 12
		},
		{
			name:   "Target far smaller than minimum possible sum",
			nums:   []int{1, 2, 3, 4, 5},
			target: -500,
			want:   6, // min triplet is 1 + 2 + 3 = 6
		},
		{
			name:   "All negative numbers with positive target",
			nums:   []int{-10, -20, -30, -40},
			target: 100,
			want:   -60, // closest is -10 + -20 + -30 = -60
		},

		// --- 5. Negative Targets ---
		{
			name:   "Negative target closest from below",
			nums:   []int{-10, -5, 0, 5, 10},
			target: -2,
			want:   0, // -5 + 0 + 5 = 0 (| -2 - 0 | = 2 vs -5 which has diff 3)
		},
		{
			name:   "Negative target with all negative elements",
			nums:   []int{-8, -7, -3, -2, -1},
			target: -7,
			want:   -6, // -3 + -2 + -1 = -6 (| -7 - -6 | = 1)
		},

		// --- 6. Duplicates Handling ---
		{
			name:   "Array with many identical numbers",
			nums:   []int{1, 1, 1, 5, 5, 5},
			target: 8,
			want:   7, // 1 + 1 + 5 = 7 (distance 1 to 8)
		},
		{
			name:   "Pairs of duplicates with target between them",
			nums:   []int{0, 2, 1, -3},
			target: 1,
			want:   0, // -3 + 1 + 2 = 0 (distance 1 to 1)
		},
		{
			name:   "Repeated zeros with non-zero target",
			nums:   []int{0, 0, 0, 0, 0},
			target: 5,
			want:   0,
		},

		// --- 7. Boundary & Constraint Values (-1000 to 1000, target -10^4 to 10^4) ---
		{
			name:   "Maximum allowed positive numbers",
			nums:   []int{1000, 1000, 1000},
			target: 10000,
			want:   3000,
		},
		{
			name:   "Maximum allowed negative numbers",
			nums:   []int{-1000, -1000, -1000},
			target: -10000,
			want:   -3000,
		},
		{
			name:   "Opposite extreme boundaries canceling near zero",
			nums:   []int{-1000, 0, 1000},
			target: 1,
			want:   0,
		},
		{
			name:   "Wide range spanning -1000 to 1000",
			nums:   []int{-1000, -500, 0, 200, 800, 1000},
			target: 650,
			want:   700, // -500 + 200 + 1000 = 700 (|650 - 700| = 50)
		},

		// --- 8. Powers of Two / Asymmetric Distributions ---
		{
			name:   "Powers of two with exact match",
			nums:   []int{1, 2, 4, 8, 16, 32, 64, 128},
			target: 82,
			want:   82, // 2 + 16 + 64 = 82
		},
		{
			name:   "Powers of two closest match",
			nums:   []int{1, 2, 4, 8, 16, 32, 64, 128},
			target: 83,
			want:   82, // 2 + 16 + 64 = 82 (distance 1)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSumClosest(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("threeSumClosest(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func BenchmarkThreeSumClosest(b *testing.B) {
	nums := []int{-1, 2, 1, -4, 5, -10, 8, 12, -7, 3, 0, -2, 9, 14, -6}
	target := 7
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = threeSumClosest(nums, target)
	}
}
