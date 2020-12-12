package array

func MergeArray(arr1, arr2 []int) []int {
	mergedArray := make([]int, len(arr1)+len(arr2))
	p1, p2, m1 := 0, 0, 0
	for p1 < len(arr1) && p2 < len(arr2) {
		if arr1[p1] < arr2[p2] {
			mergedArray[m1] = arr1[p1]
			m1++
			p1++
		} else {
			mergedArray[m1] = arr2[p2]
			m1++
			p2++
		}
	}

	for p1 < len(arr1) {
		mergedArray[m1] = arr1[p1]
		m1++
		p1++
	}

	for p2 < len(arr2) {
		mergedArray[m1] = arr2[p2]
		m1++
		p2++
	}
	return mergedArray
}
