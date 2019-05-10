//	Time: O(n*log(n))		Space:	O(1)
//	The built-in Sort method takes O(n*log(n)) time, which then gets
//	added to O(n) time, so we end up with O(n*log(n)) time.

package program
import "sort"

func TwoNumberSum(array []int, target int) []int {
	sort.Ints(array)
	left := 0
	right := len(array) - 1
	var sum int
	for left < right {		
		sum = array[left] + array[right]
		if sum == target {
			return []int{array[left], array[right],}
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}
	return []int{}
}

