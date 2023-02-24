package binarysearch

func BinarySearch(arr []int, target int) (index int) {
	low, high := 0, len(arr)-1

	for high >= low {
		mid := (low + high) / 2

		if arr[mid] > target {
			high = mid - 1
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
