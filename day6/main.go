package day6

import (
	"../utils"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
)

func Run() {
	PrintMemUsage()
	data := []byte(strings.Join(strings.Split(utils.ReadFileInput("day6/input")[0], ","), ""))
	println(data)
	for i := 0; i < 80; i++ {
		data = processFish(data)
		fmt.Print(i)
		PrintMemUsage()
	}
	log.Println(len(data))
}

func processFish(fish []byte) []byte {
	result := fish
	for index, value := range fish {
		intValue, _ := strconv.Atoi(string(value))
		if intValue == 0 {
			result = []byte(string(result) + "8")
			result[index] = 6
		} else {
			result[index] = strconv.Itoa(intValue - 1)[0]
		}
	}
	return result
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}