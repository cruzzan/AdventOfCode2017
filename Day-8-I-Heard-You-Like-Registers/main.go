package main

import (
	"os"
	"bufio"
	"regexp"
	"strconv"
	"fmt"
)

type condition struct {
	register string
	operator string
	value int
}


type instruction struct {
	register string
	action string
	value int
	condition condition
}

func main() {
	instructionSets := readInputFile()
	registers, instructions := parseInstructionSets(instructionSets)

	c1 := make(chan int)
	c2 := make(chan int)

	go task1(registers, instructions, c1)
	go task2(registers, instructions, c2)

	value1 := <-c1
	value2 := <-c2

	fmt.Println("Task 1: ", value1)
	fmt.Println("Task 2: ", value2)
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

func findValueInString(pattern string, haystack string) [][]string {
	regex := regexp.MustCompile(pattern)
	return regex.FindAllStringSubmatch(haystack, -1)
}

func parseInstructionSets(lines []string) (map[string]int, map[int]instruction) {
	var registers = make(map[string]int)
	var instructions = make(map[int]instruction)

	for i, line := range lines {
		register := findValueInString("^(\\w+)", line)[0][1]
		action := findValueInString("^\\w+ (\\w+)", line)[0][1]
		value, err := strconv.Atoi(findValueInString("^[\\s+[:alpha:]+]+(-*\\d*)", line)[0][1])
		check(err)
		condition := parseCondition(findValueInString("if\\s(.*)", line)[0][1])

		registers[register] = 0
		instructions[i] = instruction{
			register: register,
			action: action,
			value: value,
			condition: condition,
		}
	}

	return registers, instructions
}

func parseCondition(input string) condition {
	register := findValueInString("^(\\w+)", input)[0][1]
	operator := findValueInString("([<|>|=|!]+)", input)[0][1]
	value, err := strconv.Atoi(findValueInString("(-*\\d+)", input)[0][1])
	check(err)

	result := condition{
		register: register,
		operator: operator,
		value: value,
	}

	return result
}

func task1(r map[string]int, instructions map[int]instruction, c1 chan int) {
	registers := copyMap(r)

	for i := 0; i < len(instructions); {
		instruction := instructions[i]
		if conditionIsMet(registers, instruction.condition) {
			registers[instruction.register] = calc(
				registers[instruction.register],
				instruction.value,
				instruction.action,
				)
		}

		i++
	}

	var largest = 0
	for _, value := range registers {
		if value > largest {
			largest = value
		}
	}

	c1<- largest
}

func task2(r map[string]int, instructions map[int]instruction, c1 chan int) {
	registers := copyMap(r)
	var largest = 0

	for i := 0; i < len(instructions); {
		instruction := instructions[i]
		if conditionIsMet(registers, instruction.condition) {
			value := calc(
				registers[instruction.register],
				instruction.value,
				instruction.action,
			)

			registers[instruction.register] = value

			if value > largest {
				largest = value
			}
		}

		i++
	}

	c1<- largest
}

func calc(value1 int, value2 int, action string) int {
	switch action {
	case "inc":
		return value1 + value2
	case "dec":
		return value1 - value2
	default:
		panic("Action unrecognized " + action)
	}
}

func conditionIsMet(r map[string]int, c condition) bool {
	switch c.operator {
	case "<":
		return r[c.register] < c.value
	case "<=":
		return r[c.register] <= c.value
	case ">":
		return r[c.register] > c.value
	case ">=":
		return r[c.register] >= c.value
	case "==":
		return r[c.register] == c.value
	case "!=":
		return r[c.register] != c.value
	default:
		panic("Operator unrecognized " + c.operator)
	}
}

func copyMap(original map[string]int) map[string]int {
	newMap := make(map[string]int)

	for k, v := range original {
		newMap[k] = v
	}

	return newMap
}
