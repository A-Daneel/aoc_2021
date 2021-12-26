package day10

import "testing"

func TestPart1(t * testing.T) {
	got := Part1("day10.test")
	want := 26397

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t * testing.T) {
	got := Part2("day10.test")
	want := 288957

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
