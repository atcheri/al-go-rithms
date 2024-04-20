package kruskal

// type Graph struct {
// 	Edges []Edge
// }

// type Node struct {
// 	Value int
// }

// type Edge struct {
// 	A      *Node
// 	B      *Node
// 	Weight int
// }

// func (g Graph) AppendEdge(from, to, weight int) Graph {
// 	edge := Edge{
// 		A:      &Node{Value: from},
// 		B:      &Node{Value: to},
// 		Weight: weight,
// 	}
// 	tmp := make([]Edge, len(g.Edges))
// 	tmp = append(tmp, edge)
// 	copy(tmp, g.Edges)

// 	return Graph{
// 		Edges: tmp,
// 	}
// }

// func FindTree(edges []Edge) []Edge {
// 	if len(edges) == 0 {
// 		return nil
// 	}

// 	if edges[0].A == nil && edges[0].B == nil {
// 		return nil
// 	}

// 	// TODO: need to sort the edges in ASC order
// 	ranks := make(map[int]int)
// 	parents := make(map[int]int)
// 	for _, edge := range edges {
// 		parents[edge.A.Value] = edge.A.Value
// 		parents[edge.B.Value] = edge.B.Value
// 		if _, ok := ranks[edge.A.Value]; !ok {
// 			ranks[edge.A.Value] = 1
// 		}
// 		if _, ok := ranks[edge.B.Value]; !ok {
// 			ranks[edge.B.Value] = 1
// 		}
// 	}

// 	tree := make([]Edge, 0)
// 	index := 0
// 	i := 0

// 	for index < len(edges) {
// 		if i >= len(edges) {
// 			break
// 		}
// 		edge := edges[i]
// 		parentA := findNodeParent(edge.A.Value, parents)
// 		parentB := findNodeParent(edge.B.Value, parents)
// 		if parentA != parentB {
// 			tree = append(tree, edge)
// 			union(edge.A, edge.B, ranks, parents)
// 			index++
// 		}

// 		i++
// 	}

// 	return tree
// }

// func findNodeParent(k int, parents map[int]int) int {
// 	if parents[k] == k {
// 		return k
// 	}

// 	return findNodeParent(parents[k], parents)
// }

// func union(a *Node, b *Node, ranks map[int]int, parents map[int]int) {
// 	parentA := parents[a.Value]
// 	parentB := parents[b.Value]
// 	if ranks[parentA] > ranks[parentB] {
// 		parents[b.Value] = parentA
// 	} else if ranks[parentA] < ranks[parentB] {
// 		parents[a.Value] = parentB
// 	} else {
// 		parents[b.Value] = parentA
// 		ranks[parentA]++
// 	}
// }
