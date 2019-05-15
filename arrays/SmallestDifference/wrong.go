-1, 3, 5, 10, 20, 28, 
15, 17, 26, 134, 135,

package program
import	(
	"sort"
	"math"
)

func absDiff(x, y int) int {
	return int(math.Abs(float64(x) - float64(y)))
}

func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)
	i1, i2 := 0, 0
	output := []int{array1[i1], array2[i2],}
	for (i1 < len(array1) || i2 < len(array2)) {
		
		if absDiff(array1[i1], array2[i2]) < absDiff(output[0], output[1]) {
			output = []int{array1[i1], array2[i2],}
		}	else if absDiff(array1[i1], array2[i2]) > absDiff(output[0], output[1]) {
			return output
		}
		if array1[i1] < array2[i2] && i1 < len(array1) - 1 {
			i1++
		} else if i2 < len(array2) - 1 {
			i2++
		}
	}
	return output	
}

