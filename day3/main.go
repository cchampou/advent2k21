package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Program started")
	var binaryValues []string
	file, err := os.OpenFile("input", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Fail to open file")
		return
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		binaryValues = append(binaryValues, line)
	}

	gamma := find(binaryValues, 0, '1', '0')
	epsilon := find(binaryValues, 0, '0', '1')
	fmt.Println(gamma, epsilon)
}

func find(table []string, index int, char uint8, otherChar uint8) string {
	if len(table) == 1 || index > len(table[0]) - 1 {
		return table[0]
	}
	totalValue := float64(len(table))
	freq := .0
	var filter uint8
	for _, currentBinary := range table {
		if currentBinary[index] == char {
			freq++
		}
	}
	if freq >= totalValue / 2 {
		filter = char
	} else {
		filter = otherChar
	}
	fmt.Println(freq, totalValue / 2, filter)
	fmt.Println(filterTable(table, filter, index))
	return find(filterTable(table, filter, index), index + 1, char, otherChar)
}

func filterTable(table []string, filter uint8, index int) []string {
	var res []string
	for _, currentValue := range table {
		if currentValue[index] == filter {
			res = append(res, currentValue)
		}
	}
	return res
}