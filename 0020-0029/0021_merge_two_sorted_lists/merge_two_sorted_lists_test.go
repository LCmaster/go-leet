package mergetwosortedlists

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

// Helper function to convert a linked list into a slice of ints.
func listToSlice(head *ListNode) []int {
	var nums []int
	for curr := head; curr != nil; curr = curr.Next {
		nums = append(nums, curr.Val)
	}
	return nums
}

func TestMergeTwoLists(t *testing.T) {
	// Pre-generate slices for max constraint testing (up to 50 nodes each)
	list50A := make([]int, 50)
	list50B := make([]int, 50)
	merged100 := make([]int, 100)
	for i := 0; i < 50; i++ {
		list50A[i] = i * 2       // 0, 2, 4, ..., 98
		list50B[i] = i*2 + 1     // 1, 3, 5, ..., 99
		merged100[i*2] = i * 2
		merged100[i*2+1] = i*2 + 1
	}

	allMinus100 := make([]int, 50)
	allPlus100 := make([]int, 50)
	mergedBoundary100 := make([]int, 100)
	for i := 0; i < 50; i++ {
		allMinus100[i] = -100
		allPlus100[i] = 100
		mergedBoundary100[i] = -100
		mergedBoundary100[50+i] = 100
	}

	tests := []struct {
		name  string
		list1 []int
		list2 []int
		want  []int
	}{
		// --- LeetCode Official Examples ---
		{
			name:  "Example 1: Interleaved lists with duplicates",
			list1: []int{1, 2, 4},
			list2: []int{1, 3, 4},
			want:  []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:  "Example 2: Both lists empty",
			list1: []int{},
			list2: []int{},
			want:  []int{},
		},
		{
			name:  "Example 3: First list empty, second list has one node",
			list1: []int{},
			list2: []int{0},
			want:  []int{0},
		},

		// --- Empty List Variations ---
		{
			name:  "First list has one node, second list empty",
			list1: []int{0},
			list2: []int{},
			want:  []int{0},
		},
		{
			name:  "First list empty, second list multi-node",
			list1: []int{},
			list2: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "First list multi-node, second list empty",
			list1: []int{1, 2, 3},
			list2: []int{},
			want:  []int{1, 2, 3},
		},
		{
			name:  "First list empty, second list all negative",
			list1: []int{},
			list2: []int{-5, -3, -1},
			want:  []int{-5, -3, -1},
		},

		// --- Single Element Lists ---
		{
			name:  "Single element each, list1 < list2",
			list1: []int{1},
			list2: []int{2},
			want:  []int{1, 2},
		},
		{
			name:  "Single element each, list2 < list1",
			list1: []int{2},
			list2: []int{1},
			want:  []int{1, 2},
		},
		{
			name:  "Single element each, identical values",
			list1: []int{5},
			list2: []int{5},
			want:  []int{5, 5},
		},
		{
			name:  "Single element each, negative and positive",
			list1: []int{-10},
			list2: []int{10},
			want:  []int{-10, 10},
		},

		// --- Disjoint / Non-Overlapping Ranges ---
		{
			name:  "Disjoint: All list1 elements smaller than list2",
			list1: []int{1, 2, 3},
			list2: []int{4, 5, 6},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "Disjoint: All list2 elements smaller than list1",
			list1: []int{4, 5, 6},
			list2: []int{1, 2, 3},
			want:  []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:  "Disjoint negative and positive ranges",
			list1: []int{-30, -20, -10},
			list2: []int{10, 20, 30},
			want:  []int{-30, -20, -10, 10, 20, 30},
		},

		// --- Asymmetric Lengths ---
		{
			name:  "Asymmetric: 1 element in list1 belongs at the beginning",
			list1: []int{0},
			list2: []int{1, 2, 3, 4, 5},
			want:  []int{0, 1, 2, 3, 4, 5},
		},
		{
			name:  "Asymmetric: 1 element in list1 belongs at the end",
			list1: []int{10},
			list2: []int{1, 2, 3, 4, 5},
			want:  []int{1, 2, 3, 4, 5, 10},
		},
		{
			name:  "Asymmetric: 1 element in list1 belongs in the middle",
			list1: []int{3},
			list2: []int{1, 2, 4, 5},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "Asymmetric: 1 element in list2 belongs in the middle",
			list1: []int{1, 2, 4, 5},
			list2: []int{3},
			want:  []int{1, 2, 3, 4, 5},
		},
		{
			name:  "Asymmetric: short list1 (2) and long list2 (7)",
			list1: []int{2, 6},
			list2: []int{1, 3, 4, 5, 7, 8, 9},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:  "Asymmetric: long list1 (7) and short list2 (2)",
			list1: []int{1, 3, 4, 5, 7, 8, 9},
			list2: []int{2, 6},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},

		// --- Duplicate Values and Identical Lists ---
		{
			name:  "Identical multi-element lists",
			list1: []int{1, 3, 5},
			list2: []int{1, 3, 5},
			want:  []int{1, 1, 3, 3, 5, 5},
		},
		{
			name:  "All elements identical across both lists",
			list1: []int{2, 2, 2},
			list2: []int{2, 2, 2, 2},
			want:  []int{2, 2, 2, 2, 2, 2, 2},
		},
		{
			name:  "Consecutive duplicates within list1",
			list1: []int{1, 1, 1, 5},
			list2: []int{2, 4},
			want:  []int{1, 1, 1, 2, 4, 5},
		},
		{
			name:  "Consecutive duplicates within list2",
			list1: []int{2, 4},
			list2: []int{3, 3, 3, 6},
			want:  []int{2, 3, 3, 3, 4, 6},
		},
		{
			name:  "Multiple duplicates interleaved across both lists",
			list1: []int{1, 2, 2, 5},
			list2: []int{2, 2, 3, 5},
			want:  []int{1, 2, 2, 2, 2, 3, 5, 5},
		},

		// --- Negative Values & Extreme Constraints (-100 <= val <= 100) ---
		{
			name:  "All negative numbers interleaved",
			list1: []int{-10, -8, -3},
			list2: []int{-9, -5, -1},
			want:  []int{-10, -9, -8, -5, -3, -1},
		},
		{
			name:  "Mixed negative, zero, and positive",
			list1: []int{-5, 0, 5},
			list2: []int{-3, 0, 3},
			want:  []int{-5, -3, 0, 0, 3, 5},
		},
		{
			name:  "Extreme constraint boundaries (-100 and 100)",
			list1: []int{-100, 100},
			list2: []int{-100, 100},
			want:  []int{-100, -100, 100, 100},
		},
		{
			name:  "Minimum bound single vs maximum bound single",
			list1: []int{-100},
			list2: []int{100},
			want:  []int{-100, 100},
		},
		{
			name:  "All zeros in both lists",
			list1: []int{0, 0},
			list2: []int{0, 0, 0},
			want:  []int{0, 0, 0, 0, 0},
		},

		// --- Alternating Patterns ---
		{
			name:  "Strictly alternating: list1 odd, list2 even",
			list1: []int{1, 3, 5, 7, 9},
			list2: []int{2, 4, 6, 8, 10},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:  "Strictly alternating: list1 even, list2 odd",
			list1: []int{2, 4, 6, 8, 10},
			list2: []int{1, 3, 5, 7, 9},
			want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},

		// --- Maximum Constraints (50 nodes per list) ---
		{
			name:  "Max constraint: list1 has 50 nodes, list2 empty",
			list1: list50A,
			list2: []int{},
			want:  list50A,
		},
		{
			name:  "Max constraint: list1 empty, list2 has 50 nodes",
			list1: []int{},
			list2: list50B,
			want:  list50B,
		},
		{
			name:  "Max constraint: 50 even nodes + 50 odd nodes -> 100 nodes merged",
			list1: list50A,
			list2: list50B,
			want:  merged100,
		},
		{
			name:  "Max constraint: 50 nodes at -100 + 50 nodes at 100",
			list1: allMinus100,
			list2: allPlus100,
			want:  mergedBoundary100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := sliceToList(tt.list1)
			l2 := sliceToList(tt.list2)
			merged := mergeTwoLists(l1, l2)
			got := listToSlice(merged)

			// Handle empty slice comparison (nil slice vs []int{})
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}

			if !slices.Equal(got, tt.want) {
				t.Errorf("mergeTwoLists(%v, %v) = %v, want %v", tt.list1, tt.list2, got, tt.want)
			}
		})
	}
}
