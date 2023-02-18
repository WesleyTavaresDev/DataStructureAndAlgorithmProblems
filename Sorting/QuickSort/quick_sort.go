package quicksort

func partition(arr []int, left, right int) int {
	pivot := arr[left]
	i := left

	for j := left + 1; j <= right; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[left], arr[i] = arr[i], arr[left]

	return i
}

func Quicksort(arr []int, left, right int) {

	if left < right {
		index_pivot := partition(arr, left, right)
		Quicksort(arr, left, index_pivot-1)
		Quicksort(arr, index_pivot+1, right)
	}

}
