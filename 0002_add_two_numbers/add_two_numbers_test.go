package addtwonumbers

import (
	"reflect"
	"testing"
)

// Helper function to convert a slice to a linked list
func sliceToList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

// Helper function to convert a linked list to a slice
func listToSlice(head *ListNode) []int {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		{
			name: "Example 1",
			l1:   []int{2, 4, 3},
			l2:   []int{5, 6, 4},
			want: []int{7, 0, 8},
		},
		{
			name: "Example 2",
			l1:   []int{0},
			l2:   []int{0},
			want: []int{0},
		},
		{
			name: "Example 3",
			l1:   []int{9, 9, 9, 9, 9, 9, 9},
			l2:   []int{9, 9, 9, 9},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := sliceToList(tt.l1)
			list2 := sliceToList(tt.l2)
			
			gotList := addTwoNumbers(list1, list2)
			got := listToSlice(gotList)

			// If the returned list is nil but we expect an empty slice, handle that
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
