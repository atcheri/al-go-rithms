package heapsort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_heapifyDown(t *testing.T) {
	type testCase struct {
		name   string
		input  []int
		output []int
	}

	cases := []testCase{
		{
			name:   "for a heap created from {1, 4, 2, 5, 13, 6, 17}, returns {}",
			input:  []int{1, 4, 2, 5, 13, 6, 17},
			output: []int{13, 5, 6, 1, 4, 2},
		},
		{
			name:   "for a heap created from {10, 5, 3, 2, 4, 15}, returns {}",
			input:  []int{10, 5, 3, 2, 4, 15, 19, 6},
			output: []int{15, 6, 10, 5, 4, 3, 2},
		},
	}
	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			// arrange
			heap := NewMaxHeapFrom(c.input)
			sut := make([]int, len(c.input))
			copy(sut, heap.ToSlice())
			sut[0], sut[len(sut)-1] = sut[len(sut)-1], sut[0]
			sut = sut[:len(sut)-1]

			// act
			res := heapifyDown(sut, 0)

			// assert
			assert.Equal(t, c.output, res)
		})
	}

}
