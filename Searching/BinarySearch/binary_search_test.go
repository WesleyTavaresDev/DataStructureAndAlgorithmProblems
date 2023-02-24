package binarysearch

import (
	"reflect"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 51}

	targetIndex := BinarySearch(arr, 0)

	want := 0

	if !reflect.DeepEqual(want, targetIndex) {
		t.Errorf("Want %d, received %d", want, targetIndex)
	}
}

func TestBinarySearchUnsorted(t *testing.T) {
	arr := []int{7, 10, 9, 12, 51, 5, 6, 2, 1, 0, 4}

	targetIndex := BinarySearch(arr, 0)

	want := -1

	if !reflect.DeepEqual(want, targetIndex) {
		t.Errorf("Want %d, received %d", want, targetIndex)
	}
}
