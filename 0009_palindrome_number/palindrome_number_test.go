package palindromenumber

import (
	"math"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want bool
	}{
		{
			name: "Example 1 - odd length palindrome",
			x:    121,
			want: true,
		},
		{
			name: "Example 2 - negative number",
			x:    -121,
			want: false,
		},
		{
			name: "Example 3 - trailing zero",
			x:    10,
			want: false,
		},
		{
			name: "Zero is palindrome",
			x:    0,
			want: true,
		},
		{
			name: "Single digit positive",
			x:    7,
			want: true,
		},
		{
			name: "Single digit negative",
			x:    -7,
			want: false,
		},
		{
			name: "Two identical digits",
			x:    11,
			want: true,
		},
		{
			name: "Two different digits",
			x:    12,
			want: false,
		},
		{
			name: "Even length palindrome",
			x:    1221,
			want: true,
		},
		{
			name: "Even length non-palindrome",
			x:    1231,
			want: false,
		},
		{
			name: "Odd length non-palindrome",
			x:    123,
			want: false,
		},
		{
			name: "Internal zeros palindrome",
			x:    1001,
			want: true,
		},
		{
			name: "Internal zeros odd length palindrome",
			x:    10001,
			want: true,
		},
		{
			name: "Multiple trailing zeros",
			x:    1000,
			want: false,
		},
		{
			name: "Large odd length palindrome",
			x:    123454321,
			want: true,
		},
		{
			name: "Large even length palindrome",
			x:    12344321,
			want: true,
		},
		{
			name: "Large 32-bit palindrome",
			x:    2147447412,
			want: true,
		},
		{
			name: "MaxInt32 boundary",
			x:    math.MaxInt32, // 2147483647
			want: false,
		},
		{
			name: "MinInt32 boundary",
			x:    math.MinInt32, // -2147483648
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.x); got != tt.want {
				t.Errorf("isPalindrome(%d) = %v, want %v", tt.x, got, tt.want)
			}
		})
	}
}
