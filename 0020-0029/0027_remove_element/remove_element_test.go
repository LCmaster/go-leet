package removeelement

import (
	"slices"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	// Pre-generate arrays for maximum constraint tests (n = 100, 0 <= nums[i] <= 50, 0 <= val <= 100)
	const maxN = 100

	// 1. 100 elements all equal to val (k = 0)
	n100AllVal := make([]int, maxN)
	for i := range n100AllVal {
		n100AllVal[i] = 25
	}

	// 2. 100 elements none equal to val (k = 100)
	n100NoVal := make([]int, maxN)
	wantN100NoVal := make([]int, maxN)
	for i := range n100NoVal {
		n100NoVal[i] = i % 51 // values 0..50
		wantN100NoVal[i] = i % 51
	}
	slices.Sort(wantN100NoVal)

	// 3. 100 elements: 50 matching val, 50 non-matching
	n100HalfVal := make([]int, maxN)
	var wantN100HalfVal []int
	for i := range n100HalfVal {
		if i%2 == 0 {
			n100HalfVal[i] = 7 // val to remove
		} else {
			val := (i / 2) % 50
			if val == 7 {
				val = 8
			}
			n100HalfVal[i] = val
			wantN100HalfVal = append(wantN100HalfVal, val)
		}
	}
	slices.Sort(wantN100HalfVal)

	tests := []struct {
		name string
		nums []int
		val  int
		want []int
	}{
		// --- 1. LeetCode Official Examples ---
		{
			name: "Example 1: nums = [3,2,2,3], val = 3",
			nums: []int{3, 2, 2, 3},
			val:  3,
			want: []int{2, 2},
		},
		{
			name: "Example 2: nums = [0,1,2,2,3,0,4,2], val = 2",
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
			want: []int{0, 0, 1, 3, 4},
		},

		// --- 2. Empty Array (n = 0) ---
		{
			name: "Empty array with val = 0",
			nums: []int{},
			val:  0,
			want: []int{},
		},
		{
			name: "Empty array with val = 50",
			nums: []int{},
			val:  50,
			want: []int{},
		},

		// --- 3. Single Element (n = 1) ---
		{
			name: "Single element matching val: [1], val = 1",
			nums: []int{1},
			val:  1,
			want: []int{},
		},
		{
			name: "Single element non-matching val: [1], val = 2",
			nums: []int{1},
			val:  2,
			want: []int{1},
		},
		{
			name: "Single element zero matching: [0], val = 0",
			nums: []int{0},
			val:  0,
			want: []int{},
		},
		{
			name: "Single element zero non-matching: [0], val = 10",
			nums: []int{0},
			val:  10,
			want: []int{0},
		},
		{
			name: "Single element max constraint: [50], val = 50",
			nums: []int{50},
			val:  50,
			want: []int{},
		},

		// --- 4. Two Elements (n = 2) ---
		{
			name: "Two elements: both match val: [3, 3], val = 3",
			nums: []int{3, 3},
			val:  3,
			want: []int{},
		},
		{
			name: "Two elements: neither matches val: [1, 2], val = 3",
			nums: []int{1, 2},
			val:  3,
			want: []int{1, 2},
		},
		{
			name: "Two elements: first matches val: [3, 1], val = 3",
			nums: []int{3, 1},
			val:  3,
			want: []int{1},
		},
		{
			name: "Two elements: second matches val: [1, 3], val = 3",
			nums: []int{1, 3},
			val:  3,
			want: []int{1},
		},

		// --- 5. All Elements Equal to val ---
		{
			name: "All elements equal to val: [4, 4, 4, 4, 4], val = 4",
			nums: []int{4, 4, 4, 4, 4},
			val:  4,
			want: []int{},
		},
		{
			name: "All zeros matching val = 0: [0, 0, 0], val = 0",
			nums: []int{0, 0, 0},
			val:  0,
			want: []int{},
		},

		// --- 6. No Elements Equal to val ---
		{
			name: "No elements match: val is not present",
			nums: []int{1, 2, 3, 4, 5},
			val:  9,
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Val exceeds maximum element value: val = 100",
			nums: []int{0, 10, 25, 50},
			val:  100,
			want: []int{0, 10, 25, 50},
		},

		// --- 7. Target val Position Variations ---
		{
			name: "Val only at head: [5, 5, 1, 2, 3], val = 5",
			nums: []int{5, 5, 1, 2, 3},
			val:  5,
			want: []int{1, 2, 3},
		},
		{
			name: "Val only in middle: [1, 2, 5, 5, 3, 4], val = 5",
			nums: []int{1, 2, 5, 5, 3, 4},
			val:  5,
			want: []int{1, 2, 3, 4},
		},
		{
			name: "Val only at tail: [1, 2, 3, 5, 5], val = 5",
			nums: []int{1, 2, 3, 5, 5},
			val:  5,
			want: []int{1, 2, 3},
		},
		{
			name: "Alternating val and non-val: [2, 1, 2, 1, 2, 1], val = 2",
			nums: []int{2, 1, 2, 1, 2, 1},
			val:  2,
			want: []int{1, 1, 1},
		},
		{
			name: "Alternating val and non-val: [1, 2, 1, 2, 1, 2], val = 1",
			nums: []int{1, 2, 1, 2, 1, 2},
			val:  1,
			want: []int{2, 2, 2},
		},

		// --- 8. Repeated Non-val Elements ---
		{
			name: "Multiple duplicates of non-val elements",
			nums: []int{4, 4, 2, 4, 4, 2, 4},
			val:  2,
			want: []int{4, 4, 4, 4, 4},
		},
		{
			name: "Mixed duplicate values: [1, 1, 2, 2, 3, 3], val = 2",
			nums: []int{1, 1, 2, 2, 3, 3},
			val:  2,
			want: []int{1, 1, 3, 3},
		},

		// --- 9. Boundary Constraints (0 <= nums[i] <= 50, 0 <= val <= 100) ---
		{
			name: "Boundary extremes: 0 and 50 with val = 0",
			nums: []int{0, 50, 0, 50, 0},
			val:  0,
			want: []int{50, 50},
		},
		{
			name: "Boundary extremes: 0 and 50 with val = 50",
			nums: []int{0, 50, 0, 50, 0},
			val:  50,
			want: []int{0, 0, 0},
		},
		{
			name: "Val = 0 with multiple zeros and non-zeros",
			nums: []int{0, 1, 0, 2, 0, 3},
			val:  0,
			want: []int{1, 2, 3},
		},

		// --- 10. Maximum Scale Constraints (n = 100) ---
		{
			name: "Max constraint n = 100: all matching val",
			nums: n100AllVal,
			val:  25,
			want: []int{},
		},
		{
			name: "Max constraint n = 100: none matching val",
			nums: n100NoVal,
			val:  99,
			want: wantN100NoVal,
		},
		{
			name: "Max constraint n = 100: half matching val",
			nums: n100HalfVal,
			val:  7,
			want: wantN100HalfVal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := slices.Clone(tt.nums)
			k := removeElement(nums, tt.val)

			if k != len(tt.want) {
				t.Fatalf("removeElement() returned k = %d; want %d", k, len(tt.want))
			}

			// LeetCode allows elements to be returned in any order.
			// The custom judge sorts the first k elements before comparing.
			got := slices.Clone(nums[:k])
			slices.Sort(got)

			wantSorted := slices.Clone(tt.want)
			slices.Sort(wantSorted)

			if !slices.Equal(got, wantSorted) {
				t.Errorf("sorted nums[:k] = %v; want %v", got, wantSorted)
			}

			// Sanity check: ensure none of the first k elements equal val
			for i, v := range nums[:k] {
				if v == tt.val {
					t.Errorf("nums[%d] = %d equals removed val %d", i, v, tt.val)
				}
			}
		})
	}
}

func BenchmarkRemoveElement(b *testing.B) {
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = i % 10
	}

	b.ResetTimer()
	for b.Loop() {
		testSlice := slices.Clone(nums)
		_ = removeElement(testSlice, 3)
	}
}
