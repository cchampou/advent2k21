package day5

func countScore(grid Grid) int {
	score := 0
	for i := range grid.data {
		for j := range grid.data[0] {
			raw := grid.data[i][j]
			if raw != '.' && raw > '1' {
				score++
			}
		}
	}

	return score
}