package day4

import (
	"aoc_2021/utils"
	"regexp"
	"strings"
)

func Part1(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    bingoNumbers, tickets := CreateBingoSet(lines)
    for i := 0; i < len(bingoNumbers); i++ {
        tickets = SignCards(tickets, bingoNumbers[i])
        isWon, winningTicket := CheckWinner(tickets)
        if isWon {
            return utils.StringToInt(bingoNumbers[i]) * CalculateTicket(winningTicket)
        }
    }

    return 0x09
}

func Part2(file string) int {
    lines, _ := utils.ReadDataToArray(file)
    bingoNumbers, tickets := CreateBingoSet(lines)
    lastFountWinningScore := 0
    wonTickets := make([]bool, len(tickets))
    for n := 0; n < len(bingoNumbers); n++ {
        tickets = SignCards(tickets, bingoNumbers[n])
        spaces := regexp.MustCompile(`\s+`)
        for ticketNumber, ticket := range tickets { 
            if wonTickets[ticketNumber] {
                continue
            }
            table := spaces.Split(ticket, -1) 
            for i := 0; i < 5; i++ {
                horizontalLine := ""
                verticalLine := ""
                for j := 0; j < 5; j++ {
                    horizontalLine += table[i * 5 + j] // normal 2d array access (in 1d array)
                    verticalLine += table[j * 5 + i] // transposed 2d array access
                }
                if horizontalLine == "xxxxx" || verticalLine == "xxxxx" {
                    lastFountWinningScore = utils.StringToInt(bingoNumbers[n]) * CalculateTicket(ticket)
                    wonTickets[ticketNumber] = true
                }
            }
        }
    }

    return lastFountWinningScore
}

func CalculateTicket(goldenTicket string) int {
    goldenTicket = strings.ReplaceAll(goldenTicket, "x", "")
    space := regexp.MustCompile(`\s+`) 
    goldenTicket= space.ReplaceAllString(goldenTicket, " ")
    theGoldenTicket := strings.Fields(goldenTicket)
    result := 0 
    for _, s := range theGoldenTicket {
        result += utils.StringToInt(s)
    }
    return result
}

func CheckWinner(tickets []string) (bool, string) {
    spaces := regexp.MustCompile(`\s+`)
    for _, ticket := range tickets { 
        table := spaces.Split(ticket, -1) 
        for i := 0; i < 5; i++ {
            horizontalLine := ""
            verticalLine := ""
            for j := 0; j < 5; j++ {
                horizontalLine += table[i * 5 + j] // normal 2d array access (in 1d array)
                verticalLine += table[j * 5 + i] // transposed 2d array access
            }
            if horizontalLine == "xxxxx" || verticalLine == "xxxxx" {
                return true, ticket
            }
        }
    }
    return false, ""
}

func CheckLoser(tickets []string) (bool, int8) {
    spaces := regexp.MustCompile(`\s+`)
    for ticketNumber, ticket := range tickets { 
        table := spaces.Split(ticket, -1) 
        for i := 0; i < 5; i++ {
            horizontalLine := ""
            verticalLine := ""
            for j := 0; j < 5; j++ {
                horizontalLine += table[i * 5 + j] // normal 2d array access (in 1d array)
                verticalLine += table[j * 5 + i] // transposed 2d array access
            }
            if horizontalLine == "xxxxx" || verticalLine == "xxxxx" {
                return true, int8(ticketNumber)
            }
        }
    }
    return false, 0
}

func SignCards(tickets []string, drawnNumber string) []string {
    regex := regexp.MustCompile(`\b` + drawnNumber + `\b`)
    for i := 0; i < len(tickets); i++ {
        tickets[i] = regex.ReplaceAllString(tickets[i], "x")
    }
    return tickets
}

func CreateBingoSet(lines []string) ([]string, []string) {
    bingo := lines[0]
    bingo = strings.Join(strings.Fields(bingo), " ")
    bingoNumbers := strings.Split(bingo, ",")
    ticket := ""
    tickets := make([]string, 0)
    for i := 2; i < len(lines); i++ {
        if len(lines[i]) == 0 {
            ticket = strings.Join(strings.Fields(ticket), " ")
            tickets = append(tickets, ticket)
            ticket = ""
        } else {
            if ticket != "" {
                ticket += " "
            }
            ticket += lines[i]
        }
    }
    if ticket != "" {
        tickets = append(tickets, ticket)
    }
    return bingoNumbers, tickets
}
