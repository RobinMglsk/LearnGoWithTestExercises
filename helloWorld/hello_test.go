package main

import "testing"

func TestHello(t *testing.T){
	got :=Hello("Robin")
	want := "Hello, Robin"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}