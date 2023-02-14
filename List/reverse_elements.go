package list

func ReverseElements(arr []int) []int {

	arrCopy := make([]int, 0, cap(arr))

	for i := len(arr); i > 0; i-- {
		arrCopy = append(arr, arr[i])
	}

	return arrCopy
}
