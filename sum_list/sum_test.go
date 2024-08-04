package sum_list

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("wanted %q but got %q, given %v", want, got, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}

		got := Sum(numbers)
		want := 21

		if got != want {
			t.Errorf("wanted %q but got %q, given %v", want, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{2, 3}, []int{1, 6, 5})
	want := []int{5, 12}

	if !reflect.DeepEqual(got,want) {
		t.Errorf("got %v want %v", got, want)
	}
}
