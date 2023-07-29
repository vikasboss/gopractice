package main

import "testing"

func TestHello(t *testing.T) {
	// Add a test case
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Gabriel")
		want := "Hello Gabriel"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("Say Hello World!", func(t *testing.T) {
		got := Hello("World!")
		want := "Hello World!"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})
}
