package day4

import (
	"../utils"
	"strings"
)

func GenerateCalledNumbers(input string) []int  {
	sepStrings := strings.Split(input, ",")
	return utils.ConvertLinesToInt(utils.RemoveEmptyLines(sepStrings))
}

func convertLineToInts(line string) []int {
	var result []int
	newLineWithStrings := strings.Split(strings.TrimSpace(line), " ")
	result = utils.ConvertLinesToInt(utils.RemoveEmptyLines(newLineWithStrings))
	return result
}