# 20. Valid Parentheses

**Difficulty:** Easy

## Description

Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

An input string is valid if:
- Open brackets must be closed by the same type of brackets.
- Open brackets must be closed in the correct order.
- Every close bracket has a corresponding open bracket of the same type.

## Solution
We can use a stack to track open brackets and match them with incoming closing brackets:

- Perform quick parity and boundary checks (e.g. odd length or starting with closing bracket returns `false`).
- Iterate through each character in `s`:
  - If it is an opening bracket (`'('`, `'['`, `'{'`), push its corresponding closing bracket (`')'`, `']'`, `'}'`) onto the stack.
  - If it is a closing bracket, check that the stack is non-empty and the top of the stack matches; pop from the stack.
- Return `true` if the stack is empty after processing the entire string.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(n)$ | $O(n)$ |
| Single pass over the string of length $n$ | Stack stores at most $n / 2$ expected closing brackets |
