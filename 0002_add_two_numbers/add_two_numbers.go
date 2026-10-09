package addtwonumbers

// 2. Add Two Numbers

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// Examples:
// - Input: l1 = [2,4,3], l2 = [5,6,4]           | Output: [7,0,8]
// - Input: l1 = [0], l2 = [0]                   | Output: [0]
// - Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9] | Output: [8,9,9,9,0,0,0,1]
//
// Constraints:
// - The number of nodes in each linked list is in the range [1, 100].
// - 0 <= Node.val <= 9
// - It is guaranteed that the list represents a number that does not have leading zeros.

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	carry := 0

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return head.Next
}
