package day09

import "testing"

func TestPart1(t * testing.T) {
	got := Part1("day9.test")
	want := 15

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t * testing.T) {
	got := Part2("day9.test")
	want := 1134

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
