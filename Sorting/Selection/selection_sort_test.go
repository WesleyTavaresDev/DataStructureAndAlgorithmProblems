package selection_test

import (
	"reflect"
	"selection"
	"testing"
)

func TestSelectionAscending(t *testing.T) {

	arr := []int{5, 4, 3, 2, 1, 0}
	selection.SelectionSort(arr, selection.Ascending)
	t.Log(arr)

	want := []int{0, 1, 2, 3, 4, 5}

	if !reflect.DeepEqual(want, arr) {
		t.Errorf("Wanting %d, received %d", want, arr)
	}
}

func TestSelectionDecreasing(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5}
	selection.SelectionSort(arr, selection.Decreasing)
	t.Log(arr)

	want := []int{5, 4, 3, 2, 1, 0}

	if !reflect.DeepEqual(want, arr) {
		t.Errorf("Wanting %d, received %d", want, arr)
	}
}
