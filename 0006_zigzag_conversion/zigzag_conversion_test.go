package zigzagconversion

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		numRows int
		want    string
	}{
		{
			name:    "Example 1",
			s:       "PAYPALISHIRING",
			numRows: 3,
			want:    "PAHNAPLSIIGYIR",
		},
		{
			name:    "Example 2",
			s:       "PAYPALISHIRING",
			numRows: 4,
			want:    "PINALSIGYAHRPI",
		},
		{
			name:    "Example 3",
			s:       "A",
			numRows: 1,
			want:    "A",
		},
		{
			name:    "numRows is 1 with longer string",
			s:       "AB",
			numRows: 1,
			want:    "AB",
		},
		{
			name:    "numRows is 2",
			s:       "PAYPALISHIRING",
			numRows: 2,
			want:    "PYAIHRNAPLSIIG",
		},
		{
			name:    "len(s) <= numRows",
			s:       "ABC",
			numRows: 5,
			want:    "ABC",
		},
		{
			name:    "Two characters with numRows 2",
			s:       "AB",
			numRows: 2,
			want:    "AB",
		},
		{
			name:    "String with comma and period",
			s:       "A,B.C",
			numRows: 2,
			want:    "ABC,.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.s, tt.numRows); got != tt.want {
				t.Errorf("convert(%q, %d) = %q, want %q", tt.s, tt.numRows, got, tt.want)
			}
		})
	}
}
