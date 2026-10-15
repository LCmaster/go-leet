package generateparentheses

import (
	"slices"
	"strings"
	"testing"
)

// sortSlice returns a sorted copy of a string slice to ensure assertions
// are order-independent.
func sortSlice(s []string) []string {
	if s == nil {
		return []string{}
	}
	cp := make([]string, len(s))
	copy(cp, s)
	slices.Sort(cp)
	return cp
}

// isValid checks that a string consists only of well-formed parentheses.
func isValid(s string) bool {
	balance := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			balance++
		case ')':
			balance--
			if balance < 0 {
				return false
			}
		default:
			return false
		}
	}
	return balance == 0
}

func TestGenerateParenthesis_ExactMatches(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		// --- Official LeetCode Examples ---
		{
			name: "Example 1: n = 3 (5 combinations)",
			n:    3,
			want: []string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			},
		},
		{
			name: "Example 2: Minimum constraint n = 1 (1 combination)",
			n:    1,
			want: []string{
				"()",
			},
		},

		// --- Additional Exact Match Verifications ---
		{
			name: "n = 2 (2 combinations)",
			n:    2,
			want: []string{
				"(())",
				"()()",
			},
		},
		{
			name: "n = 4 (14 combinations, C_4)",
			n:    4,
			want: []string{
				"(((())))",
				"((()()))",
				"((())())",
				"((()))()",
				"(()(()))",
				"(()()())",
				"(()())()",
				"(())(())",
				"(())()()",
				"()((()))",
				"()(()())",
				"()(())()",
				"()()(())",
				"()()()()",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.n)
			sortedGot := sortSlice(got)
			sortedWant := sortSlice(tt.want)

			if !slices.Equal(sortedGot, sortedWant) {
				t.Errorf("generateParenthesis(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestGenerateParenthesis_CatalanCounts(t *testing.T) {
	// The number of valid combinations for n pairs corresponds to the n-th Catalan number:
	// C_n = (1 / (n + 1)) * (2n choose n)
	catalanCounts := []struct {
		n     int
		count int
	}{
		{n: 1, count: 1},
		{n: 2, count: 2},
		{n: 3, count: 5},
		{n: 4, count: 14},
		{n: 5, count: 42},
		{n: 6, count: 132},
		{n: 7, count: 429},
		{n: 8, count: 1430}, // Maximum constraint (1 <= n <= 8)
	}

	for _, tt := range catalanCounts {
		t.Run(strings.ReplaceAll(strings.TrimSpace(strings.Repeat(" ", 0)), "", "")+
			"Catalan count for n = "+string(rune('0'+tt.n)), func(t *testing.T) {
			got := generateParenthesis(tt.n)
			if len(got) != tt.count {
				t.Fatalf("generateParenthesis(%d) produced %d combinations, want %d (Catalan C_%d)",
					tt.n, len(got), tt.count, tt.n)
			}
		})
	}
}

func TestGenerateParenthesis_Invariants(t *testing.T) {
	// Exhaustively verify structural and mathematical properties for all valid n from 1 to 8:
	for n := 1; n <= 8; n++ {
		got := generateParenthesis(n)
		seen := make(map[string]struct{}, len(got))

		fullyNested := strings.Repeat("(", n) + strings.Repeat(")", n)
		fullySequential := strings.Repeat("()", n)
		foundNested := false
		foundSequential := false

		for idx, str := range got {
			// 1. Length invariant: length must be exactly 2 * n
			if len(str) != 2*n {
				t.Errorf("n=%d [%d]: string %q has length %d, want %d", n, idx, str, len(str), 2*n)
			}

			// 2. Character counts: exactly n '(' and n ')'
			if strings.Count(str, "(") != n || strings.Count(str, ")") != n {
				t.Errorf("n=%d [%d]: string %q does not contain exactly %d '(' and %d ')'", n, idx, str, n, n)
			}

			// 3. Well-formedness invariant: prefix balance must never drop below 0
			if !isValid(str) {
				t.Errorf("n=%d [%d]: string %q is not a valid well-formed parentheses sequence", n, idx, str)
			}

			// 4. Uniqueness invariant: all generated strings must be distinct
			if _, exists := seen[str]; exists {
				t.Errorf("n=%d [%d]: duplicate combination found: %q", n, idx, str)
			}
			seen[str] = struct{}{}

			if str == fullyNested {
				foundNested = true
			}
			if str == fullySequential {
				foundSequential = true
			}
		}

		// 5. Canonical extreme combinations must always be generated
		if !foundNested {
			t.Errorf("n=%d: expected to find fully nested combination %q", n, fullyNested)
		}
		if !foundSequential {
			t.Errorf("n=%d: expected to find fully sequential combination %q", n, fullySequential)
		}
	}
}

func BenchmarkGenerateParenthesis_N3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generateParenthesis(3)
	}
}

func BenchmarkGenerateParenthesis_N8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generateParenthesis(8)
	}
}
