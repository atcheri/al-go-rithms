package mergesort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	nums     []int
	expected []int
}

type testCaseMergeArrays struct {
	name     string
	a        []int
	b        []int
	expected []int
}

func Test_MergeSort(t *testing.T) {
	testCases := []testCase{
		{
			name:     "returns the first element for a 1 element slice",
			nums:     []int{1},
			expected: []int{1},
		},
		{
			name:     "returns {1, 2} for {2, 1}",
			nums:     []int{2, 1},
			expected: []int{1, 2},
		},
		{
			name:     "returns {1, 2, 3, 4, 5, 6} for {2, 1, 3, 6, 5, 4}",
			nums:     []int{2, 1, 3, 6, 5, 4},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tc := range testCases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			sorted := MergeSort(c.nums)
			assert.Equal(t, c.expected, sorted)
		})
	}
}

func Test_MergeArrays(t *testing.T) {
	testCases := []testCaseMergeArrays{
		{
			name:     "returns {1, 2} for {1} and {2}",
			a:        []int{1},
			b:        []int{2},
			expected: []int{1, 2},
		},
		{
			name:     "returns {1, 2, 3} for {1, 3} and {2}",
			a:        []int{1, 3},
			b:        []int{2},
			expected: []int{1, 2, 3},
		},
		{
			name:     "returns {1, 2, 3, 4} for {1, 3} and {2, 4}",
			a:        []int{1, 3},
			b:        []int{2, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "returns {1, 2, 3, 4, 5} for {1, 3, 5} and {2, 4}",
			a:        []int{1, 3, 5},
			b:        []int{2, 4},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range testCases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			sorted := MergeArrays(c.a, c.b)
			assert.Equal(t, c.expected, sorted)
		})
	}
}
