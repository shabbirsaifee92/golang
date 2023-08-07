package main

import (
	"math"
	"testing"
)

func TestPerimter(t *testing.T) {
	t.Run("Reactange Perimieter", func(t *testing.T) {
		rect := Rectangle{10, 5}
		got := rect.Perimieter()
		expected := 30.0

		if got != expected {
			t.Errorf("expected %.2f got %.2f", expected, got)
		}
	})
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, math.Pi*10*10)
	})

}
