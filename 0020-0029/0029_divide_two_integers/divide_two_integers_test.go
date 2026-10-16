package dividetwointegers

import (
	"math"
	"testing"
)

func TestDivide(t *testing.T) {
	const (
		minInt32 = math.MinInt32 // -2147483648 (-2^31)
		maxInt32 = math.MaxInt32 //  2147483647 ( 2^31 - 1)
	)

	tests := []struct {
		name     string
		dividend int
		divisor  int
		want     int
	}{
		// --- 1. LeetCode Official Examples ---
		{
			name:     "Example 1: 10 / 3 = 3",
			dividend: 10,
			divisor:  3,
			want:     3,
		},
		{
			name:     "Example 2: 7 / -3 = -2",
			dividend: 7,
			divisor:  -3,
			want:     -2,
		},

		// --- 2. 32-Bit Integer Overflow & Boundary Clamping ---
		{
			name:     "Overflow case: -2^31 / -1 clamps to 2^31 - 1",
			dividend: minInt32,
			divisor:  -1,
			want:     maxInt32,
		},
		{
			name:     "Boundary: -2^31 / 1 = -2^31",
			dividend: minInt32,
			divisor:  1,
			want:     minInt32,
		},
		{
			name:     "Boundary: (2^31 - 1) / 1 = 2^31 - 1",
			dividend: maxInt32,
			divisor:  1,
			want:     maxInt32,
		},
		{
			name:     "Boundary: (2^31 - 1) / -1 = -(2^31 - 1)",
			dividend: maxInt32,
			divisor:  -1,
			want:     -maxInt32,
		},
		{
			name:     "Both extremes: -2^31 / -2^31 = 1",
			dividend: minInt32,
			divisor:  minInt32,
			want:     1,
		},
		{
			name:     "Both extremes: (2^31 - 1) / (2^31 - 1) = 1",
			dividend: maxInt32,
			divisor:  maxInt32,
			want:     1,
		},
		{
			name:     "Opposite extremes: -2^31 / (2^31 - 1) = -1",
			dividend: minInt32,
			divisor:  maxInt32,
			want:     -1,
		},
		{
			name:     "Opposite extremes: (2^31 - 1) / -2^31 = 0",
			dividend: maxInt32,
			divisor:  minInt32,
			want:     0,
		},

		// --- 3. Zero Dividend ---
		{
			name:     "Zero dividend with positive divisor: 0 / 1 = 0",
			dividend: 0,
			divisor:  1,
			want:     0,
		},
		{
			name:     "Zero dividend with negative divisor: 0 / -1 = 0",
			dividend: 0,
			divisor:  -1,
			want:     0,
		},
		{
			name:     "Zero dividend with large divisor: 0 / 12345 = 0",
			dividend: 0,
			divisor:  12345,
			want:     0,
		},
		{
			name:     "Zero dividend with minInt32: 0 / -2^31 = 0",
			dividend: 0,
			divisor:  minInt32,
			want:     0,
		},

		// --- 4. Dividend Magnitude Smaller Than Divisor (|dividend| < |divisor|) ---
		{
			name:     "Positive: 1 / 2 = 0",
			dividend: 1,
			divisor:  2,
			want:     0,
		},
		{
			name:     "Negative dividend: -1 / 2 = 0",
			dividend: -1,
			divisor:  2,
			want:     0,
		},
		{
			name:     "Negative divisor: 1 / -2 = 0",
			dividend: 1,
			divisor:  -2,
			want:     0,
		},
		{
			name:     "Both negative: -1 / -2 = 0",
			dividend: -1,
			divisor:  -2,
			want:     0,
		},
		{
			name:     "Smaller magnitude: 5 / 10 = 0",
			dividend: 5,
			divisor:  10,
			want:     0,
		},
		{
			name:     "Smaller magnitude negative: -5 / 10 = 0",
			dividend: -5,
			divisor:  10,
			want:     0,
		},

		// --- 5. Exact Divisions (No Remainder) ---
		{
			name:     "Positive / Positive: 6 / 3 = 2",
			dividend: 6,
			divisor:  3,
			want:     2,
		},
		{
			name:     "Negative / Positive: -6 / 3 = -2",
			dividend: -6,
			divisor:  3,
			want:     -2,
		},
		{
			name:     "Positive / Negative: 6 / -3 = -2",
			dividend: 6,
			divisor:  -3,
			want:     -2,
		},
		{
			name:     "Negative / Negative: -6 / -3 = 2",
			dividend: -6,
			divisor:  -3,
			want:     2,
		},
		{
			name:     "Multiples of 10: 100 / 10 = 10",
			dividend: 100,
			divisor:  10,
			want:     10,
		},
		{
			name:     "Multiples of 25: 100 / 25 = 4",
			dividend: 100,
			divisor:  25,
			want:     4,
		},

		// --- 6. Truncation Toward Zero (Fractional part discarded) ---
		{
			name:     "Truncate positive/positive: 15 / 4 = 3",
			dividend: 15,
			divisor:  4,
			want:     3,
		},
		{
			name:     "Truncate negative/positive: -15 / 4 = -3 (not -4)",
			dividend: -15,
			divisor:  4,
			want:     -3,
		},
		{
			name:     "Truncate positive/negative: 15 / -4 = -3 (not -4)",
			dividend: 15,
			divisor:  -4,
			want:     -3,
		},
		{
			name:     "Truncate negative/negative: -15 / -4 = 3",
			dividend: -15,
			divisor:  -4,
			want:     3,
		},
		{
			name:     "Truncate: 19 / 5 = 3",
			dividend: 19,
			divisor:  5,
			want:     3,
		},
		{
			name:     "Truncate: -19 / 5 = -3",
			dividend: -19,
			divisor:  5,
			want:     -3,
		},

		// --- 7. Powers of Two (tests bit-shifting mechanics) ---
		{
			name:     "Power of 2: 1024 / 2 = 512",
			dividend: 1024,
			divisor:  2,
			want:     512,
		},
		{
			name:     "Power of 2: 1024 / 4 = 256",
			dividend: 1024,
			divisor:  4,
			want:     256,
		},
		{
			name:     "Power of 2: 1024 / 8 = 128",
			dividend: 1024,
			divisor:  8,
			want:     128,
		},
		{
			name:     "Power of 2: 1024 / 1024 = 1",
			dividend: 1024,
			divisor:  1024,
			want:     1,
		},
		{
			name:     "Power of 2: 1024 / 2048 = 0",
			dividend: 1024,
			divisor:  2048,
			want:     0,
		},

		// --- 8. Divisor is 1 or -1 ---
		{
			name:     "Positive / 1: 42 / 1 = 42",
			dividend: 42,
			divisor:  1,
			want:     42,
		},
		{
			name:     "Negative / 1: -42 / 1 = -42",
			dividend: -42,
			divisor:  1,
			want:     -42,
		},
		{
			name:     "Positive / -1: 42 / -1 = -42",
			dividend: 42,
			divisor:  -1,
			want:     -42,
		},
		{
			name:     "Negative / -1: -42 / -1 = 42",
			dividend: -42,
			divisor:  -1,
			want:     42,
		},

		// --- 9. Large Boundary Numbers with Various Divisors ---
		{
			name:     "maxInt32 / 2: 2147483647 / 2 = 1073741823",
			dividend: maxInt32,
			divisor:  2,
			want:     1073741823,
		},
		{
			name:     "maxInt32 / 3: 2147483647 / 3 = 715827882",
			dividend: maxInt32,
			divisor:  3,
			want:     715827882,
		},
		{
			name:     "minInt32 / 2: -2147483648 / 2 = -1073741824",
			dividend: minInt32,
			divisor:  2,
			want:     -1073741824,
		},
		{
			name:     "minInt32 / -2: -2147483648 / -2 = 1073741824",
			dividend: minInt32,
			divisor:  -2,
			want:     1073741824,
		},
		{
			name:     "minInt32 / 3: -2147483648 / 3 = -715827882",
			dividend: minInt32,
			divisor:  3,
			want:     -715827882,
		},
		{
			name:     "minInt32 / -3: -2147483648 / -3 = 715827882",
			dividend: minInt32,
			divisor:  -3,
			want:     715827882,
		},
		{
			name:     "minInt32 / 4: -2147483648 / 4 = -536870912",
			dividend: minInt32,
			divisor:  4,
			want:     -536870912,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := divide(tt.dividend, tt.divisor)
			if got != tt.want {
				t.Errorf("divide(%d, %d) = %d; want %d", tt.dividend, tt.divisor, got, tt.want)
			}
		})
	}
}

func BenchmarkDivide(b *testing.B) {
	b.ResetTimer()
	for b.Loop() {
		_ = divide(math.MaxInt32, 3)
		_ = divide(math.MinInt32, -3)
		_ = divide(1000000000, 7)
	}
}
