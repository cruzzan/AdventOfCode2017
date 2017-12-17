package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"time"
)

type layer struct {
	depth int
	ranges int
	scannerPosition int
	scannerDirection int
	scannerBlindCycleSize int
}

func main() {
	layers, deepestPoint := readInputFile()

	timerStart := time.Now()
	task1 := task1(layers, deepestPoint)
	task1Duration := time.Since(timerStart)

	timerStart = time.Now()
	task2 := task2(layers)
	task2Duration := time.Since(timerStart)

	fmt.Println("Task 1: ", task1, " in: ", task1Duration)
	fmt.Println("Task 2: ", task2, " in: ", task2Duration)
}
func check(err error)  {
	if err != nil {
		panic(err)
	}
}
func readInputFile() (layers map[int]layer, deepestPoint int) {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	var layerMap = make(map[int]layer)
	var deepest = 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": ")

		depth, err := strconv.Atoi(line[0])
		check(err)
		ranges, err := strconv.Atoi(line[1])
		check(err)

		layerMap[depth] = layer{
			depth: depth,
			ranges: ranges,
			scannerPosition: 0,
			scannerDirection: 1,
			scannerBlindCycleSize: (2 * ranges) - 2,
		}

		if depth > deepest {
			deepest = depth
		}
	}

	return layerMap, deepest
}

func task1(layersMap map[int]layer, deepestPoint int) int {
	layers := cloneMap(layersMap)
	var severity = 0

	for depth := 0; depth <= deepestPoint; {
		severity += calculateSeverity(layers[depth])

		layers = moveScanners(layers)

		depth++
	}

	return severity
}

func task2(layersMap map[int]layer) int {
	delay := 0
	layers := cloneMap(layersMap)

	for {
		var caught = false
		for key, layer := range layers {
			currentTick := delay + key
			if currentTick % layer.scannerBlindCycleSize == 0 {
				caught = true
				break
			}
		}

		if !caught {
			break
		}

		delay++
	}

	return delay
}

func cloneMap(originalMap map[int]layer) map[int]layer {
	var newMap = make(map[int]layer)
	for key, value := range originalMap{
		newMap[key] = value
	}

	return newMap
}

func calculateSeverity(l layer) int {
	if l.scannerPosition == 0 && l.ranges > 0 {
		return l.depth * l.ranges
	}

	return 0
}

func moveScanners(layers map[int]layer) map[int]layer {
	for i, layer := range layers {
		layer.scannerPosition += layer.scannerDirection

		if layer.scannerPosition <= 0 && layer.scannerDirection < 0 {
			layer.scannerDirection = 1
		}else if layer.scannerPosition >= layer.ranges -1 && layer.scannerDirection > 0 {
			layer.scannerDirection = -1
		}

		layers[i] = layer
	}

	return layers
}
