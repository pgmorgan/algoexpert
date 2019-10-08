package program

import (
	"fmt"
)

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

func (tree *BST) findTargetAndParent(value int) (cur *BST, pLeft **BST, pRight **BST) {
	cur = tree
	if cur.value == value {
		fmt.Println(cur.value, value)
		pLeft, pRight = &tree, &tree
		return
	}
	for cur.value != value {
		if value < cur.value {
			pLeft = &cur
			pRight = nil
			cur = cur.left
		} else if value > cur.value {
			pLeft = nil
			pRight = &cur
			cur = cur.right
		}
	}
	return
}

func (tree *BST) Remove(value int) *BST {
	cur, parentWB, parentEB := tree.findTargetAndParent(value)
	if cur.left == nil && cur.right == nil {
		// Target has no children
		fmt.Println(cur.value, value)
		tree = nil
	} else if cur.left != nil && cur.right == nil {
		// Target has one left child
		if parentWB == parentEB {
			// And Target is Root
			tree = cur.left
		} else if parentWB != nil && parentEB == nil {
			// And target has one parent to its right (WB)
			(*parentWB).left = cur.left
		} else if parentWB == nil && parentEB != nil {
			// And target has one parent to its left (EB)
			(*parentEB).right = cur.left
		}
	} else if cur.left == nil && cur.right != nil {
		if parentWB == parentEB {
			tree = cur.right
		} else if parentWB != nil && parentEB == nil {
			(*parentWB).left = cur.right
		} else if parentWB == nil && parentEB != nil {
			(*parentEB).right = cur.right
		}
	} else if cur.left != nil && cur.right != nil {
		cur.value = cur.right.getMinValue()
		cur.right.Remove(cur.value)
	}
	fmt.Println(tree.InOrderTraverse([]int{}))
	return tree
}

func (tree *BST) getMinValue() int {
	cur := tree
	for cur.left != nil {
		cur = cur.left
	}
	return cur.value
}
