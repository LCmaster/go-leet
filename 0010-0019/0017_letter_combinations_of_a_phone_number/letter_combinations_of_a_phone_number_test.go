package lettercombinationsofaphonenumber

import (
	"reflect"
	"slices"
	"testing"
)

// sortSlice returns a sorted copy of a string slice to ensure test assertions
// are order-independent (LeetCode specifies: "Return the answer in any order").
func sortSlice(s []string) []string {
	if s == nil {
		return []string{}
	}
	cp := make([]string, len(s))
	copy(cp, s)
	slices.Sort(cp)
	return cp
}

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		// --- 1. Empty Input (Edge Case) ---
		{
			name:   "Empty string returns empty array",
			digits: "",
			want:   []string{},
		},

		// --- 2. Official LeetCode Examples ---
		{
			name:   "Example 1: digits 23",
			digits: "23",
			want:   []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:   "Example 2: single digit 2",
			digits: "2",
			want:   []string{"a", "b", "c"},
		},

		// --- 3. Single Digit (Length 1) ---
		{
			name:   "Single digit 3 (def)",
			digits: "3",
			want:   []string{"d", "e", "f"},
		},
		{
			name:   "Single digit 7 with 4 letters (pqrs)",
			digits: "7",
			want:   []string{"p", "q", "r", "s"},
		},
		{
			name:   "Single digit 8 (tuv)",
			digits: "8",
			want:   []string{"t", "u", "v"},
		},
		{
			name:   "Single digit 9 with 4 letters (wxyz)",
			digits: "9",
			want:   []string{"w", "x", "y", "z"},
		},

		// --- 4. Two Digits (Length 2) ---
		{
			name:   "Identical digits 22",
			digits: "22",
			want:   []string{"aa", "ab", "ac", "ba", "bb", "bc", "ca", "cb", "cc"},
		},
		{
			name:   "3-letter and 4-letter digits 27 (12 combinations)",
			digits: "27",
			want: []string{
				"ap", "aq", "ar", "as",
				"bp", "bq", "br", "bs",
				"cp", "cq", "cr", "cs",
			},
		},
		{
			name:   "Two 4-letter digits 79 (16 combinations)",
			digits: "79",
			want: []string{
				"pw", "px", "py", "pz",
				"qw", "qx", "qy", "qz",
				"rw", "rx", "ry", "rz",
				"sw", "sx", "sy", "sz",
			},
		},

		// --- 5. Three Digits (Length 3) ---
		{
			name:   "Three 3-letter digits 234 (27 combinations)",
			digits: "234",
			want: []string{
				"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi",
				"bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi",
				"cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi",
			},
		},
		{
			name:   "Three digits with 4-letter digit 237 (36 combinations)",
			digits: "237",
			want: []string{
				"adp", "adq", "adr", "ads", "aep", "aeq", "aer", "aes", "afp", "afq", "afr", "afs",
				"bdp", "bdq", "bdr", "bds", "bep", "beq", "ber", "bes", "bfp", "bfq", "bfr", "bfs",
				"cdp", "cdq", "cdr", "cds", "cep", "ceq", "cer", "ces", "cfp", "cfq", "cfr", "cfs",
			},
		},

		// --- 6. Maximum Constraint (Length 4) ---
		{
			name:   "Max length 4 with four 3-letter digits 2345 (81 combinations)",
			digits: "2345",
			want: func() []string {
				m := map[byte]string{
					'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
				}
				var res []string
				for _, c1 := range m['2'] {
					for _, c2 := range m['3'] {
						for _, c3 := range m['4'] {
							for _, c4 := range m['5'] {
								res = append(res, string([]rune{c1, c2, c3, c4}))
							}
						}
					}
				}
				return res
			}(),
		},
		{
			name:   "Max length 4 with four 4-letter digits 7777 (256 combinations)",
			digits: "7777",
			want: func() []string {
				letters := "pqrs"
				var res []string
				for _, c1 := range letters {
					for _, c2 := range letters {
						for _, c3 := range letters {
							for _, c4 := range letters {
								res = append(res, string([]rune{c1, c2, c3, c4}))
							}
						}
					}
				}
				return res
			}(),
		},
		{
			name:   "Max length 4 with mixed 3 and 4-letter digits 7892 (144 combinations)",
			digits: "7892",
			want: func() []string {
				m := map[byte]string{
					'7': "pqrs", '8': "tuv", '9': "wxyz", '2': "abc",
				}
				var res []string
				for _, c1 := range m['7'] {
					for _, c2 := range m['8'] {
						for _, c3 := range m['9'] {
							for _, c4 := range m['2'] {
								res = append(res, string([]rune{c1, c2, c3, c4}))
							}
						}
					}
				}
				return res
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := letterCombinations(tt.digits)

			// Both empty slice and nil are acceptable for empty input
			if len(tt.want) == 0 {
				if len(got) != 0 {
					t.Errorf("letterCombinations(%q) = %v (len %d), want empty slice", tt.digits, got, len(got))
				}
				return
			}

			sortedGot := sortSlice(got)
			sortedWant := sortSlice(tt.want)

			if len(sortedGot) != len(sortedWant) {
				t.Fatalf("letterCombinations(%q) returned %d combinations, want %d", tt.digits, len(sortedGot), len(sortedWant))
			}

			if !reflect.DeepEqual(sortedGot, sortedWant) {
				t.Errorf("letterCombinations(%q)\ngot:  %v\nwant: %v", tt.digits, sortedGot, sortedWant)
			}
		})
	}
}

func BenchmarkLetterCombinations(b *testing.B) {
	digits := "7892" // 144 combinations
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = letterCombinations(digits)
	}
}
