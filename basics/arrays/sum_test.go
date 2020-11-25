package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	given := []int{1, 2, 3, 4, 5}
	got := Sum(given)
	want := 15
	if got != want {
		t.Errorf("got %d, expected %d, given %v", got, want, given)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		givenA := []int{1, 2, 3}
		givenB := []int{1, 2, 3}
		got := SumAll(givenA, givenB)
		want := []int{2, 4, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, expected %v", got, want)
		}
	})

	t.Run("input have different length", func(t *testing.T) {
		// givenA := []int{1, 2, 3}
		// givenB := []int{1, 2, 3, 4}
		// got := SumAll(givenA, givenB)
		// want := make([]int, 0)
		// if !reflect.DeepEqual(got, want) {
		// 	t.Errorf("got %v, expected %v", got, want)
		// }
	})
}
