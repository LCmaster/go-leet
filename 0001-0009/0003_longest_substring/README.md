# 3. Longest Substring Without Repeating Characters

**Difficulty:** Medium

## Description

Given a string `s`, find the length of the longest substring without duplicate characters.

## Solution
We can solve this in a single pass using a sliding window and a hash map:

- Maintain a map `charIndexMap` storing each character's latest index and a `start` pointer for the window boundary.
- For each character `char` at index `i`, if it was previously seen at or after `start`, move `start` to `index + 1`.
- Update `charIndexMap[char] = i` and update `maxLength` with `i - start + 1`.
- Return `maxLength`.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(n)$ | $O(\min(n, m))$ |
| Single pass over the string with $O(1)$ map operations | Stores unique characters in the map, bounded by string length $n$ and alphabet size $m$ |
