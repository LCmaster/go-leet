package longestcommonprefix

import (
	"strings"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		// --- LeetCode Official Examples ---
		{
			name: "Example 1 - partial common prefix",
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "Example 2 - no common prefix",
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},

		// --- Single Element Arrays (Length = 1) ---
		{
			name: "Single word in array",
			strs: []string{"alone"},
			want: "alone",
		},
		{
			name: "Single empty string in array",
			strs: []string{""},
			want: "",
		},
		{
			name: "Single character in array",
			strs: []string{"a"},
			want: "a",
		},

		// --- Empty Strings in Array (String Length = 0) ---
		{
			name: "Empty string at beginning",
			strs: []string{"", "b", "c"},
			want: "",
		},
		{
			name: "Empty string in middle",
			strs: []string{"a", "", "c"},
			want: "",
		},
		{
			name: "Empty string at end",
			strs: []string{"a", "b", ""},
			want: "",
		},
		{
			name: "All empty strings",
			strs: []string{"", "", ""},
			want: "",
		},

		// --- Entire String is Prefix ---
		{
			name: "First string is full prefix",
			strs: []string{"ab", "abc", "abcd"},
			want: "ab",
		},
		{
			name: "Middle string is full prefix",
			strs: []string{"interstellar", "inter", "internal"},
			want: "inter",
		},
		{
			name: "Last string is full prefix",
			strs: []string{"prefixes", "prefixing", "pre"},
			want: "pre",
		},

		// --- All Strings Identical ---
		{
			name: "Multiple identical strings",
			strs: []string{"hello", "hello", "hello"},
			want: "hello",
		},
		{
			name: "Two identical strings",
			strs: []string{"same", "same"},
			want: "same",
		},
		{
			name: "Single repeating character across all strings",
			strs: []string{"z", "z", "z", "z"},
			want: "z",
		},

		// --- No Common Prefix Cases ---
		{
			name: "Immediate first character mismatch",
			strs: []string{"apple", "banana", "cherry"},
			want: "",
		},
		{
			name: "Prefix shared by some but not all strings",
			strs: []string{"cat", "caterpillar", "dog", "cattle"},
			want: "",
		},
		{
			name: "First and last share prefix but middle differs",
			strs: []string{"prefix_a", "other", "prefix_b"},
			want: "",
		},

		// --- Single Character Prefix ---
		{
			name: "Two elements sharing only first character",
			strs: []string{"cir", "car"},
			want: "c",
		},
		{
			name: "Multiple words sharing only first letter",
			strs: []string{"zebra", "zoo", "zone", "zero"},
			want: "z",
		},

		// --- Varying Word Lengths ---
		{
			name: "First element is single character",
			strs: []string{"a", "ab", "abc", "abcd"},
			want: "a",
		},
		{
			name: "First element is very long, others shorter",
			strs: []string{"interspecies", "interstellar", "interstate"},
			want: "inters",
		},
		{
			name: "Common prefix across two long words",
			strs: []string{"antidisestablishmentarianism", "antidisestablishment"},
			want: "antidisestablishment",
		},

		// --- Late Mismatch (Subtle Divergence) ---
		{
			name: "Divergence at the last character",
			strs: []string{"abcdefghij", "abcdefghik", "abcdefghil"},
			want: "abcdefghi",
		},
		{
			name: "Divergence near the end of long strings",
			strs: []string{"supercalifragilistic", "supercalifragilist", "supercalifragile"},
			want: "supercalifragil",
		},

		// --- Repeated Characters Pattern ---
		{
			name: "All same characters with different lengths",
			strs: []string{"aaaa", "aa", "aaaaaa"},
			want: "aa",
		},
		{
			name: "Alternating characters mismatching early",
			strs: []string{"abab", "aba"},
			want: "aba",
		},

		// --- Maximum Constraints Scale ---
		{
			name: "Many strings (50 elements) with common prefix",
			strs: func() []string {
				arr := make([]string, 50)
				for i := range arr {
					arr[i] = "commonprefix" + strings.Repeat("x", i)
				}
				return arr
			}(),
			want: "commonprefix",
		},
		{
			name: "Long string (200 characters) common prefix",
			strs: func() []string {
				base := strings.Repeat("a", 195)
				return []string{base + "bcde", base + "bcfg", base + "bchj"}
			}(),
			want: strings.Repeat("a", 195) + "bc",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonPrefix(tt.strs)
			if got != tt.want {
				t.Errorf("longestCommonPrefix(%v) = %q, want %q", tt.strs, got, tt.want)
			}
		})
	}
}
