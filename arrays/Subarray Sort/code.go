package program

func SubarraySort(array []int) []int {
	misfits := make([]int, 0)
	for i, v := range array {
		if i == 0 {
			if v > array[i+1] {
				misfits = append(misfits, array[i])
			}
		} else if i == len(array)-1 {
			if v < array[i-1] {
				misfits = append(misfits, array[i])
			}
		} else if v > array[i+1] || v < array[i-1] {
			misfits = append(misfits, array[i])
		}
	}
	if len(misfits) == 0 {
		return []int{-1, -1}
	}
	low, high := minAndMax(misfits)
	i := findPosition(array, low)
	j := findPosition(array, high)
	return []int{i, j}
}

func minAndMax(a []int) (low, high int) {
	low, high = a[0], a[0]
	for _, v := range a {
		if v < low {
			low = v
		}
		if v > high {
			high = v
		}
	}
	return low, high
}

func findPosition(a []int, val int) (index int) {
	if val <= a[0] {
		return 0
	} else if val >= a[len(a)-1] {
		return len(a) - 1
	}
	for i := 1; i <= len(a)-2; i++ {
		if val >= a[i-1] && val <= a[i+1] {
			if val < a[i] {
				index = i
			} else if val >= a[i] {
				index = i + 1
			}
			break
		}
	}
	return index
}
