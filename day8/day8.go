package day8

import (
	"aoc_2021/utils"
	"fmt"
	"strings"
)

func Part1(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    count := 0
    for _, line := range lines {
        temp := strings.Split(line, " | ")
        rightSide := strings.Fields(temp[1])
        for _, digit := range rightSide {
            if len(digit) == 2 || len(digit) == 4 || len(digit) == 3 || len(digit) == 7 {
                count++
            }
        }
    }
    return count
}

func DeduceNumbers(line string) []int {
    split := strings.Split(line, " | ")
    left := split[0]
    leftArray := strings.Fields(left)
    for i, v := range leftArray {
        if len(v) == 2 {
            leftArray[i] = "1"
        }
        if len(v) == 3 {
            leftArray[i] = "7"
        }
        if len(v) == 4 {
            leftArray[i] = "4"
        }
        if len(v) == 7 {
            leftArray[i] = "8"
        }
    }
    fmt.Println(leftArray)

    return []int{}
}

func Part2(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    DeduceNumbers(lines[7])
    return len(lines)
}
