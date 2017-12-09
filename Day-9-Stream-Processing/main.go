package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
)

func main() {
	input := readInputFile()

	chan1 := make(chan int)
	chan2 := make(chan int)

	go task1(input, chan1)
	go task2(input, chan2)

	result1 := <-chan1
	result2 := <-chan2

	fmt.Println("Task 1: ", result1, " Task 2: ", result2)
}

func task1(rawInput string, c chan<- int) {
	input := stripCancelled(rawInput)
	input = stripGarbage(input)
	input = stripCommas(input)

	c<- countPointsInGroup(input)
}

func task2(rawInput string, c chan<- int)  {
	var input = stripCancelled(rawInput)
	var garbageCleaned = 0
	regex, err := regexp.Compile("\\<(.*?)\\>")
	check(err)

	garbageCollections := regex.FindAllStringSubmatch(input, -1)

	for _, g := range garbageCollections {
		garbageCleaned += len(g[1])
	}

	c<- garbageCleaned
}

func check(err error)  {
	if err != nil {
		panic(err)
	}
}

func readInputFile() string {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	var text string

	for scanner.Scan() {
		text += scanner.Text()
	}

	return text
}

func stripCancelled(input string) string {
	regex, err := regexp.Compile("\\!.")
	check(err)

	return regex.ReplaceAllString(input, "")
}

func stripGarbage(input string) string {
	regex, err := regexp.Compile("\\<.*?\\>")
	check(err)

	return regex.ReplaceAllString(input, "")
}

func stripCommas(input string) string {
	regex, err := regexp.Compile(",")
	check(err)

	return regex.ReplaceAllString(input, "")
}

func countPointsInGroup(input string) int {
	var heap = 0
	var points = 0
	for _, char := range input {
		if string(char) == "{" {
			heap++
		}else {
			points += heap
			heap--
		}
	}

	return points
}
