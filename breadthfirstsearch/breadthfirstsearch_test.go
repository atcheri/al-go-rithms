package breadthfirstsearch_test

import (
	"al-go-rithms/breadthfirstsearch"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	tree     breadthfirstsearch.BinaryTree
	target   int
	expected bool
}

var (
	biggerBinaryTree = breadthfirstsearch.BinaryTree{
		Root: breadthfirstsearch.Leaf{
			Value: 5,
			Children: []breadthfirstsearch.Leaf{
				{
					Value: 2,
					Children: []breadthfirstsearch.Leaf{
						{
							Value: 3,
							Children: []breadthfirstsearch.Leaf{
								{
									Value: 4,
								},
								{
									Value: 7,
								},
							},
						},
						{
							Value: 8,
							Children: []breadthfirstsearch.Leaf{
								{
									Value: 9,
								},
								{
									Value: 10,
								},
							},
						},
					},
				},
				{
					Value: 6,
					Children: []breadthfirstsearch.Leaf{
						{
							Value: 11,
							Children: []breadthfirstsearch.Leaf{
								{
									Value: 12,
								},
								{
									Value: 13,
								},
							},
						},
					},
				},
			},
		},
	}
)

func Test_BinarySearchIterative(t *testing.T) {
	testCases := []testCase{
		{
			name: "finds the node in a tree with only one root node",
			tree: breadthfirstsearch.BinaryTree{
				Root: breadthfirstsearch.Leaf{
					Value: 5,
				},
			},
			target:   5,
			expected: true,
		},
		{
			name: "does not finds the node in a tree with only one root node",
			tree: breadthfirstsearch.BinaryTree{
				Root: breadthfirstsearch.Leaf{
					Value: 5,
				},
			},
			target:   0,
			expected: false,
		},
		{
			name: "does not finds the node in a tree of depth 1",
			tree: breadthfirstsearch.BinaryTree{
				Root: breadthfirstsearch.Leaf{
					Value: 5,
					Children: []breadthfirstsearch.Leaf{
						{
							Value: 2,
						},
						{
							Value: 6,
						},
					},
				},
			},
			target:   0,
			expected: false,
		},
		{
			name: "finds the node in a tree of depth 1",
			tree: breadthfirstsearch.BinaryTree{
				Root: breadthfirstsearch.Leaf{
					Value: 5,
					Children: []breadthfirstsearch.Leaf{
						{
							Value: 2,
						},
						{
							Value: 6,
						},
					},
				},
			},
			target:   6,
			expected: true,
		},
		{
			name:     "finds the node in a tree of depth 3",
			tree:     biggerBinaryTree,
			target:   13,
			expected: true,
		},
	}

	for _, tc := range testCases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := breadthfirstsearch.BreadthFirstSearchIterative(c.tree, c.target)
			assert.Equal(t, c.expected, res)
		})
	}
}
