package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func getSum(a []int, b int) int {
    return a[b - 2] + a[b - 1] + a[b]
}

func main() {
    var measures []int
    total := 0
    f, err := os.OpenFile("input", os.O_RDONLY, os.ModePerm)

    if err != nil {
        log.Fatalf("Open file errorL %v", err)
        return
    }

    defer f.Close()

    sc := bufio.NewScanner(f)

    for sc.Scan() {
        text := sc.Text()
        newMeasure, err := strconv.Atoi(text)
        if err != nil {
            log.Fatalf("Line could not be converted to number: %v", text)
            return
        }
        measures = append(measures, newMeasure)
    }

    for index, _ := range measures {
        if index > 2 && getSum(measures, index) > getSum(measures, index - 1) {
            total++
        }
    }

    fmt.Println(total)
}