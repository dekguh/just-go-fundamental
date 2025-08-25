package main

import (
	"learn/just-go-fundamental/helpers"
	"testing"
)

func TestSumTotal(t *testing.T) {
	t.Run("should return 0 if the array is empty", func(t *testing.T) {
		expected := 0
		input := []int{}
		result := helpers.SumTotal((input))
		if result != expected {
			t.Errorf("expected is %d but got %d", expected, result)
		}
	})
}
