package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		sum := Sum(numbers)
		expected := 15

		if sum != expected {
			t.Errorf("expected %d got %d given %v", expected, sum, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{1, 1, 1})
	expected := []int{6, 3}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
