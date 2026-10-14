package romantointeger

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		// --- LeetCode Official Examples ---
		{
			name: "Example 1 - III",
			s:    "III",
			want: 3,
		},
		{
			name: "Example 2 - LVIII",
			s:    "LVIII",
			want: 58,
		},
		{
			name: "Example 3 - MCMXCIV",
			s:    "MCMXCIV",
			want: 1994,
		},

		// --- Boundary Constraints ---
		{
			name: "Minimum constraint - I (1)",
			s:    "I",
			want: 1,
		},
		{
			name: "Maximum constraint - MMMCMXCIX (3999)",
			s:    "MMMCMXCIX",
			want: 3999,
		},
		{
			name: "Maximum length Roman numeral - MMMDCCCLXXXVIII (3888, len 15)",
			s:    "MMMDCCCLXXXVIII",
			want: 3888,
		},

		// --- Single Roman Symbols ---
		{
			name: "Single symbol I (1)",
			s:    "I",
			want: 1,
		},
		{
			name: "Single symbol V (5)",
			s:    "V",
			want: 5,
		},
		{
			name: "Single symbol X (10)",
			s:    "X",
			want: 10,
		},
		{
			name: "Single symbol L (50)",
			s:    "L",
			want: 50,
		},
		{
			name: "Single symbol C (100)",
			s:    "C",
			want: 100,
		},
		{
			name: "Single symbol D (500)",
			s:    "D",
			want: 500,
		},
		{
			name: "Single symbol M (1000)",
			s:    "M",
			want: 1000,
		},

		// --- Subtractive Cases ---
		{
			name: "Subtractive IV (4)",
			s:    "IV",
			want: 4,
		},
		{
			name: "Subtractive IX (9)",
			s:    "IX",
			want: 9,
		},
		{
			name: "Subtractive XL (40)",
			s:    "XL",
			want: 40,
		},
		{
			name: "Subtractive XC (90)",
			s:    "XC",
			want: 90,
		},
		{
			name: "Subtractive CD (400)",
			s:    "CD",
			want: 400,
		},
		{
			name: "Subtractive CM (900)",
			s:    "CM",
			want: 900,
		},

		// --- Repeated Symbols (Powers of 10) ---
		{
			name: "Repetition II (2)",
			s:    "II",
			want: 2,
		},
		{
			name: "Repetition XX (20)",
			s:    "XX",
			want: 20,
		},
		{
			name: "Repetition XXX (30)",
			s:    "XXX",
			want: 30,
		},
		{
			name: "Repetition CC (200)",
			s:    "CC",
			want: 200,
		},
		{
			name: "Repetition CCC (300)",
			s:    "CCC",
			want: 300,
		},
		{
			name: "Repetition MM (2000)",
			s:    "MM",
			want: 2000,
		},
		{
			name: "Repetition MMM (3000)",
			s:    "MMM",
			want: 3000,
		},

		// --- Additive Combinations with 5 / 50 / 500 ---
		{
			name: "Additive VI (6)",
			s:    "VI",
			want: 6,
		},
		{
			name: "Additive VII (7)",
			s:    "VII",
			want: 7,
		},
		{
			name: "Additive VIII (8)",
			s:    "VIII",
			want: 8,
		},
		{
			name: "Additive LX (60)",
			s:    "LX",
			want: 60,
		},
		{
			name: "Additive LXX (70)",
			s:    "LXX",
			want: 70,
		},
		{
			name: "Additive LXXX (80)",
			s:    "LXXX",
			want: 80,
		},
		{
			name: "Additive DC (600)",
			s:    "DC",
			want: 600,
		},
		{
			name: "Additive DCC (700)",
			s:    "DCC",
			want: 700,
		},
		{
			name: "Additive DCCC (800)",
			s:    "DCCC",
			want: 800,
		},

		// --- Multi-Digit Pure Additive ---
		{
			name: "Additive XII (12)",
			s:    "XII",
			want: 12,
		},
		{
			name: "Additive XXVII (27)",
			s:    "XXVII",
			want: 27,
		},
		{
			name: "Additive CLXI (161)",
			s:    "CLXI",
			want: 161,
		},
		{
			name: "All symbols in descending order - MDCLXVI (1666)",
			s:    "MDCLXVI",
			want: 1666,
		},

		// --- Numbers with Zero Placeholders ---
		{
			name: "Zero in tens place - CI (101)",
			s:    "CI",
			want: 101,
		},
		{
			name: "Zero in ones place - MX (1010)",
			s:    "MX",
			want: 1010,
		},
		{
			name: "Zeros in hundreds and tens places - MI (1001)",
			s:    "MI",
			want: 1001,
		},
		{
			name: "Zero in hundreds place - MMXXIV (2024)",
			s:    "MMXXIV",
			want: 2024,
		},
		{
			name: "Zeros in hundreds and tens - MMMVIII (3008)",
			s:    "MMMVIII",
			want: 3008,
		},

		// --- Combined Subtractive Patterns ---
		{
			name: "All-fours pattern - CDXLIV (444)",
			s:    "CDXLIV",
			want: 444,
		},
		{
			name: "All-nines pattern - CMXCIX (999)",
			s:    "CMXCIX",
			want: 999,
		},
		{
			name: "Mixed subtractive - MMMDCCXLIX (3749)",
			s:    "MMMDCCXLIX",
			want: 3749,
		},
		{
			name: "Year 1492 - MCDXCII",
			s:    "MCDXCII",
			want: 1492,
		},
		{
			name: "Year 1776 - MDCCLXXVI",
			s:    "MDCCLXXVI",
			want: 1776,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := romanToInt(tt.s)
			if got != tt.want {
				t.Errorf("romanToInt(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
