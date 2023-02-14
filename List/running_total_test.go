package list

import (
	"reflect"
	"testing"
)

func TestTotal(t *testing.T) {

	elements := []int{2, 4, 6, 8, 20}

	total := RunTotal(elements)
	want := []int{2, 6, 12, 20, 40}

	if !reflect.DeepEqual(want, total) {
		t.Errorf("Wanting %d, received %d", want, total)
	}
}

func TestTotalNegative(t *testing.T) {
	elements := []int{-2, -4, -6, -8, 20}

	total := RunTotal(elements)
	want := []int{-2, -6, -12, -20, 0}

	if !reflect.DeepEqual(want, total) {
		t.Errorf("Wanting %d, received %d", want, total)
	}
}
