# 23. Merge k Sorted Lists

**Difficulty:** Hard

## Description

You are given an array of `k` linked-lists `lists`, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.

## Solution
We can use a Min-Heap (Priority Queue) to always pick the smallest available node:

- Insert the head node of each non-empty list into the min-heap.
- Use a dummy head node and pointer `curr = dummy`.
- Repeatedly extract the minimum node from the heap:
  - Attach the extracted node to `curr.Next` and advance `curr`.
  - If the extracted node has a next node (`minNode.Next != nil`), insert `minNode.Next` into the heap.
- Return `dummy.Next`.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(N \log k)$ | $O(k)$ |
| Where $N$ is total nodes and $k$ is number of lists; each heap operation takes $O(\log k)$ | Heap holds at most $k$ elements simultaneously |
