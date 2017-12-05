package main

import (
	"flag"
	"strings"
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main()  {
	flag.Parse()
	task := flag.Arg(0)
	lines := readInputFile()
	var result int

	switch strings.ToLower(task) {
	case "task1":
		result = task1(stringListToIntList(lines))
		break
	case "task2":
		result = task2(stringListToIntList(lines))
	}

	fmt.Println(result)
}

func check(err error)  {
	if err != nil {
		panic(err)
	}
}

func readInputFile() []string {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func stringListToIntList(input []string) []int {
	var intList = make([]int, len(input))
	for i, value := range input {
		intValue, err := strconv.Atoi(value)
		check(err)

		intList[i] = intValue
	}

	return intList
}

func task1(input []int) int {
	var jumps = 0
	var exitFound = false
	var cursor = 0

	for exitFound == false {
		if cursor > len(input) -1 {
			exitFound = true
			break
		}else {
			currentValue := input[cursor]
			input[cursor] += 1
			cursor = cursor + currentValue
		}
		jumps += 1
	}

	return jumps
}

func task2(input []int) int {
	var jumps = 0
	var exitFound = false
	var cursor = 0

	for exitFound == false {
		if cursor > len(input) -1 {
			exitFound = true
			break
		}else {
			offsetChange := 1
			currentValue := input[cursor]

			if currentValue >= 3 {
				offsetChange = -1
			}

			input[cursor] += offsetChange
			cursor = cursor + currentValue
		}
		jumps += 1
	}

	return jumps
}
