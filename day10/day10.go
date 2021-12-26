package day10

import (
	"aoc_2021/utils"
	"sort"
)

func resolveLines(lines []string) (int, int) {
    highscore := 0 
    incompletes := []string{}
    for _, line := range lines {
        stack := []rune{}
        corrupt := rune(0)
        // do while plz
        loop:
        for _, char := range line {
            switch char {
            case '(', '[', '{', '<':
                stack = append(stack, char)
            case ')', ']', '}', '>':
                pop := stack[len(stack)-1]
                stack = stack[:len(stack)-1]
                if pop == '(' && char == ')' ||
                pop == '[' && char == ']' ||
                pop == '{' && char == '}' ||
                pop == '<' && char == '>'{
                    continue
                } else {
                    corrupt = char
                    break loop
                }
            }
        }
        switch corrupt {
        case ')':
            highscore += 3
        case ']':
            highscore += 57
        case '}':
            highscore += 1197
        case '>':
            highscore += 25137
        default:
            incompletes = append(incompletes, string(stack))
        }
    }

    scores := make([]int, len(incompletes))
    for i, line := range incompletes {
        for char := len(line)-1; char >= 0; char-- {
            scores[i] *= 5
            switch line[char] {
            case '(':
                scores[i] += 1
            case '[':
                scores[i] += 2
            case '{':
                scores[i] += 3
            case '<':
                scores[i] += 4
            }
        }
    }
    sort.Ints(scores)
    return highscore, scores[len(scores)/2]
}

func Part1(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    part1, _ := resolveLines(lines)
    return part1
}


func Part2(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    _, part2 := resolveLines(lines)
    return part2
} 
