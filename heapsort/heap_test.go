package heapsort_test

import (
	"al-go-rithms/heapsort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewHeap(t *testing.T) {
	// arrange
	// act
	heap := heapsort.NewMaxHeap()

	// assert
	assert.NotNil(t, heap)
	assert.Equal(t, heap.Size(), 0)
}

func Test_Insert(t *testing.T) {
	type testCase struct {
		name   string
		input  []int
		output []int
	}
	cases := []testCase{
		{
			name:   "inserting 5, shows a one item slice",
			input:  []int{5},
			output: []int{5},
		},
		{
			name:   "inserting 5 and 6, shows a 2 items slice",
			input:  []int{5, 6},
			output: []int{6, 5},
		},
		{
			name:   "inserting 5, 6 4, and 8, returns 8 6 5 4",
			input:  []int{5, 6, 4, 8},
			output: []int{8, 6, 4, 5},
		},
		{
			name:   "inserting 10, 5, 3, 2, 4 and 15, returns 15 5 10 2 4 3",
			input:  []int{10, 5, 3, 2, 4, 15},
			output: []int{15, 5, 10, 2, 4, 3},
		},
		{
			name:   "inserting 1, 4, 2, 5, 13, 6, 17, returns 17, 5, 13, 5, 1, 4, 2, 6",
			input:  []int{1, 4, 2, 5, 13, 6, 17},
			output: []int{17, 5, 13, 1, 4, 2, 6},
		},
	}

	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			// arrange
			heap := heapsort.NewMaxHeap()

			// act
			for _, item := range c.input {
				heap.Insert(item)
			}

			// assert
			assert.Equal(t, heap.Size(), len(c.output))
			assert.Equal(t, heap.ToSlice(), c.output)
		})
	}
}

func Test_Pull(t *testing.T) {
	type testCase struct {
		name   string
		input  []int
		output int
	}
	cases := []testCase{
		{
			name:   "returns the 1st element for a 1 element heap",
			input:  []int{1},
			output: 1,
		}, {
			name:   "returns 3 for a heap with created from {1, 2, 3}",
			input:  []int{1, 2, 3},
			output: 3,
		},
		{
			name:   "returns 15 for a heap created from {10, 5, 3, 2, 4, 15}",
			input:  []int{10, 5, 3, 2, 4, 15},
			output: 15,
		},
		{
			name:   "returns 17 for a heap created from {1, 4, 2, 5, 13, 6, 17}",
			input:  []int{1, 4, 2, 5, 13, 6, 17},
			output: 17,
		},
	}

	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			// arrange
			heap := heapsort.NewMaxHeapFrom(c.input)

			// act
			max := heap.Pull()

			// assert
			assert.Equal(t, len(c.input)-1, heap.Size())
			assert.Equal(t, c.output, max)
		})
	}
}

func Test_Sort(t *testing.T) {
	type testCase struct {
		name   string
		input  []int
		output []int
	}

	cases := []testCase{
		{
			name:   "for a heap created from {5}, sorts and returns {5}",
			input:  []int{5},
			output: []int{5},
		},
		{
			name:   "for a heap created from {1, 2}, sorts and returns {1, 2}",
			input:  []int{1, 2},
			output: []int{1, 2},
		},
		{
			name:   "for a heap created from {10, 5, 3, 2, 4, 15}, sorts and returns {2, 3, 4, 5, 10, 15}",
			input:  []int{10, 5, 3, 2, 4, 15},
			output: []int{2, 3, 4, 5, 10, 15},
		},
		{
			name:   "for a heap created from {1, 4, 2, 5, 13, 6, 17}, sorts and returns {1, 2, 4, 5, 6, 13, 17}",
			input:  []int{1, 4, 2, 5, 13, 6, 17},
			output: []int{1, 2, 4, 5, 6, 13, 17},
		},
	}

	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			// arrange
			heap := heapsort.NewMaxHeapFrom(c.input)

			// act
			sorted := heap.Sort()

			// assert
			assert.Equal(t, len(c.input), heap.Size())
			assert.Equal(t, c.output, sorted)
		})
	}

}
