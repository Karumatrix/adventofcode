package main

import (
	"bufio"
	"fmt"
	"os"
)

func filepathToString(filepath string) string {
    content, err := os.ReadFile(filepath)
    if err != nil {
		panic(err)
    }
    return string(content)
}

func filepathToStringArray(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

func countVowels(line string) bool {
	nbVowels := 0
	vowels := "aeiou"

	for _, letter := range line {
		for _, vowel := range vowels {
			if vowel == letter {
				nbVowels++
			}
		}
	}
	if nbVowels >= 3 {
		return true
	}
	return false
}

func countDoubleLetters(line string) bool {
	for i := 0; i < len(line) - 1; i++ {
		if line[i] == line[i + 1] {
			return true
		}
	}
	return false
}

func checkSubstrings(line string) bool {
	stringsToCheck := []string{
		"ab",
		"cd",
		"pq",
		"xy",
	}
	for _, str := range stringsToCheck {
		for i := 0; i < len(line) - 1; i++ {
			if line[i] == str[0] && line[i+1] == str[1] {
				return true
			}
		}
	}
	return false
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	result := 0

	for _, line := range file {
		if countVowels(line) && countDoubleLetters(line) && !checkSubstrings(line) {
			result++
		}
	}
	fmt.Println(result)
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("USAGE: go run <file> <input-file>")
	} else {
		aoc(os.Args[1])
	}
}
