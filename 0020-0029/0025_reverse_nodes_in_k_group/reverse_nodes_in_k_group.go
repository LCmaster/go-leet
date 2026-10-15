package reversenodesinkgroup

// 25. Reverse Nodes in k-Group
//
// Given the `head` of a linked list, reverse the nodes of the list `k` at a time, and return *the modified list*.
//
// `k` is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of `k` then left-out nodes, in the end, should remain as it is.
//
// You may not alter the values in the list's nodes, only nodes themselves may be changed.
//
// Example 1:
//
// ```text
// Input: head = [1,2,3,4,5], k = 2
// Output: [2,1,4,3,5]
// ```
//
// Example 2:
//
// ```text
// Input: head = [1,2,3,4,5], k = 3
// Output: [3,2,1,4,5]
// ```
//
// **Constraints:**
// - The number of nodes in the list is `n`.
// - `1 <= k <= n <= 5000`
// - `0 <= Node.val <= 1000`
//
// **Follow-up:** Can you solve the problem in `O(1)` extra memory space?

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}
	dummy := &ListNode{Next: head}
	prevGroup := dummy
	for {
		fast := prevGroup
		for i := 0; i < k && fast != nil; i++ {
			fast = fast.Next
		}
		if fast == nil {
			break
		}
		groupStart := prevGroup.Next
		nextGroup := fast.Next

		prev := nextGroup
		curr := groupStart
		for curr != nextGroup {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}
		prevGroup.Next = fast
		prevGroup = groupStart
	}
	return dummy.Next
}
