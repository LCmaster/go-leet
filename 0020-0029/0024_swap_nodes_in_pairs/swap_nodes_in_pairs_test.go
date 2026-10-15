package swapnodesinpairs

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

func TestSwapPairs(t *testing.T) {
	// Pre-generate data for maximum constraint tests (0 <= length <= 100)
	// 1. Length 100 (even): [1, 2, 3, 4, ..., 100] -> [2, 1, 4, 3, ..., 100, 99]
	len100 := make([]int, 100)
	want100 := make([]int, 100)
	for i := 0; i < 100; i += 2 {
		len100[i] = i + 1
		len100[i+1] = i + 2
		want100[i] = i + 2
		want100[i+1] = i + 1
	}

	// 2. Length 99 (odd): [1, 2, ..., 98, 99] -> [2, 1, ..., 98, 97, 99]
	len99 := make([]int, 99)
	want99 := make([]int, 99)
	for i := 0; i < 98; i += 2 {
		len99[i] = i + 1
		len99[i+1] = i + 2
		want99[i] = i + 2
		want99[i+1] = i + 1
	}
	len99[98] = 99
	want99[98] = 99

	tests := []struct {
		name string
		head []int
		want []int
	}{
		// --- 1. Official LeetCode Examples ---
		{
			name: "Example 1: Standard 4-node list [1,2,3,4]",
			head: []int{1, 2, 3, 4},
			want: []int{2, 1, 4, 3},
		},
		{
			name: "Example 2: Empty list []",
			head: []int{},
			want: []int{},
		},
		{
			name: "Example 3: Single-node list [1]",
			head: []int{1},
			want: []int{1},
		},
		{
			name: "Example 4: 3-node list with trailing odd element [1,2,3]",
			head: []int{1, 2, 3},
			want: []int{2, 1, 3},
		},

		// --- 2. Parity Variations (Even vs Odd Lengths) ---
		{
			name: "Length 2: Single pair [1,2]",
			head: []int{1, 2},
			want: []int{2, 1},
		},
		{
			name: "Length 5: Two pairs + 1 trailing node [1,2,3,4,5]",
			head: []int{1, 2, 3, 4, 5},
			want: []int{2, 1, 4, 3, 5},
		},
		{
			name: "Length 6: Three full pairs [1,2,3,4,5,6]",
			head: []int{1, 2, 3, 4, 5, 6},
			want: []int{2, 1, 4, 3, 6, 5},
		},
		{
			name: "Length 7: Three full pairs + 1 trailing node",
			head: []int{1, 2, 3, 4, 5, 6, 7},
			want: []int{2, 1, 4, 3, 6, 5, 7},
		},

		// --- 3. Duplicate and Constant Values ---
		{
			name: "All nodes with identical values [5,5,5,5]",
			head: []int{5, 5, 5, 5},
			want: []int{5, 5, 5, 5},
		},
		{
			name: "Alternating pairs of identical values [1,1,2,2]",
			head: []int{1, 1, 2, 2},
			want: []int{1, 1, 2, 2},
		},
		{
			name: "All zeros [0,0,0]",
			head: []int{0, 0, 0},
			want: []int{0, 0, 0},
		},

		// --- 4. Value Ordering and Boundary Constraints (0 <= Node.val <= 100) ---
		{
			name: "Decreasing initial order [4,3,2,1]",
			head: []int{4, 3, 2, 1},
			want: []int{3, 4, 1, 2},
		},
		{
			name: "Boundary constraint values 0 and 100 [0, 100, 100, 0]",
			head: []int{0, 100, 100, 0},
			want: []int{100, 0, 0, 100},
		},

		// --- 5. Maximum Constraint Tests (Length = 99 and 100) ---
		{
			name: "Maximum constraint length = 100 (50 pairs)",
			head: len100,
			want: want100,
		},
		{
			name: "Maximum constraint length = 99 (49 pairs + 1 trailing node)",
			head: len99,
			want: want99,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputList := sliceToList(tt.head)
			resultList := swapPairs(inputList)
			got := listToSlice(resultList)

			if len(got) == 0 && len(tt.want) == 0 {
				return
			}

			if !slices.Equal(got, tt.want) {
				t.Errorf("swapPairs(%v) = %v, want %v", tt.head, got, tt.want)
			}
		})
	}
}

func BenchmarkSwapPairs_4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := sliceToList([]int{1, 2, 3, 4})
		_ = swapPairs(list)
	}
}

func BenchmarkSwapPairs_100(b *testing.B) {
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list := sliceToList(nums)
		_ = swapPairs(list)
	}
}
