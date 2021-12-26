package day09

import (
    "aoc_2021/utils"
    "sort"
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
    count := 0
    chart := createHightMap(lines)

    for y, row := range chart {
        for x := range row {
            if checkNeighbores(chart, y, x) {
                count += chart[y][x] + 1
            }
        }
    }

    return count
}


func Part2(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    count := 1 //were doing multiplications. if we start with 0 we keep zero
    chart := createHightMap(lines)
    chart2 := make([][]rune, len(lines))
    charCount := map[rune]int{}
    for i, row := range lines {
        chart2[i] = make([]rune, len(row))
    }

    currentBasin := 'A' //acii value of A comes before a
    for y, row := range chart {
        for x := range row {
            if chart2[y][x] != 0 {
                continue
            }
            if chart[y][x] == 9 {
                chart2[y][x] = '9'
                continue
            }

            type Point struct {
                x, y int
            }
            queue := []Point{{x,y}}

            for len(queue) > 0 {
                currentPoint := queue[0]
                //pops currentPoint off queue
                queue = queue[1:]

                chart2[currentPoint.y][currentPoint.x] = currentBasin

                if currentPoint.y > 0 &&
                chart[currentPoint.y-1][currentPoint.x] != 9 &&
                chart2[currentPoint.y-1][currentPoint.x] == 0 {
                    queue = append(queue, Point{currentPoint.x, currentPoint.y-1})
                }
                if currentPoint.y < len(chart)-1 &&
                chart[currentPoint.y+1][currentPoint.x] != 9 &&
                chart2[currentPoint.y+1][currentPoint.x] == 0 {
                    queue = append(queue, Point{currentPoint.x, currentPoint.y+1})
                }
                if currentPoint.x > 0 &&
                chart[currentPoint.y][currentPoint.x-1] != 9 &&
                chart2[currentPoint.y][currentPoint.x-1] == 0 {
                    queue = append(queue, Point{currentPoint.x-1, currentPoint.y})
                }
                if currentPoint.x < len(chart[0])-1 &&
                chart[currentPoint.y][currentPoint.x+1] != 9 &&
                chart2[currentPoint.y][currentPoint.x+1] == 0 {
                    queue = append(queue, Point{currentPoint.x+1, currentPoint.y})
                }
            }
            currentBasin++
        }
    }

    for y, row := range chart2 {
        for x := range row {
            if chart2[y][x] == '9' {
                continue
            }
            current := chart2[y][x]
            charCount[current]++
        }
    }

    var sortedBasins []int
    for _, value := range charCount {
        if value == '9' {
            continue
        }
        sortedBasins = append(sortedBasins, value)
    }

    sort.Ints(sortedBasins)
    for i := 1; i <= 3; i++ {
        count *= sortedBasins[len(sortedBasins)-i]
    }
    return count
} 
