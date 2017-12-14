package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"math"
)

func main() {
	input := readInputFile()

	fmt.Println("Task 1: ", task1(input), " Task 2: ", task2(input))
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
	var steps []string

	for scanner.Scan() {
		steps = strings.Split(scanner.Text(), ",")
	}

	return steps
}

func task1(steps []string) int {
	cords := make(map[string]int)
	cords["x"] = 0
	cords["y"] = 0
	cords["z"] = 0

	for _, step := range steps {
		cords = makeStep(step, cords)
	}

	x := int(math.Abs(float64(cords["x"])))
	y := int(math.Abs(float64(cords["y"])))
	z := int(math.Abs(float64(cords["z"])))

	return cubeManhattanDistance(x, y, z)
}

func task2(steps []string) int {
	var farthest = 0
	cords := make(map[string]int)
	cords["x"] = 0
	cords["y"] = 0
	cords["z"] = 0

	for _, step := range steps {
		cords = makeStep(step, cords)
		x := int(math.Abs(float64(cords["x"])))
		y := int(math.Abs(float64(cords["y"])))
		z := int(math.Abs(float64(cords["z"])))

		currentDistance := cubeManhattanDistance(x, y, z)

		if currentDistance > farthest {
			farthest = currentDistance
		}
	}

	return farthest
}

func makeStep(direction string, currentCords map[string]int) map[string]int {
	switch direction {
	case "n":
		currentCords["y"] += 1
		currentCords["z"] -= 1
		break
	case "ne":
		currentCords["x"] += 1
		currentCords["z"] -= 1
		break
	case "se":
		currentCords["y"] -= 1
		currentCords["x"] += 1
		break
	case "s":
		currentCords["y"] -= 1
		currentCords["z"] += 1
		break
	case "sw":
		currentCords["x"] -= 1
		currentCords["z"] += 1
		break
	case "nw":
		currentCords["y"] += 1
		currentCords["x"] -= 1
		break
	}
	
	
	return currentCords
}

func cubeManhattanDistance(x, y, z int) int {
	return (x + y + z) / 2
}
