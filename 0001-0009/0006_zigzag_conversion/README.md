# 6. Zigzag Conversion

**Difficulty:** Medium

## Description

The string `"PAYPALISHIRING"` is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

```text
P   A   H   N
A P L S I I G
Y   I   R
```

And then read line by line: `"PAHNAPLSIIGYIR"`

Write the code that will take a string and make this conversion given a number of rows:

```text
string convert(string s, int numRows);
```

## Solution
We can calculate the row indices mathematically without simulating the grid:

- If `numRows == 1` or `numRows >= len(s)`, return `s` directly.
- Each full zigzag cycle spans `cycle = 2 * numRows - 2` characters.
- Iterate through each row from `0` to `numRows - 1` and jump by `cycle`:
  - Append the primary character at `i`.
  - For inner rows (`0 < row < numRows - 1`), append the diagonal character at `i + (cycle - 2 * row)` if within bounds.
- Return the reconstructed string.

**Complexity:**
| Time | Space |
| :--- | :--- |
| $O(n)$ | $O(1)$ |
| Visits each character of the string of length $n$ exactly once | Auxiliary space ($O(n)$ to construct the returned string) |
