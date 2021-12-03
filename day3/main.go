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
	var gamma = ""
	var epsilon = ""

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

	dataLength := len(binaryValues[0])
	totalValue := len(binaryValues)

	for i := 0; i < dataLength; i++ {
		freq := 0
		for _, currentBinary := range binaryValues {
			if currentBinary[i] == '1' {
				freq++
			}
		}
		if freq > totalValue / 2 {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		} else {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		}
	}
	fmt.Println(gamma, epsilon)

}
