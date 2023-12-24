package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

var digitStrings = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

func findStringDigit(line string, idx int, lineSize int, w io.Writer) (int, bool) {
	sub := ""
	if idx+5 <= lineSize {
		sub = line[idx : idx+5]
	} else {
		sub = line[idx:]
	}
	for k, v := range digitStrings {
		if strings.HasPrefix(sub, k) {
			fmt.Fprintf(w, " found match for sub %s", sub)
			return v, true
		}
	}
	return -1, false
}

func main() {
	sum := 0
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result, err := os.Create("result.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	scanner := bufio.NewScanner(file)
	lines := 1
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Fprintf(result, "Found empty string at index %d\n", lines)
			continue
		}
		fmt.Fprintf(result, "lines %d", lines)
		fmt.Fprintf(result, " line %s ", line)
		firstDigit := -1
		lastDigit := -1
		lineSize := len(line)
		for i, c := range line {
			if unicode.IsDigit(c) {
				firstDigit = int(c - '0')
				break
			} else if sd, found := findStringDigit(line, i, lineSize, result); found {
				firstDigit = sd
				break
			}
		}

		for i := range line {
			j := lineSize - 1 - i
			c := rune(line[j])
			if unicode.IsDigit(c) {
				lastDigit = int(c - '0')
				break
			} else if sd, found := findStringDigit(line, j, lineSize, result); found {
				lastDigit = sd
				break
			}
		}

		fmt.Fprintf(result, "\nfirstDigit %d lastDigit %d\n", firstDigit, lastDigit)
		val := firstDigit*10 + lastDigit
		fmt.Fprintf(result, "Val %d \n", val)
		sum += val
		lines++
		fmt.Fprintln(result)
	}

	fmt.Println("sum", sum)
}
