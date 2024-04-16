package depthfirstsearch

type BinaryTree struct {
	Root Leaf
}

type Leaf struct {
	Value    int
	Children []Leaf
}
