package list

func IsContains(arr []int, target int) bool {

	for _, v := range arr {
		if v == target {
			return true
		}
	}

	return false
}
