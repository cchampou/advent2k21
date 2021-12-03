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

	gamma := findGamma(binaryValues, 0, "")
	epsilon := findEpsilon(gamma)
	fmt.Println(gamma, epsilon)
}

func findGamma(table []string, index int, currentValue string) string {
	if index > len(table[0]) - 1 {
		return currentValue
	}
	totalValue := len(table)
	freq := 0
	for _, currentBinary := range table {
		if currentBinary[index] == '1' {
			freq++
		}
	}
	if freq > totalValue / 2 {
		return findGamma(table, index + 1, currentValue + "1")
	} else {
		return findGamma(table, index + 1, currentValue + "0")
	}
}

func findEpsilon(gamma string) string  {
	epsilon := ""
	for _, currentChar := range gamma {
		if currentChar == '1' {
			epsilon += "0"
		} else {
			epsilon += "1"
		}
	}
	return epsilon
}
