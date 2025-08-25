package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("world")
		want := "hello"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
