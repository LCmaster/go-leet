# 21. Merge Two Sorted Lists

**Difficulty:** Easy

## Description

You are given the heads of two sorted linked lists `list1` and `list2`.

Merge the two lists into one **sorted** list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.

## Solution
We can iterate through both lists using a dummy head:

- Initialize a `dummy` head node and a pointer `curr = dummy`.
- While both `list1` and `list2` are non-nil:
  - Compare `list1.Val` and `list2.Val`, attach the smaller node to `curr.Next`, and advance that list's pointer.
  - Advance `curr = curr.Next`.
- Splice the remaining non-nil list onto `curr.Next`.
- Return `dummy.Next`.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(n + m)$ | $O(1)$ |
| Single pass where $n$ and $m$ are the lengths of the two lists | Splices existing nodes in-place without new node allocations |
