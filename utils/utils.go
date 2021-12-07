package utils

import (
	"bufio"
	"log"
	"os"
)

func ReadDataToArray(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	values := []string{}
	for scanner.Scan() {
		val := scanner.Text()
		values = append(values, val)
	}
	return values, scanner.Err()
}

func StringToInt(s string) int {
    i := 0
    for _, ch := range s {
        i *= 10
        i += int(ch) - 0x30
    }
    return i
}
