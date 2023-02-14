package list

func RunTotal(elements []int) []int {

	if len(elements) == 1 {
		return elements
	}

	for i := 1; i < len(elements); i++ {
		elements[i] += elements[i-1]
	}

	return elements
}
