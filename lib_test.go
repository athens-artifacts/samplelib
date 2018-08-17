package samplelib

import "testing"

func TestHello(t *testing.T) {
	want:= "rawr"
	got := Hello()
	if got !=want{
		t.Fatalf("expected %q got %q", want, got)
	}
}
