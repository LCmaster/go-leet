package validparentheses

import (
	"strings"
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		// --- LeetCode Official Examples ---
		{
			name: "Example 1: Single pair parentheses",
			s:    "()",
			want: true,
		},
		{
			name: "Example 2: All three types in sequence",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "Example 3: Mismatched types",
			s:    "(]",
			want: false,
		},
		{
			name: "Example 4: Nested parentheses and brackets",
			s:    "([])",
			want: true,
		},
		{
			name: "Example 5: Interleaved / improperly ordered brackets",
			s:    "([)]",
			want: false,
		},

		// --- Odd Length Strings (Cannot possibly be balanced) ---
		{
			name: "Odd length: Single open parenthesis",
			s:    "(",
			want: false,
		},
		{
			name: "Odd length: Single close parenthesis",
			s:    ")",
			want: false,
		},
		{
			name: "Odd length: Single square open",
			s:    "[",
			want: false,
		},
		{
			name: "Odd length: Single curly close",
			s:    "}",
			want: false,
		},
		{
			name: "Odd length: Three open parentheses",
			s:    "(((",
			want: false,
		},
		{
			name: "Odd length: Valid pair plus extra open at end",
			s:    "()(",
			want: false,
		},
		{
			name: "Odd length: Valid pair plus extra close at end",
			s:    "())",
			want: false,
		},

		// --- Starting with Closing Brackets ---
		{
			name: "Reversed pair )(",
			s:    ")(",
			want: false,
		},
		{
			name: "Reversed square ][",
			s:    "][",
			want: false,
		},
		{
			name: "Reversed curly }{",
			s:    "}{",
			want: false,
		},
		{
			name: "Close curly before valid pair",
			s:    "}()",
			want: false,
		},
		{
			name: "Close bracket before valid nested pair",
			s:    "]()",
			want: false,
		},

		// --- Unclosed / Leftover Opening Brackets on Stack ---
		{
			name: "Unclosed open bracket in front of valid pair",
			s:    "(()",
			want: false,
		},
		{
			name: "Unclosed curly bracket in front",
			s:    "{()[]",
			want: false,
		},
		{
			name: "Unclosed open bracket at the end",
			s:    "()(){",
			want: false,
		},
		{
			name: "Valid sequence with trailing open curly",
			s:    "()[]{}{",
			want: false,
		},
		{
			name: "All opening brackets",
			s:    "(((([[[[{{{{",
			want: false,
		},

		// --- All Closing Brackets ---
		{
			name: "All closing brackets",
			s:    "))))]]]]}}}}",
			want: false,
		},

		// --- Deeply Nested Valid Structures ---
		{
			name: "Deep nesting same type: ((()))",
			s:    "((()))",
			want: true,
		},
		{
			name: "Deep nesting mixed types: {[()]}",
			s:    "{[()]}",
			want: true,
		},
		{
			name: "Triple nesting all types: {[{()}]}",
			s:    "{[{()}]}",
			want: true,
		},
		{
			name: "Complex symmetric nesting: {{{[[[((()))]]]}}}",
			s:    "{{{[[[((()))]]]}}}",
			want: true,
		},

		// --- Deeply Nested Invalid Structures ---
		{
			name: "Deep nesting with wrong innermost closing: {[([)]]}",
			s:    "{[([)]]}",
			want: false,
		},
		{
			name: "Deep nesting with mismatched closing at bottom: {[(])}",
			s:    "{[(])}",
			want: false,
		},
		{
			name: "Deep nesting with mismatched outer bracket: {[()]>] (bracket mismatch)",
			s:    "{[()]",
			want: false,
		},
		{
			name: "Nested valid structure with mismatched closing at end",
			s:    "((()))]",
			want: false,
		},

		// --- Sequential and Compound Structures ---
		{
			name: "Multiple identical pairs: ()()()()()",
			s:    "()()()()()",
			want: true,
		},
		{
			name: "Multiple nested pairs in sequence: (())[][]",
			s:    "(())[][]",
			want: true,
		},
		{
			name: "Compound valid mixture: ()[]{}([{}])",
			s:    "()[]{}([{}])",
			want: true,
		},
		{
			name: "Compound mixture with one mismatch in middle: ()[]{}([)()])",
			s:    "()[]{}([)()])",
			want: false,
		},

		// --- Large / Boundary Test Cases (Constraint: 1 <= s.length <= 10^4) ---
		{
			name: "Large valid: 5000 nested '((...))' pairs (len = 10000)",
			s:    strings.Repeat("(", 5000) + strings.Repeat(")", 5000),
			want: true,
		},
		{
			name: "Large invalid: 5000 '(' followed by 4999 ')' (len = 9999)",
			s:    strings.Repeat("(", 5000) + strings.Repeat(")", 4999),
			want: false,
		},
		{
			name: "Large valid: 5000 sequential '()' pairs (len = 10000)",
			s:    strings.Repeat("()", 5000),
			want: true,
		},
		{
			name: "Large invalid: 5000 sequential '()' with mismatch at position 5000",
			s:    strings.Repeat("()", 2500) + "(]" + strings.Repeat("()", 2499),
			want: false,
		},
		{
			name: "Large valid: 2500 repeating '{[]}' sequences (len = 10000)",
			s:    strings.Repeat("{[]}", 2500),
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValid(tt.s)
			if got != tt.want {
				t.Errorf("isValid(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
