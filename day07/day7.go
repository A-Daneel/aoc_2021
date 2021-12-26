package day07

import (
	"aoc_2021/utils"
	"math"
)

func FindMax(input []int) int {
    max := 0
    for _, i := range input {
        if i > max {
            max = i
        }
    }
    return max
}

func FindMin(input []int) int {
    min := input[0]
    for _, i := range input {
        if i < min {
            min = i
        }
    }
    return min
}

// triangle numbers, should read up.
// think I remember this correctly.
func MovementCost(i int) int {
    return (i * (i + 1)) / 2
}

func Part1(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    crabsPositions := utils.ReadStringArrayToIntArray(lines[0])
    max := FindMax(crabsPositions)
    fuelUsage := make([]int, max+1)
    for i := 0; i <= max; i++ {
        for _, v := range crabsPositions {
            fuelUsage[i] += int(math.Abs(float64(v-i)))
        }
    }
    return  FindMin(fuelUsage)
}

func Part2(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    crabsPositions := utils.ReadStringArrayToIntArray(lines[0])
    max := FindMax(crabsPositions)
    fuelUsage := make([]int, max+1)
    for i := 0; i <= max; i++ {
        for _, v := range crabsPositions {
            fuelUsage[i] += MovementCost(int(math.Abs(float64(v-i))))
        }
    }
    return  FindMin(fuelUsage)
}
