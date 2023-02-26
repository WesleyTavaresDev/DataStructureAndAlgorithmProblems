package ternarysearch

func TernarySearch(arr []int, l, r, target int) (index int) {
	if l < r {
		mid1 := l + (r-l)/3
		mid2 := r - (r-l)/3

		if target == arr[mid1] {
			return mid1
		} else if target == arr[mid2] {
			return mid2
		} else if target < arr[mid1] {
			return TernarySearch(arr, l, mid1-1, target)
		} else if target > arr[mid2] {
			return TernarySearch(arr, mid2+1, r, target)
		} else {
			return TernarySearch(arr, mid1+1, mid2-1, target)
		}
	}

	return -1
}
