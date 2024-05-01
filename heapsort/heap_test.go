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
