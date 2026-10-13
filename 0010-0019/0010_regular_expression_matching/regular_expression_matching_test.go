package regularexpressionmatching

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want bool
	}{
		// --- LeetCode Official Examples ---
		{
			name: "Example 1 - pattern shorter without star",
			s:    "aa",
			p:    "a",
			want: false,
		},
		{
			name: "Example 2 - star repeats preceding character",
			s:    "aa",
			p:    "a*",
			want: true,
		},
		{
			name: "Example 3 - dot star matches any string",
			s:    "ab",
			p:    ".*",
			want: true,
		},

		// --- Exact Character Matching (No '.' or '*') ---
		{
			name: "Exact single character match",
			s:    "a",
			p:    "a",
			want: true,
		},
		{
			name: "Single character mismatch",
			s:    "a",
			p:    "b",
			want: false,
		},
		{
			name: "Multi-character exact match",
			s:    "abc",
			p:    "abc",
			want: true,
		},
		{
			name: "Pattern longer than string without star",
			s:    "abc",
			p:    "abcd",
			want: false,
		},
		{
			name: "String longer than pattern without star",
			s:    "abcd",
			p:    "abc",
			want: false,
		},

		// --- Dot '.' Wildcard Only ---
		{
			name: "Dot matches single character",
			s:    "a",
			p:    ".",
			want: true,
		},
		{
			name: "Multiple dots match characters",
			s:    "ab",
			p:    "..",
			want: true,
		},
		{
			name: "Dot in middle of characters",
			s:    "abc",
			p:    "a.c",
			want: true,
		},
		{
			name: "Dot matches but neighboring character fails",
			s:    "abc",
			p:    "a.d",
			want: false,
		},

		// --- Star '*' Matching 0 Times ---
		{
			name: "Star matches zero times at end",
			s:    "a",
			p:    "ab*",
			want: true,
		},
		{
			name: "Star matches zero times at beginning",
			s:    "a",
			p:    "b*a",
			want: true,
		},
		{
			name: "Multiple stars matching zero times",
			s:    "aab",
			p:    "c*a*b",
			want: true,
		},
		{
			name: "Star matches zero times followed by failure",
			s:    "a",
			p:    "ab*b",
			want: false,
		},

		// --- Star '*' Matching Multiple Times ---
		{
			name: "Star matches multiple times",
			s:    "aaa",
			p:    "a*",
			want: true,
		},
		{
			name: "Star matches multiple times with trailing identical char",
			s:    "aaaa",
			p:    "a*a",
			want: true,
		},
		{
			name: "Multiple interleaved star patterns",
			s:    "aaa",
			p:    "ab*a*c*a",
			want: true,
		},

		// --- Dot and Star Combinations ('.*') ---
		{
			name: "Dot star matches any prefix",
			s:    "aab",
			p:    ".*b",
			want: true,
		},
		{
			name: "Dot star matches middle",
			s:    "abcdef",
			p:    "a.*f",
			want: true,
		},
		{
			name: "Dot star but required trailing character missing",
			s:    "ab",
			p:    ".*c",
			want: false,
		},
		{
			name: "Multiple dot stars in sequence",
			s:    "abcd",
			p:    ".*.*.*",
			want: true,
		},

		// --- Classic Backtracking & Edge Cases ---
		{
			name: "Mississippi failing match",
			s:    "mississippi",
			p:    "mis*is*p*.",
			want: false,
		},
		{
			name: "Mississippi successful match",
			s:    "mississippi",
			p:    "mis*is*ip*.",
			want: true,
		},
		{
			name: "Backtracking with dot star and trailing characters",
			s:    "bbbba",
			p:    ".*a*a",
			want: true,
		},
		{
			name: "Pattern requires more characters than string length",
			s:    "a",
			p:    ".*..a*",
			want: false,
		},
		{
			name: "Multiple consecutive skipped elements",
			s:    "a",
			p:    "b*c*d*a*e*",
			want: true,
		},
		{
			name: "Long exact match with repeated groups",
			s:    "aabcbcbcaccbcaabc",
			p:    ".*a*aa*.*b*.c*.*",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.s, tt.p); got != tt.want {
				t.Errorf("isMatch(%q, %q) = %v, want %v", tt.s, tt.p, got, tt.want)
			}
		})
	}
}
