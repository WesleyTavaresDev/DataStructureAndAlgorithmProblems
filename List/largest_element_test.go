package list

import "testing"

func TestLargest(t *testing.T) {
	shouldReturn25 := [...]int{0, 0, 9, 1, 25, 6, 13}
	want := 25
	returned := LargestElement(shouldReturn25[:])

	if returned != want {
		t.Errorf("result %d, waiting %d", returned, want)
	}
}
