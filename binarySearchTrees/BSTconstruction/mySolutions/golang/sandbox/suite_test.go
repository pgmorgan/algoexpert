// Add, edit, or remove tests in this file.
// Treat it as your playground!

package program

import (
	"fmt"
	"reflect"
	"testing"
)

func NewBST(root int, values ...int) *BST {
	tree := &BST{value: root}
	for _, value := range values {
		tree.Insert(value)
	}
	return tree
}

func (tree *BST) InOrderTraverse(array []int) []int {
	if tree.left != nil {
		array = tree.left.InOrderTraverse(array)
	}
	array = append(array, tree.value)
	if tree.right != nil {
		array = tree.right.InOrderTraverse(array)
	}
	return array
}

var test1 = NewBST(10, 5, 15, 5, 2, 14, 22)
var test2 = NewBST(10, 15, 11, 22).Remove(10)
var test3 = NewBST(10, 5, 7, 2).Remove(10)
var test4 = NewBST(10, 5, 15, 22, 17, 34, 7, 2, 5, 1, 35, 27, 16, 30).Remove(22).Remove(17)

func TestCase1(t *testing.T) {
	output := test1.left.value
	expected := 5
	if output != expected {
		t.Fail()
	}
}

func TestCase2(t *testing.T) {
	output := test1.right.right.value
	expected := 22
	if output != expected {
		t.Fail()
	}
}

func TestCase3(t *testing.T) {
	output := test1.right.left.value
	expected := 14
	if output != expected {
		t.Fail()
	}
}

func TestCase4(t *testing.T) {
	output := test1.left.right.value
	expected := 5
	if output != expected {
		t.Fail()
	}
}

func TestCase5(t *testing.T) {
	output := test1.left.left.value
	expected := 2
	if output != expected {
		t.Fail()
	}
}

func TestCase6(t *testing.T) {
	output := test1.left.left.left
	if output != nil {
		t.Fail()
	}
}

func TestCase7(t *testing.T) {
	output := test1.right.left.right
	if output != nil {
		t.Fail()
	}
}

func TestCase8(t *testing.T) {
	output := test1.Contains(15)
	expected := true
	if output != expected {
		t.Fail()
	}
}

func TestCase9(t *testing.T) {
	output := test1.Contains(2)
	expected := true
	if output != expected {
		t.Fail()
	}
}

func TestCase10(t *testing.T) {
	output := test1.Contains(5)
	expected := true
	if output != expected {
		t.Fail()
	}
}

func TestCase11(t *testing.T) {
	output := test1.Contains(10)
	expected := true
	if output != expected {
		t.Fail()
	}
}

func TestCase12(t *testing.T) {
	output := test1.Contains(22)
	expected := true
	if output != expected {
		t.Fail()
	}
}

func TestCase13(t *testing.T) {
	output := test1.Contains(23)
	expected := false
	if output != expected {
		t.Fail()
	}
}

func TestCase14(t *testing.T) {
	output := test2.InOrderTraverse([]int{})
	expected := []int{11, 15, 22}
	if !reflect.DeepEqual(output, expected) {
		t.Fail()
	}
}

func TestCase15(t *testing.T) {
	output := test3.InOrderTraverse([]int{})
	expected := []int{2, 5, 7}
	if !reflect.DeepEqual(output, expected) {
		t.Fail()
	}
}

func TestCase16(t *testing.T) {
	output := test4.InOrderTraverse([]int{})
	expected := []int{1, 2, 5, 5, 7, 10, 15, 16, 27, 30, 34, 35}
	fmt.Println(output)
	fmt.Println(expected)
	if !reflect.DeepEqual(output, expected) {
		t.Fail()
	}
}

func TestCase17(t *testing.T) {
	output := test4.right.right.value
	expected := 27
	if !reflect.DeepEqual(output, expected) {
		t.Fail()
	}
}

func TestCase18(t *testing.T) {
	output := test4.right.right.left.value
	expected := 16
	if !reflect.DeepEqual(output, expected) {
		t.Fail()
	}
}
