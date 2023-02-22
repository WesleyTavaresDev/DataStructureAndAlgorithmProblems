package linearseach

import (
	"reflect"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	arr := []int{100, 98, 11, 61, 30, 47, 42, 77, 54, 91, 99}

	targetIndex := LinearSearch(arr, 99)
	want := 10

	if !reflect.DeepEqual(want, targetIndex) {
		t.Errorf("Want %d, received %d", want, targetIndex)
	}
}

func TestLinearSearchNotFound(t *testing.T) {
	arr := []int{100, 98, 11, 61, 30, 47, 42, 77, 54, 91, 99}

	targetIndex := LinearSearch(arr, 250230)

	want := -1

	if !reflect.DeepEqual(want, targetIndex) {
		t.Errorf("Want %d, received %d", want, targetIndex)
	}
}

func TestBigLinearSearch(t *testing.T) {
	arr := []int{100, 98, 11, 61, 30, 47, 42, 77, 54, 91, 99, 52, 95, 49, 45, 15, 23, 3, 60, 72, 102, 59, 44, 31, 51, 41}
	targetIndex := LinearSearch(arr, 41)

	want := len(arr) - 1

	if !reflect.DeepEqual(want, targetIndex) {
		t.Errorf("Want %d, received %d", want, targetIndex)
	}
}
