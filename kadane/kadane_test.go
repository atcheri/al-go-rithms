package kadane_test

import (
	"al-go-rithms/kadane"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	nums     []int
	expected []int
}

func Test_MaxSum(t *testing.T) {
	testCases := []testCase{
		{
			name:     "returns a slice with the 1st item for a slice with 1 item",
			nums:     []int{1},
			expected: []int{1},
		},
		{
			name:     "returns a slice with the 1st item for [1, -1]",
			nums:     []int{1, -1},
			expected: []int{1},
		},
		{
			name:     "returns a slice with the 1st item for [-1, 1]",
			nums:     []int{-1, 1},
			expected: []int{1},
		},
		{
			name:     "returns a slice with the 2 itemx for [1, -1, 3]",
			nums:     []int{1, -1, 3},
			expected: []int{1, -1, 3},
		},
		{
			name:     "returns {4, -1, 2, 1} for {-2, 1, -3, 4, -1, 2, 1, -5, 4}",
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: []int{4, -1, 2, 1},
		},
		{
			name:     "returns {4, -1, 2, 1} for {-2, 1, -3, 4, -1, 2, 1, -5, 4}",
			nums:     []int{1, 3, 8, -2, 6, -8, 5},
			expected: []int{1, 3, 8, -2, 6},
		},
		{
			name:     "returns {} for {-2, -3, 4, -1, -2, 1, -5, -3}",
			nums:     []int{-2, -3, 4, -1, -2, 1, 5, -3},
			expected: []int{4, -1, -2, 1, 5},
		},
		{
			name:     "returns {} for {4, -1, 2, -7, 3, 4}",
			nums:     []int{4, -1, 2, -7, 3, 4},
			expected: []int{3, 4},
		},
	}

	for _, tc := range testCases {
		sum := kadane.MaxSum(tc.nums)
		assert.Equal(t, tc.expected, sum)

	}
}
