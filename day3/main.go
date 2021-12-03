package main

import (
	"../utils"
	"fmt"
)

func main() {
	binaryValues := utils.ReadFileInput()
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
	if char == '0' {
		if freq > totalValue / 2 {
			filter = otherChar
		} else {
			filter = char
		}
	} else {
		if freq >= totalValue / 2 {
			filter = char
		} else {
			filter = otherChar
		}
	}
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