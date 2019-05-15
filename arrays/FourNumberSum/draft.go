package program
import (
	"fmt"
)

func	FourNumberSum(array []int, target int) [][]int {
	output := [][]int{}
	m := map[int][]int
	table := [][]int{}
	
	for i := 0; i < len(array) - 1; i++ {
		for j := i + 1; j < len(array); j++ {
			table = append(table, []int{array[i], array[j], array[i] + array[j],}
		}
	}

	for _, x := range table {
		if y, ok := m[x[3]]; ok {
			output = append(output, []int{x[0], x[1], y[0], y[1]})
		} else {
			m[target - x[3]] = x
		}
	}
	return output
}
