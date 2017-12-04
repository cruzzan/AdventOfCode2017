package main

import (
	"flag"
	"strings"
	"fmt"
	"os"
	"bufio"
	"reflect"
)

func main()  {
	flag.Parse()
	task := flag.Arg(0)
	lines := readInputFile()
	var result int

	switch strings.ToLower(task) {
	case "task1":
		result = task1(lines)
		break
	case "task2":
		result = task2(lines)
	}

	fmt.Println(result)
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
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func task1(passphrases []string) int {
	var validCount = 0

	for _, passphrase := range passphrases {
		var isValid = true
		words := strings.Split(passphrase, " ")

		for i, word := range words {
			var foundAt = sliceContains(words, word, false)
			if foundAt != -1 && foundAt != i {
				isValid = false
			}
		}

		if isValid {
			validCount++
		}
	}

	return validCount
}

func task2(passphrases []string) int {
	var validCount = 0

	for _, passphrase := range passphrases {
		var isValid = true
		words := strings.Split(passphrase, " ")

		for i, word := range words {
			var foundAt = sliceContains(words, word, true)
			if foundAt != -1 && foundAt != i {
				isValid = false
			}
		}

		if isValid {
			validCount++
		}
	}

	return validCount
}

func sliceContains(slice []string, needle string, checkAnagram bool) int {
	for i, item := range slice{
		var hit bool

		if item == needle {
			hit = true
		}

		if checkAnagram {
			if isAnagram(item, needle) {
				hit = true
			}
		}

		if hit {
			return i
		}
	}

	return -1
}

func isAnagram(word1, word2 string) bool {
	wordMap1 := mapWord(word1)
	wordMap2 := mapWord(word2)

	return reflect.DeepEqual(wordMap1, wordMap2)
}

func mapWord(word string) map[string]int {
	var wordMap =  make(map[string]int)

	for _, char := range word{
		wordMap[string(char)] += 1
	}

	return wordMap
}
