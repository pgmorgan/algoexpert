package program

import (
	"fmt"
	"math"
)

type BST struct {
	value int

	left  *BST
	right *BST
}

func abs(value int) int {
	return int(math.Abs(float64(value)))
}

func (tree *BST) FindClosestValue(target int) int {
	cur := tree
	diff := cur.value - target
	min := abs(diff)
	closest := cur.value
	for {
		diff = cur.value - target
		if abs(diff) < min {
			min = abs(diff)
			closest = cur.value
		}
		if diff == 0 {
			fmt.Println(closest)
			return closest
		} else if diff < 0 && cur.right != nil {
			cur = cur.right
			fmt.Println(cur.value)
		} else if diff > 0 && cur.left != nil {
			cur = cur.left
			fmt.Println(cur.value)
		} else {
			fmt.Println("break")
			break
		}
	}
	fmt.Println(closest)
	return closest
}
