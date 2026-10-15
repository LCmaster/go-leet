package removenthnodefromendoflist

// 19. Remove Nth Node From End of List
//
// Given the `head` of a linked list, remove the `n^th` node from the end of the list and return its head.
//
// Example 1:
//
// ```text
// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]
// ```
//
// Example 2:
//
// ```text
// Input: head = [1], n = 1
// Output: []
// ```
//
// Example 3:
//
// ```text
// Input: head = [1,2], n = 1
// Output: [1]
// ```
//
// **Constraints:**
// - The number of nodes in the list is `sz`.
// - `1 <= sz <= 30`
// - `0 <= Node.val <= 100`
// - `1 <= n <= sz`
//
// **Follow up:** Could you do this in one pass?

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	for range n {
		fast = fast.Next
	}
	slow := &head
	for fast != nil {
		fast = fast.Next
		slow = &(*slow).Next
	}
	// Unlink the target node directly (even if it's the head!)
	*slow = (*slow).Next
	return head
}
