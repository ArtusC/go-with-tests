package main

import "testing"

func TestHello(t *testing.T){
	got := Hello("Artus")
	want := "Hello, super uper world and Artus"

	if got != want{
		t.Errorf("got %q | want %q", got, want)
	}
}