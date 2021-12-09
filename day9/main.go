package day9

import (
	"../utils"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	level		int
	top			int
	right		int
	bottom		int
	left		int
	isLow		bool
	isBaseIn	bool
}

func Run() {
	var sizes []int
	data := parse(utils.ReadFileInput("day9/input"))
	processed := process(data, &sizes)
	fmt.Println(score(processed))
	log.Println(scoreWith3BaseIn(sizes))
}

func scoreWith3BaseIn(sizes []int) int {
	score := 1
	sort.Ints(sizes)
	last3 := sizes[len(sizes) - 3:]
	for _, value := range last3 {
		score *= value
	}
	return score
}

func score(data [][]Point) int {
	score := 0
	for _, line := range data {
		for _, current := range line {
			if current.isLow {
				score += current.level + 1
			}
		}
	}
	return score
}

func process(data [][]Point, sizes *[]int) [][]Point {
	for i, line := range data {
		for j := range line {
			if i - 1 >= 0 {
				data[i][j].top = data[i - 1][j].level
			} else {
				data[i][j].top = -1
			}
			if j + 1 < len(line) {
				data[i][j].right = data[i][j + 1].level
			} else {
				data[i][j].right = -1
			}
			if i + 1 < len(data) {
				data[i][j].bottom = data[i + 1][j].level
			} else {
				data[i][j].bottom = -1
			}
			if j - 1 >= 0 {
				data[i][j].left = data[i][j - 1].level
			} else {
				data[i][j].left = -1
			}
			if (data[i][j].right == -1 || data[i][j].level < data[i][j].right) &&
				(data[i][j].bottom == -1 || data[i][j].level < data[i][j].bottom) &&
				(data[i][j].left == -1 || data[i][j].level < data[i][j].left) &&
				(data[i][j].top == -1 || data[i][j].level < data[i][j].top) {
				data[i][j].isLow = true
				//log.Println("Found low, checking baseIn ", i, j, data[i][j].level)
				baseInSize := 0
				findBaseIn(&data, i, j, &baseInSize)
				//log.Println("baseInSize ", baseInSize)
				*sizes = append(*sizes, baseInSize)
			} else {
				data[i][j].isLow = false
			}
		}
	}
	return data
}
func findBaseIn(data *[][]Point, i int, j int, size *int) {
	//log.Print("Checking baseIn for point ", i, j, (*data)[i][j].level)
	if (*data)[i][j].isBaseIn || (*data)[i][j].level == 9 {
		//log.Println("Already baseIn or stop")
		return
	}
	//fmt.Scanln()
	if (*data)[i][j].level != 9 {
		(*data)[i][j].isBaseIn = true
		*size++
	}
	if i < len(*data) - 1 {
		findBaseIn(data, i + 1, j, size)
	}
	if j < len((*data)[0]) - 1 {
		findBaseIn(data, i, j + 1, size)
	}
	if i > 0 {
		findBaseIn(data, i - 1, j, size)
	}
	if j > 0 {
		findBaseIn(data, i, j - 1, size)
	}
	return
}

func parse(lines []string) [][]Point {
	var output [][]Point
	for _, line := range lines {
		var pointLine []Point
		asStrings := strings.Split(line, "")
		for _, current := range asStrings {
			converted, _ := strconv.Atoi(current)
			pointLine = append(pointLine, Point{level: converted})
		}
		output = append(output, pointLine)
	}
	return output
}

func printMap(points [][]Point) {
	for _, line := range points {
		for _, current := range line {
			fmt.Print(current.level)
			if current.isBaseIn {
				fmt.Print(".")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}