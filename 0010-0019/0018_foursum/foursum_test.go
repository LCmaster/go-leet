package foursum

import (
	"cmp"
	"reflect"
	"slices"
	"testing"
)

// normalizeQuadruplets sorts each individual quadruplet and then sorts the list
// of quadruplets lexicographically, ensuring test comparisons are order-independent
// (LeetCode specifies: "You may return the answer in any order").
func normalizeQuadruplets(quads [][]int) [][]int {
	if len(quads) == 0 {
		return [][]int{}
	}
	result := make([][]int, len(quads))
	for i, q := range quads {
		cp := make([]int, len(q))
		copy(cp, q)
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

func TestFourSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   [][]int
	}{
		// --- 1. Official LeetCode Examples ---
		{
			name:   "Example 1: standard mixed numbers with multiple quadruplets",
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			want: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			name:   "Example 2: repeated identical numbers",
			nums:   []int{2, 2, 2, 2, 2},
			target: 8,
			want: [][]int{
				{2, 2, 2, 2},
			},
		},

		// --- 2. Minimum Array Length & Base Cases (len < 4 and len == 4) ---
		{
			name:   "Empty array",
			nums:   []int{},
			target: 0,
			want:   [][]int{},
		},
		{
			name:   "Length 1 array",
			nums:   []int{1},
			target: 1,
			want:   [][]int{},
		},
		{
			name:   "Length 3 array (fewer than 4 elements)",
			nums:   []int{1, 2, 3},
			target: 6,
			want:   [][]int{},
		},
		{
			name:   "Length 4 exact match",
			nums:   []int{1, 2, 3, 4},
			target: 10,
			want: [][]int{
				{1, 2, 3, 4},
			},
		},
		{
			name:   "Length 4 no match",
			nums:   []int{1, 2, 3, 4},
			target: 11,
			want:   [][]int{},
		},

		// --- 3. Zero Handling & Duplicates ---
		{
			name:   "Exactly four zeros summing to zero",
			nums:   []int{0, 0, 0, 0},
			target: 0,
			want: [][]int{
				{0, 0, 0, 0},
			},
		},
		{
			name:   "Many zeros with target zero (must return single [0,0,0,0])",
			nums:   []int{0, 0, 0, 0, 0, 0, 0},
			target: 0,
			want: [][]int{
				{0, 0, 0, 0},
			},
		},
		{
			name:   "Four zeros with non-zero target",
			nums:   []int{0, 0, 0, 0},
			target: 1,
			want:   [][]int{},
		},

		// --- 4. Strict Duplicate Suppression ---
		{
			name:   "Heavy duplicate numbers with single quadruplet solution",
			nums:   []int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3},
			target: 7,
			want: [][]int{
				{1, 1, 2, 3},
				{1, 2, 2, 2},
			},
		},
		{
			name:   "Pairs of positive and negative duplicates around zero",
			nums:   []int{-2, -2, -1, -1, 0, 0, 1, 1, 2, 2},
			target: 0,
			want: [][]int{
				{-2, -2, 2, 2},
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-2, 0, 1, 1},
				{-1, -1, 0, 2},
				{-1, -1, 1, 1},
				{-1, 0, 0, 1},
			},
		},

		// --- 5. All Positive / All Negative ---
		{
			name:   "All positive numbers with valid quadruplet",
			nums:   []int{1, 2, 3, 4, 5, 6},
			target: 10,
			want: [][]int{
				{1, 2, 3, 4},
			},
		},
		{
			name:   "All positive numbers with target too small",
			nums:   []int{1, 2, 3, 4, 5},
			target: 5,
			want:   [][]int{},
		},
		{
			name:   "All negative numbers with valid quadruplet",
			nums:   []int{-6, -5, -4, -3, -2, -1},
			target: -14,
			want: [][]int{
				{-6, -5, -2, -1},
				{-6, -4, -3, -1},
				{-5, -4, -3, -2},
			},
		},
		{
			name:   "All negative numbers with target too large",
			nums:   []int{-4, -3, -2, -1},
			target: 0,
			want:   [][]int{},
		},

		// --- 6. Negative Targets & Asymmetric Distributions ---
		{
			name:   "Negative target with mixed signs",
			nums:   []int{1, -2, -5, -4, -3, 3, 3, 5},
			target: -11,
			want: [][]int{
				{-5, -4, -3, 1},
			},
		},
		{
			name:   "Multiple quadruplets with negative target",
			nums:   []int{-3, -1, 0, 2, 4, 5},
			target: 0,
			want: [][]int{
				{-3, -1, 0, 4},
			},
		},

		// --- 7. Large Numbers & 32-bit Integer Overflow Prevention ---
		// In 32-bit arithmetic, 10^9 * 4 wraps around into negative values.
		{
			name:   "Extreme values summing to zero",
			nums:   []int{-1000000000, -1000000000, 1000000000, 1000000000},
			target: 0,
			want: [][]int{
				{-1000000000, -1000000000, 1000000000, 1000000000},
			},
		},
		{
			name:   "Famous LeetCode overflow testcase (sum = 4*10^9, target = -294967296)",
			nums:   []int{1000000000, 1000000000, 1000000000, 1000000000},
			target: -294967296,
			want:   [][]int{}, // In 32-bit integer, 4*10^9 overflows to -294967296; must NOT match
		},
		{
			name:   "Negative overflow prevention (sum = -4*10^9, target = 294967296)",
			nums:   []int{-1000000000, -1000000000, -1000000000, -1000000000},
			target: 294967296,
			want:   [][]int{},
		},

		// --- 8. Wide Range with Multiple Combinations ---
		{
			name:   "Symmetric range -3 to 3",
			nums:   []int{-3, -2, -1, 0, 1, 2, 3},
			target: 0,
			want: [][]int{
				{-3, -2, 2, 3},
				{-3, -1, 1, 3},
				{-3, 0, 1, 2},
				{-2, -1, 0, 3},
				{-2, -1, 1, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fourSum(tt.nums, tt.target)
			normGot := normalizeQuadruplets(got)
			normWant := normalizeQuadruplets(tt.want)

			if len(normWant) == 0 {
				if len(normGot) != 0 {
					t.Errorf("fourSum(%v, %d) = %v, want empty result", tt.nums, tt.target, normGot)
				}
				return
			}

			if !reflect.DeepEqual(normGot, normWant) {
				t.Errorf("fourSum(%v, %d)\ngot:  %v\nwant: %v", tt.nums, tt.target, normGot, normWant)
			}
		})
	}
}

func BenchmarkFourSum(b *testing.B) {
	nums := []int{-3, -2, -1, 0, 0, 1, 2, 3, 4, -4, 5, -5, 2, -2, 1, -1}
	target := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fourSum(nums, target)
	}
}
