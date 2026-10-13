package longestpalindromicsubstring

// 5. Longest Palindromic Substring
//
// Given a string `s`, return *the longest* *palindromic* *substring* in `s`.
//
// Example 1:
//
// ```text
// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.
// ```
//
// Example 2:
//
// ```text
// Input: s = "cbbd"
// Output: "bb"
// ```
//
// **Constraints:**
// - `1 <= s.length <= 1000`
// - `s` consist of only digits and English letters.

func longestPalindrome(str string) string {
	if len(str) == 1 {
		return str
	}

	start, maxLength := 0, 0
	expand := func(l, r int) {
		for l >= 0 && r < len(str) && str[l] == str[r] {
			l--
			r++
		}

		length := r - l - 1
		if length > maxLength {
			maxLength = length
			start = l + 1
		}
	}
	for i := 0; i < len(str); i++ {
		expand(i, i)   // Odd length center
		expand(i, i+1) // Even length center
	}
	return str[start : start+maxLength]
}
