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

func operandValue(operand, A, B, C int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return A
	case 5:
		return B
	case 6:
		return C
	default:
		panic(fmt.Sprintf("Invalid combo operand"))
	}
}

func aoc(filepath string) {
	file := filepathToStringArray(filepath)
	registerA, registerB, registerC := 0, 0, 0
	fmt.Sscanf(file[0], "Register A: %d", &registerA)
	parts := strings.Split(file[4], " ")
	program := strings.Split(parts[1], ",")
	ptr := 0
	output := []string{}

	for ptr < len(program) {
		opcode, _ := strconv.Atoi(program[ptr])
		operand, _ := strconv.Atoi(program[ptr + 1])
		ptr += 2

		switch opcode {
		case 0:
			denominator := 1 << operandValue(operand, registerA, registerB, registerC)
			if denominator != 0 {
				registerA = registerA / denominator
			}
		case 1:
			registerB ^= operand
		case 2:
			registerB = operandValue(operand, registerA, registerB, registerC) % 8
		case 3:
			if registerA != 0 {
				ptr = operand
			}
		case 4:
			registerB ^= registerC
		case 5:
			output = append(output, fmt.Sprintf("%d", operandValue(operand, registerA, registerB, registerC) % 8))
		case 6:
			denominator := 1 << operandValue(operand, registerA, registerB, registerC)
			if denominator != 0 {
				registerB = registerA / denominator
			}
		case 7:
			denominator := 1 << operandValue(operand, registerA, registerB, registerC)
			if denominator != 0 {
				registerC = registerA / denominator
			}
		}
	}
	fmt.Println(strings.Join(output, ","))
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("USAGE: go run <file> <input-file>")
	} else {
		aoc(os.Args[1])
	}
}
