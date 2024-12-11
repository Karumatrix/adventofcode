package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func concatNumbers(nb1, nb2 int) (result int) {
	nbStr1, nbStr2 := strconv.Itoa(nb1), strconv.Itoa(nb2)

	result, _ = strconv.Atoi(nbStr1 + nbStr2)
	return result
}

func checkOperations(values []int, resultWanted, currValue, index int) bool {
	if index == len(values) {
		return currValue == resultWanted
	}
	if checkOperations(values, resultWanted, currValue + values[index], index + 1) {
		return true
	}
	if checkOperations(values, resultWanted, currValue * values[index], index + 1) {
		return true
	}
	if checkOperations(values, resultWanted, concatNumbers(currValue, values[index]), index + 1) {
		return true
	}
	return false
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	result := 0

	for _, line := range file {
		parts := strings.Split(line, ":")
		resultWanted := 0
		tmpValue := 0
		fmt.Sscanf(parts[0], "%d", &resultWanted)
		var values []int
		subParts := strings.Split(parts[1], " ")
		for _, part := range subParts {
			fmt.Sscanf(part, "%d", &tmpValue)
			values = append(values, tmpValue)
		}
		if checkOperations(values, resultWanted, values[0], 1) {
			result += resultWanted
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
