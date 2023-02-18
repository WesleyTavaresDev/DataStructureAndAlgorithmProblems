package quicksort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {

	arr := []int{5, 8, 6, 2, 0, 78, 10, 11, 3, 54}

	Quicksort(arr, 0, len(arr)-1)

	want := []int{0, 2, 3, 5, 6, 8, 10, 11, 54, 78}

	if !reflect.DeepEqual(want, arr) {
		t.Errorf("Wanted %d, received %d", want, arr)
	}
}
