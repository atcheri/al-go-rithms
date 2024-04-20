package kruskal

import (
	"sort"
)

type Edge struct {
	u, v int
	w    int
}

type Graph struct {
	V int
	E []Edge
}

func (g *Graph) AppendEdge(u, v, w int) {
	g.E = append(g.E, Edge{u, v, w})
}

func (g *Graph) FindMST() []Edge {
	sort.Slice(g.E, func(i, j int) bool {
		return g.E[i].w < g.E[j].w
	})

	parents := make([]int, g.V)
	for i := range parents {
		parents[i] = i
	}

	ranks := make([]int, g.V)

	mst := make([]Edge, 0)
	for _, edge := range g.E {
		if find(edge.u, parents) != find(edge.v, parents) {
			union(edge.u, edge.v, parents, ranks)
			mst = append(mst, edge)
		}
	}

	return mst
}

func find(u int, parents []int) int {
	// if parents[u] != u {
	// 	parents[u] = find(parents[u], parents)
	// }

	// return parents[u]

	if parents[u] == u {
		return u
	}

	return find(parents[u], parents)
}

func union(u, v int, parents, ranks []int) {
	rootU := find(u, parents)
	rootV := find(v, parents)

	if ranks[rootU] > ranks[rootV] {
		parents[rootV] = rootU
	} else if ranks[rootU] < ranks[rootV] {
		parents[rootU] = rootV
	} else {
		parents[rootV] = rootU
		ranks[rootU]++
	}
}
