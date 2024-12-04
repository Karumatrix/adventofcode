package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFileContent(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func checkChar(file []string, char byte, x, y int) bool {
	if x >= 0 && x < len(file) && y >= 0 && y < len(file[0]) && file[x][y] == char {
		return true
	}
	return false
}

func main() {
	file := readFileContent("./input.txt")
	result := 0

	for indexX, line := range file {
		for indexY, letter := range line {
			if letter == 'A' {
				/*
					M.M
					.A.
					S.S
				*/
				if checkChar(file, 'M', indexX - 1, indexY - 1) && checkChar(file, 'M', indexX - 1, indexY + 1) && checkChar(file, 'S', indexX + 1, indexY + 1) && checkChar(file, 'S', indexX + 1, indexY - 1) {
					result++
				}
				/*
					S.M
					.A.
					S.M
				*/
				if checkChar(file, 'S', indexX - 1, indexY - 1) && checkChar(file, 'M', indexX - 1, indexY + 1) && checkChar(file, 'M', indexX + 1, indexY + 1) && checkChar(file, 'S', indexX + 1, indexY - 1) {
					result++
				}
				/*
					S.S
					.A.
					M.M
				*/
				if checkChar(file, 'S', indexX - 1, indexY - 1) && checkChar(file, 'S', indexX - 1, indexY + 1) && checkChar(file, 'M', indexX + 1, indexY + 1) && checkChar(file, 'M', indexX + 1, indexY - 1) {
					result++
				}
				/*
					M.S
					.A.
					M.S
				*/
				if checkChar(file, 'M', indexX - 1, indexY - 1) && checkChar(file, 'S', indexX - 1, indexY + 1) && checkChar(file, 'S', indexX + 1, indexY + 1) && checkChar(file, 'M', indexX + 1, indexY - 1) {
					result++
				}
			}
		}
	}
	fmt.Println(result)
}