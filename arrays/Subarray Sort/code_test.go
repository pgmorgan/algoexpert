package program

import (
	"reflect"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := []int{-1, -1}
	output := SubarraySort([]int{1, 2})
	if !reflect.DeepEqual(output, expected) {
		t.Fail()
	}
}

func TestCase2(t *testing.T) {
	expected := []int{0, 1}
	output := SubarraySort([]int{2, 1})
	if !reflect.DeepEqual(output, expected) {
		t.Fail()
	}
}

func TestCase3(t *testing.T) {
	expected := []int{3, 9}
	output := SubarraySort([]int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19})
	if !reflect.DeepEqual(output, expected) {
		t.Fail()
	}
}
