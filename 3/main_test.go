package integer

import "testing"

func TestAdd(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		sum := Add(3, 4)
		expected := 7
		if sum != expected {
			t.Errorf("expected %d got %d", expected, sum)
		}
	})
}
