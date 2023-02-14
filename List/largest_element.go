package list

// Write a function that returns the largest element in a list.

func LargestElement() int {
	largest := 0

	arr := []int{3, 10, 20, 30, 1, 4, 25, 0, 5, -1}

	for i := 0; i < len(arr); i++ {

		if arr[i] > largest {
			largest = arr[i]
		}
	}

	return largest
}
