package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("world")
	want := "hello"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
