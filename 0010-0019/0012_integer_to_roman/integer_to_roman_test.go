package integertoroman

import (
	"testing"
)

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		// --- LeetCode Official Examples ---
		{
			name: "Example 1 - 3749",
			num:  3749,
			want: "MMMDCCXLIX",
		},
		{
			name: "Example 2 - 58",
			num:  58,
			want: "LVIII",
		},
		{
			name: "Example 3 - 1994",
			num:  1994,
			want: "MCMXCIV",
		},

		// --- Boundary Cases ---
		{
			name: "Minimum constraint (1)",
			num:  1,
			want: "I",
		},
		{
			name: "Maximum constraint (3999)",
			num:  3999,
			want: "MMMCMXCIX",
		},

		// --- Single Roman Symbols ---
		{
			name: "Single symbol I (1)",
			num:  1,
			want: "I",
		},
		{
			name: "Single symbol V (5)",
			num:  5,
			want: "V",
		},
		{
			name: "Single symbol X (10)",
			num:  10,
			want: "X",
		},
		{
			name: "Single symbol L (50)",
			num:  50,
			want: "L",
		},
		{
			name: "Single symbol C (100)",
			num:  100,
			want: "C",
		},
		{
			name: "Single symbol D (500)",
			num:  500,
			want: "D",
		},
		{
			name: "Single symbol M (1000)",
			num:  1000,
			want: "M",
		},

		// --- Subtractive Cases ---
		{
			name: "Subtractive IV (4)",
			num:  4,
			want: "IV",
		},
		{
			name: "Subtractive IX (9)",
			num:  9,
			want: "IX",
		},
		{
			name: "Subtractive XL (40)",
			num:  40,
			want: "XL",
		},
		{
			name: "Subtractive XC (90)",
			num:  90,
			want: "XC",
		},
		{
			name: "Subtractive CD (400)",
			num:  400,
			want: "CD",
		},
		{
			name: "Subtractive CM (900)",
			num:  900,
			want: "CM",
		},

		// --- Repeated Symbols (Powers of 10) ---
		{
			name: "Repetition II (2)",
			num:  2,
			want: "II",
		},
		{
			name: "Repetition III (3)",
			num:  3,
			want: "III",
		},
		{
			name: "Repetition XX (20)",
			num:  20,
			want: "XX",
		},
		{
			name: "Repetition XXX (30)",
			num:  30,
			want: "XXX",
		},
		{
			name: "Repetition CC (200)",
			num:  200,
			want: "CC",
		},
		{
			name: "Repetition CCC (300)",
			num:  300,
			want: "CCC",
		},
		{
			name: "Repetition MM (2000)",
			num:  2000,
			want: "MM",
		},
		{
			name: "Repetition MMM (3000)",
			num:  3000,
			want: "MMM",
		},

		// --- Additive Forms with 5 / 50 / 500 ---
		{
			name: "Additive VI (6)",
			num:  6,
			want: "VI",
		},
		{
			name: "Additive VII (7)",
			num:  7,
			want: "VII",
		},
		{
			name: "Additive VIII (8)",
			num:  8,
			want: "VIII",
		},
		{
			name: "Additive LX (60)",
			num:  60,
			want: "LX",
		},
		{
			name: "Additive LXX (70)",
			num:  70,
			want: "LXX",
		},
		{
			name: "Additive LXXX (80)",
			num:  80,
			want: "LXXX",
		},
		{
			name: "Additive DC (600)",
			num:  600,
			want: "DC",
		},
		{
			name: "Additive DCC (700)",
			num:  700,
			want: "DCC",
		},
		{
			name: "Additive DCCC (800)",
			num:  800,
			want: "DCCC",
		},

		// --- Numbers with Zero Placeholders ---
		{
			name: "Zero in tens place (101)",
			num:  101,
			want: "CI",
		},
		{
			name: "Zero in ones place (1010)",
			num:  1010,
			want: "MX",
		},
		{
			name: "Zeros in hundreds and tens places (1001)",
			num:  1001,
			want: "MI",
		},
		{
			name: "Zeros in hundreds place (2024)",
			num:  2024,
			want: "MMXXIV",
		},
		{
			name: "Zeros in hundreds and tens (3008)",
			num:  3008,
			want: "MMMVIII",
		},

		// --- Combined Subtractive Patterns ---
		{
			name: "All-fours pattern (444)",
			num:  444,
			want: "CDXLIV",
		},
		{
			name: "All-nines pattern (999)",
			num:  999,
			want: "CMXCIX",
		},
		{
			name: "Year 1492",
			num:  1492,
			want: "MCDXCII",
		},
		{
			name: "Year 1776",
			num:  1776,
			want: "MDCCLXXVI",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intToRoman(tt.num)
			if got != tt.want {
				t.Errorf("intToRoman(%d) = %q, want %q", tt.num, got, tt.want)
			}
		})
	}
}
