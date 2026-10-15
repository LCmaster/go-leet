package mergeksortedlists

import "math"

// 23. Merge k Sorted Lists
//
// You are given an array of `k` linked-lists `lists`, each linked-list is sorted in ascending order.
//
// *Merge all the linked-lists into one sorted linked-list and return it.*
//
// Example 1:
//
// ```text
// Input: lists = [[1,4,5],[1,3,4],[2,6]]
// Output: [1,1,2,3,4,4,5,6]
// Explanation: The linked-lists are:
// [
//   1->4->5,
//   1->3->4,
//   2->6
// ]
// merging them into one sorted linked list:
// 1->1->2->3->4->4->5->6
// ```
//
// Example 2:
//
// ```text
// Input: lists = []
// Output: []
// ```
//
// Example 3:
//
// ```text
// Input: lists = [[]]
// Output: []
// ```
//
// **Constraints:**
// - `k == lists.length`
// - `0 <= k <= 10^4`
// - `0 <= lists[i].length <= 500`
// - `-10^4 <= lists[i][j] <= 10^4`
// - `lists[i]` is sorted in **ascending order**.
// - The sum of `lists[i].length` will not exceed `10^4`.

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	m := len(lists)
	head := &ListNode{}

	if m == 0 {
		return nil
	}
	if m == 1 {
		return lists[0]
	}

	curr := head
	for true {
		minIndex := -1
		minValue := math.MaxInt32
		for i, l := range lists {
			if l == nil {
				continue
			}
			if l.Val < minValue {
				minIndex = i
				minValue = l.Val
			}
		}
		if minIndex < 0 {
			break
		}

		curr.Next = lists[minIndex]
		lists[minIndex] = lists[minIndex].Next
		curr = curr.Next
	}
	return head.Next
}
