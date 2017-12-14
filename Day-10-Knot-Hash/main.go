package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"math"
)

func main() {
	var input = readInputFile()
	var cryptoList [256]int

	for i := 0; i < 256; i++{
		cryptoList[i] = i
	}

	fmt.Println("Task 1: ", task1(cryptoList, input), " Task 2: ", task2(cryptoList, input))
}

func task1(list [256]int, input string) int {
	var pivots = parseNumberString(input)
	list, _, _ = shuffleList(pivots, list, 0, 0)

	return list[0] * list[1]
}

func task2(list [256]int, input string) string {
	pivots := parseInputToAsciiValueList(input)
	pivots = addStandardLegthSuffixes(pivots)

	var skip = 0
	var position = 0

	for i := 0; i < 64 ; i++ {
		list, skip, position = shuffleList(pivots, list, skip, position)
	}

	var denseHash [16]int
	var hash string

	for i := 0; i < len(list); {
		var value int
		for _, item := range list[i:i+16] {
			value = value ^ item
			i++
		}
		index := int(math.Floor(float64(i-1)/16))
		denseHash[index] = value
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

func check(err error)  {
	if err != nil {
		panic(err)
	}
}

func readInputFile() string {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	var input string

	for scanner.Scan() {
		input += scanner.Text()
	}

	return input
}

func parseNumberString(input string) []int {
	var numbers []int
	for _, value := range strings.Split(input, ",") {
		numeric, err := strconv.Atoi(value)
		check(err)
		numbers = append(numbers, numeric)
	}

	return numbers
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
