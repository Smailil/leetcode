package main

import (
	"testing"
)

func TestIsReflected(t *testing.T) {
	testCases := []struct {
		name   string
		points [][]int
		want   bool
	}{
		{
			name:   "basic_symmetric",
			points: [][]int{{1, 1}, {-1, 1}},
			want:   true,
		},
		{
			name:   "example_with_multiple_points",
			points: [][]int{{3, 1}, {2, 1}, {0, 1}, {-1, 1}},
			want:   true,
		},
		{
			name:   "not_symmetric",
			points: [][]int{{1, 1}, {-1, 1}, {2, 2}},
			want:   false,
		},
		{
			name:   "all_points_same",
			points: [][]int{{0, 0}, {0, 0}, {0, 0}},
			want:   true,
		},
		{
			name:   "asymmetric_even_count",
			points: [][]int{{0, 0}, {1, 0}, {3, 0}, {4, 0}},
			want:   true,
		},
		{
			name:   "symmetric_with_duplicates",
			points: [][]int{{1, 1}, {1, 1}, {-1, 1}, {-1, 1}},
			want:   true,
		},
		{
			name:   "odd_count_symmetric",
			points: [][]int{{-1, 1}, {0, 1}, {1, 1}},
			want:   true,
		},
		{
			name:   "different_y_coordinates",
			points: [][]int{{1, 2}, {2, 2}, {-1, 2}, {-2, 2}, {1, 3}, {-1, 3}},
			want:   true,
		},
		{
			name:   "no_symmetry_different_y",
			points: [][]int{{1, 1}, {-1, 2}},
			want:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isReflected(tc.points)
			if got != tc.want {
				t.Errorf("isReflected(%v) = %v, want %v", tc.points, got, tc.want)
			}
		})
	}
}
