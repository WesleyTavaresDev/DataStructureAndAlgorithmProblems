package bubble

func BubbleSort(elements []int) []int {

	for i := len(elements) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if elements[j] > elements[j+1] {
				elements[j], elements[j+1] = elements[j+1], elements[j]
			}
		}
	}

	return elements
}
