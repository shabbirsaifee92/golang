package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeat := Repeat("a")
	expected := "aaaaa"

	if repeat != expected {
		t.Errorf("expected %q got %q", expected, repeat)
	}
}

func BenchmarkRepeat(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Repeat("a")
	}
}
