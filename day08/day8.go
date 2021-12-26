package day08

import (
	"aoc_2021/utils"
	"sort"
	"strings"
)

func Part1(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    count := 0
    for _, line := range lines {
        temp := strings.Split(line, " | ")
        rightSide := strings.Fields(temp[1])
        for _, digit := range rightSide {
            if len(digit) == 2 || len(digit) == 4 || len(digit) == 3 || len(digit) == 7 {
                count++
            }
        }
    }
    return count
}

func deduceNumbers(line []string) map[string]int {
    var one, four, six, seven, eight string
    var length5 []string
    var length6 []string
    for _, v := range line {
        if len(v) == 2 {
            one = normalize(v)
        } else if len(v) == 3 {
            seven = normalize(v)
        } else if len(v) == 4 {
            four = normalize(v)
        } else if len(v) == 7 {
            eight = normalize(v)
        } else if len(v) == 5 {
            length5 = append(length5, normalize(v))
        } else if len(v) == 6 {
            length6 = append(length6, normalize(v))
        }
    }
    
    returnMap := map[string]int {
        one: 1,
        four: 4,
        seven: 7,
        eight: 8,
    }

    // find 6
    for i, d := range length6 {
        for j := range one {
            if !strings.Contains(d, string(one[j])) {
                returnMap[d] = 6
                six = d
                length6 = append(length6[:i], length6[i+1:]...)
                break
            }
        }
    }

    //find 0
    for i, d := range length6 {
        for j := range four {
            if !strings.Contains(d, string(four[j])) {
                returnMap[d] = 0
                length6 = append(length6[:i], length6[i+1:]...)
                break
            }
        }
    }

    // last one is 9
    for i, d := range length6 {
        for j := range four {
            if strings.Contains(d, string(four[j])) {
                returnMap[d] = 9
                length6 = append(length6[:i], length6[i+1:]...)
                break
            }
        }
    }

    //find 3
    for i, d := range length5 {
        if strings.Contains(d, string(one[0])) &&
        strings.Contains(d, string(one[1])) {
            returnMap[d] = 3
            length5 = append(length5[:i], length5[i+1:]...)
            break
        }
    }

    //find 5
    for i, d := range length5 {
        if strings.Contains(six, string(d[0])) &&
        strings.Contains(six, string(d[1])) &&
        strings.Contains(six, string(d[2])) &&
        strings.Contains(six, string(d[3])) &&
        strings.Contains(six, string(d[4])) {
            returnMap[d] = 5
            length5 = append(length5[:i], length5[i+1:]...)
            break
        }
    }

    // last one is 2
    // for fun, maybe try another better way?
    returnMap[length5[0]] = 2

    return returnMap
}

func normalize(input string) string {
    array := strings.Split(input, "")
    sort.Strings(array)
    return strings.Join(array, "")
}

func Part2(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    total := 0
    for _, line := range lines {
        lineArray := strings.Split(line, " | ")
        right := strings.Split(lineArray[1], " ")
        
        output := deduceNumbers(strings.Split(lineArray[0], " "))

        asDigits := output[normalize(right[0])]* 1000 +
        output[normalize(right[1])]* 100 +
        output[normalize(right[2])]* 10 +
        output[normalize(right[3])]
        
        total += asDigits
    }
    return total
}
