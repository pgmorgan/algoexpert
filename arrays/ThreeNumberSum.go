// Sort takes O(nlog(n))
// Then there are two loops O(n^2)
// Time:	O(n^2)		Space:	O(1)

package program
import "sort"

func ThreeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)
	triplets := [][]int{}
	//FOR SOME BIZARRE REASON IF I DECLARE TRIPLETS HERE AS
	//var triplets [][]int
	//IT FAILS THE TESTS ON ALGOEXPERT
	for i := 0; i < len(array) - 2; i++	{
		left := i + 1
		right := len(array) - 1
		for left < right {
			sum := array[i] + array[left] + array[right]
			if sum == target {
				triplets = append(triplets, []int{array[i], array[left], array[right],})
				left++
				right--
			} else if sum < target {
				left++
			} else if sum > target {
				right--
			}
		}
	}
	return triplets
}

