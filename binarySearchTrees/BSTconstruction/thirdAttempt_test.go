package program

import (
	"reflect"
	"testing"
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
	cur := tree
	if cur == nil {
		return nil
	}
	parentWB, parentEB := &tree, &tree

	if cur.value != value {
		for cur.value != value {
			if value < cur.value {
				parentWB = &cur
				parentEB = nil
				cur = cur.left
			} else if value > cur.value {
				parentWB = nil
				parentEB = &cur
				cur = cur.right
			}
		}	
		cur.Remove(value)
	}

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

var test2 = NewBST(10, 15, 11, 22).Remove(15)
//var test2 = NewBST(10, 15, 11, 22)

func NewBST(root int, values ...int) *BST {
	tree := &BST{value: root}
	for _, value := range values {
		tree.Insert(value)
	}
	return tree
}

func (tree *BST) InOrderTraverse(array []int) []int {
	if tree == nil {
		return array
	}
	if tree.left != nil {
		array = tree.left.InOrderTraverse(array)
	}
	array = append(array, tree.value)
	if tree.right != nil {
		array = tree.right.InOrderTraverse(array)
	}
	return array
}

func TestCase14(t *testing.T) {
	output := test2.InOrderTraverse([]int{})
	expected := []int{10, 11, 22}
	fmt.Println(output)
	fmt.Println(expected)
	if !reflect.DeepEqual(output, expected) {
		t.Fail()
	}
}
