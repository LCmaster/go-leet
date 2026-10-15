package mergeksortedlists

import (
	"slices"
	"testing"
)

// Helper function to convert a slice of ints into a linked list.
func sliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, val := range nums {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}
	return dummy.Next
}

// Helper function to convert a 2D slice of ints into a slice of linked lists.
func slicesToLists(lists [][]int) []*ListNode {
	if lists == nil {
		return nil
	}
	res := make([]*ListNode, len(lists))
	for i, l := range lists {
		res[i] = sliceToList(l)
	}
	return res
}

// Helper function to convert a linked list into a slice of ints.
func listToSlice(head *ListNode) []int {
	var nums []int
	for curr := head; curr != nil; curr = curr.Next {
		nums = append(nums, curr.Val)
	}
	return nums
}

func TestMergeKLists(t *testing.T) {
	// Pre-generate data for large scale tests
	// 1. 500 lists of 1 element each (k = 500, total = 500 nodes)
	k500Single := make([][]int, 500)
	wantK500Single := make([]int, 500)
	for i := 0; i < 500; i++ {
		// interleaved values: e.g. 500 - i to ensure diverse insertion
		val := i - 250 // spans negative to positive: -250 to 249
		k500Single[i] = []int{val}
		wantK500Single[i] = val
	}
	slices.Sort(wantK500Single)

	// 2. 10 lists of 500 elements each (k = 10, total = 5,000 nodes, max lists[i].length = 500)
	k10Long := make([][]int, 10)
	wantK10Long := make([]int, 5000)
	for i := 0; i < 10; i++ {
		k10Long[i] = make([]int, 500)
		for j := 0; j < 500; j++ {
			val := j*10 + i // strictly interleaved
			k10Long[i][j] = val
			wantK10Long[j*10+i] = val
		}
	}

	// 3. 500 empty lists
	k500Empty := make([][]int, 500)
	for i := range k500Empty {
		k500Empty[i] = []int{}
	}

	tests := []struct {
		name  string
		lists [][]int
		want  []int
	}{
		// --- 1. Official LeetCode Examples ---
		{
			name:  "Example 1: 3 interleaved lists with duplicates",
			lists: [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			want:  []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			name:  "Example 2: Empty outer array (k = 0)",
			lists: [][]int{},
			want:  []int{},
		},
		{
			name:  "Example 3: Single empty list (k = 1, lists[0] = nil)",
			lists: [][]int{{}},
			want:  []int{},
		},

		// --- 2. Nil and Empty List Edge Cases ---
		{
			name:  "Multiple empty lists (k = 3)",
			lists: [][]int{{}, {}, {}},
			want:  []int{},
		},
		{
			name:  "Empty list at start of array",
			lists: [][]int{{}, {1, 3}, {2, 4}},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "Empty list in middle of array",
			lists: [][]int{{1, 3}, {}, {2, 4}},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "Empty list at end of array",
			lists: [][]int{{1, 3}, {2, 4}, {}},
			want:  []int{1, 2, 3, 4},
		},
		{
			name:  "Multiple empty lists interspersed between valid lists",
			lists: [][]int{{}, {5}, {}, {1, 2}, {}, {3, 4}, {}},
			want:  []int{1, 2, 3, 4, 5},
		},

		// --- 3. Single List (k = 1) ---
		{
			name:  "Single list with 1 element",
			lists: [][]int{{42}},
			want:  []int{42},
		},
		{
			name:  "Single list with multiple elements",
			lists: [][]int{{1, 2, 3, 4, 5}},
			want:  []int{1, 2, 3, 4, 5},
		},

		// --- 4. Two Lists (k = 2) ---
		{
			name:  "Two lists of equal length",
			lists: [][]int{{1, 2, 4}, {1, 3, 4}},
			want:  []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:  "Two lists: single elements reversed order",
			lists: [][]int{{10}, {5}},
			want:  []int{5, 10},
		},

		// --- 5. Many Lists with Single Elements ---
		{
			name:  "5 lists of 1 element each (scrambled order)",
			lists: [][]int{{5}, {3}, {1}, {4}, {2}},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "6 lists of 1 element each with duplicates",
			lists: [][]int{{3}, {1}, {3}, {2}, {1}, {2}},
			want:  []int{1, 1, 2, 2, 3, 3},
		},

		// --- 6. Disjoint / Non-Overlapping Ranges ---
		{
			name:  "Sequential non-overlapping lists",
			lists: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:  "Reverse non-overlapping lists (each list strictly smaller than preceding)",
			lists: [][]int{{7, 8}, {5, 6}, {3, 4}, {1, 2}},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8},
		},

		// --- 7. Highly Asymmetric Lengths ---
		{
			name: "One long list with several single-element lists",
			lists: [][]int{
				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				{0},
				{11},
				{5},
			},
			want: []int{0, 1, 2, 3, 4, 5, 5, 6, 7, 8, 9, 10, 11},
		},
		{
			name: "Various lengths: 1, 4, 2, 3",
			lists: [][]int{
				{10},
				{1, 5, 9, 12},
				{2, 6},
				{3, 4, 8},
			},
			want: []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12},
		},

		// --- 8. Duplicates and Identical Lists ---
		{
			name:  "3 identical lists",
			lists: [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}},
			want:  []int{1, 1, 1, 2, 2, 2, 3, 3, 3},
		},
		{
			name:  "All elements identical across all lists",
			lists: [][]int{{7, 7}, {7, 7, 7}, {7}},
			want:  []int{7, 7, 7, 7, 7, 7},
		},
		{
			name:  "All zeros in all lists",
			lists: [][]int{{0, 0}, {0}, {0, 0, 0}},
			want:  []int{0, 0, 0, 0, 0, 0},
		},

		// --- 9. Negative Numbers & Boundary Values (-10^4 <= val <= 10^4) ---
		{
			name:  "All negative numbers across multiple lists",
			lists: [][]int{{-10, -5}, {-8, -2}, {-7, -1}},
			want:  []int{-10, -8, -7, -5, -2, -1},
		},
		{
			name: "Negative, zero, and positive crossing extremes",
			lists: [][]int{
				{-100, 0, 100},
				{-50, 0, 50},
				{-10000, 10000},
			},
			want: []int{-10000, -100, -50, 0, 0, 50, 100, 10000},
		},
		{
			name:  "Extreme constraint boundaries (-10000, 0, 10000)",
			lists: [][]int{{-10000}, {10000}, {0}},
			want:  []int{-10000, 0, 10000},
		},

		// --- 10. Large Scale / Stress Testing ---
		{
			name:  "500 lists of 1 element each (k = 500, range -250..249)",
			lists: k500Single,
			want:  wantK500Single,
		},
		{
			name:  "10 lists of 500 elements each (5000 total nodes)",
			lists: k10Long,
			want:  wantK10Long,
		},
		{
			name:  "500 empty lists",
			lists: k500Empty,
			want:  []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputLists := slicesToLists(tt.lists)
			resultHead := mergeKLists(inputLists)
			got := listToSlice(resultHead)

			if len(got) == 0 && len(tt.want) == 0 {
				return
			}

			if !slices.Equal(got, tt.want) {
				if len(got) <= 20 {
					t.Errorf("mergeKLists(%v) = %v, want %v", tt.lists, got, tt.want)
				} else {
					t.Errorf("mergeKLists(%s) returned len=%d, want len=%d; first 10 got=%v, want=%v",
						tt.name, len(got), len(tt.want), got[:10], tt.want[:10])
				}
			}
		})
	}
}

func BenchmarkMergeKLists_Small(b *testing.B) {
	raw := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := slicesToLists(raw)
		_ = mergeKLists(input)
	}
}

func BenchmarkMergeKLists_500x1(b *testing.B) {
	raw := make([][]int, 500)
	for i := 0; i < 500; i++ {
		raw[i] = []int{i}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := slicesToLists(raw)
		_ = mergeKLists(input)
	}
}

func BenchmarkMergeKLists_10x500(b *testing.B) {
	raw := make([][]int, 10)
	for i := 0; i < 10; i++ {
		raw[i] = make([]int, 500)
		for j := 0; j < 500; j++ {
			raw[i][j] = j*10 + i
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		input := slicesToLists(raw)
		_ = mergeKLists(input)
	}
}
