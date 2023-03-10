package insertion_test

import (
	"reflect"
	insertion "sorting/Insertion"
	"testing"
)

func TestInsertionSort(t *testing.T) {

	arr := []int{7, 4, 2, 5, 8, 0, 1}

	insertion.Insertion(arr)

	want := []int{0, 1, 2, 4, 5, 7, 8}

	if !reflect.DeepEqual(want, arr) {
		t.Errorf("Wanted %d, received %d", want, arr)
	}
}

func TestInsertSortWithEmptyList(t *testing.T) {
	arr := make([]int, 0)

	if arr, e := insertion.Insertion(arr); e == nil {
		t.Errorf("%d is an empty list", arr)
	}
}
