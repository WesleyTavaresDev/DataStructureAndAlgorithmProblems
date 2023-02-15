package selection

func SelectionSort(elements []int) []int {

	for i := 0; i < len(elements)-1; i++ {
		minIndex := i

		for j := i + 1; j <= len(elements)-1; j++ {
			if elements[j] < elements[minIndex] {
				minIndex = j
			}
		}
		elements[i], elements[minIndex] = elements[minIndex], elements[i]
	}

	return elements
}
