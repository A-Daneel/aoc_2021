package day08

import "testing"

func TestPart1(t * testing.T) {
	got := Part1("day8.test")
	want := 26

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t * testing.T) {
	got := Part2("day8.test")
	want := 61229

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
