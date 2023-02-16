package merge

func merge(elements []int, left int, middle int, right int) {

	helper := make([]int, len(elements))

	copy(helper, elements)

	i := left
	j := middle + 1
	k := left

	for ; i <= middle && j <= right; k++ {
		if helper[i] <= helper[j] {
			elements[k] = helper[i]
			i++
		} else {
			elements[k] = helper[j]
			j++
		}
	}

	for ; i <= middle; i++ {
		elements[k] = helper[i]
		k++
	}

	for ; j <= right; j++ {
		elements[k] = helper[j]
		k++
	}
}

func MergeSort(elements []int, left int, right int) {

	if left >= right {
		return
	} else {

		middle := (left + right) / 2
		MergeSort(elements, left, middle)
		MergeSort(elements, middle+1, right)

		merge(elements, left, middle, right-1)
	}
}
