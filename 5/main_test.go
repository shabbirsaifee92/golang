package main

import "testing"

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
