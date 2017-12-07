package main

import (
	"os"
	"bufio"
	"strings"
	"strconv"
)

type bankHistory struct {
	snapshot []string
}

func (bh *bankHistory) addSnapshot(snapshot string) {
	bh.snapshot = append(bh.snapshot, snapshot)
}

func (bh bankHistory) snapshotExists(snapshot string) bool {
	for _, snap := range bh.snapshot {
		if snapshot == snap {
			return true
		}
	}

	return false
}


func main()  {
	input := readFromFile()

	if len(input) != 16 {
		println("Error: expected 16 memory banks, not ", len(input))
		os.Exit(2)
	}

	task1Result, task2Result := tasks(input)

	println("Task 1: ", task1Result, " Task 2: ", task2Result)

	os.Exit(0)
}

func checkError(err error)  {
	if err != nil {
		panic(err)
	}
}

func readFromFile() []int {
	var values []int
	file, err := os.Open("input.txt")
	checkError(err)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		for _, number := range strings.Split(line, "\t") {
			number, err := strconv.Atoi(number)
			checkError(err)
			values = append(values, number)
		}
	}

	return values
}

func tasks(input []int) (task1Result int, task2Result int){

	return task1(input), task2(input)
}

func task1(input []int) int {
	var infiniteLoopDetected = false
	var history bankHistory
	var redistributions = 0

	for !infiniteLoopDetected {
		bankSnapshot := banksToString(input)

		if history.snapshotExists(bankSnapshot) {
			infiniteLoopDetected = true
			break
		}

		history.addSnapshot(bankSnapshot)

		biggestBankIndex := findBiggestBank(input)

		input = redistributeBank(input, biggestBankIndex)

		redistributions += 1
	}

	return redistributions
}

func task2(input []int) int {
	var infiniteLoopDetectedCount = 0
	var history bankHistory
	var redistributions = 0

	for infiniteLoopDetectedCount < 2 {
		bankSnapshot := banksToString(input)

		if history.snapshotExists(bankSnapshot) {
			infiniteLoopDetectedCount += 1
			continue
		}

		history.addSnapshot(bankSnapshot)

		biggestBankIndex := findBiggestBank(input)

		input = redistributeBank(input, biggestBankIndex)

		redistributions += 1
	}

	return redistributions
}

func findBiggestBank(banks []int) int {
	var index = 0

	for i, bank := range banks {
		if bank > banks[index] {
			index = i
		}
	}

	return index
}

func redistributeBank(banks []int, index int) []int {
	var done = false
	for i := index + 1; !done; {
		if i == len(banks) {
			i = 0
			continue
		}

		if i == index {
			done = true
			break
		}

		if banks[index] > 0 {
			banks[i] += 1
			banks[index] -= 1
		}

		i++
	}

	return banks
}

func banksToString(banks []int) string {
	var banksState = ""

	for _, bankSize := range banks {
		banksState = banksState + " " + strconv.Itoa(bankSize)
	}

	return banksState
}
