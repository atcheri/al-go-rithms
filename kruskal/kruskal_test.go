package kruskal_test

import (
	"al-go-rithms/kruskal"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name     string
	graph    *kruskal.Graph
	expected []kruskal.Edge
}

func buildGraph(vals [][]int, size int) *kruskal.Graph {
	graph := &kruskal.Graph{E: make([]kruskal.Edge, 0), V: size}
	for _, v := range vals {
		graph.AppendEdge(v[0], v[1], v[2])
	}

	return graph
}

var (
	fromExampleTestCase = testCase{
		name: "",
		graph: buildGraph([][]int{
			{2, 4, 5}, // in
			{0, 3, 5}, // in
			{3, 5, 6}, // in
			{1, 4, 7}, // in
			{0, 1, 7}, // in
			{1, 2, 8},
			{4, 5, 8},
			{4, 6, 9}, // in
			{1, 3, 9},
			{5, 6, 11},
			{3, 4, 15},
		}, 7),
		expected: []kruskal.Edge{},
	}
)

func Test_Kruskal(t *testing.T) {
	testCases := []testCase{
		// {
		// 	name:  "returns an empty slice of edges when no edges",
		// 	edges: []kruskal.Edge{},
		// },
		// {
		// 	name: "returns an empty slice of edges given one invalid edge",
		// 	edges: []kruskal.Edge{
		// 		{
		// 			A: nil,
		// 			B: nil,
		// 		},
		// 	},
		// },
		// {
		// 	name: "returns the same slice of edge given on edge",
		// 	edges: []kruskal.Edge{
		// 		{
		// 			A: &kruskal.Node{
		// 				Value:  1,
		// 				Parent: 1,
		// 			},
		// 			B: &kruskal.Node{
		// 				Value:  2,
		// 				Parent: 1,
		// 			},
		// 		},
		// 	},
		// 	expected: []kruskal.Edge{
		// 		{
		// 			A: &kruskal.Node{
		// 				Value:  1,
		// 				Parent: 1,
		// 			},
		// 			B: &kruskal.Node{
		// 				Value:  2,
		// 				Parent: 1,
		// 			},
		// 		},
		// 	},
		// },
		// twoEdgesTestCase,
		// threeEdgesTestCase,
		fromExampleTestCase,
	}

	for _, tc := range testCases {
		c := tc
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			mst := c.graph.FindMST()
			assert.Equal(t, c.expected, mst)
		})
	}
	assert.True(t, true)
}
