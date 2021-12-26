package day01

import (
	"aoc_2021/utils"
	"log"
)

func Part1(file string) int {
    isDeeper := 0
   	depths, err := utils.ReadDataToArray(file)
	if err != nil {
		log.Fatal(err)
	} 
    for i := 0; i < len(depths)-1; i++ {
        if utils.StringToInt(depths[i]) < utils.StringToInt(depths[i+1]) {
            isDeeper++
        }
    }
    return isDeeper
}

func Part2(file string) int {
    isDeeper := 0
   	depths, err := utils.ReadDataToArray(file)
	if err != nil {
		log.Fatal(err)
	} 
    for i := 0; i < len(depths)-3; i++ {
        first := utils.StringToInt(depths[i]) + utils.StringToInt(depths[i+1]) + utils.StringToInt(depths[i+2])
        second := utils.StringToInt(depths[i+1]) + utils.StringToInt(depths[i+2]) + utils.StringToInt(depths[i+3])
        if first < second {
            isDeeper++
        }
    }
    return isDeeper
}
