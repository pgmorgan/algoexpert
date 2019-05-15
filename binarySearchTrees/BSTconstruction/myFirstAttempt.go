(tree *BST) Insert(value int) *BST {
	// Write your code here.
	// Do not edit the return statement of this method.

	current := tree
	for true {
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
		if value < current.value{
			current = current.left
		} else if value > current.value {
			current = current.right
		} else if current.value == value {
			return true
		}
	}
	return false
}

func findTargetAndParent(tree *BST, value int) (node *BST, parent *BST, right bool) {
	node = tree
	parent = nil
	right = false

	for node != nil {
		if node.value < value {
			parent = node
			node = node.right
			right = true
		} else if node.value > value {
			parent = node
			node = node.left
			right = false
		} else if node.value == value {
			return node, parent, right
		}
	}
	return nil, nil, false
}	

func findSmallestAndParent(toBeDel *BST) (node *BST, parent *BST) {
	parent = toBeDel
	node = toBeDel.right

	if node == nil {
		return nil, nil
	}
	for node.left != nil {
		parent = node
		node = node.left
	}
	return node, parent
}

func (tree *BST) Remove(value int) *BST {
	target, parent, right := findTargetAndParent(tree, value)
	if target == nil {
		return tree
	}
	smallest, smallestParent := findSmallestAndParent(target)
	if smallest == nil && right {
		parent.right = nil
		//	} else if smallest == nil && !right {
		//		parent.left = nil
	} else if smallest != nil {
		target.value = smallest.value
		smallestParent.left = smallest.right
	}
	return tree
}
