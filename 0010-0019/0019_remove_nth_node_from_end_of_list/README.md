# 19. Remove Nth Node From End of List

**Difficulty:** Medium

## Description

Given the `head` of a linked list, remove the $n^{\text{th}}$ node from the end of the list and return its head.

## Solution
We can use fast and slow pointers with a pointer-to-pointer:

- Advance a `fast` pointer $n$ steps into the list.
- Initialize `slow` as a pointer to the head pointer (`slow := &head`).
- Advance both `fast` and `slow` together until `fast == nil`.
- At this point, `*slow` points directly to the target node's address. Unlink it via `*slow = (*slow).Next`.
- Return `head`.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(L)$ | $O(1)$ |
| Single pass over the linked list of length $L$ | Uses only pointer references without extra allocations |
