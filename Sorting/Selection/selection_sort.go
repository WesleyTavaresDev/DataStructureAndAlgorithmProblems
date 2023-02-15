package selection

type Order int

const (
	Ascending  Order = iota
	Decreasing       = iota
)

func SelectionSort(elements []int, order Order) []int {

	for i := 0; i < len(elements)-1; i++ {
		minIndex := i
		for j := i + 1; j <= len(elements)-1; j++ {
			if order == Ascending && elements[j] < elements[minIndex] || order == Decreasing && elements[j] > elements[minIndex] {
				minIndex = j
			}
		}
		elements[i], elements[minIndex] = elements[minIndex], elements[i]
	}

	return elements
}
