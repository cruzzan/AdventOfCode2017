package main

import (
	"fmt"
	"flag"
	"strconv"
	"time"
)

func main() {
	flag.Parse()
	offset, err := strconv.Atoi(flag.Arg(0))
	check(err)

	task1TimeStart := time.Now()
	task1result := task1(offset)
	fmt.Println("Task 1:", task1result, "took:", execTime(task1TimeStart))

	task2TimeStart := time.Now()
	task2result := task2(offset)
	fmt.Println("Task 2:", task2result, "took:", execTime(task2TimeStart))

}

func check(err error)  {
	if err != nil {
		panic(err)
	}
}

func execTime(start time.Time) string {
	return time.Now().Sub(start).String()
}

func task1(offset int) int {
	var buffer = []int{0}
	var currentPosition = 0

	for i := 0; i < 2017; i++ {
		buffer, currentPosition = whirl(buffer, currentPosition, offset)
	}

	return buffer[currentPosition+1]
}

func task2(offset int) int {
	var currentPosition = 0
	var posOne = 0

	for i := 0; i < 50000000; i++ {
		value := i + 1
		currentPosition = ((currentPosition + offset) % value) +1
		if currentPosition == 1 {
			posOne = value
		}
	}

	return posOne
}

func whirl(buffer []int, currentPosition int, offset int) (buff []int, key int) {
	key = ((currentPosition + offset) % len(buffer)) +1

	beginning := buffer[0:key]
	var rest []int
	if key < len(buffer) {
		rest = buffer[key:]
	}

	buff = append(buff, beginning...)
	buff = append(buff, len(buffer))
	buff = append(buff, rest...)

	if key >= len(buff) {
		key = 0
	}

	return
}
