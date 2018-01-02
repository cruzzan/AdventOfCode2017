package main

import (
	"flag"
	"fmt"
	"strconv"
)

type generator struct {
	base int
	factor int
	criteria int
}

func (g *generator) generateNextValue(mod int) int {
	g.base = (g.base * g.factor) % mod
	return g.base
}

func (g *generator) generateList(size, mod int, withCriteria bool) []int {
	var list []int
	for len(list) < size {
		val := g.generateNextValue(mod)
		if withCriteria {
			if val % g.criteria == 0  {
				list = append(list, val)
			}
		} else {
			list = append(list, val)
		}
	}
	return list
}

func main() {
	var genA, genB generator
	const mod = 2147483647

	flag.Parse()

	genABase, err := strconv.Atoi(flag.Arg(0))
	check(err)
	genA.base = genABase
	genA.factor = 16807
	genA.criteria = 4

	genBBase, err := strconv.Atoi(flag.Arg(1))
	check(err)
	genB.base = genBBase
	genB.factor = 48271
	genB.criteria = 8

	fmt.Println("Task 1: ", task1(genA, genB, mod), " Task 2: ", task2(genA, genB, mod))
}

func check(err error)  {
	if err != nil {
		panic(err)
	}
}

func task1(genA generator, genB generator, mod int) int {
	var matchCount = 0
	var listA = genA.generateList(40000000, mod, false)
	var listB = genB.generateList(40000000, mod, false)

	for i := 0; i < len(listA); i++ {
		if compareLastTwoBytes(listA[i], listB[i]) {
			matchCount++
		}
	}

	return matchCount
}

func task2(genA generator, genB generator, mod int) int {
	var matchCount = 0
	var listA = genA.generateList(5000000, mod, true)
	var listB = genB.generateList(5000000, mod, true)

	for i := 0; i < len(listA); i++ {
		if compareLastTwoBytes(listA[i], listB[i]) {
			matchCount++
		}
	}

	return matchCount
}

func compareLastTwoBytes(val1, val2 int) bool {
	// Check the results against eachothers bitwise comparisons with `1111111111111111`
	if val1 & 65535 == val2 & 65535 {
		return true
	}
	return false
}
