package program

// Do not edit the class below except for
// the insert, contains, and remove methods.
// Feel free to add new properties and methods
// to the class.
type BST struct {
	value int

	left  *BST
	right *BST
}

func (tree *BST) Insert(value int) *BST {
	var parent *BST
	var right bool
	current := tree
	for true {
		parent = current
		if value < current.value {
			current = current.left
			right = false
		} else {
			current = current.right
			right = true
		}
		if current == nil {
			break
		}
	}

	if right == true {
		parent.right = &BST{value: value}
	} else {
		parent.left = &BST{value: value}
	}
	return tree
}

func (tree *BST) Contains(value int) bool {
	for true {
		if value < tree.value {
			if tree.left == nil {
				break
			}
			tree = tree.left
		} else if value > tree.value {
			if tree.right == nil {
				break
			}
			tree = tree.right
		} else if value == tree.value {
			return true
		}
	}
	return false
}

func (tree *BST) Remove(value int) *BST {
	if tree.value == value {
		if tree.right == nil {
			return tree.left
		} else {
			tree.value = tree.right.retrieveAndRemoveMin(tree)
		}
	} else {
		tree.removeChild(value, nil)
	}
	return tree
}

func (tree *BST) retrieveAndRemoveMin(parent *BST) int {
	for tree.left != nil {
		parent = tree
		tree = tree.left
	}
	if parent.right == tree {
		parent.right = tree.right
	} else {
		parent.left = tree.right
	}
	return tree.value
}

func (tree *BST) removeChild(value int, parent *BST) {
	for true {
		if value < tree.value {
			if tree.left == nil {
				return
			}
			parent = tree
			tree = tree.left
		} else if value > tree.value {
			if tree.right == nil {
				return
			}
			parent = tree
			tree = tree.right
		} else {
			if tree.right == nil {
				if parent.left == tree {
					parent.left = tree.left
					break
				} else {
					parent.right = tree.left
					break
				}
			} else {
				tree.value = tree.retrieveAndRemoveMin(parent)
				break
			}
		}
	}
}

// func (tree *BST) Remove(value int) *BST {
// 	if tree.value == value {
// 		tree.removeRoot(value, &tree)
// 	} else {
// 		tree.removeChildNode(value)
// 	}
// 	return tree
// }

// func (tree *BST) removeRoot(value int, ad_tree **BST) {
// 	if tree.left == nil {
// 		*ad_tree = tree.right
// 	} else if tree.right != nil {
// 		tree.value = tree.right.retrieveAndRemoveMinValue(tree)
// 	} else {
// 		*ad_tree = tree.left
// 	}
// }

// func (tree *BST) retrieveAndRemoveMinValue(parent *BST) int {
// 	for tree.left != nil {
// 		parent = tree
// 		tree = tree.left
// 	}
// 	parent.left = tree.right
// 	// value := tree.value

// 	// tree = tree.right
// 	// parent.right = tree.right
// 	// parent.left = nil
// 	return tree.value
// }

// func (tree *BST) removeChildNode(value int) {
// 	var child, parent *BST
// 	current := tree

// 	for true {
// 		if value < current.value {
// 			child = current.left
// 		} else if value > current.value {
// 			child = current.right
// 		} else {
// 			break
// 		}
// 		if child == nil {
// 			_ = parent
// 			return
// 		}
// 		parent = current
// 		current = child
// 	}
// 	if current.right != nil {
// 		current.value = current.right.retrieveAndRemoveMinValue(current)
// 	} else if parent.left == current {
// 		parent.left = current.left
// 	} else if parent.right == current {
// 		parent.right = current.left
// 	}
// 	return
// }
