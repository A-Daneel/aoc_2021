package day5

import (
	"aoc_2021/utils"
	"fmt"
	"math"
)

type operation struct {
    x1 int
    y1 int 
    x2 int
    y2 int
}

func CreateOperations(lines []string) ([]operation, []operation, int, int) { 
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
    return operations, validOperations, maxX+1, maxY+1
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

func SignHorizonalAndVertical(ventMap[][]int, op []operation, considerDiagonal bool) [][]int {
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
        } else if l.y1 == l.y2{
            x1 := l.x1
            x2 := l.x2
            if x2 < x1 {
                x1 = l.x2
                x2 = l.x1
            }
            for x := x1; x <= x2; x++ {
                ventMap[x][l.y1]++
            }
        } else {
            if !considerDiagonal { continue } 
            fmt.Println("help")
            deltaX := int(float64(l.x2-l.x1)/ math.Abs(float64(l.x2-l.x1)))
            deltaY := int(float64(l.y2-l.y1)/ math.Abs(float64(l.y2-l.y1)))
            x := l.x1
            y := l.y1

            for y != l.y2+deltaY {
                ventMap[x][y]++
                x += deltaX
                y += deltaY
            }
        }
    }
    return ventMap
}

func Part1(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    _, operations, maxX, maxY := CreateOperations(lines)
    ventMap := make([][]int, maxY)
    for i := 0; i < maxY; i++ {
        ventMap[i] = make([]int, maxX)
    }
    
    ventMap = SignHorizonalAndVertical(ventMap, operations, false)
    count := CountHeatmap(ventMap)
    return count
}

func Part2(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    operations, _, maxX, maxY := CreateOperations(lines)
    ventMap := make([][]int, maxY)
    for i := 0; i < maxY; i++ {
        ventMap[i] = make([]int, maxX)
    }
    
    ventMap = SignHorizonalAndVertical(ventMap, operations, true)
    count := CountHeatmap(ventMap)
    return count
}
