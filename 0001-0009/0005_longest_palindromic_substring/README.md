# 5. Longest Palindromic Substring

**Difficulty:** Medium

## Description

Given a string `s`, return the longest palindromic substring in `s`.

## Solution
We can expand around each possible center:

- A palindrome can center on a single character (odd length) or between two characters (even length), yielding $2n - 1$ possible centers.
- For each index `i`, expand outward using helper pointers `(i, i)` and `(i, i + 1)` while characters match.
- Track the starting position `start` and maximum length `maxLength` whenever a longer palindrome is found.
- Return the substring `s[start : start + maxLength]`.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(n^2)$ | $O(1)$ |
| Expanding around each of the $2n - 1$ centers takes up to $O(n)$ time | Uses only pointer variables for expansion |
