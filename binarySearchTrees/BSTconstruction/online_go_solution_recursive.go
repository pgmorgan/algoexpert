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
	if value < tree.value {
		if tree.left == nil {
			tree.left = &BST{value: value}
		} else {
			tree.left.Insert(value)
		}
	} else {
		if tree.right == nil {
			tree.right = &BST{value: value}
		} else {
			tree.right.Insert(value)
		}
	}
	return tree
}

func (tree *BST) Contains(value int) bool {
	if value < tree.value {
		if tree.left == nil {
			return false
		} else {
			return tree.left.Contains(value)
		}
	} else if value > tree.value {
		if tree.right == nil {
			return false
		} else {
			return tree.right.Contains(value)
		}
	}
	return true
}

func (tree *BST) Remove(value int) *BST {
	tree.remove(value, nil, &tree)
	return tree
}

func (tree *BST) remove(value int, parent *BST, ad_tree **BST) {
	current := tree
	for current != nil {
		if value < current.value {
			parent = current
			current = current.left
		} else if value > current.value {
			parent = current
			current = current.right
		} else {
			if current.left != nil && current.right != nil {
				current.value = current.right.getMinValue()
				current.right.remove(current.value, current, &(current.right))
			// 	CASE 1 - TARGET HAS TWO CHILDREN
			} else if parent == nil {
			/*	CASE 2 - TARGET HAS ONE OR ZERO CHILDREN
						 AND TARGET IS THE ROOT NODE */
				if current.left != nil {
					*ad_tree = current.left
				} else if current.right != nil {
					*ad_tree = current.right
				} else {
					current.value = 0
				}
			} else if parent.left == current {
			/*	CASE 3 - TARGET HAS ONE OR ZERO CHILDREN AND
						 TARGET IS NOT THE ROOT NODE AND
						 TARGET IS A LEFT CHILD */
				if current.left != nil {
					parent.left = current.left
				} else {
					parent.left = current.right
				}
			} else if parent.right == current {
				if current.left != nil {
					parent.right = current.left
				} else {
					parent.right = current.right
				}
			}
			break
		}
	}
}

func (tree *BST) getMinValue() int {
	if tree.left == nil {
		return tree.value
	}
	return tree.left.getMinValue()
}
