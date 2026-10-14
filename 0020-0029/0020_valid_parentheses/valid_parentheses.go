package validparentheses

// 20. Valid Parentheses
//
// Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.
//
// An input string is valid if:
// - Open brackets must be closed by the same type of brackets.
// - Open brackets must be closed in the correct order.
// - Every close bracket has a corresponding open bracket of the same type.
//
// Example 1:
//
// **Input:** s = "()"
//
// **Output:** true
//
// Example 2:
//
// **Input:** s = "()[]{}"
//
// **Output:** true
//
// Example 3:
//
// **Input:** s = "(]"
//
// **Output:** false
//
// Example 4:
//
// **Input:** s = "([])"
//
// **Output:** true
//
// Example 5:
//
// **Input:** s = "([)]"
//
// **Output:** false
//
// **Constraints:**
// - `1 <= s.length <= 10^4`
// - `s` consists of parentheses only `'()[]{}'`.

func isValid(s string) bool {
	m := len(s)
	buff := make([]int, 0, m)
	value := []int{'(': 1, ')': -1, '[': 2, ']': -2, '{': 3, '}': -3}

	if m < 2 {
		return false
	}
	if value[s[0]] < 0 && value[s[m-1]] > 0 {
		return false
	}

	for _, c := range s {
		if value[c] > 0 {
			buff = append(buff, value[c])
		} else {
			if len(buff) > 0 && value[c]+buff[len(buff)-1] == 0 {
				buff = buff[:len(buff)-1]
			} else {
				return false
			}
		}
	}

	return len(buff) == 0
}
