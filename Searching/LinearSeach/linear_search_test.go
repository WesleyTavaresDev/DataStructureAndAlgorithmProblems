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
