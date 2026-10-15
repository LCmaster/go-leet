package generateparentheses

// 22. Generate Parentheses
//
// Given `n` pairs of parentheses, write a function to *generate all combinations of well-formed parentheses*.
//
// Example 1:
//
// ```text
// Input: n = 3
// Output: ["((()))","(()())","(())()","()(())","()()()"]
// ```
//
// Example 2:
//
// ```text
// Input: n = 1
// Output: ["()"]
// ```
//
// **Constraints:**
// - `1 <= n <= 8`

func generateParenthesis(n int) []string {
	result := []string{}
	var backtrack func(string, int, int)
	backtrack = func(str string, i, j int) {
		if i == n && j == n {
			result = append(result, str)
			return
		}
		if i < n {
			backtrack(str+"(", i+1, j)
		}
		if j < i {
			backtrack(str+")", i, j+1)
		}
	}

	backtrack("(", 1, 0)
	return result
}
