package removenthnodefromendoflist

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

func TestRemoveNthFromEnd(t *testing.T) {
	// Build a 30-element slice for boundary constraint testing (1 <= sz <= 30)
	sz30 := make([]int, 30)
	for i := 0; i < 30; i++ {
		sz30[i] = i + 1
	}

	tests := []struct {
		name string
		head []int
		n    int
		want []int
	}{
		// --- LeetCode Official Examples ---
		{
			name: "Example 1: Remove second from end in 5-node list",
			head: []int{1, 2, 3, 4, 5},
			n:    2,
			want: []int{1, 2, 3, 5},
		},
		{
			name: "Example 2: Remove the only node in 1-node list",
			head: []int{1},
			n:    1,
			want: []int{},
		},
		{
			name: "Example 3: Remove last node in 2-node list",
			head: []int{1, 2},
			n:    1,
			want: []int{1},
		},

		// --- Removing the Head Node (n == length of list) ---
		{
			name: "Remove head from 2-node list (n = 2)",
			head: []int{1, 2},
			n:    2,
			want: []int{2},
		},
		{
			name: "Remove head from 3-node list (n = 3)",
			head: []int{1, 2, 3},
			n:    3,
			want: []int{2, 3},
		},
		{
			name: "Remove head from 5-node list (n = 5)",
			head: []int{1, 2, 3, 4, 5},
			n:    5,
			want: []int{2, 3, 4, 5},
		},

		// --- Removing the Tail Node (n == 1) ---
		{
			name: "Remove tail from 3-node list (n = 1)",
			head: []int{1, 2, 3},
			n:    1,
			want: []int{1, 2},
		},
		{
			name: "Remove tail from 5-node list (n = 1)",
			head: []int{1, 2, 3, 4, 5},
			n:    1,
			want: []int{1, 2, 3, 4},
		},

		// --- Removing Intermediate Nodes ---
		{
			name: "Remove exact middle of 3-node list (n = 2)",
			head: []int{1, 2, 3},
			n:    2,
			want: []int{1, 3},
		},
		{
			name: "Remove second node of 4-node list (n = 3)",
			head: []int{10, 20, 30, 40},
			n:    3,
			want: []int{10, 30, 40},
		},
		{
			name: "Remove third node of 4-node list (n = 2)",
			head: []int{10, 20, 30, 40},
			n:    2,
			want: []int{10, 20, 40},
		},

		// --- Duplicate Values and Zeros (0 <= Node.val <= 100) ---
		{
			name: "List with all zero values",
			head: []int{0, 0, 0},
			n:    1,
			want: []int{0, 0},
		},
		{
			name: "List with duplicate values remove middle",
			head: []int{7, 7, 7, 7},
			n:    2,
			want: []int{7, 7, 7},
		},
		{
			name: "List with repeated pattern",
			head: []int{1, 2, 1, 2, 1},
			n:    3,
			want: []int{1, 2, 2, 1},
		},

		// --- Maximum Constraint (sz = 30) ---
		{
			name: "Max constraint sz=30: remove head (n = 30)",
			head: sz30,
			n:    30,
			want: sz30[1:],
		},
		{
			name: "Max constraint sz=30: remove tail (n = 1)",
			head: sz30,
			n:    1,
			want: sz30[:29],
		},
		{
			name: "Max constraint sz=30: remove middle node (n = 15)",
			head: sz30,
			n:    15,
			// Removing 15th from end means index 30 - 15 = 15 (value 16)
			want: append(append([]int{}, sz30[:15]...), sz30[16:]...),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listInput := sliceToList(tt.head)
			resultList := removeNthFromEnd(listInput, tt.n)
			got := listToSlice(resultList)

			// Handle empty slice comparison (nil slice vs []int{})
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}

			if !slices.Equal(got, tt.want) {
				t.Errorf("removeNthFromEnd(%v, %d) = %v, want %v", tt.head, tt.n, got, tt.want)
			}
		})
	}
}
