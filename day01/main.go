package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/aliusman3/aoc/util"
)

var digitStrings = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

func main() {

	inputLines := util.ReadInput("input")
	logFile := util.CreateLogFile()
	defer logFile.Close()
	logOutput := util.GetLogger(logFile)

	var findStringDigit = func(line string, idx int, lineSize int) (int, bool) {
		sub := ""
		if idx+5 <= lineSize {
			sub = line[idx : idx+5]
		} else {
			sub = line[idx:]
		}
		for k, v := range digitStrings {
			if strings.HasPrefix(sub, k) {
				logOutput(" found match for sub %s", sub)
				return v, true
			}
		}
		return -1, false
	}

	sum := 0
	for i, line := range inputLines {
		lineNumber := i + 1
		logOutput("lineNumber %d", lineNumber)
		logOutput(" line %s ", line)
		firstDigit := -1
		lastDigit := -1
		lineSize := len(line)
		for i, c := range line {
			if unicode.IsDigit(c) {
				firstDigit = int(c - '0')
				break
			} else if sd, found := findStringDigit(line, i, lineSize); found {
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
			} else if sd, found := findStringDigit(line, j, lineSize); found {
				lastDigit = sd
				break
			}
		}

		logOutput("\nfirstDigit %d lastDigit %d\n", firstDigit, lastDigit)
		val := firstDigit*10 + lastDigit
		logOutput("Val %d \n", val)
		sum += val
		logOutput("\n")
	}

	fmt.Println("sum", sum)
}
