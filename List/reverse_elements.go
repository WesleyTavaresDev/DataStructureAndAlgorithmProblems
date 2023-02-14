package list

func ReverseElements(arr []int) []int {

	arrCopy := make([]int, 0, cap(arr))

	for i, j := 0, len(arrCopy)-1; i < j; i, j = i+1, j-1 {
		arrCopy[i], arrCopy[j] = arrCopy[j], arrCopy[i]
	}

	return arrCopy
}
