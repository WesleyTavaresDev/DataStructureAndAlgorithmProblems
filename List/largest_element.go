package list

// Write a function that returns the largest element in a list.

func LargestElement(arr []int) int {
	largest := 0

	for i := 0; i < len(arr); i++ {

		if arr[i] > largest {
			largest = arr[i]
		}
	}

	return largest
}
