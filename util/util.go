package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadInput(filePath string) []string {
	input, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	inputArray := []string{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		inputArray = append(inputArray, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return inputArray
}

func CreateLogFile() *os.File {
	output, err := os.Create("output.log")
	if err != nil {
		log.Fatal(err)
	}
	return output
}

func GetLogger(w *os.File) func(s string, ags ...any) {
	return func(s string, ags ...any) {
		fmt.Fprintf(w, s, ags...)
	}
}
