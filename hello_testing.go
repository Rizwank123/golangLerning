package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Riwan")
	want := "hello Rizwan"
	if got != want {
		t.Errorf("got %q wnat %q", got, want)
	}
}
