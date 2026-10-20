# 28. Find the Index of the First Occurrence in a String

**Difficulty:** Easy

## Description

Given two strings `needle` and `haystack`, return the index of the first occurrence of `needle` in `haystack`, or `-1` if `needle` is not part of `haystack`.

## Solution
We can use a sliding window over `haystack` to find `needle`:

- Let $h = \text{len}(haystack)$ and $n = \text{len}(needle)$.
- Slide a window of size $n$ from index `0` up to $h - n$.
- Check if the substring slice `haystack[i:i+n]` equals `needle`.
- Return the starting index `i` on the first match, or `-1` if no match is found.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O((h - n + 1) \cdot n)$ | $O(1)$ |
| Slides window of size $n$ across $h - n + 1$ possible positions | Constant extra space (string slicing in Go does not allocate new buffers) |
