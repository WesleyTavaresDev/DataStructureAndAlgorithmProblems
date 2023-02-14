package list

import (
	"reflect"
	"testing"
)

func TestReverseElements(t *testing.T) {

	elements := []int{5, 4, 2}
	t.Log(elements)

	want := ReverseElements(elements)

	t.Log(want)
	if reflect.DeepEqual(elements, want) {
		t.Errorf("original %d, waiting %d", elements[:], want[:])
	}
}
