package search

import "golang.org/x/exp/constraints"

// BinarySearch выполняет бинарный поиск в отсортированном слайсе.
func BinarySearch[Comp constraints.Ordered](sortedData []Comp, target Comp) int {
	n := len(sortedData)

	if n < 3 { // linear search
		for i, val := range sortedData {
			if val == target {
				return i
			}
		}
	} else { // binary search
		leftIdx, rightIdx := int(1), n-1
		for leftIdx <= rightIdx {
			midIdx := int(leftIdx + (rightIdx-leftIdx)/2)
			if sortedData[midIdx] == target {
				return midIdx // success
			}

			if sortedData[midIdx] > target {
				rightIdx = midIdx - 1
			} else {
				leftIdx = midIdx + 1
			}
		}
	}

	return -1
}
