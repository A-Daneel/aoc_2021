package day9

import "testing"

func TestPart1(t * testing.T) {
	got := Part1("day9.test")
	want := 15

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
