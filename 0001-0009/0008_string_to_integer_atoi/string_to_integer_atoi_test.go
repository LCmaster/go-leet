package stringtointegeratoi

import (
	"math"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1 - basic positive integer",
			s:    "42",
			want: 42,
		},
		{
			name: "Example 2 - leading whitespace and negative sign with leading zero",
			s:    "   -042",
			want: -42,
		},
		{
			name: "Example 3 - stops reading at non-digit characters",
			s:    "1337c0d3",
			want: 1337,
		},
		{
			name: "Example 4 - zero followed by minus",
			s:    "0-1",
			want: 0,
		},
		{
			name: "Example 5 - leading words before number",
			s:    "words and 987",
			want: 0,
		},
		{
			name: "Empty string",
			s:    "",
			want: 0,
		},
		{
			name: "Only spaces",
			s:    "    ",
			want: 0,
		},
		{
			name: "Only sign",
			s:    "+",
			want: 0,
		},
		{
			name: "Explicit plus sign",
			s:    "+1",
			want: 1,
		},
		{
			name: "Plus sign with spaces",
			s:    "   +42",
			want: 42,
		},
		{
			name: "Double signs",
			s:    "+-12",
			want: 0,
		},
		{
			name: "Words after digits",
			s:    "4193 with words",
			want: 4193,
		},
		{
			name: "Decimal dot stops reading",
			s:    "3.14159",
			want: 3,
		},
		{
			name: "Leading zeros",
			s:    "  0000000000012345",
			want: 12345,
		},
		{
			name: "Exact 32-bit max boundary",
			s:    "2147483647",
			want: math.MaxInt32,
		},
		{
			name: "Exact 32-bit min boundary",
			s:    "-2147483648",
			want: math.MinInt32,
		},
		{
			name: "32-bit positive overflow clamps to MaxInt32",
			s:    "2147483648",
			want: math.MaxInt32,
		},
		{
			name: "32-bit negative overflow clamps to MinInt32",
			s:    "-2147483649",
			want: math.MinInt32,
		},
		{
			name: "Large positive overflow clamps to MaxInt32",
			s:    "91283472332",
			want: math.MaxInt32,
		},
		{
			name: "Large negative overflow clamps to MinInt32",
			s:    "-91283472332",
			want: math.MinInt32,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.s); got != tt.want {
				t.Errorf("myAtoi(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
