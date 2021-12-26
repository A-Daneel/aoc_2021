package day11

import (
	"aoc_2021/utils"
)

func step(input[][]int) int {
    flashes := 0
    flashed := [10][10]bool{}
    //increment all by one, that happens always
    for i := range input {
        for j := 0; j < 10; j++ {
            input[i][j]++
        }
    }

    // loopy loop
    updateHappend := true
    for updateHappend {
        updateHappend = false

        for y := range input {
            for x := 0; x < 10; x++ {
                if flashed[y][x] {
                    continue
                }

                if input[y][x] > 9 {
                    updateHappend = true
                    flashed[y][x] = true

                    //increment neighbores
                    if y - 1 >= 0 {
                        if x - 1 >= 0 {
                            input[y-1][x-1]++
                        }
                        input[y-1][x]++
                        if x + 1 < 10 {
                            input[y-1][x+1]++
                        }

                    }
                    if x - 1 >= 0 {
                        input[y][x-1]++
                    }
                    if x + 1 < 10 {
                        input[y][x+1]++
                    }
                    if y + 1 < 10 {
                        if x - 1 >= 0 {
                            input[y+1][x-1]++
                        }
                        input[y+1][x]++
                        if x + 1 < 10 {
                            input[y+1][x+1]++
                        }
                    }
                }
            }
        }
    }
    
    for y := range input {
        for x := 0; x < 10; x++ {
            if flashed[y][x] {
                input[y][x] = 0
                flashes++
            }
        }
    }
    return flashes
}

func createGrid(input []string) [][]int {
    grid := make([][]int, len(input))
    for i := range grid {
        grid[i] = make([]int, 10)
        for j := 0; j < 10; j++ {
            grid[i][j] = int(input[i][j] - byte('0'))
        }
    }
    return grid
}

func Part1(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    flashes := 0
    grid := createGrid(lines)
    for i := 0; i < 100; i++ {
        flashes += step(grid)
    }
    return flashes
}


func Part2(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    flashes := 0
    count := 0
    grid := createGrid(lines)
    for {
        flashes = step(grid)
        count++
        if flashes == 100 {
            break
        }
    }
    return count
} 
