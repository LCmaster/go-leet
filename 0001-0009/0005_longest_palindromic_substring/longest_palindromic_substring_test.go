package longestpalindromicsubstring

import (
	"slices"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		possible []string // Some inputs (like "babad") have multiple valid answers ("bab" or "aba")
	}{
		{
			name:     "Example 1",
			input:    "babad",
			possible: []string{"bab", "aba"},
		},
		{
			name:     "Example 2",
			input:    "cbbd",
			possible: []string{"bb"},
		},
		{
			name:     "Single character",
			input:    "a",
			possible: []string{"a"},
		},
		{
			name:     "Two same characters",
			input:    "aa",
			possible: []string{"aa"},
		},
		{
			name:     "Two different characters",
			input:    "ac",
			possible: []string{"a", "c"},
		},
		{
			name:     "Entire string is palindrome",
			input:    "racecar",
			possible: []string{"racecar"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalindrome(tt.input)
			if !slices.Contains(tt.possible, got) {
				t.Errorf("longestPalindrome(%q) = %q, want one of %v", tt.input, got, tt.possible)
			}
		})
	}
}
