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

func checkHor(file []string, indX, indY int) (result int) {
	if indY + 1 < len(file[0]) && file[indX][indY + 1] == 'M' &&
		indY + 2 < len(file[0]) && file[indX][indY + 2] == 'A' &&
		indY + 3 < len(file[0]) && file[indX][indY + 3] == 'S' {
			result += 1
		}
	if indY - 1 >= 0 && file[indX][indY - 1] == 'M' &&
		indY - 2 >= 0 && file[indX][indY - 2] == 'A' &&
		indY - 3 >= 0 && file[indX][indY - 3] == 'S' {
			result += 1
	}
	return result
}

func checkVer(file []string, indX, indY int) (result int) {
	if indX + 1 < len(file) && file[indX + 1][indY] == 'M' &&
		indX + 2 < len(file) && file[indX + 2][indY] == 'A' &&
		indX + 3 < len(file) && file[indX + 3][indY] == 'S' {
			result += 1
	}
	if indX - 1 >= 0 && file[indX - 1][indY] == 'M' &&
		indX - 2 >= 0 && file[indX - 2][indY] == 'A' &&
		indX - 3 >= 0 && file[indX - 3][indY] == 'S' {
			result += 1
	}
	return result
}

func checkDia(file []string, indX, indY int) (result int) {
	if indX + 1 < len(file) && indY + 1 < len(file[0]) && file[indX + 1][indY + 1] == 'M' &&
		indX + 2 < len(file) && indY + 2 < len(file[0]) && file[indX + 2][indY + 2] == 'A' &&
		indX + 3 < len(file) && indY + 3 < len(file[0]) && file[indX + 3][indY + 3] == 'S' {
			result += 1
	}
	if indX - 1 >= 0 && indY - 1 >= 0 && file[indX - 1][indY - 1] == 'M' &&
		indX - 2 >= 0 && indY - 2 >= 0 && file[indX - 2][indY - 2] == 'A' &&
		indX - 3 >= 0 && indY - 3 >= 0 && file[indX - 3][indY - 3] == 'S' {
			result += 1
	}
	if indX + 1 < len(file) && indY - 1 >= 0 && file[indX + 1][indY - 1] == 'M' &&
		indX + 2 < len(file) && indY - 2 >= 0 && file[indX + 2][indY - 2] == 'A' &&
		indX + 3 < len(file) && indY - 3 >= 0 && file[indX + 3][indY - 3] == 'S' {
			result += 1
	}
	if indX - 1 >= 0 && indY + 1 < len(file[0]) && file[indX - 1][indY + 1] == 'M' &&
		indX - 2 >= 0 && indY + 2 < len(file[0]) && file[indX - 2][indY + 2] == 'A' &&
		indX - 3 >= 0 && indY + 3 < len(file[0]) && file[indX - 3][indY + 3] == 'S' {
			result += 1
	}
	return result
}

func main() {
	file := readFileContent("./input2.txt")
	result := 0

	for indexX, line := range file {
		for indexY, letter := range line {
			if letter == 'X' {
				result += checkHor(file, indexX, indexY)
				result += checkVer(file, indexX, indexY)
				result += checkDia(file, indexX, indexY)
			}
		}
	}
	fmt.Println(result)
}