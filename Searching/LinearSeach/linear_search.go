package linearseach

func LinearSearch(arr []int, target int) (index int) {
	for index := 0; index < len(arr); index++ {
		if arr[index] == target {
			return index
		}
	}

	return -1
}
