package main

import (
	"fmt"
	"log"
	"strconv"
	"unicode"

	"github.com/aliusman3/aoc/util"
)

func isSymbol(c rune) bool {
	return !unicode.IsDigit(c) && c != '.'
}

func main() {
	inputLines := util.ReadInput("input")

	logFile := util.CreateLogFile()
	logOutput := util.GetLogger(logFile)

	directions := [][]int{
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}
	width := len(inputLines[0])
	height := len(inputLines)
	logOutput("width %d height %d\n\n", width, height)

	partNumbers := make([][]bool, height)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			partNumbers[i] = append(partNumbers[i], false)
		}
	}

	for i, line := range inputLines {
		for j := 0; j < len(line); j++ {
			c := rune(line[j])
			if unicode.IsDigit(c) {
				for _, direction := range directions {
					newX := i + direction[0]
					newY := j + direction[1]
					if newX >= 0 && newX < width && newY >= 0 && newY < height && isSymbol(rune(inputLines[newX][newY])) {
						partNumbers[i][j] = true
					}
				}
			}
		}
	}

	partSum := 0
	for i, line := range inputLines {
		for j := 0; j < len(line); j++ {
			c := rune(line[j])
			if unicode.IsDigit(c) {
				isPart := false
				num := ""
				for ; j < len(line) && unicode.IsDigit(rune(line[j])); j++ {
					num += string(line[j])
					if partNumbers[i][j] {
						isPart = true
					}
				}
				if isPart {
					n, err := strconv.Atoi(num)
					if err != nil {
						log.Fatal(err)
					}
					partSum += n
				}
			}
		}
	}

	fmt.Println("partSum ", partSum)
}
