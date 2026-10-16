package substringwithconcatenationofallwords

import (
	"slices"
	"strings"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	// Pre-generate large dataset for maximum constraint testing (len(s) up to 10^4)
	largeSAllA := strings.Repeat("a", 1000)
	largeWordsAllA := make([]string, 500)
	for i := range largeWordsAllA {
		largeWordsAllA[i] = "a"
	}
	wantLargeAllA := make([]int, 501)
	for i := range wantLargeAllA {
		wantLargeAllA[i] = i
	}

	tests := []struct {
		name  string
		s     string
		words []string
		want  []int
	}{
		// --- 1. LeetCode Official Examples ---
		{
			name:  "Example 1: barfoothefoobarman with [foo, bar]",
			s:     "barfoothefoobarman",
			words: []string{"foo", "bar"},
			want:  []int{0, 9},
		},
		{
			name:  "Example 2: wordgoodgoodgoodbestword with [word, good, best, word]",
			s:     "wordgoodgoodgoodbestword",
			words: []string{"word", "good", "best", "word"},
			want:  []int{},
		},
		{
			name:  "Example 3: barfoofoobarthefoobarman with [bar, foo, the]",
			s:     "barfoofoobarthefoobarman",
			words: []string{"bar", "foo", "the"},
			want:  []int{6, 9, 12},
		},

		// --- 2. String Shorter than Total Concatenated Length ---
		{
			name:  "String shorter than word length: len(s) = 2, word length = 3",
			s:     "ab",
			words: []string{"abc"},
			want:  []int{},
		},
		{
			name:  "String shorter than total words length: len(s) = 3, 2 words of length 2",
			s:     "abc",
			words: []string{"ab", "cd"},
			want:  []int{},
		},

		// --- 3. Exact Length Matching (len(s) == total words length) ---
		{
			name:  "Exact match single permutation",
			s:     "abcdef",
			words: []string{"ab", "cd", "ef"},
			want:  []int{0},
		},
		{
			name:  "Exact match permuted order",
			s:     "efabcd",
			words: []string{"ab", "cd", "ef"},
			want:  []int{0},
		},
		{
			name:  "Exact length mismatch by one character",
			s:     "abcdef",
			words: []string{"ab", "cd", "zz"},
			want:  []int{},
		},

		// --- 4. Single Word in Words Array ---
		{
			name:  "Single word with single occurrence",
			s:     "hello world",
			words: []string{"world"},
			want:  []int{6},
		},
		{
			name:  "Single word with multiple occurrences",
			s:     "foobarfoo",
			words: []string{"foo"},
			want:  []int{0, 6},
		},
		{
			name:  "Single 1-character word with overlapping occurrences",
			s:     "aaaa",
			words: []string{"a"},
			want:  []int{0, 1, 2, 3},
		},
		{
			name:  "Single word absent from string",
			s:     "abcdef",
			words: []string{"xyz"},
			want:  []int{},
		},

		// --- 5. Duplicate Words in Words List ---
		{
			name:  "Duplicate words required: s has matching frequencies",
			s:     "wordgoodgoodgoodbestword",
			words: []string{"word", "good", "best", "good"},
			want:  []int{8},
		},
		{
			name:  "All identical 1-character words: [a, a] in 'aaa'",
			s:     "aaa",
			words: []string{"a", "a"},
			want:  []int{0, 1},
		},
		{
			name:  "All identical 2-character words: [aa, aa] in 'aaaaaa'",
			s:     "aaaaaa",
			words: []string{"aa", "aa"},
			want:  []int{0, 1, 2},
		},

		// --- 6. Non-zero Alignment Offsets (Testing all offset phases mod wordLen) ---
		{
			name:  "Match starts at offset 1: wordLen = 3",
			s:     "xfoobar",
			words: []string{"foo", "bar"},
			want:  []int{1},
		},
		{
			name:  "Match starts at offset 2: wordLen = 3",
			s:     "xxfoobar",
			words: []string{"foo", "bar"},
			want:  []int{2},
		},
		{
			name:  "Multiple matches at different offset modulos",
			s:     "abacba",
			words: []string{"a", "b"},
			want:  []int{0, 1, 4},
		},

		// --- 7. Overlapping & Cascading Matches ---
		{
			name:  "LeetCode classic benchmark: lingmindraboofooowingdingbarrwingmonkeypoundcake",
			s:     "lingmindraboofooowingdingbarrwingmonkeypoundcake",
			words: []string{"fooo", "barr", "wing", "ding", "wing"},
			want:  []int{13},
		},

		// --- 8. Words with Unknown Tokens or Interspersed Garbage ---
		{
			name:  "Interspersed non-matching word resets window",
			s:     "barfooxxxfoobar",
			words: []string{"foo", "bar"},
			want:  []int{0, 9},
		},
		{
			name:  "Word appears more times in window than required",
			s:     "foofoofoo",
			words: []string{"foo", "foo"},
			want:  []int{0, 3},
		},
		{
			name:  "Characters in words never appear in s",
			s:     "abcdefgh",
			words: []string{"xy", "zz"},
			want:  []int{},
		},

		// --- 9. Long Words (wordLen up to 30) ---
		{
			name: "Long words: wordLen = 10, total window = 30",
			s:    "prefix0123456789abcdefghijqwertyuiopxyz",
			words: []string{
				"0123456789",
				"abcdefghij",
				"qwertyuiop",
			},
			want: []int{6},
		},

		// --- 10. Large Scale / Constraint Boundaries ---
		{
			name:  "Large scale: 1000 'a's with 500 words of 'a'",
			s:     largeSAllA,
			words: largeWordsAllA,
			want:  wantLargeAllA,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findSubstring(tt.s, tt.words)

			if len(got) == 0 && len(tt.want) == 0 {
				return
			}

			gotSorted := slices.Clone(got)
			slices.Sort(gotSorted)

			wantSorted := slices.Clone(tt.want)
			slices.Sort(wantSorted)

			if !slices.Equal(gotSorted, wantSorted) {
				t.Errorf("findSubstring(%q, %v) = %v; want %v", tt.s, tt.words, got, tt.want)
			}
		})
	}
}

func BenchmarkFindSubstring(b *testing.B) {
	s := strings.Repeat("barfoothefoobarman", 50)
	words := []string{"foo", "bar", "the"}

	b.ResetTimer()
	for b.Loop() {
		_ = findSubstring(s, words)
	}
}
