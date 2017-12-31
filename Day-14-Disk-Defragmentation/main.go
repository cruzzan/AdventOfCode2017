package main

import (
	"flag"
	"fmt"
	"strconv"
)

type cell struct {
	value int
	checked bool
}

type cord struct {
	i int
	j int
}

func main() {
	flag.Parse()
	input := flag.Arg(0)
	var matrix [128]string

	for i := 0; i < 128 ; i++ {
		hash := knotHash(input + "-" + strconv.Itoa(i))
		matrix[i] = stringToBin(hash)
	}

	fmt.Println("Task 1: ", task1(matrix), " Task 2: ", task2(matrix))
}

func check(err error)  {
	if err != nil {
		panic(err)
	}
}

func task1(matrix [128]string) (used int) {
	used = 0
	for _, row := range matrix {
		for _, c := range row {
			val, err := strconv.Atoi(string(c))
			check(err)
			used += val
		}
	}
	return
}

// https://en.wikipedia.org/wiki/Connected-component_labeling
// https://en.wikipedia.org/wiki/Depth-first_search

func task2(matrix [128]string) (groupCount int) {
	count := 0
	var groupMatrix [128][128]cell
	for i, row := range matrix {
		for j, c := range row {
			val, err := strconv.Atoi(string(c))
			check(err)
			var node cell
			node.value = val
			node.checked = false
			groupMatrix[i][j] = node
			if val == 1 {
				count++
			}
		}
	}

	groupCount = 0
	for i := 0; i < 128; i++ {
		for j := 0; j < 128; j++ {
			if groupMatrix[i][j].value == 1 && !groupMatrix[i][j].checked {
				groupCount += 1
				groupMatrix = visitGroup(groupMatrix, i, j)
			}
		}
	}

	return
}

func visitGroup(matrix [128][128]cell, i int, j int) [128][128]cell {
	matrix[i][j].checked = true
	neighbours := findNeighbours(matrix, i, j)

	for _, neighbour := range neighbours {
		matrix = visitGroup(matrix, neighbour.i, neighbour.j)
	}
	return matrix
}

func findNeighbours(matrix [128][128]cell, i int, j int) map[int]cord {
	var neighbours = make(map[int]cord)
	if i > 0 && matrix[i-1][j].value == 1 && !matrix[i-1][j].checked {
		neighbours[1] = cord{i-1, j}
	}
	if i < 127 && matrix[i+1][j].value == 1 && !matrix[i+1][j].checked {
		neighbours[2] = cord{i+1, j}
	}
	if j > 0 && matrix[i][j-1].value == 1 && !matrix[i][j-1].checked {
		neighbours[3] = cord{i, j-1}
	}
	if j < 127 && matrix[i][j+1].value == 1 && !matrix[i][j+1].checked {
		neighbours[4] = cord{i, j+1}
	}

	return neighbours
}

func stringToBin(s string) string {
	var bin string

	for i := 0; i < len(s); {
		val, _ := strconv.ParseInt(s[i:i+2], 16, 64)
		bin += fmt.Sprintf("%.8b", val)
		i += 2
	}
	return bin
}

func knotHash(input string) string {
	var cryptoList [256]int

	for i := 0; i < 256; i++{
		cryptoList[i] = i
	}

	pivots := parseInputToAsciiValueList(input)
	pivots = addStandardLegthSuffixes(pivots)

	var skip = 0
	var position = 0

	for i := 0; i < 64 ; i++ {
		cryptoList, skip, position = shuffleList(pivots, cryptoList, skip, position)
	}

	var hash string

	for i := 0; i < len(cryptoList); {
		var value int
		for _, item := range cryptoList[i:i+16] {
			value = value ^ item
			i++
		}
		hash += fmt.Sprintf("%02x", value)
	}

	return hash
}

func shuffleList(pivots []int, list [256]int, skip int, position int) (shuffledList [256]int, skipSize int, currentPosition int) {
	for _, pivotPoint := range pivots{
		if (position + pivotPoint) > len(list) {
			var end = (position + pivotPoint - 1) % len(list)
			list = flipRange(list, position, end)
		}else {
			var end = position + pivotPoint -1
			list = flipRange(list, position, end)
		}

		if position + pivotPoint + skip > len(list) {
			position = (position + pivotPoint + skip) % len(list)
		}else {
			position += pivotPoint + skip
		}

		skip++
	}

	return list, skip, position
}

func flipRange(list [256]int, start int, end int) [256]int {
	var i = start
	var temp []int

	for {
		if i == end +1 {
			break
		}

		if i >= len(list) {
			i = 0
		}

		temp = append([]int{list[i]}, temp...)

		i++
	}

	i = start
	for _, item := range temp{
		if i >= len(list) {
			i = 0
		}
		list[i] = item

		i++
	}

	return list
}

func parseInputToAsciiValueList(input string) []int {
	var valueList []int

	for _, char := range input{
		valueList = append(valueList, int(char))
	}

	return valueList
}

func addStandardLegthSuffixes(list []int) []int {
	list = append(list, 17)
	list = append(list, 31)
	list = append(list, 73)
	list = append(list, 47)
	list = append(list, 23)

	return list
}
