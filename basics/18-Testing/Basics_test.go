package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Nazar")
		want := "Hello Nazar! :)"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("Adding two numbers", func(t *testing.T) {
		got := Add(2, 3)
		want := 5

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestMultiple(t *testing.T) {
	t.Run("Multiplying two numbers", func(t *testing.T) {
		got := Multiple(2, 3)
		want := 6

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
