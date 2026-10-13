package regularexpressionmatching

// 10. Regular Expression Matching
//
// Given an input string `s`and a pattern `p`, implement regular expression matching with support for `'.'` and `'*'` where:
// - `'.'` Matches any single character.​​​​
// - `'*'` Matches zero or more of the preceding element.
//
// Return a boolean indicating whether the matching covers the entire input string (not partial).
//
// Example 1:
//
// ```text
// Input: s = "aa", p = "a"
// Output: false
// Explanation: "a" does not match the entire string "aa".
// ```
//
// Example 2:
//
// ```text
// Input: s = "aa", p = "a*"
// Output: true
// Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
// ```
//
// Example 3:
//
// ```text
// Input: s = "ab", p = ".*"
// Output: true
// Explanation: ".*" means "zero or more (*) of any character (.)".
// ```
//
// **Constraints:**
// - `1 <= s.length<= 20`
// - `1 <= p.length<= 20`
// - `s` contains only lowercase English letters.
// - `p` contains only lowercase English letters, `'.'`, and`'*'`.
// - It is guaranteed for each appearance of the character `'*'`, there will be a previous valid character to match.

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	memo := make([][]bool, m+1)

	for i := range memo {
		memo[i] = make([]bool, n+1)
		for j := range memo[i] {
			if i == 0 && j == 0 {
				memo[i][j] = true
				continue
			}
			if i == 0 || j == 0 {
				if i == 0 && j >= 2 && p[j-1] == '*' {
					memo[i][j] = memo[i][j-2]
				} else {
					memo[i][j] = false
				}

				continue
			}

			if s[i-1] == p[j-1] && memo[i-1][j-1] {
				memo[i][j] = true
			} else if p[j-1] == '.' {
				memo[i][j] = memo[i-1][j-1]
			} else if p[j-1] == '*' {
				memo[i][j] = memo[i][j-2]
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					memo[i][j] = memo[i][j] || memo[i-1][j]
				}
			}
		}
	}

	return memo[m][n]
}
