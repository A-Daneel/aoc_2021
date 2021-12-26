package day04

import "testing"

func TestPart1(t * testing.T) {
	got := Part1("day4.test")
	want := 4512

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t * testing.T) {
	got := Part2("day4.test")
	want := 1924

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
