package day5

import (
	"fmt"
	"log"
)

type Grid struct {
	width	int
	height	int
	data	[][]rune
}

func initGrid(vectors []Vector) Grid {
	grid := Grid{
		width:  0,
		height: 0,
		data:   nil,
	}
	for _, currentVector := range vectors {
		if currentVector.start.x > grid.width {
			grid.width = currentVector.start.x
		}
		if currentVector.end.x > grid.width {
			grid.width = currentVector.end.x
		}
		if currentVector.start.y > grid.height {
			grid.height = currentVector.start.y
		}
		if currentVector.end.y > grid.height {
			grid.height = currentVector.end.y
		}
	}
	grid.width = grid.width + 3
	grid.height = grid.height + 3
	log.Printf("Grid generated with width %d and height %d", grid.width, grid.height)

	return drawEmptyGrid(grid)
}

func draw(vectors []Vector, grid Grid) Grid {
	for _, currentVector := range vectors {
		grid = drawVector(currentVector, grid)
	}
	return grid
}

func drawVector(vector Vector, grid Grid) Grid {
	log.Printf("Drawing vector %v", vector)
	grid = drawPoint(vector.start, drawPoint(vector.end, grid))
	if vector.start.y == vector.end.y {
		if vector.start.x < vector.end.x {
			for i := vector.start.x + 1; i < vector.end.x; i++ {
				grid = drawPoint(Point{x: i, y: vector.start.y}, grid)
			}
		} else {
			for i := vector.end.x + 1; i < vector.start.x; i++ {
				grid = drawPoint(Point{x: i, y: vector.start.y}, grid)
			}
		}
	} else {
		if vector.start.y < vector.end.y {
			for i := vector.start.y + 1; i < vector.end.y; i++ {
				grid = drawPoint(Point{x: vector.start.x, y: i}, grid)
			}
		} else {
			for i := vector.end.y + 1; i < vector.start.y; i++ {
				grid = drawPoint(Point{x: vector.start.x, y: i}, grid)
			}
		}
	}
	printGrid(grid)
	fmt.Scanln()
	return grid
}

func drawPoint(point Point, grid Grid) Grid {
	currentValue := grid.data[point.y][point.x]
	if currentValue == '.' {
		grid.data[point.y][point.x] = '1'
	} else {
		newValue := string(currentValue + 1)
		grid.data[point.y][point.x] = rune(newValue[0])
	}
	return grid
}

func drawEmptyGrid(grid Grid) Grid  {
	for i := 0; i < grid.width; i++ {
		grid.data = append(grid.data, make([]rune, grid.height))
		for j := 0; j < grid.height; j++ {
			grid.data[i][j] = '.'
		}
	}
	return grid
}

func printGrid(grid Grid)  {
	for i := range grid.data {
		for j := range grid.data[i] {
			fmt.Print(string(grid.data[i][j]))
		}
		fmt.Print(string('\n'))
	}
}