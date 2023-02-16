package merge

func MergeSort(elements []int) []int {
	if len(elements) < 2 {
		return elements
	}

	first := MergeSort(elements[:len(elements)/2])
	last := MergeSort(elements[len(elements)/2:])

	return merge(first, last)
}

func merge(left, right []int) []int {
	merged := []int{}

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	for ; i < len(left); i++ {
		merged = append(merged, left[i])
	}

	for ; j < len(right); j++ {
		merged = append(merged, right[j])
	}

	return merged
}
