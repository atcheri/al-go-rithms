package binarysearch_test

import (
	"al-go-rithms/binarysearch"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	nums     []int
	target   int
	expected int
}

func Test_BinarySearch(t *testing.T) {
	testCases := []testCase{
		{
			name:     "empty slice returns -1",
			nums:     []int{},
			target:   3,
			expected: -1,
		},
		{
			name:     "a slice of size 1 returns 0 if target is found",
			nums:     []int{1},
			target:   2,
			expected: -1,
		},
		{
			name:     "a slice of size 1 returns 0 if target is found",
			nums:     []int{1},
			target:   1,
			expected: 0,
		},
		{
			name:     "a slice of sorted random numbers returns -1 if target is not found",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:   11,
			expected: -1,
		},
		{
			name:     "a slice of sorted random numbers returns the correct index if target is found",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:   6,
			expected: 5,
		}, {
			name:     "a slice of sorted random numbers returns the correct index if target is found",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:   3,
			expected: 2,
		},
		{
			name:     "a slice of sorted random numbers returns the correct index if target is found",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:   9,
			expected: 8,
		},
	}

	for _, tc := range testCases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := binarysearch.BinarySearch(c.nums, c.target)
			assert.Equal(t, c.expected, res)
		})
	}

}
