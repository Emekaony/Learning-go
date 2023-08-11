package main

import "testing"

func TestAdd(t *testing.T) {
	x := AddNer(4)
	n := x(5)

	if n != 9 {
		t.Fatalf("expected n = 10, got %d", n)
	}
}
