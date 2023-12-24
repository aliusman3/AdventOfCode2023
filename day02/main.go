package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/aliusman3/aoc/util"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Cube struct {
	amount int
	color  string
}

func extractCubes(cubesString string) []*Cube {
	cubeSets := []string{}
	for _, cubeSet := range strings.Split(cubesString, ";") {
		cubeSets = append(cubeSets, strings.Trim(cubeSet, " "))
	}

	cubes := []*Cube{}
	for _, cubeSet := range cubeSets {
		for _, cube := range strings.Split(cubeSet, ",") {
			cubeConfig := strings.Split(strings.Trim(cube, " "), " ")
			cubeAmount, err := strconv.Atoi(cubeConfig[0])
			if err != nil {
				log.Fatal(err)
			}
			cubeColor := cubeConfig[1]
			if _, found := maxCubes[cubeColor]; !found {
				log.Fatalf("Could not find match for cube color '%s'", cubeColor)
			}
			cubes = append(cubes, &Cube{
				amount: cubeAmount,
				color:  cubeColor,
			})
		}
	}
	return cubes
}

var maxCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	inputLines := util.ReadInput("input")

	output := util.CreateOutput()
	defer output.Close()
	logOutput := util.GetLogger(output)

	logOutput("============ PART 1 ============\n")
	part1(inputLines, logOutput)
	logOutput("\n============ PART 2 ============\n")
	part2(inputLines, logOutput)
}

func part1(inputLines []string, logOutput func(string, ...any)) {
	idSum := 0

	for i, line := range inputLines {
		lineNumber := i + 1
		splitLine := strings.Split(line, ": ")
		gameId, err := strconv.Atoi(strings.Split(splitLine[0], " ")[1])
		if err != nil {
			log.Fatal(err)
		}
		logOutput("Line %d ", lineNumber)
		logOutput(" Game ID: %d ", gameId)

		impossibleConfig := false
		cubes := extractCubes(splitLine[1])
		for _, cube := range cubes {
			if cube.amount > maxCubes[cube.color] {
				logOutput(" game impossible because cube %s exceeded max allowed cubes %d ", cube.color, maxCubes[cube.color])
				impossibleConfig = true
			}
		}

		if !impossibleConfig {
			logOutput(" game possible ")
			idSum += gameId
		}

		logOutput("\n")
	}

	fmt.Println("Id Sum", idSum)
}

func part2(inputLines []string, logOutput func(string, ...any)) {
	powerSum := 0
	for i, line := range inputLines {
		lineNumber := i + 1
		splitLine := strings.Split(line, ": ")
		gameId, err := strconv.Atoi(strings.Split(splitLine[0], " ")[1])
		if err != nil {
			log.Fatal(err)
		}
		logOutput("Line %d ", lineNumber)
		logOutput(" Game ID: %d ", gameId)

		maxAmounts := map[string]int{}
		cubes := extractCubes(splitLine[1])
		for _, cube := range cubes {
			maxAmounts[cube.color] = max(maxAmounts[cube.color], cube.amount)
		}

		cubePower := 1
		for _, v := range maxAmounts {
			cubePower *= v
		}
		logOutput(" cube power %d ", cubePower)
		powerSum += cubePower

		logOutput("\n")
	}

	fmt.Println("Power Sum", powerSum)
}
