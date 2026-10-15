package reversenodesinkgroup

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

func TestReverseKGroup(t *testing.T) {
	// Pre-generate data for maximum constraint tests (n = 5000)
	n5000 := make([]int, 5000)
	for i := 0; i < 5000; i++ {
		n5000[i] = i + 1
	}

	// 1. n = 5000, k = 2 (2500 pairs swapped)
	wantK2 := make([]int, 5000)
	for i := 0; i < 5000; i += 2 {
		wantK2[i] = i + 2
		wantK2[i+1] = i + 1
	}

	// 2. n = 5000, k = 5000 (entire list reversed)
	wantK5000 := make([]int, 5000)
	for i := 0; i < 5000; i++ {
		wantK5000[i] = 5000 - i
	}

	// 3. n = 5000, k = 7 (714 groups of 7 reversed + 2 trailing nodes)
	wantK7 := make([]int, 5000)
	copy(wantK7, n5000)
	for i := 0; i+7 <= 5000; i += 7 {
		slices.Reverse(wantK7[i : i+7])
	}

	tests := []struct {
		name string
		head []int
		k    int
		want []int
	}{
		// --- 1. Official LeetCode Examples ---
		{
			name: "Example 1: head = [1,2,3,4,5], k = 2 (2 pairs + 1 leftover)",
			head: []int{1, 2, 3, 4, 5},
			k:    2,
			want: []int{2, 1, 4, 3, 5},
		},
		{
			name: "Example 2: head = [1,2,3,4,5], k = 3 (1 group of 3 + 2 leftovers)",
			head: []int{1, 2, 3, 4, 5},
			k:    3,
			want: []int{3, 2, 1, 4, 5},
		},

		// --- 2. Boundary Values of k ---
		{
			name: "k = 1: list unchanged",
			head: []int{1, 2, 3, 4, 5},
			k:    1,
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "k = n: entire list reversed",
			head: []int{1, 2, 3, 4, 5},
			k:    5,
			want: []int{5, 4, 3, 2, 1},
		},
		{
			name: "Single node list: n = 1, k = 1",
			head: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "Two nodes: n = 2, k = 1",
			head: []int{1, 2},
			k:    1,
			want: []int{1, 2},
		},
		{
			name: "Two nodes: n = 2, k = 2",
			head: []int{1, 2},
			k:    2,
			want: []int{2, 1},
		},

		// --- 3. Exact Multiples of k (No Leftovers) ---
		{
			name: "Exact multiple: n = 4, k = 2",
			head: []int{1, 2, 3, 4},
			k:    2,
			want: []int{2, 1, 4, 3},
		},
		{
			name: "Exact multiple: n = 6, k = 3",
			head: []int{1, 2, 3, 4, 5, 6},
			k:    3,
			want: []int{3, 2, 1, 6, 5, 4},
		},
		{
			name: "Exact multiple: n = 8, k = 4",
			head: []int{1, 2, 3, 4, 5, 6, 7, 8},
			k:    4,
			want: []int{4, 3, 2, 1, 8, 7, 6, 5},
		},

		// --- 4. Varying Leftover Counts ---
		{
			name: "n = 3, k = 2 (1 group of 2 + 1 leftover)",
			head: []int{1, 2, 3},
			k:    2,
			want: []int{2, 1, 3},
		},
		{
			name: "n = 5, k = 4 (1 group of 4 + 1 leftover)",
			head: []int{1, 2, 3, 4, 5},
			k:    4,
			want: []int{4, 3, 2, 1, 5},
		},
		{
			name: "n = 7, k = 3 (2 groups of 3 + 1 leftover)",
			head: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			want: []int{3, 2, 1, 6, 5, 4, 7},
		},
		{
			name: "n = 8, k = 3 (2 groups of 3 + 2 leftovers)",
			head: []int{1, 2, 3, 4, 5, 6, 7, 8},
			k:    3,
			want: []int{3, 2, 1, 6, 5, 4, 7, 8},
		},

		// --- 5. k Greater Than n (Safety Check) ---
		{
			name: "k > n: list remains unchanged (no complete group)",
			head: []int{1, 2},
			k:    3,
			want: []int{1, 2},
		},
		{
			name: "k > n with 4 nodes: k = 5",
			head: []int{1, 2, 3, 4},
			k:    5,
			want: []int{1, 2, 3, 4},
		},

		// --- 6. Duplicate and Identical Node Values ---
		{
			name: "All identical values",
			head: []int{7, 7, 7, 7, 7},
			k:    2,
			want: []int{7, 7, 7, 7, 7},
		},
		{
			name: "Repeated pairs with k = 2",
			head: []int{1, 1, 2, 2, 3, 3},
			k:    2,
			want: []int{1, 1, 2, 2, 3, 3},
		},
		{
			name: "All zeros",
			head: []int{0, 0, 0, 0},
			k:    3,
			want: []int{0, 0, 0, 0},
		},

		// --- 7. Boundary Constraint Values (0 <= Node.val <= 1000) ---
		{
			name: "Extreme values 0 and 1000 with k = 3",
			head: []int{0, 1000, 0, 1000, 0},
			k:    3,
			want: []int{0, 1000, 0, 1000, 0},
		},
		{
			name: "Alternating 0 and 1000 with k = 2",
			head: []int{0, 1000, 0, 1000},
			k:    2,
			want: []int{1000, 0, 1000, 0},
		},

		// --- 8. Maximum Constraint Tests (n = 5000) ---
		{
			name: "Max constraint n = 5000, k = 1 (identity)",
			head: n5000,
			k:    1,
			want: n5000,
		},
		{
			name: "Max constraint n = 5000, k = 2 (2500 pairs swapped)",
			head: n5000,
			k:    2,
			want: wantK2,
		},
		{
			name: "Max constraint n = 5000, k = 5000 (full reverse)",
			head: n5000,
			k:    5000,
			want: wantK5000,
		},
		{
			name: "Max constraint n = 5000, k = 7 (714 groups + 2 leftovers)",
			head: n5000,
			k:    7,
			want: wantK7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputList := sliceToList(tt.head)
			resultList := reverseKGroup(inputList, tt.k)
			got := listToSlice(resultList)

			if len(got) == 0 && len(tt.want) == 0 {
				return
			}

			if !slices.Equal(got, tt.want) {
				if len(got) <= 20 {
					t.Errorf("reverseKGroup(%v, %d) = %v, want %v", tt.head, tt.k, got, tt.want)
				} else {
					t.Errorf("reverseKGroup(%s) failed: len(got)=%d, len(want)=%d; first 10 got=%v, want=%v",
						tt.name, len(got), len(tt.want), got[:10], tt.want[:10])
				}
			}
		})
	}
}

func BenchmarkReverseKGroup_Small(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := sliceToList([]int{1, 2, 3, 4, 5})
		_ = reverseKGroup(list, 2)
	}
}

func BenchmarkReverseKGroup_5000_k2(b *testing.B) {
	nums := make([]int, 5000)
	for i := range nums {
		nums[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list := sliceToList(nums)
		_ = reverseKGroup(list, 2)
	}
}

func BenchmarkReverseKGroup_5000_k5000(b *testing.B) {
	nums := make([]int, 5000)
	for i := range nums {
		nums[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list := sliceToList(nums)
		_ = reverseKGroup(list, 5000)
	}
}
