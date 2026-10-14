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
	'1': "",
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
	m := len(digits)
	total := 1
	for i := range m {
		total *= len(dic[digits[i]])
	}
	buff := make([]rune, m)
	result := make([]string, 0, total)

	var backtrack func(int)
	backtrack = func(index int) {
		if index == m {
			result = append(result, string(buff))
			return
		}
		for _, c := range dic[digits[index]] {
			buff[index] = c
			backtrack(index + 1)
		}
	}

	backtrack(0)
	return result
}
