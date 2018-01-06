package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"strconv"
)

type command struct {
	name string
	positions []string
}

func main() {
	var commands = readInputFile()
	var dancers = createDancers()

	fmt.Println("Task 1:")
	printDancers(preformDance(commands, dancers))
	fmt.Println()
	fmt.Println("Task 2: ")
	printDancers(massDancePerformance(commands, dancers, 100))
}

func massDancePerformance(commands []command, dancers [16]string, nrDances int) [16]string {
	var cycles = make(map[string]string)
	var runDances = 0
	for {
		implodedDancers := implodeDancers(dancers)
		if cycles[implodedDancers] != "" {
			dancers = preformDance(commands, dancers)
			break
		}
		cycles[implodedDancers] = implodedDancers
		dancers = preformDance(commands, dancers)
		runDances++
	}
	var offset = nrDances % runDances
	for i := 0; i < offset - 1; i++ {
		dancers = preformDance(commands, dancers)
	}
	return dancers
}

func preformDance(commands []command, dancers [16]string) [16]string {
	for _, command := range commands {
		dancers = doALittleDance(command, dancers)
	}
	return dancers
}

func check(err error)  {
	if err != nil {
		panic(err)
	}
}

func readInputFile() []command {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)

	var commands []command

	for scanner.Scan() {
		rawCommands := strings.Split(scanner.Text(), ",")
		for _, rawCommand := range rawCommands {
			var cmd command
			switch string(rawCommand[:1]) {
			case "s":
				cmd.name = "Split"
				break
			case "x":
				cmd.name = "Exchange"
				break
			case "p":
				cmd.name = "Partner"
				break
			}

			points := strings.Split(string(rawCommand[1:]), "/")

			if len(points) > 1 {
				cmd.positions = append(cmd.positions, points[0], points[1])
			} else {
				cmd.positions = append(cmd.positions, points[0])
			}

			commands = append(commands, cmd)
		}
	}

	return commands
}

func createDancers() [16]string {
	var dancers [16]string
	for i := 0; i < 16; i++ {
		dancers[i] = string(i + 97)
	}

	return dancers
}

func doALittleDance(command command, dancers [16]string) [16]string {
	switch command.name {
	case "Split":
		return split(dancers, command.positions[0])
	case "Exchange":
		return exchange(dancers, command.positions[0], command.positions[1])
	case "Partner":
		return partner(dancers, command.positions[0], command.positions[1])
	default:
		return dancers
	}
}

func split(dancers [16]string, piv string) [16]string {
	var result [16]string
	pivot, err := strconv.Atoi(piv)
	check(err)

	var key = 0
	for _, dancer := range dancers[16-pivot:] {
		result[key] = dancer
		key++
	}
	for _, dancer := range dancers[0:16-pivot] {
		result[key] = dancer
		key++
	}

	return result
}

func exchange(dancers [16]string, pos1 string, pos2 string) [16]string {
	dancer1, err := strconv.Atoi(pos1)
	check(err)
	dancer2, err := strconv.Atoi(pos2)
	check(err)

	temp := dancers[dancer1]
	dancers[dancer1] = dancers[dancer2]
	dancers[dancer2] = temp

	return dancers
}

func partner(dancers [16]string, name1 string, name2 string) [16]string {
	var dancer1 = -1
	var dancer2 = -1
	for key, dancer := range dancers {
		if dancer == name1 {
			dancer1 = key
		}

		if dancer == name2 {
			dancer2 = key
		}

		if dancer1 != -1 && dancer2 != -1 {
			break
		}
	}

	temp := dancers[dancer1]
	dancers[dancer1] = dancers[dancer2]
	dancers[dancer2] = temp

	return dancers
}

func printDancers(dancers [16]string) {
	for _, dancer := range dancers{
		fmt.Print(dancer)
	}
	fmt.Println()
}

func implodeDancers(dancers [16]string) string {
	var result string
	for _, dancer := range dancers{
		result += dancer
	}

	return result
}
