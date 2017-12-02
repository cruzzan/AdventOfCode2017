package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
	"flag"
)

func main() {
	flag.Parse()
	task := flag.Arg(0)
	input := flag.Arg(1)
	brokenDown := strings.Split(input, "")
	var result int

	switch strings.ToLower(task) {
	case "task1":
		result = task1(brokenDown)
		break
	case "task2":
		result = task2(brokenDown)
	}
	
	fmt.Println(result)
}

func task1(input []string) int  {
	var sum = 0

	for i, char := range input {
		var next = ""

		if i != len(input)-1 {
			next = input[i+1]
		}else{
			next = input[0]
		}

		if char == next {
			number, err := strconv.Atoi(char)

			if err != nil {
				fmt.Println("Something went wrong when converting to int")
				os.Exit(1)
			}

			sum += number
		}
	}

	return sum
}

func task2(input []string) int  {
	var sum = 0

	shift := len(input)/2
	length := len(input)

	for i, char := range input {
		var next = ""

		if shift + i <= length - 1 {
			next = input[i + shift]
		}else{
			point := (i + shift) - length
			next = input[point]
		}

		if char == next {
			number, err := strconv.Atoi(char)

			if err != nil {
				fmt.Println("Something went wrong when converting to int")
				os.Exit(1)
			}

			sum += number
		}
	}

	return sum
}
