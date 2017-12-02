package main

import (
	"flag"
	"strings"
	"fmt"
	"os"
	"bufio"
	"strconv"
	"math"
)

func main()  {
	flag.Parse()
	task := flag.Arg(0)
	rows := readInputFile()
	var result int

	switch strings.ToLower(task) {
	case "task1":
		result = task1(rows)
		break
	case "task2":
		result = task2(rows)
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

func task1(rows []string) int {
	var diff int
	diff = 0
	for _, row := range rows {
		items := strings.Split(row, "\t")

		highest, err := strconv.Atoi(items[0])
		check(err)
		lowest, err := strconv.Atoi(items[0])
		check(err)

		for _, item := range items{
			number, err := strconv.Atoi(item)
			check(err)

			if number > highest {
				highest = number
			}

			if number < lowest {
				lowest = number
			}
		}
		diff += highest - lowest
	}
	return diff
}

func task2(rows []string) int {
	var diff int
	diff = 0

	for i, row := range rows {
		items := strings.Split(row, "\t")
		var numerator, denominator int

		for _, item1 := range items{
			number1, err := strconv.Atoi(item1)
			check(err)

			for _, item2 := range items{
				number2, err := strconv.Atoi(item2)
				check(err)

				if number1 > number2 {
					if isDivisible(number1, number2) {
						numerator = number1
						denominator = number2
						break
					}
				}else if number1 < number2 {
					if isDivisible(number2, number1) {
						numerator = number2
						denominator = number1
						break
					}
				}
			}
		}
		fmt.Println(i, " numerator: ", numerator, " denominator: ", denominator)
		diff += numerator / denominator
	}
	return diff
}

func isDivisible(numerator, denominator int) bool {
	return math.Mod(float64(numerator), float64(denominator)) == 0
}
