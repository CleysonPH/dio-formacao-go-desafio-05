package main

import "testing"

func TestShouldReturnCorretValueWhenSumFunctionIsCalled(t *testing.T) {
	got := sum(1, 2, 3, 4, 5)
	want := 15
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestShouldReturnCorretValueWhenSubFunctionIsCalled(t *testing.T) {
	got := sub(1, 2, 3, 4, 5)
	want := -15
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestShouldReturnCorretValueWhenMulFunctionIsCalled(t *testing.T) {
	got := mul(1, 2, 3, 4, 5)
	want := 120
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestShouldReturnCorretValueWhenDivFunctionIsCalled(t *testing.T) {
	got := div(1, 2, 3, 4, 5)
	want := 0
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
