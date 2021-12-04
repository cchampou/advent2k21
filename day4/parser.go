package day4

import (
	"log"
	"strconv"
	"strings"
)

func generateCalledNumbers(input string) []int  {
	var res []int
	sepStrings := strings.Split(input, ",")
	for _, current := range sepStrings {
		converted, err := strconv.Atoi(current)
		if err != nil {
			log.Fatalf("Fail to convert called number %s", current)
		}
		res = append(res, converted)
	}
	return res
}

func convertLineToInts(line string) []int {
	var result []int
	newLineWithStrings := strings.Split(strings.TrimSpace(line), " ")
	for _, currentValue := range newLineWithStrings {
		if len(currentValue) > 0 {
			converted, err := strconv.Atoi(currentValue)
			if err != nil {
				log.Fatalf("Exception while converting to int: %s", currentValue)
			}
			result = append(result, converted)
		}
	}
	return result
}