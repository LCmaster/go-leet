package findtheindexofthefirstoccurrenceinastring

import (
	"strings"
	"testing"
)

func TestStrStr(t *testing.T) {
	// Pre-generate large strings for boundary constraint tests (length up to 10^4)
	const maxN = 10000

	// 1. Max constraint: 10,000 identical 'a's
	largeAllA := strings.Repeat("a", maxN)
	largeNeedleA := strings.Repeat("a", 500) // matches at 0

	// 2. Worst-case scenario for naive matching: 9,999 'a's + 'b', needle is 1,000 'a's + 'b'
	largeNearMissHaystack := strings.Repeat("a", maxN-1) + "b"
	largeNearMissNeedle := strings.Repeat("a", 1000) + "b" // matches at 9000-1 = 8999

	// 3. Worst-case mismatch: 10,000 'a's, needle is 5,000 'a's + 'b' (not found)
	largeNotFoundNeedle := strings.Repeat("a", 5000) + "b"

	tests := []struct {
		name     string
		haystack string
		needle   string
		want     int
	}{
		// --- 1. LeetCode Official Examples ---
		{
			name:     "Example 1: sadbutsad, needle = sad",
			haystack: "sadbutsad",
			needle:   "sad",
			want:     0,
		},
		{
			name:     "Example 2: leetcode, needle = leeto",
			haystack: "leetcode",
			needle:   "leeto",
			want:     -1,
		},

		// --- 2. Exact Length Matches (len(haystack) == len(needle)) ---
		{
			name:     "Single char match: 'a' in 'a'",
			haystack: "a",
			needle:   "a",
			want:     0,
		},
		{
			name:     "Single char mismatch: 'b' in 'a'",
			haystack: "a",
			needle:   "b",
			want:     -1,
		},
		{
			name:     "Exact match multi-char: 'hello' in 'hello'",
			haystack: "hello",
			needle:   "hello",
			want:     0,
		},
		{
			name:     "Exact length mismatch at end: 'hellp' in 'hello'",
			haystack: "hello",
			needle:   "hellp",
			want:     -1,
		},

		// --- 3. Needle Longer than Haystack ---
		{
			name:     "Needle 2 chars, haystack 1 char",
			haystack: "a",
			needle:   "aa",
			want:     -1,
		},
		{
			name:     "Needle substantially longer than haystack",
			haystack: "short",
			needle:   "muchlongerneedle",
			want:     -1,
		},

		// --- 4. Position Variations ---
		{
			name:     "Needle at the very beginning (index 0)",
			haystack: "prefixandcontent",
			needle:   "prefix",
			want:     0,
		},
		{
			name:     "Needle in the exact middle",
			haystack: "startmidend",
			needle:   "mid",
			want:     5,
		},
		{
			name:     "Needle at the very end",
			haystack: "contentandsuffix",
			needle:   "suffix",
			want:     10,
		},

		// --- 5. Multiple Occurrences (Must return FIRST occurrence) ---
		{
			name:     "Duplicate occurrences at 0 and 6: 'sadbutsad'",
			haystack: "sadbutsad",
			needle:   "sad",
			want:     0,
		},
		{
			name:     "Three occurrences: return index of first",
			haystack: "abc_abc_abc",
			needle:   "abc",
			want:     0,
		},
		{
			name:     "Overlapping pattern: 'aaaaa' with needle 'aaa'",
			haystack: "aaaaa",
			needle:   "aaa",
			want:     0,
		},
		{
			name:     "Overlapping periodic pattern: 'abababa' with needle 'aba'",
			haystack: "abababa",
			needle:   "aba",
			want:     0,
		},

		// --- 6. Partial Matches and Backtracking Traps ---
		{
			name:     "Classic partial match trap: 'mississippi' with 'issip'",
			haystack: "mississippi",
			needle:   "issip",
			want:     4,
		},
		{
			name:     "Repeating prefix fails at last char: 'aaabaaabaaae' with 'aaae'",
			haystack: "aaabaaabaaae",
			needle:   "aaae",
			want:     8,
		},
		{
			name:     "Multiple false starts: 'bbaabbbaaa' with 'baaa'",
			haystack: "bbaabbbaaa",
			needle:   "baaa",
			want:     6,
		},
		{
			name:     "Reset and resume: 'abcabcdabcde' with 'abcde'",
			haystack: "abcabcdabcde",
			needle:   "abcde",
			want:     7,
		},

		// --- 7. Single Character Needles ---
		{
			name:     "Single char at index 0",
			haystack: "abcdef",
			needle:   "a",
			want:     0,
		},
		{
			name:     "Single char in the middle",
			haystack: "abcdef",
			needle:   "d",
			want:     3,
		},
		{
			name:     "Single char at the end",
			haystack: "abcdef",
			needle:   "f",
			want:     5,
		},
		{
			name:     "Single char absent from haystack",
			haystack: "abcdef",
			needle:   "z",
			want:     -1,
		},

		// --- 8. Repetitive / Periodic Strings ---
		{
			name:     "All identical characters: 10 'a's with needle 'aaaa'",
			haystack: "aaaaaaaaaa",
			needle:   "aaaa",
			want:     0,
		},
		{
			name:     "All identical except last character: 'aaaaab' with 'b'",
			haystack: "aaaaab",
			needle:   "b",
			want:     5,
		},
		{
			name:     "Pattern not found in alternating sequence: 'abababab' with 'baaa'",
			haystack: "abababab",
			needle:   "baaa",
			want:     -1,
		},

		// --- 9. Maximum Scale Constraints (Length = 10^4) ---
		{
			name:     "Max constraint 10^4: all 'a's with needle of 500 'a's",
			haystack: largeAllA,
			needle:   largeNeedleA,
			want:     0,
		},
		{
			name:     "Max constraint 10^4: near-miss match at the end",
			haystack: largeNearMissHaystack,
			needle:   largeNearMissNeedle,
			want:     maxN - 1 - 1000,
		},
		{
			name:     "Max constraint 10^4: needle not found after full scan",
			haystack: largeAllA,
			needle:   largeNotFoundNeedle,
			want:     -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strStr(tt.haystack, tt.needle)
			if got != tt.want {
				t.Errorf("strStr(%q, %q) = %d; want %d", tt.haystack, tt.needle, got, tt.want)
			}
		})
	}
}

func BenchmarkStrStr(b *testing.B) {
	haystack := strings.Repeat("a", 10000) + "b"
	needle := strings.Repeat("a", 500) + "b"

	b.ResetTimer()
	for b.Loop() {
		_ = strStr(haystack, needle)
	}
}
