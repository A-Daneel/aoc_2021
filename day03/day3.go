package day03

import (
	"aoc_2021/utils"
	"fmt"
	"strconv"
)



func Part1(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    gamma, epsilon := "", ""    
    for i := 0; i < len(lines[0]); i++ {
        most, least := CalculateMostAndLeast(lines, i)
        gamma += string(most) 
        epsilon += string(least)
    }

    return ReadBinary(gamma) * ReadBinary(epsilon)
}

func Part2(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    oxygen := FilterList(lines, 0, true)
    co2 := FilterList(lines, 0, false)
    
    return ReadBinary(oxygen) * ReadBinary(co2)
}

func FilterList(lines []string, i int, mostCommon bool) string {
    if len(lines) == 1 {
        return lines[0]
    }
    most, least := CalculateMostAndLeast(lines, i)
    returnList := make([]string, 0)
    comparator := least
    if mostCommon {
        comparator = most
    }

    for _, line := range lines {
        if comparator == line[i] {
            returnList = append(returnList, line)

        }
    }
    
    return FilterList(returnList, i+1, mostCommon)
}

func CalculateMostAndLeast(lines []string, i int) (uint8, uint8) {
    countZero, countOne := 0, 0
    for _, line := range lines {
        if line[i] == '0' {
            countZero++
        } else {
            countOne++
        }
    }
    if countZero > countOne {
        return '0', '1'
    }
    return '1', '0'
}

func ReadBinary(binary string) int {
    i, err := strconv.ParseInt(binary, 2, 64)
    if err != nil {
        fmt.Println(err)
    }
    return int(i)
}
