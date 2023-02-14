package list

// Write a function that returns the largest element in a list.

func LargestElement(arr []int) int {
	largest := 0

	for _, number := range arr {

		if number > largest {
			largest = number
		}
	}

	return largest
}
