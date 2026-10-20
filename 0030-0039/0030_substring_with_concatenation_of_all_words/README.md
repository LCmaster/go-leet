# 30. Substring with Concatenation of All Words

**Difficulty:** Hard

## Description

You are given a string `s` and an array of strings `words`. All strings in `words` are of the same length.

A concatenated string is a string that exactly contains all the strings of any permutation of `words` concatenated.

Return an array of the starting indices of all concatenated substrings in `s`. You can return the answer in any order.

## Solution
We can use a fixed-size sliding window with frequency map comparison:

- Build a `frequencyMap` of expected word counts.
- Let `subStringLength = len(words) * len(words[0])`. If `len(s) < subStringLength`, return nil.
- For each starting index `i` from `0` to `len(s) - subStringLength`:
  - Extract words of length `wordLength` from `s[i : i+subStringLength]` into a reusable `localMap`.
  - Compare frequencies between `localMap` and `frequencyMap`.
  - If they match, record `i` as a valid starting index.
- Return `indices`.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O((n - k \cdot L) \cdot k)$ | $O(k \cdot L)$ |
| Where $n = \text{len}(s)$, $k = \text{len}(words)$, and $L$ is word length | Stores word frequency maps containing up to $k$ words |
