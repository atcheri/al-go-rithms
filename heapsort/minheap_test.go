package heapsort_test

import (
	"al-go-rithms/heapsort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewMinHeapFrom(t *testing.T) {
	type testCase struct {
		name   string
		input  []int
		output []int
	}

	cases := []testCase{
		{
			name:   "for an input of {10, 7, 11, 5, 4, 13}, creates a MinHeap with {4, 5, 11, 10, 7, 13}",
			input:  []int{10, 7, 11, 5, 4, 13},
			output: []int{4, 5, 11, 10, 7, 13},
		},
		{
			name:   "for an input of {10, 5, 15, 2, 20, 30}, creates a MinHeap with {2, 5, 15, 10, 20, 30}",
			input:  []int{10, 5, 15, 2, 20, 30},
			output: []int{2, 5, 15, 10, 20, 30},
		},
	}

	for _, tc := range cases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			// arrange and act
			sut := heapsort.NewMinHeapFrom(c.input)

			// assert
			assert.Equal(t, c.output, sut.ToSlice())
		})
	}
}
