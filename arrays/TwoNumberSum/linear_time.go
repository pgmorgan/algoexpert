//	Time: O(n)	Space: O(n)

package program

func TwoNumberSum(array []int, target int) []int {
	m := make(map[int]int)
	
	for _, x := range array {
		if y, ok := m[x]; ok {
			if x < y {
				return []int{x, y}
			} else {
				return []int{y, x}
			}
		} else {
			m[target - x] = x
		}
	}
	return []int{}
}
