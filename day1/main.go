package main

import (
    "../utils"
    "fmt"
    "log"
    "strconv"
)

func getSum(a []int, b int) int {
    return a[b - 2] + a[b - 1] + a[b]
}

func main() {
    measures := convertLinesToInt(utils.ReadFileInput())

    total := 0
    for index, _ := range measures {
        if index > 2 && getSum(measures, index) > getSum(measures, index - 1) {
            total++
        }
    }

    fmt.Println(total)
}

func convertLinesToInt(raw []string) []int {
    var output []int
    for _, currentValue := range raw{
        newMeasure, err := strconv.Atoi(currentValue)
        if err != nil {
            log.Fatalf("Line could not be converted to number: %v", currentValue)
        }
        output = append(output, newMeasure)
    }
    return output
}