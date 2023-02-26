package ternarysearch

import (
	"reflect"
	"testing"
)

func TestTernarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	want := 8

	targetIndex := TernarySearch(arr, 0, len(arr)-1, 9)

	if !reflect.DeepEqual(want, targetIndex) {
		t.Errorf("Want %d, received %d", want, targetIndex)
	}
}
