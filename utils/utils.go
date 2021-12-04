package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

func StringToInt(value string) (val int) {
    val, err := strconv.Atoi(value)
    if err != nil {
        log.Fatal(err)
    }
    return
}
