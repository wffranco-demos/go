package main

import "testing"

func TestMax(t *testing.T) {
	got := Max(1, 2)
	want := 2

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

	got = Max(3, 2)
	want = 3
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
