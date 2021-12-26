package day07

import "testing"

func TestPart1(t * testing.T) {
	got := Part1("day7.test")
	want := 37

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t * testing.T) {
	got := Part2("day7.test")
	want := 168

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
