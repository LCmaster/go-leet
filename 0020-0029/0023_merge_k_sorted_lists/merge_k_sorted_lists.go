package mergeksortedlists

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

type MinHeap struct {
	elements []*ListNode
}

func (h *MinHeap) Insert(key *ListNode) {
	h.elements = append(h.elements, key)
	h.heapifyUp(len(h.elements) - 1)
}

func (h *MinHeap) ExtractMin() *ListNode {
	if len(h.elements) == 0 {
		return nil
	}

	min := h.elements[0]
	lastIndex := len(h.elements) - 1
	h.elements[0] = h.elements[lastIndex]
	h.elements = h.elements[:lastIndex]
	h.heapifyDown(0)

	return min
}

func (h *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.elements[index].Val >= h.elements[parentIndex].Val {
			break
		}
		h.elements[index], h.elements[parentIndex] = h.elements[parentIndex], h.elements[index]
		index = parentIndex
	}
}

func (h *MinHeap) heapifyDown(index int) {
	lastIndex := len(h.elements) - 1

	for {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		smallest := index

		if leftChild <= lastIndex && h.elements[leftChild].Val < h.elements[smallest].Val {
			smallest = leftChild
		}
		if rightChild <= lastIndex && h.elements[rightChild].Val < h.elements[smallest].Val {
			smallest = rightChild
		}
		if smallest == index {
			break
		}
		h.elements[index], h.elements[smallest] = h.elements[smallest], h.elements[index]
		index = smallest
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	m := len(lists)

	if m == 0 {
		return nil
	}
	if m == 1 {
		return lists[0]
	}

	heap := &MinHeap{}
	for _, l := range lists {
		if l != nil {
			heap.Insert(l)
		}
	}
	head := &ListNode{}
	curr := head
	for len(heap.elements) > 0 {
		minNode := heap.ExtractMin()
		curr.Next = minNode
		curr = curr.Next
		if minNode.Next != nil {
			heap.Insert(minNode.Next)
		}
	}

	return head.Next
}
