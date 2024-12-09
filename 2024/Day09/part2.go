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
		if letter == -1 {
			isFree = true
		}
		if !isFree && letter == -1 {
			return false
		}
		if isFree && letter != -1 {
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
	pos := 0
	for posInv > 0 {
		for posInv > 0 && disk[posInv] == -1 {
			posInv--
		}
		size := 0
		for posInv - size >= 0 && disk[posInv - size] == disk[posInv] {
			size++
		}
		free := 0
		pos = 0
		for free < size {
			pos++
			if disk[pos] != -1 {
				continue
			}
			if pos > posInv {
				break
			}
			free = 0
			for pos + free < len(disk) && disk[pos + free] == -1 {
				free++
			}
		}
		if pos > posInv {
			posInv -= size
			continue
		}
		for i := 0; i < size; i++ {
			disk[pos], disk[posInv - i] = disk[posInv - i], disk[pos]
			pos++
		}
		posInv--
	}
	for i, block := range disk {
		if block != -1 {
			result += i * block
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
