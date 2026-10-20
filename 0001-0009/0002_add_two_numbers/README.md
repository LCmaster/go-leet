# 2. Add Two Numbers

**Difficulty:** Medium

## Description

You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order**, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number `0` itself.

## Solution
We can simulate column-by-column addition:
- Use a dummy head node to build the result list, and maintain a `carry` variable.
- Traverse both lists as long as either list has nodes or `carry > 0`.
- At each step, add non-nil node values and `carry`, append a new node with `sum % 10`, and update `carry = sum / 10`.
- Return `dummy.Next`.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(\max(m, n))$ | $O(1)$ |
| Iterates at most $\max(m, n) + 1$ times, where $m$ and $n$ are the lengths of `l1` and `l2` | Auxiliary space ($O(\max(m, n))$ if counting the output list) |
