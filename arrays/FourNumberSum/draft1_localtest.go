package main
import (
	"fmt"
)

func	FourNumberSum(array []int, target int) [][]int {
	output := [][]int{}
	m := make(map[int][]int)
	table := [][]int{}
	
	//	Time: O(n^2)	Space:	O(n)
	for i := 0; i < len(array) - 1; i++ {
		for j := i + 1; j < len(array); j++ {
			table = append(table, []int{array[i], array[j], array[i] + array[j],})
		}
	}

	//	Time: O(n)		Space:	O(n)
	for _, x := range table {
		if y, ok := m[x[2]]; ok {
			output = append(output, []int{x[0], x[1], y[0], y[1]})
		} else {
			m[target - x[2]] = x
		}
	}
	return output
}

func	sumIntSlice(s []int) int {
	var total int
	for _, v := range s {
		total += v
	}
	return total
}

func	main() {
	for _, a := range FourNumberSum([]int{7, 6, 4, -1, 1, 2}, 16) {
		fmt.Println(a, "\t", sumIntSlice(a))
	}
}
