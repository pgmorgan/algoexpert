func (tree *BST) Insert(value int) *BST {
	current := tree
	for {
		if value < current.value {
			if current.left == nil {
				current.left = &BST{value: value}
				break
			} else {
				current = current.left
			}
		} else {
			if current.right == nil {
				current.right = &BST{value: value}
				break
			} else {
				current = current.right
			}
		}
	}
	return tree
}

func (tree *BST) Contains(value int) bool {
	current := tree
	for current != nil {
		if value < current.value {
			current = current.left
		} else if value > current.value {
			current = current.right
		} else {
			return true
		}
	}
	return false
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
				current.right.remove(current.value, current)
			// 	CASE 1 - TARGET HAS TWO CHILDREN 
			} else if parent == nil {
			/*	CASE 2 - TARGET HAS ONE OR ZERO CHILDREN
						 AND TARGET IS THE ROOT NODE */
				if current.left != nil {
				/*	CASE 2A - ROOT NODE TARGET HAS LEFT CHILD ONLY */
					*ad_tree = current.left
				} else if current.right != nil {
				/*	CASE 2B - ROOT NODE TARGET HAS RIGHT CHILD ONLY */
					*ad_tree = current.right
				} else {
				/*	CASE 2C - ROOT NODE AND BOTH CHILDREN ARE NIL */
					current.value = 0
				}
			} else if parent.left == current {
			/*	CASE 3 - TARGET HAS ONE OR ZERO CHILDREN AND 
						 TARGET IS NOT THE ROOT NODE AND
						 TARGET IS A LEFT CHILD */
				if current.left != nil {
				//	TARGET HAS LEFT CHILD ONLY
					parent.left = current.left
					// SKIP OVER CURRENT.  CONNECT PARENT.LEFT TO CURRENT.LEFT
				} else {
				//	TARGET HAS RIGHT CHILD ONLY
					parent.left = current.right
					// SKIP OVER CURRENT. CONNECT PARENT.LEFT 
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
