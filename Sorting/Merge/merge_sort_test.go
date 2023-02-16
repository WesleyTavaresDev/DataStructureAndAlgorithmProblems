package merge_test

import (
	"reflect"
	merge "sorting/Merge"
	"testing"
)

func TestMergeSortOdd(t *testing.T) {

	arr := []int{42, 4, 67, 30, 45, 2, 7, 9, 10, 8}

	arr = merge.MergeSort(arr)

	want := []int{2, 4, 7, 8, 9, 10, 30, 42, 45, 67}

	if !reflect.DeepEqual(want, arr) {
		t.Errorf("Wanted %d, received %d", want, arr)
	}
}
