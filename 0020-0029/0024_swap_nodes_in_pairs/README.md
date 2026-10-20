# 24. Swap Nodes in Pairs

**Difficulty:** Medium

## Description

Given a linked list, swap every two adjacent nodes and return its head. You must solve the problem without modifying the values in the list's nodes (i.e., only nodes themselves may be changed).

## Solution
We can iterate using a pointer-to-pointer to rewire nodes in-place:

- Initialize `currP := &head` to point to the incoming pointer of the current pair.
- While `*currP != nil && (*currP).Next != nil`:
  - Let `curr = *currP` and `nxt = curr.Next`.
  - Rewire: `curr.Next = nxt.Next`, `nxt.Next = curr`, and `*currP = nxt`.
  - Advance `currP = &curr.Next` to the next pair.
- Return `head`.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(n)$ | $O(1)$ |
| Single pass through the linked list of length $n$ | In-place pointer updates with no extra allocations |
