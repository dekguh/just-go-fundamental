package helpers

import "testing"

func TestAreaFormula(t *testing.T) {
	t.Run("Box", func(t *testing.T) {
		box := Box{2, 2, 2}
		got := box.AreaFormula()
		expected := 8.0

		if got != expected {
			t.Errorf("expected is %.2f but got %.2f", expected, got)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{2}
		got := circle.AreaFormula()
		expected := 12.56

		if got != expected {
			t.Errorf("expected is %.2f but got %.2f", got, expected)
		}
	})
}
