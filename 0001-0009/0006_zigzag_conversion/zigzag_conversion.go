package zigzagconversion

// 6. Zigzag Conversion
//
// The string `"PAYPALISHIRING"` is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)
//
// ```text
// P   A   H   N
// A P L S I I G
// Y   I   R
// ```
//
// And then read line by line: `"PAHNAPLSIIGYIR"`
//
// Write the code that will take a string and make this conversion given a number of rows:
//
// ```text
// string convert(string s, int numRows);
// ```
//
// Example 1:
//
// ```text
// Input: s = "PAYPALISHIRING", numRows = 3
// Output: "PAHNAPLSIIGYIR"
// ```
//
// Example 2:
//
// ```text
// Input: s = "PAYPALISHIRING", numRows = 4
// Output: "PINALSIGYAHRPI"
// Explanation:
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
// ```
//
// Example 3:
//
// ```text
// Input: s = "A", numRows = 1
// Output: "A"
// ```
//
// **Constraints:**
// - `1 <= s.length <= 1000`
// - `s` consists of English letters (lower-case and upper-case), `','` and `'.'`.
// - `1 <= numRows <= 1000`

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	index := 0
	result := make([]byte, len(s))

	cycle := (numRows*2 - 2)
	for row := range numRows {
		step := cycle - 2*row
		for i := row; i < len(s); i += cycle {
			result[index] = s[i]
			index++
			if row > 0 && row < numRows-1 && i+step < len(s) {
				result[index] = s[i+step]
				index++
			}
		}
	}

	return string(result)
}
