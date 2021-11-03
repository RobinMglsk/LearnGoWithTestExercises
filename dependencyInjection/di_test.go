package di

import (
	"bytes"
	"testing"
)

func TestGreeter(t *testing.T){
	buffer := bytes.Buffer{}
	Greet(&buffer, "Robin")

	got := buffer.String()
	want := "Hello, Robin"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}