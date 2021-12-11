package day6

import (
	"aoc_2021/utils"
)

func CreateNextGeneration(fishes []int) []int {
    nextGeneration := make([]int, len(fishes))
    for i, fish := range fishes {
        nextGeneration[i] = fish-1
        if (nextGeneration[i] == -1) {
            nextGeneration[i] = 6
            nextGeneration = append(nextGeneration, 8)
        }
    }

    return nextGeneration
}

func FishArray(fishy []int) []int {
    fishArray := make([]int, 9)
    for _, fish := range fishy {
        fishArray[fish]++
    }

    return fishArray
}

func NextGen(fishes []int) []int {
    nextGen := make([]int, 9)
    for i := 1; i < 9; i++ {
        nextGen[i-1] += fishes[i]
        
    }
    nextGen[6] += fishes[0]
    nextGen[8] += fishes[0]
    return nextGen
}

func CountFish(allTheFish []int) int {
    sum := 0
    for fish := range allTheFish {
        sum += allTheFish[fish]
    }
    return sum
}

func Part1(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    fishyLifeTime := utils.ReadStringArrayToIntArray(lines[0])
    for i := 0; i < 80; i++ {
        fishyLifeTime = CreateNextGeneration(fishyLifeTime) 
    }
    return len(fishyLifeTime)
}

func Part2(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    fishyLifeTime := utils.ReadStringArrayToIntArray(lines[0])
    rayFish := FishArray(fishyLifeTime)
    for i := 0; i < 256; i++ {
        rayFish = NextGen(rayFish)
    }
    return CountFish(rayFish)
}
