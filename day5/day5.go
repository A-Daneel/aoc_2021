package day5

import (
	"aoc_2021/utils"
	"fmt"
)

type operation struct {
    x1 int
    y1 int 
    x2 int
    y2 int
}

func CreateOperations(lines []string) ([]operation, int, int) { 
    operations := make([]operation, len(lines)) 
    validOperations := make([]operation, 0)
    maxX := 0
    maxY := 0
    for i, line := range lines {
        fmt.Sscanf(line, "%d,%d -> %d,%d", 
        &operations[i].x1,
        &operations[i].y1,
        &operations[i].x2,
        &operations[i].y2)

        if operations[i].x1 > maxX {
            maxX = operations[i].x1
        }
        if operations[i].x2 > maxX {
            maxX = operations[i].x2
        }
        if operations[i].y1 > maxY {
            maxY = operations[i].y1
        }
        if operations[i].y2 > maxY {
            maxY = operations[i].y2
        }
        if operations[i].x1 == operations[i].x2 || operations[i].y1 == operations[i].y2 {
            validOperations = append(validOperations, operations[i])
        }
    }
    return validOperations, maxX+1, maxY+1
}

func CountHeatmap(ventMap [][]int) int {
    count := 0
    for _, row := range ventMap {
        for _, value := range row {
            if value >= 2 {
                count++
            }
        }
    }
    return count
}

func SignHorizonalAndVertical(ventMap[][]int, op []operation) [][]int {
    for _, l := range op {
        if l.x1 == l.x2 {
            y1 := l.y1
            y2 := l.y2
            if y2 < y1 {
                y1 = l.y2
                y2 = l.y1
            }
            for y := y1; y <= y2; y++ {
                ventMap[l.x1][y]++
            }
        } else {
            x1 := l.x1
            x2 := l.x2
            if x2 < x1 {
                x1 = l.x2
                x2 = l.x1
            }
            for x := x1; x <= x2; x++ {
                ventMap[x][l.y1]++
            }
        }
    }
    return ventMap
}

func Part1(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    operations, maxX, maxY := CreateOperations(lines)
    ventMap := make([][]int, maxY)
    for i := 0; i < maxY; i++ {
        ventMap[i] = make([]int, maxX)
    }
    
    ventMap = SignHorizonalAndVertical(ventMap, operations)
    count := CountHeatmap(ventMap)
    return count
}

func Part2(file string) int {
    return 0x09
}
