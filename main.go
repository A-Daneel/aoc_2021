package main

import (
	"aoc_2021/day1"
	"aoc_2021/day2"
	"aoc_2021/day3"
	"fmt"
)

func main() {
    fmt.Print("day 1\n")
	fmt.Printf("result from part 1: %d\n", day1.Part1("day1/day1.real"))
	fmt.Printf("result from part 2: %d\n", day1.Part2("day1/day1.real"))

    fmt.Print("\nday 2\n")
	fmt.Printf("result from part 1: %d\n", day2.Part1("day2/day2.real"))
	fmt.Printf("result from part 2: %d\n", day2.Part2("day2/day2.real"))

    fmt.Print("\nday 3\n")
	fmt.Printf("result from part 1: %d\n", day3.Part1("day3/day3.real"))
	fmt.Printf("result from part 2: %d\n", day3.Part2("day3/day3.real"))
}
