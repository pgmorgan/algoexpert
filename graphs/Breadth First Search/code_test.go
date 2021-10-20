package program

import (
	"reflect"
	"testing"
)

func NewNode(name string) *Node {
	return &Node{
		Name:     name,
		Children: []*Node{},
	}
}

func (n *Node) AddChildren(names ...string) *Node {
	for _, name := range names {
		child := Node{Name: name}
		n.Children = append(n.Children, &child)
	}
	return n
}

var test1 = NewNode("A").AddChildren("B", "C")
var test2 = NewNode("A").AddChildren("B", "C", "D", "E")
var test3 = NewNode("A").AddChildren("B")
var test4 = NewNode("A").AddChildren("B", "C", "D")
var test5 = NewNode("A").AddChildren("B", "C", "D", "L", "M")

func init() {
	test1.Children[0].AddChildren("D")
	test2.Children[1].AddChildren("F")
	test3.Children[0].AddChildren("C")
	test3.Children[0].Children[0].AddChildren("D").AddChildren("E")
	test3.Children[0].Children[0].Children[0].AddChildren("F")
	test4.Children[0].AddChildren("E").AddChildren("F")
	test4.Children[2].AddChildren("G").AddChildren("H")
	test4.Children[0].Children[1].AddChildren("I").AddChildren("J")
	test4.Children[2].Children[0].AddChildren("K")
	test5.Children[0].AddChildren("E").AddChildren("F").AddChildren("O")
	test5.Children[1].AddChildren("P")
	test5.Children[2].AddChildren("G").AddChildren("H")
	test5.Children[0].Children[0].AddChildren("Q").AddChildren("R")
	test5.Children[0].Children[1].AddChildren("I").AddChildren("J")
	test5.Children[2].Children[0].AddChildren("K")
	test5.Children[4].AddChildren("S").AddChildren("T").AddChildren("U").AddChildren("V")
	test5.Children[4].Children[0].AddChildren("W").AddChildren("X")
	test5.Children[4].Children[0].Children[1].AddChildren("Y").AddChildren("Z")
}

func TestCase1(t *testing.T) {
	output := test1.BreadthFirstSearch([]string{})
	expected := []string{"A", "B", "C", "D"}
	if !reflect.DeepEqual(output, expected) {
		t.Fail()
	}
}

func TestCase2(t *testing.T) {
	output := test2.BreadthFirstSearch([]string{})
	expected := []string{"A", "B", "C", "D", "E", "F"}
	if !reflect.DeepEqual(output, expected) {
		t.Fail()
	}
}

func TestCase3(t *testing.T) {
	output := test3.BreadthFirstSearch([]string{})
	expected := []string{"A", "B", "C", "D", "E", "F"}
	if !reflect.DeepEqual(output, expected) {
		t.Fail()
	}
}

func TestCase4(t *testing.T) {
	output := test4.BreadthFirstSearch([]string{})
	expected := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}
	if !reflect.DeepEqual(output, expected) {
		t.Fail()
	}
}

func TestCase5(t *testing.T) {
	output := test5.BreadthFirstSearch([]string{})
	expected := []string{"A", "B", "C", "D", "L", "M", "E", "F", "O", "P", "G", "H", "S", "T", "U", "V", "Q", "R", "I", "J", "K", "W", "X", "Y", "Z"}
	if !reflect.DeepEqual(output, expected) {
		t.Fail()
	}
}
