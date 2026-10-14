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
	buff := make([]byte, 0, m/2)

	if m%2 != 0 || s[0] == ')' || s[0] == ']' || s[0] == '}' || s[m-1] == '(' || s[m-1] == '[' || s[m-1] == '{' {
		return false
	}

	matchClose := func(b byte) (byte, bool) {
		switch b {
		case '(':
			return ')', true
		case '[':
			return ']', true
		case '{':
			return '}', true
		}
		return 0, false
	}

	for i := range m {
		if res, ok := matchClose(s[i]); ok {
			buff = append(buff, res)
		} else {
			if len(buff) == 0 || buff[len(buff)-1] != s[i] {
				return false
			}
			buff = buff[:len(buff)-1]
		}
	}

	return len(buff) == 0
}
