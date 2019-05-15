
package program

func FourNumberSum(array []int, target int) [][]int {
	m := make(map[int][][]int)
	var quadruplets [][]int
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if hpairs, ok := m[target - array[i] - array[j]]; ok {
				for _, pair := range hpairs {
					quadruplets = append(quadruplets, []int{pair[0],
															pair[1],
															array[i],
															array[j],})
				}
			}
		}

		for k := i - 1; k >= 0; k-- {
			if _, ok := m[array[i] + array[k]]; ok {
				m[array[i] + array[k]] = append(m[array[i] + array[k]], []int{array[i], array[k],})
			} else {
				m[array[i] + array[k]] = [][]int{[]int{array[i], array[k],}}
			}
		}
	}
	return quadruplets
}
