package day6

import "testing"

func TestPart1(t * testing.T) {
	got := Part1("day6.test")
	want := 5934

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t * testing.T) {
	got := Part2("day6.test")
	want := 26984457539

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
