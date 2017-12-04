package main

import (
"flag"
"strings"
"fmt"
"strconv"
	"math"
)

func main()  {
	flag.Parse()
	task := flag.Arg(0)
	input, err := strconv.Atoi(flag.Arg(1))
	check(err)
	var result int

	switch strings.ToLower(task) {
	case "task1":
		result = task1(input)
		break
	case "task2":
		result = task2(input)
	}

	fmt.Println(result)
}

func check(err error)  {
	if err != nil {
		panic(err)
	}
}

func task1(position int) int {
	var x = 0
	var y = 0
	var turns = 0

	for item := 1; item <= position; { // Keep making sides while the number is not reached
		var sideLength = (turns / 2) + 1

		for i := 0; i < sideLength; i++ {
			if item == position { // In case the number is in this line
				item++
				break
			}

			var direction = turns % 4
			if direction == 0 {
				x += 1 // East
			}else if direction == 1 {
				y += 1 // North
			}else if direction == 2 {
				x -= 1 // West
			}else {
				y -= 1 // South
			}

			item++
		}

		turns++
	}

	fmt.Println("x: ", x, " y: ", y)

	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func task2(input int) int {
	var x= 0
	var y= 0
	var turns= 0
	passed := make(map[string]int)
	var top = 0

	for top < input { // Keep making sides while the number is not reached
		var sideLength = (turns / 2) + 1

		for i := 0; i < sideLength; i++ {
			var direction = turns % 4
			top = itemValue(x, y, passed)

			passed[strconv.Itoa(x)+","+strconv.Itoa(y)] = top

			if direction == 0 {
				x += 1 // East
			} else if direction == 1 {
				y += 1 // North
			} else if direction == 2 {
				x -= 1 // West
			} else {
				y -= 1 // South
			}

			if top >= input { // In case the number reached in side
				break
			}
		}

		turns++
	}

	return top

}

func itemValue(x, y int, passed map[string]int) int {
	var value = 0

	value += passed[strconv.Itoa(x+1)+","+strconv.Itoa(y)]
	value += passed[strconv.Itoa(x+1)+","+strconv.Itoa(y+1)]
	value += passed[strconv.Itoa(x)+","+strconv.Itoa(y+1)]
	value += passed[strconv.Itoa(x-1)+","+strconv.Itoa(y+1)]
	value += passed[strconv.Itoa(x-1)+","+strconv.Itoa(y)]
	value += passed[strconv.Itoa(x-1)+","+strconv.Itoa(y-1)]
	value += passed[strconv.Itoa(x)+","+strconv.Itoa(y-1)]
	value += passed[strconv.Itoa(x+1)+","+strconv.Itoa(y-1)]

	if value == 0 {
		value = 1
	}

	return value
}
