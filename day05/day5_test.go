package day05

import "testing"

func TestPart1(t * testing.T) {
	got := Part1("day5.test")
	want := 5

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t * testing.T) {
	got := Part2("day5.test")
	want := 12

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
