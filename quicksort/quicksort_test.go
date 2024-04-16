package quicksort_test

import (
	"al-go-rithms/quicksort"
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
		{
			name:     "returns {-9, -6, -3, 1, 2, 3, 5, 6, 8, 9} for {9, -3, 5, 2, -9, 6, 8, -6, 1, 3}",
			nums:     []int{9, -3, 5, 2, -9, 6, 8, -6, 1, 3},
			expected: []int{-9, -6, -3, 1, 2, 3, 5, 6, 8, 9},
		},
	}

	for _, tc := range testCases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			sorted := quicksort.QuickSortRecursive(c.nums)
			assert.Equal(t, c.expected, sorted)
		})
	}
}
