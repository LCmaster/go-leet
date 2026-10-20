# 25. Reverse Nodes in k-Group

**Difficulty:** Hard

## Description

Given the `head` of a linked list, reverse the nodes of the list `k` at a time, and return the modified list.

`k` is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of `k` then left-out nodes, in the end, should remain as is.

You may not alter the values in the list's nodes, only nodes themselves may be changed.

## Solution
We can reverse nodes in groups of $k$ using iterative pointer manipulation:

- Use a `dummy` node pointing to `head` and maintain `prevGroup = dummy`.
- For each group:
  - Advance a `fast` pointer $k$ steps ahead from `prevGroup`. If fewer than $k$ nodes remain (`fast == nil`), stop.
  - Record `groupStart = prevGroup.Next` and `nextGroup = fast.Next`.
  - Reverse the $k$ nodes between `groupStart` and `nextGroup` in-place.
  - Connect `prevGroup.Next = fast` and advance `prevGroup = groupStart`.
- Return `dummy.Next`.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(n)$ | $O(1)$ |
| Each node is visited and reversed at most twice | Modifies pointers in-place with $O(1)$ extra memory |
