package day03

import "testing"

func TestPart1(t * testing.T) {
	got := Part1("day3.test")
	want := 198

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t * testing.T) {
	got := Part2("day3.test")
	want := 230

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
