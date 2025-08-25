package main

import (
	"learn/just-go-fundamental/helpers"
	"testing"
)

func TestSumTotal(t *testing.T) {
	t.Run("should return 0 if the array is empty", func(t *testing.T) {
		expected := float32(0)
		input := []float32{}
		result := helpers.SumTotal((input))
		if result != expected {
			t.Errorf("expected is %d but got %d", expected, result)
		}
	})
}

func TestVolumeFormula(t *testing.T) {
	t.Run("should return 0 if the height, width, and length are 0", func(t *testing.T) {
		expected := float64(0)
		result := helpers.VolumeFormula(0, 0, 0)
		if result != expected {
			t.Errorf("expected is %.2f but got %.2f", expected, result)
		}
	})

	t.Run("should return the volume of the box", func(t *testing.T) {
		expected := float64(100)
		result := helpers.VolumeFormula(10.5, 2.5, 10.5)
		if result != expected {
			t.Errorf("expected is %.2f but got %.2f", expected, result)
		}
	})
}
