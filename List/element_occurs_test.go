package list

import "testing"

func TestIsContained(t *testing.T) {

	elements := []int{1, 2, 3, 4, 5, 70, 7776, 20}

	target := 7776

	if !IsContains(elements, target) {
		t.Errorf("Element %d not found in %d", target, elements)
	}
}
