package day02

import "testing"

func TestPart1 (t *testing.T) {
	got := Part1("day2.test")
	want := 150

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2 (t *testing.T) {
	got := Part2("day2.test")
	want := 900

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
