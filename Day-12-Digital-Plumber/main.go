package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"regexp"
)

type node struct {
	id int
	connections []int
}


func main() {
	input := readInputFile()

	fmt.Println("Task 1: ", task1(input), " Task 2: ", task2(input))
}

func task1(nodes map[int]node) int {
	var visited []int

	traversGroup(nodes, 0, &visited)

	return len(visited)
}

func task2(nodes map[int]node) int {
	var visited []int
	var groupCount= 0

	for len(visited) < len(nodes) {
		var notInGroup int
		for nodeId := range nodes {
			if !pointIsVisited(nodeId, &visited) {
				notInGroup = nodeId
				break
			}
		}
		traversGroup(nodes, notInGroup, &visited)
		groupCount++
	}

	return groupCount
}

func traversGroup(nodes map[int]node, start int, visited *[]int) {
	*visited = append(*visited, start)

	for _, connectionId := range nodes[start].connections {
		if !pointIsVisited(connectionId, visited) {
			traversGroup(nodes, connectionId, visited)
		}
	}
}

func pointIsVisited(id int, visited *[]int) bool {
	for _, point := range *visited {
		if id == point {
			return true
		}
	}

	return false
}


func check(err error)  {
	if err != nil {
		panic(err)
	}
}

func readInputFile() map[int]node {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	var input = make(map[int]node)

	for scanner.Scan() {
		line := scanner.Text()
		id, err :=strconv.Atoi(findAllSubstrings(line, "^(\\d+)")[0][1])
		var connections []int
		check(err)
		connectionsString := findAllSubstrings(line, "<->\\s(.*)")
		if len(connectionsString) > 0 {
			for _, connection := range findAllSubstrings(connectionsString[0][1], "(\\d+)") {
				connectionId, err := strconv.Atoi(connection[1])
				check(err)
				connections = append(connections, connectionId)
			}
		}

		input[id] = node{
			id: id,
			connections: connections,
		}

	}

	return input
}

func findAllSubstrings(haystack string, pattern string) [][]string {
	regex := regexp.MustCompile(pattern)
	return regex.FindAllStringSubmatch(haystack, -1)
}
