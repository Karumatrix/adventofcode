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

func checkNewDisk(newDisk string) bool {
	isFree := false
	for _, letter := range newDisk {
		if letter == '.' {
			isFree = true
		}
		if !isFree && letter == '.' {
			return false
		}
		if isFree && letter != '.' {
			return false
		}
	}
	return true
}

func aoc(filepath string) {
	file := filepathToString(filepath)
	result := 0
	isFree := false
	var disk []int
	ID := 0

	for _, letter := range file {
		if letter == '\n' {
			break
		}
		value := int(letter - '0')
		if isFree {
			for i := 0; i < value; i++ {
				disk = append(disk, -1)
			}
			isFree = false
		} else {
			for i := 0; i < value; i++ {
				disk = append(disk, ID)
			}
			isFree = true
			ID++
		}
	}
	posInv := len(disk) - 1
	for pos := 0; pos < len(disk); pos++ {
		if disk[pos] == -1 {
			for disk[posInv] == -1 {
				posInv--
			}
			if posInv <= pos {
				break
			}
			disk[pos], disk[posInv] = disk[posInv], disk[pos]
		}
	}
	for i, block := range disk {
		if block == -1 {
			break
		}
		result += i * block
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
