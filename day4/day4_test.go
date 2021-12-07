package day4

import "testing"

func TestPart1(t * testing.T) {
	got := Part1("day4.test")
	want := 4512

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

