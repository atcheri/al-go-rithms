package depthfirstsearch

func DepthFirstSearchIterative(tree BinaryTree, target int) bool {

	root := tree.Root

	if root.Children == nil {
		return root.Value == target
	}

	queue := &Stack[Leaf]{}
	queue.Push(root)
	for !queue.IsEmpty() {
		curr := queue.Pop()
		if curr.Value == target {
			return true
		}

		for _, c := range curr.Children {
			queue.Push(c)
		}
	}

	return false
}
