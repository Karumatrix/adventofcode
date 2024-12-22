package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

var (
	numericKeypad = map[string][2]int{
		"7": {0, 0},
		"8": {0, 1},
		"9": {0, 2},
		"4": {1, 0},
		"5": {1, 1},
		"6": {1, 2},
		"1": {2, 0},
		"2": {2, 1},
		"3": {2, 2},
		"0": {3, 1},
		"A": {3, 2},
	}
	directionalKeypad = map[string][2]int{
		"^": {0, 1},
		"A": {0, 2},
		"<": {1, 0},
		"v": {1, 1},
		">": {1, 2},
	}

	DIRECTIONS = map[string][2]int{
		"v": {1, 0},
		"^": {-1, 0},
		">": {0, 1},
		"<": {0, -1},
	}
)

func getPermutations(s string) []string {
	uniquePerms := make(map[string]bool)
	var result []string
	permute([]rune(s), 0, uniquePerms)

	for perm := range uniquePerms {
		result = append(result, perm)
	}
	return result
}

func permute(runes []rune, start int, uniquePerms map[string]bool) {
	if start == len(runes)-1 {
		uniquePerms[string(runes)] = true
		return
	}

	for i := start; i < len(runes); i++ {
		runes[start], runes[i] = runes[i], runes[start]
		permute(runes, start+1, uniquePerms)
		runes[start], runes[i] = runes[i], runes[start]
	}
}

type Position [2]int

func getSequence(sequence string, depth int, useDir bool, current Position) int {
	fmt.Println(sequence)
	if sequence == "" {
		return 0
	}
	var keypad map[string][2]int
	if useDir {
		keypad = directionalKeypad
	} else {
		keypad = numericKeypad
	}
	row, col := 0, 0
	if current == [2]int{-1, -1} {
		row, col = keypad["A"][0], keypad["A"][1]
	} else {
		row, col = current[0], current[1]
	}
	targetRow, targetCol := keypad[string(sequence[0])][0], keypad[string(sequence[0])][1]

	dr, dc := targetRow - row, targetCol - col
	var moves string
	if dr > 0 {
		moves += strings.Repeat("v", dr)
	} else {
		moves += strings.Repeat("^", -dr)
	}
	if dc > 0 {
		moves += strings.Repeat(">", dc)
	} else {
		moves += strings.Repeat("<", -dc)
	}
	if depth <= 0 {
		return len(moves) + getSequence(sequence[1:], depth, useDir, [2]int{targetRow, targetCol}) + 1
	}
	var candidates []int
	for _, permutation := range getPermutations(moves) {
		tmpRow, tmpCol := row, col
		validPermutation := true
		for _, move := range permutation {
			dr, dc := DIRECTIONS[string(move)][0], DIRECTIONS[string(move)][1]
			tmpRow, tmpCol = tmpRow + dr, tmpCol + dc
			isValid := false
			for _, dirs := range keypad {
				if tmpRow == dirs[0] && tmpCol == dirs[1] {
					isValid = true
					break
				}
			}
			if !isValid {
				validPermutation = false
				break
			}
		}
		if validPermutation {
			candidates = append(candidates, getSequence(permutation + "A", depth - 1, true, [2]int{-1, -1}))
		}
	}
	if len(moves) == 0 {
		candidates = append(candidates, getSequence("A", depth - 1, true, [2]int{-1, -1}))
	}
	min_len := math.MaxInt64
	for _, candidate := range candidates {
		if candidate < min_len {
			min_len = candidate
		}
	}
	if min_len < 0 {
		panic("Invalid sequence: " + sequence)
	}
	return min_len + getSequence(sequence[1:], depth, useDir, [2]int{targetRow, targetCol})
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	result := 0

	for _, line := range file {
		var id int
		fmt.Sscanf(line, "%dA", &id)
		fmt.Println(line)
		length := getSequence(line, 25, false, [2]int{-1, -1})
		result += id * length
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
