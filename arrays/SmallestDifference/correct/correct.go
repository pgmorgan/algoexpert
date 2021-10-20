//	Time:	O(n*log(n))		Space:	O(1)

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(SmallestDifference([]int{-1, 5, 10, 20, 28, 3}, []int{26, 134, 135, 15, 17}))
}

func absDiff(x, y int) int {
	return int(math.Abs(float64(x) - float64(y)))
}

func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1) //Time: n*log(n)
	sort.Ints(array2) //Time: n*log(n)
	i1, i2 := 0, 0
	output := []int{array1[i1], array2[i2]}
	for i1 < len(array1) || i2 < len(array2) { //Time: O(n)

		if absDiff(array1[i1], array2[i2]) < absDiff(output[0], output[1]) {
			output = []int{array1[i1], array2[i2]}
		}
		if array1[i1] < array2[i2] && i1 < len(array1)-1 {
			i1++
		} else if i2 < len(array2)-1 {
			i2++
		} else {
			break
		}
	}
	return output
}
