package insertion

import "errors"

func Insertion(elements []int) ([]int, error) {

	if len(elements) <= 0 {
		return elements, errors.New("Empty list")
	}

	for i := 1; i < len(elements); i++ {
		for j := i; j > 0; j-- {
			if elements[j] < elements[j-1] {
				elements[j], elements[j-1] = elements[j-1], elements[j]
			}
		}
	}

	return elements, nil
}
