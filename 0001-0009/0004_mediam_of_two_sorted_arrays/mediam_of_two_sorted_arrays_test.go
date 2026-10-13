package mediamoftwosortedarrays

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			name:  "Example 1",
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.00000,
		},
		{
			name:  "Example 2",
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.50000,
		},
		{
			name:  "nums1 is empty",
			nums1: []int{},
			nums2: []int{1},
			want:  1.00000,
		},
		{
			name:  "nums2 is empty",
			nums1: []int{2, 3},
			nums2: []int{},
			want:  2.50000,
		},
		{
			name:  "Negative numbers",
			nums1: []int{-5, 3, 6, 12, 15},
			nums2: []int{-12, -10, -6, -3, 4, 10},
			want:  3.00000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays(%v, %v) = %v, want %v", tt.nums1, tt.nums2, got, tt.want)
			}
		})
	}
}
