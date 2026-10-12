package reverseinteger

import (
	"math"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "Example 1",
			x:    123,
			want: 321,
		},
		{
			name: "Example 2",
			x:    -123,
			want: -321,
		},
		{
			name: "Example 3 - trailing zeros",
			x:    120,
			want: 21,
		},
		{
			name: "Zero",
			x:    0,
			want: 0,
		},
		{
			name: "Single positive digit",
			x:    7,
			want: 7,
		},
		{
			name: "Single negative digit",
			x:    -5,
			want: -5,
		},
		{
			name: "Multiple trailing zeros",
			x:    900000,
			want: 9,
		},
		{
			name: "Valid large positive within [-(2^31), 2^31 - 1]",
			x:    1463847412,
			want: 2147483641, // 2147483641 <= math.MaxInt32 (2147483647)
		},
		{
			name: "Valid large negative within [-(2^31), 2^31 - 1]",
			x:    -1463847412,
			want: -2147483641, // -2147483641 >= math.MinInt32 (-2147483648)
		},
		{
			name: "Positive overflow (> 2^31 - 1) returns 0",
			x:    1534236469, // reverses to 9646324351 > 2^31 - 1
			want: 0,
		},
		{
			name: "Negative overflow (< -(2^31)) returns 0",
			x:    -1534236469, // reverses to -9646324351 < -(2^31)
			want: 0,
		},
		{
			name: "Min int32 input (-2^31) overflow returns 0",
			x:    math.MinInt32, // -2147483648 reverses to -8463847412 < math.MinInt32
			want: 0,
		},
		{
			name: "Max int32 input (2^31 - 1) overflow returns 0",
			x:    math.MaxInt32, // 2147483647 reverses to 7463847412 > math.MaxInt32
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.x); got != tt.want {
				t.Errorf("reverse(%d) = %d, want %d", tt.x, got, tt.want)
			}
		})
	}
}
