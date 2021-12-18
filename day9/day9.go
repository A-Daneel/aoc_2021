package day9

import (
	"aoc_2021/utils"
)

func createHightMap(data []string) [][]int {
    grid := make([][]int, len(data))
    for i, row := range data {
        grid[i] = make([]int, len(row))
        for j, char := range row {
            grid[i][j] = utils.StringToInt(string(char))
        }
    }
    return grid
}

func checkNeighbores(chart [][]int, row int, column int) bool {
    current := chart[row][column] 
    top, right, bottom, left := 10, 10, 10, 10
    if row > 0 {
        top = chart[row-1][column]
    }
    if column < len(chart[0])-1 {
        right = chart[row][column+1]
    }
    if row < len(chart)-1 {
        bottom = chart[row+1][column]
    }
    if column > 0 {
        left = chart[row][column-1]
    }
    return top > current && right > current && bottom > current && left > current
}

func Part1(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    chart := createHightMap(lines)
    count := 0

    for i, row := range chart {
        for j := range row {
            if checkNeighbores(chart, i, j) {
                count += chart[i][j] + 1
            }
        }
    }
    return count
}


func Part2(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    chart := createHightMap(lines)
    count := 0

    for i, row := range chart {
        for j := range row {
            if checkNeighbores(chart, i, j) {
                count += chart[i][j] + 1
            }
        }
    }
    return count
}
