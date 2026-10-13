package longestsubstring

/*
3. Longest Substring Without Repeating Characters

Given a string s, find the length of the longest substring without duplicate characters.

Examples:
- Input: s = "abcabcbb" | Output: 3
- Input: s = "bbbbb"    | Output: 1
- Input: s = "pwwkew"   | Output: 3

Constraints:
- 0 <= s.length <= 10^5
- s consists of English letters, digits, symbols and spaces.
*/

func lengthOfLongestSubstring(s string) int {
	charIndexMap := make(map[rune]int)
	maxLength := 0
	start := 0

	for i, char := range s {
		if index, ok := charIndexMap[char]; ok && index >= start {
			start = index + 1
		}
		charIndexMap[char] = i
		
		currentLength := i - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}
