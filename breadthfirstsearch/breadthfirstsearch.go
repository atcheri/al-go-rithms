package breadthfirstsearch

func BreadthFirstSearchIterative(tree BinaryTree, target int) bool {
	root := tree.Root

	if root.Children == nil {
		return root.Value == target
	}

	queue := &Queue[Leaf]{}
	queue.Enqueue(root)
	for !queue.IsEmpty() {
		curr := queue.Dequeue()
		if curr.Value == target {
			return true
		}

		for _, c := range curr.Children {
			queue.Enqueue(c)
		}
	}

	return false
}
