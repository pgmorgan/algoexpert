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
	tree.remove(value, nil)
	return tree
}

func (tree *BST) remove(value int, parent *BST) {
	if value < tree.value {
		if tree.left != nil {
			tree.left.remove(value, tree)
		}
	} else if value > tree.value {
		if tree.right != nil {
			tree.right.remove(value, tree)
		}
	} else {
		if tree.left != nil && tree.right != nil {
			tree.value = tree.right.getMinValue()
			tree.right.remove(tree.value, tree)
		} else if parent == nil {
			if tree.left != nil {
				tree.value = tree.left.value
				tree.right = tree.left.right
				tree.left = tree.left.left
			} else if tree.right != nil {
				tree.value = tree.right.value
				tree.left = tree.right.left
				tree.right = tree.right.right
			} else {
				tree.value = 0
			}
		} else if parent.left == tree {
			if tree.left != nil {
				parent.left = tree.left
			} else {
				parent.left = tree.right
			}
		} else if parent.right == tree {
			if tree.left != nil {
				parent.right = tree.left
			} else {
				parent.right = tree.right
			}
		}
	}
}

func (tree *BST) getMinValue() int {
	if tree.left == nil {
		return tree.value
	}
	return tree.left.getMinValue()
}
