package bubble_test

import (
	"bubble"
	"reflect"
	"testing"
)

func TestBubble(t *testing.T) {
	arr := []int{9, 8, 6, 2, 7, 3}

	bubble.BubbleSort(arr)

	want := []int{2, 3, 6, 7, 8, 9}

	if !reflect.DeepEqual(want, arr) {
		t.Errorf("Error: waiting %d, received %d", want, arr)
	}
}
