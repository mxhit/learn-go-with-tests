package arraysslices

import (
	"reflect"
	"testing"
)

var checkSums = func(t testing.TB, got, want []int) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSum(t *testing.T) {
	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Sum of all slices passed", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{0, 9})
		want := []int{6, 9}

		checkSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("Sum of tails of all slices passed", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("Safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})

	t.Run("Passing a slice of length 1", func(t *testing.T) {
		got := SumAllTails([]int{1}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})

}
