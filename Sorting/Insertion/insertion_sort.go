package insertion

func Insertion(elements []int) []int {

	for i := 1; i < len(elements); i++ {
		for j := i; j > 0; j-- {
			if elements[j] < elements[j-1] {
				elements[j], elements[j-1] = elements[j-1], elements[j]
			}
		}
	}

	return elements
}
