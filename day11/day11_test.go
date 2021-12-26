package day11

import "testing"

func TestPart1(t * testing.T) {
	got := Part1("day11.test")
	want := 1656

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t * testing.T) {
	got := Part2("day11.test")
	want := 195

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
