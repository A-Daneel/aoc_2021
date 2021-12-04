package day2

import (
	"aoc_2021/utils"
	"log"
	"strings"
)

func Part1(file string) int {
    directions, err := utils.ReadDataToArray(file)
	if err != nil {
		log.Fatal(err)
	} 
    horizontal, vertical := 0, 0
    for i := 0; i < len(directions); i++ {
        directions, value := SplitLine(directions[i])
        switch directions {
            case "forward":
                horizontal += value
            case "down":
                vertical += value
            case "up":
                vertical -= value
        }
    }
    return horizontal * vertical
}

func Part2(file string) int {
    directions, err := utils.ReadDataToArray(file)
	if err != nil {
		log.Fatal(err)
	} 
    horizontal, vertical, aim := 0, 0, 0
    for i := 0; i < len(directions); i++ {
        directions, value := SplitLine(directions[i])
        switch directions {
            case "forward":
                vertical += value * aim
                horizontal += value
            case "down":
                aim += value
            case "up":
                aim -= value
        }
    }
    return horizontal * vertical
}

func SplitLine(line string) (string, int) {
    splitted := strings.Fields(line)
    direction := splitted[0]
    value := utils.StringToInt(splitted[1])
    return direction, value
}
