package lettercombinationsofaphonenumber

// 17. Letter Combinations of a Phone Number
//
// Given a string containing digits from `2-9` inclusive, return all possible letter combinations that the number could represent. Return the answer in **any order**.
//
// A mapping of digits to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.
//
// Example 1:
//
// ```text
// Input: digits = "23"
// Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
// ```
//
// Example 2:
//
// ```text
// Input: digits = "2"
// Output: ["a","b","c"]
// ```
//
// **Constraints:**
// - `1 <= digits.length <= 4`
// - `digits[i]` is a digit in the range `['2', '9']`.

var dic = []string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	return bfs("", 0, digits)
}

func bfs(root string, index int, digits string) []string {
	m := len(digits)
	result := []string{}
	for _, c := range dic[digits[index]] {
		if index < m-1 {
			result = append(result, bfs(root+string(c), index+1, digits)...)
		} else {
			result = append(result, root+string(c))
		}
	}
	return result
}
