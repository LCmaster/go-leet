package removeduplicatesfromsortedarray

import (
	"slices"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	// Pre-generate large dataset for constraint boundaries (n = 30000, -100 <= val <= 100)
	const maxN = 30000

	// 1. All 30,000 identical elements
	largeAllSame := make([]int, maxN)
	for i := range largeAllSame {
		largeAllSame[i] = 42
	}
	wantLargeAllSame := []int{42}

	// 2. 30,000 elements containing all 201 distinct values [-100..100] repeated ~149-150 times each
	largeAllValues := make([]int, maxN)
	wantLargeAllValues := make([]int, 201)
	for val := -100; val <= 100; val++ {
		wantLargeAllValues[val+100] = val
	}
	for i := range largeAllValues {
		// Map index evenly across the 201 values: -100 to 100
		largeAllValues[i] = -100 + (i * 201 / maxN)
	}

	tests := []struct {
		name string
		nums []int
		want []int
	}{
		// --- 1. LeetCode Official Examples ---
		{
			name: "Example 1: nums = [1,1,2]",
			nums: []int{1, 1, 2},
			want: []int{1, 2},
		},
		{
			name: "Example 2: nums = [0,0,1,1,1,2,2,3,3,4]",
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: []int{0, 1, 2, 3, 4},
		},

		// --- 2. Minimal Length Arrays (n = 1) ---
		{
			name: "Single element: zero",
			nums: []int{0},
			want: []int{0},
		},
		{
			name: "Single element: minimum constraint -100",
			nums: []int{-100},
			want: []int{-100},
		},
		{
			name: "Single element: maximum constraint 100",
			nums: []int{100},
			want: []int{100},
		},

		// --- 3. Two Elements (n = 2) ---
		{
			name: "Two distinct elements",
			nums: []int{1, 2},
			want: []int{1, 2},
		},
		{
			name: "Two duplicate elements",
			nums: []int{1, 1},
			want: []int{1},
		},
		{
			name: "Two negative duplicates",
			nums: []int{-50, -50},
			want: []int{-50},
		},
		{
			name: "Two distinct crossing zero: [-1, 1]",
			nums: []int{-1, 1},
			want: []int{-1, 1},
		},

		// --- 4. Already Unique Arrays (No Duplicates) ---
		{
			name: "All strictly increasing positive",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "All strictly increasing negative",
			nums: []int{-10, -8, -6, -4, -2},
			want: []int{-10, -8, -6, -4, -2},
		},
		{
			name: "Spanning negative to positive without duplicates",
			nums: []int{-3, -2, -1, 0, 1, 2, 3},
			want: []int{-3, -2, -1, 0, 1, 2, 3},
		},
		{
			name: "Extreme values unique: [-100, 0, 100]",
			nums: []int{-100, 0, 100},
			want: []int{-100, 0, 100},
		},

		// --- 5. All Identical Elements (Maximum Duplicate Density) ---
		{
			name: "All zeros: [0, 0, 0, 0, 0]",
			nums: []int{0, 0, 0, 0, 0},
			want: []int{0},
		},
		{
			name: "All negative duplicates: [-5, -5, -5, -5]",
			nums: []int{-5, -5, -5, -5},
			want: []int{-5},
		},
		{
			name: "All positive duplicates: [7, 7, 7, 7, 7, 7]",
			nums: []int{7, 7, 7, 7, 7, 7},
			want: []int{7},
		},

		// --- 6. Uniform Frequency Groups ---
		{
			name: "Pairs of duplicates: [1,1, 2,2, 3,3]",
			nums: []int{1, 1, 2, 2, 3, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "Triplets of duplicates: [1,1,1, 2,2,2, 3,3,3]",
			nums: []int{1, 1, 1, 2, 2, 2, 3, 3, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "Quadruplets with negatives and zero",
			nums: []int{-2, -2, -2, -2, 0, 0, 0, 0, 2, 2, 2, 2},
			want: []int{-2, 0, 2},
		},

		// --- 7. Duplicate Position Variations ---
		{
			name: "Many duplicates at head: [1,1,1,1,1, 2, 3]",
			nums: []int{1, 1, 1, 1, 1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "Many duplicates in middle: [1, 2, 2, 2, 2, 2, 3]",
			nums: []int{1, 2, 2, 2, 2, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "Many duplicates at tail: [1, 2, 3, 3, 3, 3, 3]",
			nums: []int{1, 2, 3, 3, 3, 3, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "Single unique element amidst long duplicate runs",
			nums: []int{-1, -1, -1, 0, 1, 1, 1},
			want: []int{-1, 0, 1},
		},
		{
			name: "Alternating singleton and duplicate runs",
			nums: []int{1, 2, 2, 3, 4, 4, 5, 6, 6},
			want: []int{1, 2, 3, 4, 5, 6},
		},

		// --- 8. Boundary Values (-100 and 100) ---
		{
			name: "Extremes with duplicates: [-100, -100, -100, 100, 100]",
			nums: []int{-100, -100, -100, 100, 100},
			want: []int{-100, 100},
		},
		{
			name: "Negative run, single zero, positive run",
			nums: []int{-100, -100, -50, -50, 0, 50, 50, 100, 100},
			want: []int{-100, -50, 0, 50, 100},
		},

		// --- 9. Large Scale / Constraint Limit Tests (n = 30000) ---
		{
			name: "Max constraint n = 30000: all identical elements",
			nums: largeAllSame,
			want: wantLargeAllSame,
		},
		{
			name: "Max constraint n = 30000: all 201 possible values [-100..100]",
			nums: largeAllValues,
			want: wantLargeAllValues,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := slices.Clone(tt.nums)
			k := removeDuplicates(nums)

			if k != len(tt.want) {
				t.Fatalf("removeDuplicates() returned k = %d; want %d", k, len(tt.want))
			}

			got := nums[:k]
			if !slices.Equal(got, tt.want) {
				t.Errorf("nums[:k] = %v; want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	const n = 30000
	nums := make([]int, n)
	for i := range nums {
		nums[i] = -100 + (i * 201 / n)
	}

	b.ResetTimer()
	for b.Loop() {
		testSlice := slices.Clone(nums)
		_ = removeDuplicates(testSlice)
	}
}
