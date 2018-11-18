package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

func main()  {
	start := time.Now()

	t1 := task1()
	t1Duration := time.Since(start)

	fmt.Printf("Task 1, the frequenzy was: %d, Duration: %s\n", t1, t1Duration)
}

func task1() int {
	il := readInput("input.txt")
	registry := createRegistryFromInstructionsList(il)
	lastFreq := 0
	done := false

	for i := 0; i <= len(il) && !done; {
		cmd := parseInstructionFromString(il[i])

		switch cmd[0] {
		case "snd":
			lastFreq = registry[cmd[1]]
			fmt.Printf("Emitted frequency %d\n", lastFreq)
			break
		case "set":
			registry[cmd[1]] = getValue(registry, cmd[2])

			break
		case "add":
			registry[cmd[1]] += getValue(registry, cmd[2])

			break
		case "mul":
			registry[cmd[1]] *= getValue(registry, cmd[2])

			break
		case "mod":
			registry[cmd[1]] %= getValue(registry, cmd[2])

			break
		case "rcv":
			if getValue(registry, cmd[1]) != 0 {
				registry[cmd[1]] = lastFreq
				done = true
			}
			break
		case "jgz":
			if getValue(registry, cmd[1]) > 0 {
				i += getValue(registry, cmd[2])
				continue
			}
			break
		}

		i++
	}

	return lastFreq
}

func createRegistryFromInstructionsList(il []string) map[string]int {
	registry := make(map[string]int)
	for _, i := range il {
		si := strings.Split(i, " ")

		if isAlpha(si[1]) {
			registry[si[1]] = 0
		}
	}

	return registry
}

func parseInstructionFromString(is string) []string {
	return strings.Split(is, " ")
}

func isAlpha(s string) bool {
	regRune, _ := utf8.DecodeRuneInString(s)

	if unicode.IsLetter(regRune) {
		return true
	}

	return false
}

func getValue(reg map[string]int, input string) int {
	if isAlpha(input) {
		return reg[input]
	} else {
		val, err := strconv.Atoi(input)
		check(err)
		return val
	}
}

func readInput(f string) [] string {
	var input []string
	file, err := os.Open(f)
	check(err)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
