package longestvalidparentheses

import (
	"strings"
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		// --- 1. LeetCode Official Examples ---
		{
			name: "Example 1: '(()' -> longest is '()' of len 2",
			s:    "(()",
			want: 2,
		},
		{
			name: "Example 2: ')()())' -> longest is '()()' of len 4",
			s:    ")()())",
			want: 4,
		},
		{
			name: "Example 3: empty string -> 0",
			s:    "",
			want: 0,
		},

		// --- 2. Minimal Length Boundary Cases (0 and 1 char) ---
		{
			name: "Boundary: single open '('",
			s:    "(",
			want: 0,
		},
		{
			name: "Boundary: single close ')'",
			s:    ")",
			want: 0,
		},

		// --- 3. Length 2 Strings (All 4 Permutations) ---
		{
			name: "Len 2: balanced pair '()'",
			s:    "()",
			want: 2,
		},
		{
			name: "Len 2: two open brackets '(('",
			s:    "((",
			want: 0,
		},
		{
			name: "Len 2: two close brackets '))'",
			s:    "))",
			want: 0,
		},
		{
			name: "Len 2: reversed brackets ')('",
			s:    ")(",
			want: 0,
		},

		// --- 4. Length 3 Strings (Exhaustive All 8 Permutations) ---
		{
			name: "Len 3: '((('",
			s:    "(((",
			want: 0,
		},
		{
			name: "Len 3: '(()'",
			s:    "(()",
			want: 2,
		},
		{
			name: "Len 3: '()('",
			s:    "()(",
			want: 2,
		},
		{
			name: "Len 3: '())'",
			s:    "())",
			want: 2,
		},
		{
			name: "Len 3: ')(( '",
			s:    ")((",
			want: 0,
		},
		{
			name: "Len 3: ')()'",
			s:    ")()",
			want: 2,
		},
		{
			name: "Len 3: '))('",
			s:    "))(",
			want: 0,
		},
		{
			name: "Len 3: ')))'",
			s:    ")))",
			want: 0,
		},

		// --- 5. Length 4 Strings (Key Permutations & Structures) ---
		{
			name: "Len 4: all open '(((('",
			s:    "((((",
			want: 0,
		},
		{
			name: "Len 4: all close '))))'",
			s:    "))))",
			want: 0,
		},
		{
			name: "Len 4: sequential pairs '()()'",
			s:    "()()",
			want: 4,
		},
		{
			name: "Len 4: nested pair '(())'",
			s:    "(())",
			want: 4,
		},
		{
			name: "Len 4: open prefix before pair '((()'",
			s:    "((()",
			want: 2,
		},
		{
			name: "Len 4: pair followed by close '()))'",
			s:    "()))",
			want: 2,
		},
		{
			name: "Len 4: close before open pair ')(()'",
			s:    ")(()",
			want: 2,
		},
		{
			name: "Len 4: pair with open and close boundary ')()('",
			s:    ")()(",
			want: 2,
		},
		{
			name: "Len 4: pair at start followed by reversed '())('",
			s:    "())(",
			want: 2,
		},

		// --- 6. Pure Nested Parentheses ---
		{
			name: "Nested depth 3: '((()))'",
			s:    "((()))",
			want: 6,
		},
		{
			name: "Nested depth 4: '(((())))'",
			s:    "(((())))",
			want: 8,
		},
		{
			name: "Nested depth 5: '((((()))))'",
			s:    "((((()))))",
			want: 10,
		},
		{
			name: "Nested with redundant outer opens: '((((()))'",
			s:    "((((()))",
			want: 6,
		},
		{
			name: "Nested with redundant outer closes: '((()))))'",
			s:    "((()))))",
			want: 6,
		},
		{
			name: "Nested with both open prefix and close suffix: '((((())))))'",
			s:    "((((())))))",
			want: 10,
		},

		// --- 7. Sequential / Concatenated Valid Pairs ---
		{
			name: "Sequential 3 pairs: '()()()'",
			s:    "()()()",
			want: 6,
		},
		{
			name: "Sequential 4 pairs: '()()()()'",
			s:    "()()()()",
			want: 8,
		},
		{
			name: "Sequential 6 pairs: '()()()()()()'",
			s:    "()()()()()()",
			want: 12,
		},

		// --- 8. Compound & Interleaved Patterns ---
		{
			name: "Compound: pair + nested + pair '()(())()'",
			s:    "()(())()",
			want: 8,
		},
		{
			name: "Compound: nested + pair + nested '((()))()(())'",
			s:    "((()))()(())",
			want: 12,
		},
		{
			name: "Compound: deeply nested inside pair '(()(()))'",
			s:    "(()(()))",
			want: 8,
		},
		{
			name: "Compound: multiple nested siblings inside parent '((())())'",
			s:    "((())())",
			want: 8,
		},
		{
			name: "Compound: long chain '()((()))()()(())'",
			s:    "()((()))()()(())",
			want: 16,
		},

		// --- 9. Separated Valid Substrings (Selecting the Longest Component) ---
		{
			name: "Separated: ')()())()()(' -> two components of len 4",
			s:    ")()())()()(",
			want: 4,
		},
		{
			name: "Separated: '())(()()' -> len 2 vs len 4 -> 4",
			s:    "())(()()",
			want: 4,
		},
		{
			name: "Separated: '()(())((' -> len 6 before trailing opens",
			s:    "()(())((",
			want: 6,
		},
		{
			name: "Separated: multiple closes '))()()()())()()()())'",
			s:    "))()()()())()()()())",
			want: 8,
		},
		{
			name: "Separated: invalid segments between valid chunks '))))()()()()))))()()((((('",
			s:    "))))()()()()))))()()(((((",
			want: 8,
		},
		{
			name: "Separated: middle component longest '())(()())(()'",
			s:    "())(()())(()",
			want: 6,
		},
		{
			name: "Separated: two equal longest components '())(()())(()())'",
			s:    "())(()())(()())",
			want: 12, // wait: note "())(()())(()())" has "(()())" followed by "(()())" which merges into "(()())(()())" of len 12
		},

		// --- 10. Tricky Unbalanced Patterns & Edge Cases ---
		{
			name: "Tricky: unmatched open in middle '()(()' -> 2",
			s:    "()(()",
			want: 2,
		},
		{
			name: "Tricky: unmatched open at start '(()()' -> 4",
			s:    "(()()",
			want: 4,
		},
		{
			name: "Tricky: multiple opens before nested '(()((()))' -> 8",
			s:    "(()((()))",
			want: 8,
		},
		{
			name: "Tricky: trailing close after valid '((()))())' -> 8",
			s:    "((()))())",
			want: 8,
		},
		{
			name: "Tricky: leading close with composite valid and unmatched open ')()((((())))()' -> 10",
			s:    ")()((((())))()",
			want: 10,
		},
		{
			name: "Tricky: complex LeetCode test case ')(((((()())()()))()(()))(' -> 22",
			s:    ")(((((()())()()))()(()))(",
			want: 22,
		},
		{
			name: "Tricky: nested with trailing open and close ')()((()()))' -> 10",
			s:    ")()((()()))",
			want: 10,
		},
		{
			name: "Tricky: open suffix '(())(()' -> 4",
			s:    "(())(()",
			want: 4,
		},
		{
			name: "Tricky: bookended by invalid brackets ')(()())(' -> 6",
			s:    ")(()())(",
			want: 6,
		},
		{
			name: "Tricky: duplicate close break '(()))())(' -> 4",
			s:    "(()))())(",
			want: 4,
		},
		{
			name: "Tricky: multiple open prefix before pairs '((()()' -> 4",
			s:    "((()()",
			want: 4,
		},
		{
			name: "Tricky: middle broken '((()())(()' -> 6",
			s:    "((()())(()",
			want: 6,
		},

		// --- 11. All-Open and All-Closed Sequences ---
		{
			name: "All open brackets: 10 '(((((((((('",
			s:    strings.Repeat("(", 10),
			want: 0,
		},
		{
			name: "All close brackets: 10 '))))))))))'",
			s:    strings.Repeat(")", 10),
			want: 0,
		},
		{
			name: "All open brackets: 100 '('",
			s:    strings.Repeat("(", 100),
			want: 0,
		},
		{
			name: "All close brackets: 100 ')'",
			s:    strings.Repeat(")", 100),
			want: 0,
		},

		// --- 12. Alternating Patterns ---
		{
			name: "Alternating ')(...)' contains inner '()()' chain: 10 pairs -> 18",
			s:    strings.Repeat(")(", 10),
			want: 18,
		},
		{
			name: "All closes followed by all opens: 50 ')' then 50 '(' -> 0",
			s:    strings.Repeat(")", 50) + strings.Repeat("(", 50),
			want: 0,
		},
		{
			name: "Alternating starting with close and ending with close: ')()()()()())'",
			s:    ")()()()()())",
			want: 10,
		},

		// --- 13. Large Constraint Stress Tests (N up to 30,000) ---
		{
			name: "Stress: 30,000 all opens",
			s:    strings.Repeat("(", 30000),
			want: 0,
		},
		{
			name: "Stress: 30,000 all closes",
			s:    strings.Repeat(")", 30000),
			want: 0,
		},
		{
			name: "Stress: 15,000 closes followed by 15,000 opens -> 0",
			s:    strings.Repeat(")", 15000) + strings.Repeat("(", 15000),
			want: 0,
		},
		{
			name: "Stress: 30,000 fully balanced nested '(((...)))'",
			s:    strings.Repeat("(", 15000) + strings.Repeat(")", 15000),
			want: 30000,
		},
		{
			name: "Stress: 30,000 fully balanced sequential '()()()'",
			s:    strings.Repeat("()", 15000),
			want: 30000,
		},
		{
			name: "Stress: 20,000 valid wrapped with invalid bounds ')' + 10,000 '()' + '('",
			s:    ")" + strings.Repeat("()", 10000) + "(",
			want: 20000,
		},
		{
			name: "Stress: two large valid blocks separated by ')(' (7000 pairs and 6000 pairs)",
			s:    strings.Repeat("()", 7000) + ")(" + strings.Repeat("()", 6000),
			want: 14000,
		},
		{
			name: "Stress: deep nested core with outer wrapper: 5000 '(' + 4000 '()' + 5000 ')'",
			s:    strings.Repeat("(", 5000) + strings.Repeat("()", 4000) + strings.Repeat(")", 5000),
			want: 18000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestValidParentheses(tt.s)
			if got != tt.want {
				t.Errorf("longestValidParentheses(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}

// --- Benchmark Suite ---

func BenchmarkLongestValidParentheses(b *testing.B) {
	benchmarks := []struct {
		name string
		s    string
	}{
		{
			name: "Small_10",
			s:    ")()((()()))",
		},
		{
			name: "Medium_1000_Sequential",
			s:    strings.Repeat("()", 500),
		},
		{
			name: "Medium_1000_Nested",
			s:    strings.Repeat("(", 500) + strings.Repeat(")", 500),
		},
		{
			name: "Medium_1000_WorstCase_AllOpens",
			s:    strings.Repeat("(", 1000),
		},
		{
			name: "Large_30000_Sequential",
			s:    strings.Repeat("()", 15000),
		},
		{
			name: "Large_30000_Nested",
			s:    strings.Repeat("(", 15000) + strings.Repeat(")", 15000),
		},
		{
			name: "Large_30000_Split",
			s:    strings.Repeat("()", 7000) + ")(" + strings.Repeat("()", 6000),
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = longestValidParentheses(bm.s)
			}
		})
	}
}
